package auth

import (
	"backend"
	"context"
	"time"
)

type AuthUsecaseInterface interface {
	Login(ctx context.Context, input backend.RequestLogin) (backend.User, error)
	Register(ctx context.Context, input backend.RegisterUser) (backend.User, error)
	ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error)
}

type AuthUsecase struct {
	repository AuthRepositoryInterface
}

func NewUsecase(repo AuthRepositoryInterface) AuthUsecaseInterface {
	return AuthUsecase{
		repository: repo,
	}
}

func (u AuthUsecase) Login(ctx context.Context, input backend.RequestLogin) (backend.User, error) {
	ctx, span := backend.Tracer.Start(ctx, backend.GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(400, 600)) * time.Millisecond)
	return u.repository.Login(ctx, input)
}
func (u AuthUsecase) Register(ctx context.Context, input backend.RegisterUser) (backend.User, error) {
	ctx, span := backend.Tracer.Start(ctx, backend.GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(130, 500)) * time.Millisecond)
	return u.repository.Register(ctx, input)
}

func (u AuthUsecase) ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error) {
	ctx, span := backend.Tracer.Start(ctx, backend.GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(20, 200)) * time.Millisecond)
	return u.repository.ForgotPassword(ctx, input)
}
