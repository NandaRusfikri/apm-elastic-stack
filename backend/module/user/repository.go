package user

import (
	"backend"
	"context"
	"fmt"
	"go.elastic.co/apm/v2"
	"strconv"
	"time"
)

type RepositoryInterface interface {
	Login(ctx context.Context, input backend.RequestLogin) (backend.User, error)
	Register(ctx context.Context, input backend.RegisterUser) (backend.User, error)
	ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error)
	GetList(ctx context.Context) ([]string, error)
	Update(ctx context.Context, input backend.User) (backend.User, error)
}

type Repository struct {
}

func NewRepository() RepositoryInterface {
	return Repository{}
}

func (r Repository) Login(ctx context.Context, input backend.RequestLogin) (backend.User, error) {

	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Repository")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)

	id, _ := strconv.Atoi(input.Email)

	isPrime := true
	for i := 2; i < id; i++ {
		if id%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		return backend.User{}, fmt.Errorf("not found")
	} else {
		return backend.User{
			Name: "nanda",
		}, nil
	}

}

func (r Repository) Register(ctx context.Context, input backend.RegisterUser) (backend.User, error) {

	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Repository")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)

	id, _ := strconv.Atoi(input.Email)

	isPrime := true
	for i := 2; i < id; i++ {
		if id%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		panic("kaget ada error Prime ga jelas")
	}

	return backend.User{
		Name:  input.Name,
		Email: input.Email,
	}, nil

}

func (r Repository) ForgotPassword(ctx context.Context, input backend.ForgotPassword) (backend.User, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Repository")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)

	return backend.User{Name: input.Email}, nil

}

func (r Repository) GetList(ctx context.Context) ([]string, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Repository")
	defer span.End()

	var ListUser []string
	//ceritanya query banyak data
	for i := 0; i < backend.RandInt(1000000, 5000000); i++ {
		ListUser = append(ListUser, fmt.Sprintf("User-%v", i))
	}
	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)

	return ListUser, nil

}

func (r Repository) Update(ctx context.Context, input backend.User) (backend.User, error) {
	span, ctx := apm.StartSpan(ctx, backend.GetCurrentFunctionName(), "Repository")
	defer span.End()

	time.Sleep(time.Duration(backend.RandInt(200, 1500)) * time.Millisecond)

	return backend.User{
		Name:  input.Name,
		Email: input.Email,
	}, nil

}
