package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"ops-monitor/internal/global"
	"ops-monitor/internal/models"
	"ops-monitor/pkg/ctx"
	"ops-monitor/pkg/logger"
	"ops-monitor/pkg/tools"
	"time"
)

type userService struct {
	ctx *ctx.Context
}

type InterUserService interface {
	Search(req interface{}) (interface{}, interface{})
	List(req interface{}) (interface{}, interface{})
	Get(req interface{}) (interface{}, interface{})
	Login(req interface{}) (interface{}, interface{})
	Update(req interface{}) (interface{}, interface{})
	Register(req interface{}) (interface{}, interface{})
	Delete(req interface{}) (interface{}, interface{})
	ChangePass(req interface{}) (interface{}, interface{})
}

func newInterUserService(ctx *ctx.Context) InterUserService {
	return &userService{
		ctx: ctx,
	}
}

func (us userService) Search(req interface{}) (interface{}, interface{}) {
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户搜索")

	r := req.(*models.MemberQuery)
	data, err := us.ctx.DB.User().Search(*r)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("用户搜索失败: %v", err))
		return nil, err
	}
	logger.Info(us.ctx.Ctx, lc.WithParams(map[string]interface{}{
		"user_search_count": len(data),
	}))
	return data, nil
}

func (us userService) List(req interface{}) (interface{}, interface{}) {
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户列表获取")

	data, err := us.ctx.DB.User().List()
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("查询用户列表失败: %v", err))
		return nil, err
	}
	logger.Info(us.ctx.Ctx, lc.WithParams(map[string]interface{}{
		"user_count": len(data),
	}))
	return data, nil
}

func (us userService) Get(req interface{}) (interface{}, interface{}) {
	r := req.(*models.MemberQuery)

	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户信息获取").
		WithParams(map[string]interface{}{
			"user_id": r.UserId,
		})

	data, _, err := us.ctx.DB.User().Get(*r)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("查询用户失败: %v", err))
		return nil, err
	}
	logger.Info(us.ctx.Ctx, lc.WithParams(map[string]interface{}{
		"user_id": data.UserId,
	}))

	return data, nil
}

func (us userService) Login(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Member)

	// 获取当前上下文
	currentCtx := us.ctx.Ctx
	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(currentCtx)).
		WithAction("用户登录").
		WithParams(map[string]interface{}{
			"username": r.UserName,
		})

	// 校验 Password
	arr := md5.Sum([]byte(r.Password))
	hashPassword := hex.EncodeToString(arr[:])

	q := models.MemberQuery{
		UserName: r.UserName,
	}
	data, _, err := us.ctx.DB.User().Get(q)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("查询用户失败: %v", err))
		return nil, err
	}

	switch data.CreateBy {
	default:
		if data.Password != hashPassword {
			err := fmt.Errorf("密码错误")
			logger.Error(us.ctx.Ctx, lc, err)
			return nil, err
		}
	}

	// 生成token
	r.UserId = data.UserId
	r.Password = hashPassword
	tokenData, err := tools.GenerateToken(r.UserId, r.UserName, r.Password)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("生成token失败: %v", err))
		return nil, err
	}

	// 保存到redis
	duration := time.Duration(global.Config.Jwt.Expire) * time.Second
	us.ctx.Redis.Redis().Set("uid-"+data.UserId, tools.JsonMarshal(r), duration)

	logger.Info(us.ctx.Ctx, lc.WithParams(map[string]interface{}{
		"user_id": data.UserId,
	}))

	return tokenData, nil
}

func (us userService) Register(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Member)

	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户注册").
		WithParams(map[string]interface{}{
			"username": r.UserName,
		})

	q := models.MemberQuery{UserName: r.UserName}
	_, ok, _ := us.ctx.DB.User().Get(q)
	if ok {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("用户已存在"))
		return nil, fmt.Errorf("用户已存在")
	}

	arr := md5.Sum([]byte(r.Password))
	hashPassword := hex.EncodeToString(arr[:])
	// 在初始化admin用户时会固定一个userid，所以这里需要做一下判断；
	if r.UserId == "" {
		r.UserId = tools.RandUid()
	}

	r.Password = hashPassword
	r.CreateAt = time.Now().Unix()

	if r.CreateBy == "" {
		r.CreateBy = "system"
	}

	err := us.ctx.DB.User().Create(*r)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("用户注册失败: %v", err))
		return nil, err
	}

	logger.Info(us.ctx.Ctx, lc.WithParams(map[string]interface{}{
		"user_id": r.UserId,
	}))

	return nil, nil
}

func (us userService) Update(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Member)

	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户更新").
		WithParams(map[string]interface{}{
			"username": r.UserName,
		})

	var dbData models.Member

	db := us.ctx.DB.DB().Model(models.Member{})
	db.Where("user_id = ?", r.UserId).First(&dbData)

	r.Password = dbData.Password
	err := us.ctx.DB.User().Update(*r)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("用户更新失败: %v", err))
		return nil, err
	}

	us.ctx.DB.User().ChangeCache(r.UserId)

	return nil, nil
}

func (us userService) Delete(req interface{}) (interface{}, interface{}) {
	r := req.(*models.MemberQuery)

	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户删除").
		WithParams(map[string]interface{}{
			"username": r.UserName,
		})

	err := us.ctx.DB.User().Delete(*r)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("用户删除失败: %v", err))
		return nil, err
	}

	us.ctx.DB.User().ChangeCache(r.UserId)

	return nil, nil
}

func (us userService) ChangePass(req interface{}) (interface{}, interface{}) {
	r := req.(*models.Member)

	lc := logger.NewLogContext().
		WithTraceID(tools.GetTraceID(us.ctx.Ctx)).
		WithAction("用户修改密码").
		WithParams(map[string]interface{}{
			"username": r.UserName,
		})

	arr := md5.Sum([]byte(r.Password))
	hashPassword := hex.EncodeToString(arr[:])
	r.Password = hashPassword

	err := us.ctx.DB.User().ChangePass(*r)
	if err != nil {
		logger.Error(us.ctx.Ctx, lc, fmt.Errorf("用户修改密码失败: %v", err))
		return nil, err
	}

	us.ctx.DB.User().ChangeCache(r.UserId)

	return nil, nil
}
