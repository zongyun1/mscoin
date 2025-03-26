package logic

import (
	"context"
	"grpc-common/ucenter/register"
	"ucenter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *RegisterLogic) RegisterByPhone(in *register.RegReq) (res *register.RegRes, err error) {
	logx.Info("ucenter rpc register by phone call")
	return &register.RegRes{}, nil
}
