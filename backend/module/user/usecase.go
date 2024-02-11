package user

import (
	"backend"
	"context"
	"go.elastic.co/apm/v2"
	"time"
)

type UsecaseInterface interface {
	Login(ctx context.Context, input backend.RequestLogin) (backend.User, error)
	Register(ctx context.Context, input backend.RegisterUser) (backend.User, error)
	ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error)
	GetList(ctx context.Context) ([]string, error)
	Update(ctx context.Context, input backend.User) (backend.User, error)
	ExportPDF(ctx context.Context) ([]string, error)
}

type Usecase struct {
	repository RepositoryInterface
}

func NewUsecase(repo RepositoryInterface) UsecaseInterface {
	return Usecase{
		repository: repo,
	}
}

func (u Usecase) Login(ctx context.Context, input backend.RequestLogin) (backend.User, error) {

	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Usecase")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)
	return u.repository.Login(ctx, input)
}
func (u Usecase) Register(ctx context.Context, input backend.RegisterUser) (backend.User, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Usecase")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)
	res, err := u.repository.Register(ctx, input)

	err = backend.SendEmail(ctx, input.Email)

	return res, err
}

func (u Usecase) ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Usecase")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)
	res, err := u.repository.ForgotPassword(ctx, input)
	err = backend.SendEmail(ctx, input.Email)

	return res, err
}

func (u Usecase) GetList(ctx context.Context) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Usecase")
	defer span.End()
	return u.repository.GetList(ctx)
}

func (u Usecase) ExportPDF(ctx context.Context) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Usecase")
	defer span.End()

	res, err := u.repository.GetList(ctx)
	res, err = backend.ExportPDF(ctx, res)

	return res, err
}

func (u Usecase) Update(ctx context.Context, input backend.User) (backend.User, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Usecase")
	defer span.End()

	return u.repository.Update(ctx, input)
}
