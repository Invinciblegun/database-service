package logic

import (
	"context"
	"database-service/internal/svc"
	"database-service/pb"
	dms_enterprise20181101 "github.com/alibabacloud-go/dms-enterprise-20181101/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetDmsListOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetDmsListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDmsListOrderLogic {
	return &GetDmsListOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetDmsListOrderLogic) GetDmsListOrder(in *pb.GetDmsListOrderReq) (*pb.GetDmsListOrderResp, error) {
	clientS := &ALiClientLogic{svcCtx: l.svcCtx}
	var orderlist []*pb.DmsOrder
	client, _ := clientS.GetAliDMSToken()
	listOrdersRequest := &dms_enterprise20181101.ListOrdersRequest{
		OrderStatus: tea.String(in.OrderStatus),
	}
	runtime := &util.RuntimeOptions{}
	resp, _err := client.ListOrdersWithOptions(listOrdersRequest, runtime)
	if _err != nil {
		//log.WithContext(l.ctx).Errorf(_err.Error())
		return nil, _err
	}
	for _, order := range resp.Body.Orders.Order {
		orderinfo := pb.DmsOrder{
			StatusDesc:     tea.StringValue(order.StatusDesc),
			Comment:        tea.StringValue(order.Comment),
			CreateTime:     tea.StringValue(order.CreateTime),
			Committer:      tea.StringValue(order.Committer),
			OrderId:        tea.Int64Value(order.OrderId),
			LastModifyTime: tea.StringValue(order.LastModifyTime),
			PluginType:     tea.StringValue(order.PluginType),
			CommitterId:    tea.Int64Value(order.CommitterId),
			StatusCode:     tea.StringValue(order.StatusCode),
		}
		orderlist = append(orderlist, &orderinfo)
	}
	return &pb.GetDmsListOrderResp{DmsOrderList: orderlist}, nil
}
