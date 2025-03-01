package provider

import (
	"context"
	"ops-monitor/internal/models"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewElasticSearchClient(t *testing.T) {
	client, err := NewElasticSearchClient(context.Background(), models.AlertDataSource{})
	if err != nil {
		logrus.Errorf("client -> %s", err.Error())
		return
	}

	client.Query(LogQueryOptions{})
}
