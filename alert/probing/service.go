package probing

//
//import (

//)
//
//const (
//	TaskProbing = "task:probing"
//)
//
//type ProbingService struct {
//	ctx    *ctx.Context
//	client *asynq.Client
//	server *asynq.Server
//}
//
//// NewProbingService 创建探测服务实例
//func NewProbingService(ctx *ctx.Context, redisOpt asynq.RedisClientOpt) (*ProbingService, error) {
//	client := asynq.NewClient(redisOpt)
//	server := asynq.NewServer(
//		redisOpt,
//		asynq.Config{
//			Concurrency: 10, // 并发处理的 worker 数量
//			Queues: map[string]int{
//				"probing": 10, // 队列优先级
//			},
//		},
//	)
//
//	return &ProbingService{
//		ctx:    ctx,
//		client: client,
//		server: server,
//	}, nil
//}
//
//// StartProducer 启动生产者
//func (s *ProbingService) StartProducer(rule models.ProbingRule) error {
//	payload, err := json.Marshal(rule)
//	if err != nil {
//		return fmt.Errorf("marshal rule failed: %v", err)
//	}
//
//	// 创建周期性任务
//	task := asynq.NewTask(TaskProbing, payload)
//	entryID, err := s.client.Enqueue(task,
//		asynq.Queue("probing"),
//		asynq.ProcessIn(time.Duration(rule.ProbingEndpointConfig.Strategy.EvalInterval)*time.Second),
//		asynq.MaxRetry(3),
//		asynq.Timeout(30*time.Second),
//	)
//
//	if err != nil {
//		return fmt.Errorf("enqueue task failed: %v", err)
//	}
//
//	logc.Infof(s.ctx.Ctx, "Enqueued task: %s", entryID)
//	return nil
//}
//
//// StartConsumer 启动消费者
//func (s *ProbingService) StartConsumer() error {
//	mux := asynq.NewServeMux()
//
//	// 注册任务处理器
//	mux.HandleFunc(TaskProbing, s.handleProbingTask)
//
//	logc.Info(s.ctx.Ctx, "Starting probing consumer...")
//	return s.server.Start(mux)
//}
//
//// handleProbingTask 处理探测任务
//func (s *ProbingService) handleProbingTask(ctx context.Context, t *asynq.Task) error {
//	var rule models.ProbingRule
//	if err := json.Unmarshal(t.Payload(), &rule); err != nil {
//		return fmt.Errorf("unmarshal task payload failed: %v", err)
//	}
//
//	logc.Infof(s.ctx.Ctx, "Processing probing task for rule: %s", rule.RuleId)
//
//	// 执行探测
//	eValue, err := s.runEvaluation(rule)
//	if err != nil {
//		return fmt.Errorf("run evaluation failed: %v", err)
//	}
//
//	// 处理探测结果
//	event := s.processDefaultEvent(rule)
//	event.Fingerprint = eValue.GetFingerprint()
//	event.Metric = eValue.GetLabels()
//
//	// 设置指标值
//	var isValue float64
//	if rule.RuleType != provider.TCPEndpointProvider {
//		event.Metric["value"] = eValue[rule.ProbingEndpointConfig.Strategy.Field].(float64)
//	} else {
//		if eValue["IsSuccessful"] == true {
//			isValue = 1
//		}
//		event.Metric["value"] = isValue
//	}
//
//	event.Annotations = tools.ParserVariables(rule.Annotations, event.Metric)
//
//	// 评估策略
//	var option EvalStrategy
//	if rule.RuleType != provider.TCPEndpointProvider {
//		option = EvalStrategy{
//			Operator:      rule.ProbingEndpointConfig.Strategy.Operator,
//			QueryValue:    eValue[rule.ProbingEndpointConfig.Strategy.Field].(float64),
//			ExpectedValue: rule.ProbingEndpointConfig.Strategy.ExpectedValue,
//		}
//	} else {
//		option = EvalStrategy{
//			Operator:      "==",
//			QueryValue:    isValue,
//			ExpectedValue: 1,
//		}
//	}
//
//	// 保存探测值
//	if err := SetProbingValueMap(event.GetProbingMappingKey(), eValue); err != nil {
//		return fmt.Errorf("set probing value failed: %v", err)
//	}
//
//	// 评估结果
//	s.evaluation(event, option)
//	return nil
//}
//
//// Shutdown 优雅关闭服务
//func (s *ProbingService) Shutdown() {
//	s.client.Close()
//	s.server.Shutdown()
//}
//
//// 以下是辅助方法...
//func (s *ProbingService) runEvaluation(rule models.ProbingRule) (provider.EndpointValue, error) {
//	// 保持原有的 runEvaluation 实现
//}
//
//func (s *ProbingService) processDefaultEvent(rule models.ProbingRule) models.ProbingEvent {
//	// 保持原有的 processDefaultEvent 实现
//}
//
//func (s *ProbingService) evaluation(event models.ProbingEvent, option EvalStrategy) {
//	// 保持原有的 evaluation 实现逻辑
//}
