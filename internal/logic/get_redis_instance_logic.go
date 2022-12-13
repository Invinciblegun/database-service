package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	"git.ickey.com.cn/infra/utils/log"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetRedisInstanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRedisInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedisInstanceLogic {
	return &GetRedisInstanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRedisInstanceLogic) GetRedisInstance(in *pb.GetRedisInstanceReq) (*pb.GetRedisInstanceResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var redislist []*pb.RedisInstance
	client, _err := clientS.GetAliRedisToken()
	if _err != nil {
		return nil, _err
	}

	redisRegions := l.svcCtx.ApolloConfig["RedisRegionIds"]
	regionIds := string_.Split(&redisRegions, tea.String(","), tea.Int(10))
	for _, regionId := range regionIds {
		describeRedisInstancesRequest := &r_kvstore20150101.DescribeInstancesRequest{
			RegionId: regionId,
		}
		runtime := &util.RuntimeOptions{}
		resp, _err := client.DescribeInstancesWithOptions(describeRedisInstancesRequest, runtime)
		if _err != nil {
			log.WithContext(l.ctx).Errorf(_err.Error())
			return nil, _err
		}
		console.Log(util.ToJSONString(tea.ToMap(resp)))
		for _, redis := range resp.Body.Instances.KVStoreInstance {
			redisinfo := pb.RedisInstance{
				InstanceId:       tea.StringValue(redis.InstanceId),
				InstanceName:     tea.StringValue(redis.InstanceName),
				InstanceType:     tea.StringValue(redis.InstanceType),
				EngineVersion:    tea.StringValue(redis.EngineVersion),
				EditionType:      tea.StringValue(redis.EditionType),
				ArchitectureType: tea.StringValue(redis.ArchitectureType),
				NodeType:         tea.StringValue(redis.NodeType),
				InstanceStatus:   tea.StringValue(redis.InstanceStatus),
				RegionId:         tea.StringValue(redis.RegionId),
				ConnectionDomain: tea.StringValue(redis.ConnectionDomain),
			}
			redislist = append(redislist, &redisinfo)
		}
	}
	return &pb.GetRedisInstanceResp{RedisList: redislist}, nil
}
