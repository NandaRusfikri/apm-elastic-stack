package auth

import (
	"backend"
	"context"
	"fmt"
	"time"
)

type AuthRepositoryInterface interface {
	Login(ctx context.Context, input backend.RequestLogin) (backend.User, error)
	Register(ctx context.Context, input backend.RegisterUser) (backend.User, error)
	ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error)
}

type Repository struct {
}

func NewAuthRepository() AuthRepositoryInterface {
	return Repository{}
}

func (r Repository) Login(ctx context.Context, input backend.RequestLogin) (backend.User, error) {
	ctx, span := backend.Tracer.Start(ctx, backend.GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(500, 800)) * time.Millisecond)

	if input.Email == "email" && input.Password == "password" {
		return backend.User{
			Name: "nanda",
		}, nil
	} else if input.Email == "error" && input.Password == "password" {
		panic("error error")
	} else {
		return backend.User{}, fmt.Errorf("gada")
	}

}

func (r Repository) Register(ctx context.Context, input backend.RegisterUser) (backend.User, error) {
	ctx, span := backend.Tracer.Start(ctx, backend.GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(300, 700)) * time.Millisecond)

	return backend.User{
		Name:  input.Name,
		Email: input.Email,
	}, nil

}

func (r Repository) ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error) {
	ctx, span := backend.Tracer.Start(ctx, backend.GetCurrentFunctionName())
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(600, 1700)) * time.Millisecond)

	return backend.User{Name: input.Email}, nil

}
