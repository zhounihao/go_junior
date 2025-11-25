package service

import (
	"context"
	"errors"
	"go-project_junior/webhook/internal/domain"
	"go-project_junior/webhook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail        = repository.ErrDuplicateEmail
	ErrInvalidUserOrPassword = errors.New("用户不存在或者密码不对")
)

// 因为要存储 所以这里集成 repository
// 也可以用大写的但是大写的
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return svc.repo.Create(ctx, u)
}
func (svc *UserService) Login(ctx context.Context, email, password string) (error, domain.User) {
	err, u := svc.repo.FindByEmail(ctx, email)
	if err == repository.ErrUserNotFound {
		return ErrInvalidUserOrPassword, domain.User{}
	}
	if err != nil {
		return err, domain.User{}
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return ErrInvalidUserOrPassword, domain.User{}
	}
	return nil, u
}
