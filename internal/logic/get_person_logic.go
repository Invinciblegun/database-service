package logic

import (
	"context"

	"database-service/internal/svc"
	"database-service/pb"
)

type GetPersonLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonLogic {
	return &GetPersonLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonLogic) GetPerson(in *pb.GetPersonReq) (*pb.GetPersonResp, error) {
	// todo: add your logic here and delete this line
	name := RandStringRunes(10)
	phone := int32(RandNumberRunes(10))
	return &pb.GetPersonResp{
		Id:    in.Id,
		Name:  name,
		Phone: phone,
	}, nil
}
