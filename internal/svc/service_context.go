package svc

import (
	"database-service/internal/config"
	apolloConfig "git.ickey.com.cn/infra/utils/config"
	"git.ickey.com.cn/infra/utils/db"
	"git.ickey.com.cn/infra/utils/es"
	"git.ickey.com.cn/infra/utils/grpc"
	"git.ickey.com.cn/infra/utils/kafka"
	"git.ickey.com.cn/infra/utils/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ApolloConfig  map[string]string
	RedisMap      map[string]redis.Client
	KafkaProducer map[string]kafka.Producer
	KafkaConsumer map[string]kafka.Consumer
	DBMap         map[string]db.DBClient
	ESMap         map[string]*es.Client
	GrpcMap       map[string]zrpc.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	s := &ServiceContext{
		ApolloConfig:  apolloConfig.AplConfig.GetConfig("application"),
		RedisMap:      redis.GetRedisMap(),
		KafkaConsumer: kafka.GetConsumerMap(),
		KafkaProducer: kafka.GetProducerMap(),
		DBMap:         db.GetDBMap(),
		ESMap:         es.GetESMap(),
		Config:        c,
		GrpcMap:       grpc.GetGRPCMap(),
	}
	go s.addApolloListener()
	return s
}

// addApolloListener apollo更新全局配置
func (s *ServiceContext) addApolloListener() {
	cn := apolloConfig.AplConfig.OnChange()
	for <-cn {
		s.ApolloConfig = apolloConfig.AplConfig.GetConfig("application")
	}
}
