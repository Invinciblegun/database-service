// Code generated by goctl. DO NOT EDIT!
// Source: database-service.proto

package server

import (
	"context"

	"database-service/internal/logic"
	"database-service/internal/svc"
	"database-service/pb"
	"encoding/json"
	"git.ickey.com.cn/sfs/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DatabaseServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedDatabaseServiceServer
}

func NewDatabaseServiceServer(svcCtx *svc.ServiceContext) *DatabaseServiceServer {
	return &DatabaseServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *DatabaseServiceServer) Ping(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	resp, err := l.Ping(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetPerson(ctx context.Context, in *pb.GetPersonReq) (*pb.GetPersonResp, error) {
	l := logic.NewGetPersonLogic(ctx, s.svcCtx)
	resp, err := l.GetPerson(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetRdsInstance(ctx context.Context, in *pb.GetRdsInstanceReq) (*pb.GetRdsInstanceResp, error) {
	l := logic.NewGetRdsInstanceLogic(ctx, s.svcCtx)
	resp, err := l.GetRdsInstance(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetRedisInstance(ctx context.Context, in *pb.GetRedisInstanceReq) (*pb.GetRedisInstanceResp, error) {
	l := logic.NewGetRedisInstanceLogic(ctx, s.svcCtx)
	resp, err := l.GetRedisInstance(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetDmsListOrder(ctx context.Context, in *pb.GetDmsListOrderReq) (*pb.GetDmsListOrderResp, error) {
	l := logic.NewGetDmsListOrderLogic(ctx, s.svcCtx)
	resp, err := l.GetDmsListOrder(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetDmsListUsers(ctx context.Context, in *pb.GetDmsListUsersReq) (*pb.GetDmsListUsersResp, error) {
	l := logic.NewGetDmsListUsersLogic(ctx, s.svcCtx)
	resp, err := l.GetDmsListUsers(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetDmsListInstance(ctx context.Context, in *pb.GetDmsListInstanceReq) (*pb.GetDmsListInstanceResp, error) {
	l := logic.NewGetDmsListInstanceLogic(ctx, s.svcCtx)
	resp, err := l.GetDmsListInstance(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetDmsSearchDatabase(ctx context.Context, in *pb.GetDmsSearchDatabaseReq) (*pb.GetDmsSearchDatabaseResp, error) {
	l := logic.NewGetDmsSearchDatabaseLogic(ctx, s.svcCtx)
	resp, err := l.GetDmsSearchDatabase(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetDbsBackupPlan(ctx context.Context, in *pb.GetDbsBackupPlanReq) (*pb.GetDbsBackupPlanResp, error) {
	l := logic.NewGetDbsBackupPlanLogic(ctx, s.svcCtx)
	resp, err := l.GetDbsBackupPlan(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetDtsSyncJobName(ctx context.Context, in *pb.GetDtsSyncJobNameReq) (*pb.GetDtsSyncJobNameResp, error) {
	l := logic.NewGetDtsSyncJobNameLogic(ctx, s.svcCtx)
	resp, err := l.GetDtsSyncJobName(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) GetRdsSlowLog(ctx context.Context, in *pb.GetRdsSlowLogReq) (*pb.GetRdsSlowLogResp, error) {
	l := logic.NewGetRdsSlowLogLogic(ctx, s.svcCtx)
	resp, err := l.GetRdsSlowLog(in)
	return resp, s.convertIckeyErrorToGRPCError(err)
}

func (s *DatabaseServiceServer) convertIckeyErrorToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(errors.Err); !ok {
		return err
	}

	b, _ := json.Marshal(err)
	errStatus := status.New(codes.Unknown, err.Error())
	ds, err := errStatus.WithDetails(&pb.IckeyError{
		Data: b,
	})
	if err != nil {
		return errStatus.Err()
	}
	return ds.Err()
}
