package usecase

import "github.com/yosukei3108/LearnCleanArchitecture/src/app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
