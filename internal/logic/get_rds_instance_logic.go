package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	"git.ickey.com.cn/infra/utils/log"
	string_ "github.com/alibabacloud-go/darabonba-string/client"
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v2/client"
	console "github.com/alibabacloud-go/tea-console/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetRdsInstanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRdsInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRdsInstanceLogic {
	return &GetRdsInstanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRdsInstanceLogic) GetRdsInstance(in *pb.GetRdsInstanceReq) (*pb.GetRdsInstanceResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var rdslist []*pb.RdsInstance
	client, _err := clientS.GetAliRDSToken()
	if _err != nil {
		return nil, _err
	}

	rdsRegions := l.svcCtx.ApolloConfig["RdsRegionIds"]
	regionIds := string_.Split(&rdsRegions, tea.String(","), tea.Int(10))
	for _, regionId := range regionIds {
		describeDBInstancesRequest := &rds20140815.DescribeDBInstancesRequest{
			RegionId: regionId,
		}
		runtime := &util.RuntimeOptions{}
		resp, _err := client.DescribeDBInstancesWithOptions(describeDBInstancesRequest, runtime)
		if _err != nil {
			log.WithContext(l.ctx).Errorf(_err.Error())
			return nil, _err
		}
		console.Log(util.ToJSONString(tea.ToMap(resp)))
		for _, db := range resp.Body.Items.DBInstance {
			dbinfo := pb.RdsInstance{
				DBInstanceId:          tea.StringValue(db.DBInstanceId),
				DBInstanceDescription: tea.StringValue(db.DBInstanceDescription),
				Engine:                tea.StringValue(db.Engine),
				EngineVersion:         tea.StringValue(db.EngineVersion),
				DBInstanceType:        tea.StringValue(db.DBInstanceType),
				DBInstanceStatus:      tea.StringValue(db.DBInstanceStatus),
				RegionId:              tea.StringValue(db.RegionId),
				ConnectionString:      tea.StringValue(db.ConnectionString),
			}
			rdslist = append(rdslist, &dbinfo)
		}
	}
	return &pb.GetRdsInstanceResp{RdsList: rdslist}, nil
}
