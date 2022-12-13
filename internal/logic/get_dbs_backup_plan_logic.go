package logic

import (
	"context"
	dbs20190306 "github.com/alibabacloud-go/dbs-20190306/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"

	"database-service/internal/svc"
	"database-service/pb"
)

type GetDbsBackupPlanLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDbsBackupPlanLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDbsBackupPlanLogic {
	return &GetDbsBackupPlanLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDbsBackupPlanLogic) GetDbsBackupPlan(in *pb.GetDbsBackupPlanReq) (*pb.GetDbsBackupPlanResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var backuplist []*pb.BackupPlan
	client, _ := clientS.GetAliDBSToken()
	describeBackupPlanListRequest := &dbs20190306.DescribeBackupPlanListRequest{
		PageNum:  tea.Int32(in.PageNum),
		PageSize: tea.Int32(10),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.DescribeBackupPlanListWithOptions(describeBackupPlanListRequest, runtime)
	if _err != nil {
		return nil, _err
	}
	for _, backup := range resp.Body.Items.BackupPlanDetail {
		//fmt.Println(backup)
		backupplan := pb.BackupPlan{
			BackupPlanId:   tea.StringValue(backup.BackupPlanId),
			BackupPlanName: tea.StringValue(backup.BackupPlanName),
			//DatabaseType:tea.StringValue(backup.DatabaseType),
			OSSBucketRegion:  tea.StringValue(backup.OSSBucketRegion),
			BackupMethod:     tea.StringValue(backup.BackupMethod),
			InstanceClass:    tea.StringValue(backup.InstanceClass),
			BackupPlanStatus: tea.StringValue(backup.BackupPlanStatus),
		}
		backuplist = append(backuplist, &backupplan)
	}
	return &pb.GetDbsBackupPlanResp{BackupPlanList: backuplist}, nil
}
