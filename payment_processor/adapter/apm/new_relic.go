package apm

import (
	"github.com/caiofernandes00/payment-gateway/util"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApm(config *util.Config) (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(config.NewRelicConfigAppName),
		newrelic.ConfigLicense(config.NewRelicConfigLicense),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
}
