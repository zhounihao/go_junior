package repository

import (
	"context"
	"go-project_junior/webhook/internal/domain"
	"go-project_junior/webhook/internal/repository/dao"
)

var (
	ErrDuplicateEmail = dao.ErrDuplicateEmail
	ErrUserNotFound   = dao.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDao
}

func NewUserRepository(dao *dao.UserDao) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {

	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (error, domain.User) {
	// 先从cache 里面找
	// 再从dao里面找
	// 再写会cache
	err, u := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return err, domain.User{}
	}
	return nil, toDomain(u)
}
func toDomain(u dao.User) domain.User {
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}
