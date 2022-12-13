package logic

import (
	"context"

	"database-service/internal/svc"
	"database-service/pb"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping(in *pb.Request) (*pb.Response, error) {
	// todo: add your logic here and delete this line

	return &pb.Response{}, nil
}
