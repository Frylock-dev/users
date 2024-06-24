package interceptor

import (
	"context"
	openlog "github.com/opentracing/opentracing-go/log"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

func LogInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		log.WithFields(log.Fields{
			"method":   info.FullMethod,
			"req":      req,
			"res":      res,
			"duration": time.Since(now),
		}).Error("request failed")
		openlog.Error(err)
	} else {
		log.WithFields(log.Fields{
			"method":   info.FullMethod,
			"req":      req,
			"res":      res,
			"duration": time.Since(now),
		}).Info("request complete")
	}

	return res, err
}
