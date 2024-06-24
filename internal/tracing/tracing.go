package tracing

import (
	"github.com/uber/jaeger-client-go/config"
	"log"
)

func Init(appName string) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: "localhost:6831",
		},
	}

	_, err := cfg.InitGlobalTracer(appName)
	if err != nil {
		log.Fatal(err)
	}
}
