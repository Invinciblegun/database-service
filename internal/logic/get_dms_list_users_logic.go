package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	dms_enterprise20181101 "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDmsListUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDmsListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDmsListUsersLogic {
	return &GetDmsListUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDmsListUsersLogic) GetDmsListUsers(in *pb.GetDmsListUsersReq) (*pb.GetDmsListUsersResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var userlist []*pb.DmsUserInfo
	client, _ := clientS.GetAliDMSToken()
	tenantID := clientS.GetDmsTid()
	runtime := &util.RuntimeOptions{}
	listUsersRequest := &dms_enterprise20181101.ListUsersRequest{
		Tid:        tenantID,
		PageNumber: tea.Int32(in.PageNumber),
		//PageSize:   tea.Int32(20),
	}

	resp, _err := client.ListUsersWithOptions(listUsersRequest, runtime)
	if _err != nil {
		return nil, _err
	}
	//fmt.Println(resp.Body.UserList)

	for _, user := range resp.Body.UserList.User {
		userinfo := pb.DmsUserInfo{
			NickName: tea.StringValue(user.NickName),
			UserId:   tea.StringValue(user.UserId),
			State:    tea.StringValue(user.State),
		}
		userlist = append(userlist, &userinfo)
	}
	return &pb.GetDmsListUsersResp{ListUser: userlist}, nil
}
