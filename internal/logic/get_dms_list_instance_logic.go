package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	dms_enterprise20181101 "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDmsListInstanceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDmsListInstanceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDmsListInstanceLogic {
	return &GetDmsListInstanceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDmsListInstanceLogic) GetDmsListInstance(in *pb.GetDmsListInstanceReq) (*pb.GetDmsListInstanceResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var instancelist []*pb.DmsInstance
	client, _ := clientS.GetAliDMSToken()
	listInstancesRequest := &dms_enterprise20181101.ListInstancesRequest{
		SearchKey:     tea.String(in.SearchKey),
		DbType:        tea.String(in.DbType),
		InstanceState: tea.String("NORMAL"),
		EnvType:       tea.String(in.EnvType),
		PageNumber:    tea.Int32(in.PageNumber),
		PageSize:      tea.Int32(20),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.ListInstancesWithOptions(listInstancesRequest, runtime)
	if _err != nil {
		return nil, _err
	}
	for _, instance := range resp.Body.InstanceList.Instance {
		instanceinfo := pb.DmsInstance{
			InstanceSource: tea.StringValue(instance.InstanceSource),
			InstanceId:     tea.StringValue(instance.InstanceId),
			InstanceAlias:  tea.StringValue(instance.InstanceAlias),
			InstanceType:   tea.StringValue(instance.InstanceType),
			Host:           tea.StringValue(instance.Host),
			Port:           tea.Int32Value(instance.Port),
			EnvType:        tea.StringValue(instance.EnvType),
			GroupMode:      tea.StringValue(instance.StandardGroup.GroupMode),
			State:          tea.StringValue(instance.State),
			EcsRegion:      tea.StringValue(instance.EcsRegion),
		}
		instancelist = append(instancelist, &instanceinfo)
	}

	return &pb.GetDmsListInstanceResp{DmsInstanceList: instancelist}, nil
}
