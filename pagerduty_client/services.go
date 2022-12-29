package pagerduty_client

import (
	"context"
	"github.com/PagerDuty/go-pagerduty"
)

func GetSessionConfig(_ context.Context, config *Config) (*pagerduty.Client, error) {
	client := pagerduty.NewClient(config.Token)
	return client, nil
}
