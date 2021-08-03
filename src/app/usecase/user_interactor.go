package usecase

import "github.com/yosukei3108/LearnCleanArchitecture/src/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u domain.User) (err error) {
	_, err := interactor.UserRepository.Store()
	return
}

func (interactor *UserInteractor) Users(user domain.Users) (err error) {
	user, err := interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (u domain.User) (err error) {
	user, err := interactor.UserRepository.FindById(identifier)
	return
}
