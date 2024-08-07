package user

import (
	"github.com/tfkhdyt/fiber-toolbox/hash"
	"github.com/tfkhdyt/fiber-toolbox/response"
)

type UserService interface {
	create(user *UserRegisterDTO) (*response.MessageResponse, error)
}

type userService struct {
	repo UserRepo
}

func NewUserService(repo UserRepo) UserService {
	return &userService{repo: repo}
}

func (s *userService) create(user *UserRegisterDTO) (*response.MessageResponse, error) {
	if err := s.repo.verifyUsernameAvailability(user.Username); err != nil {
		return nil, err
	}

	if err := s.repo.verifyEmailAvailability(user.Email); err != nil {
		return nil, err
	}

	hashedPwd, err := hash.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPwd

	if err := s.repo.create(&User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return nil, err
	}

	return response.NewMessageResponse(true, "user created successfully"), nil
}
