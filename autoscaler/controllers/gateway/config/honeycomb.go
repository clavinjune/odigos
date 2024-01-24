package config

import (
	odigosv1 "github.com/keyval-dev/odigos/api/odigos/v1alpha1"
	commonconf "github.com/keyval-dev/odigos/autoscaler/controllers/common"
	"github.com/keyval-dev/odigos/common"
)

type Honeycomb struct{}

func (h *Honeycomb) DestType() common.DestinationType {
	return common.HoneycombDestinationType
}

func (h *Honeycomb) ModifyConfig(dest *odigosv1.Destination, currentConfig *commonconf.Config) {
	if isTracingEnabled(dest) {
		exporterName := "otlp/honeycomb-" + dest.Name
		currentConfig.Exporters[exporterName] = commonconf.GenericMap{
			"endpoint": "api.honeycomb.io:443",
			"headers": commonconf.GenericMap{
				"x-honeycomb-team": "${HONEYCOMB_API_KEY}",
			},
		}

		tracePipelineName := "traces/honeycomb-" + dest.Name
		currentConfig.Service.Pipelines[tracePipelineName] = commonconf.Pipeline{
			Receivers:  []string{"otlp"},
			Processors: []string{"batch"},
			Exporters:  []string{exporterName},
		}
	}
}
