package logic

import (
	"context"

	"database-service/internal/svc"
	"database-service/pb"
	dms_enterprise20181101 "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDmsSearchDatabaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDmsSearchDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDmsSearchDatabaseLogic {
	return &GetDmsSearchDatabaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDmsSearchDatabaseLogic) GetDmsSearchDatabase(in *pb.GetDmsSearchDatabaseReq) (*pb.GetDmsSearchDatabaseResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var searchDatabase []*pb.DmsDatabase
	client, _ := clientS.GetAliDMSToken()
	tid := clientS.GetDmsTid()
	runtime := &util.RuntimeOptions{}
	searchDatabaseRequest := &dms_enterprise20181101.SearchDatabaseRequest{
		PageNumber: tea.Int32(in.PageNumber),
		Tid:        tid,
		SearchKey:  tea.String(in.SearchKey),
		DbType:     tea.String(in.DbType),
		EnvType:    tea.String(in.EnvType),
	}
	resp, _err := client.SearchDatabaseWithOptions(searchDatabaseRequest, runtime)
	if _err != nil {
		return nil, _err
	}

	for _, db := range resp.Body.SearchDatabaseList.SearchDatabase {
		database := pb.DmsDatabase{
			SchemaName: tea.StringValue(db.SchemaName),
			DatabaseId: tea.StringValue(db.DatabaseId),
			Encoding:   tea.StringValue(db.Encoding),
			DbType:     tea.StringValue(db.DbType),
			EnvType:    tea.StringValue(db.EnvType),
			Alias:      tea.StringValue(db.Alias),
			Host:       tea.StringValue(db.Host),
		}
		searchDatabase = append(searchDatabase, &database)
	}

	return &pb.GetDmsSearchDatabaseResp{SearchDatabase: searchDatabase}, nil
}
