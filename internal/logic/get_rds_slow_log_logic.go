package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	"git.ickey.com.cn/infra/utils/log"
	rds20140815 "github.com/alibabacloud-go/rds-20140815/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetRdsSlowLogLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRdsSlowLogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRdsSlowLogLogic {
	return &GetRdsSlowLogLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRdsSlowLogLogic) GetRdsSlowLog(in *pb.GetRdsSlowLogReq) (*pb.GetRdsSlowLogResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	client, _ := clientS.GetAliRDSToken()
	if CheckDate(in.StartTime, in.EndTime) != nil {
		err := log.Error("error date parameters")
		return nil, err
	} else {
		startTime, endTime, _ := Local2UTC(in.StartTime, in.EndTime)
		startTime = startTime + "Z"
		endTime = endTime + "Z"
		describeSlowLogsRequest := &rds20140815.DescribeSlowLogsRequest{
			DBInstanceId: tea.String(in.DBInstanceId),
			StartTime:    tea.String(startTime),
			EndTime:      tea.String(endTime),
			DBName:       tea.String(in.DBName),
			PageNumber:   tea.Int32(in.PageNumber),
		}
		runtime := &util.RuntimeOptions{}
		resp, _err := client.DescribeSlowLogsWithOptions(describeSlowLogsRequest, runtime)
		if _err != nil {
			return nil, _err
		}
		var slowloglists []*pb.SQLSlowLog
		for _, slowlog := range resp.Body.Items.SQLSlowLog {
			slowloglist := pb.SQLSlowLog{
				SQLHASH:                   tea.StringValue(slowlog.SQLHASH),
				DBName:                    tea.StringValue(slowlog.DBName),
				SQLText:                   tea.StringValue(slowlog.SQLText),
				MySQLTotalExecutionCounts: tea.Int64Value(slowlog.MySQLTotalExecutionCounts),
				MySQLTotalExecutionTimes:  tea.Int64Value(slowlog.MySQLTotalExecutionTimes),
				MaxExecutionTime:          tea.Int64Value(slowlog.MaxExecutionTime),
				TotalLockTimes:            tea.Int64Value(slowlog.TotalLockTimes),
				MaxLockTime:               tea.Int64Value(slowlog.MaxLockTime),
				ParseTotalRowCounts:       tea.Int64Value(slowlog.ParseTotalRowCounts),
				ParseMaxRowCount:          tea.Int64Value(slowlog.ParseMaxRowCount),
				ReturnTotalRowCounts:      tea.Int64Value(slowlog.ReturnTotalRowCounts),
				ReturnMaxRowCount:         tea.Int64Value(slowlog.ReturnMaxRowCount),
			}
			slowloglists = append(slowloglists, &slowloglist)
		}
		slowcount := resp.Body
		return &pb.GetRdsSlowLogResp{
			TotalRecordCount: tea.Int32Value(slowcount.TotalRecordCount),
			DBInstanceId:     tea.StringValue(slowcount.DBInstanceId),
			StartTime:        tea.StringValue(slowcount.StartTime),
			EndTime:          tea.StringValue(slowcount.EndTime),
			SlowLogList:      slowloglists,
			Engine:           tea.StringValue(slowcount.Engine),
		}, nil
	}
}
