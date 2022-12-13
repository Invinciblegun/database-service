package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	dts20200101 "github.com/alibabacloud-go/dts-20200101/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDtsSyncJobNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDtsSyncJobNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDtsSyncJobNameLogic {
	return &GetDtsSyncJobNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDtsSyncJobNameLogic) GetDtsSyncJobName(in *pb.GetDtsSyncJobNameReq) (*pb.GetDtsSyncJobNameResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	client, _ := clientS.GetAliDTSToken()
	describeSynchronizationJobsRequest := &dts20200101.DescribeSynchronizationJobsRequest{
		PageNum: tea.Int32(in.PageNum),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.DescribeSynchronizationJobsWithOptions(describeSynchronizationJobsRequest, runtime)
	if _err != nil {
		return nil, _err
	}
	var syncInstances []*pb.SyncInstance
	for _, syncjob := range resp.Body.SynchronizationInstances {
		syncInstance := pb.SyncInstance{
			SyncJobId:    tea.StringValue(syncjob.SynchronizationJobId),
			SyncJobName:  tea.StringValue(syncjob.SynchronizationJobName),
			SyncJobClass: tea.StringValue(syncjob.SynchronizationJobClass),
			Status:       tea.StringValue(syncjob.Status),
			Delay:        tea.StringValue(syncjob.Delay),
		}
		syncInstances = append(syncInstances, &syncInstance)
	}

	return &pb.GetDtsSyncJobNameResp{SyncJobList: syncInstances}, nil
}
