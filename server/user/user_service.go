package user

import (
	"github.com/tfkhdyt/chat-app/server/lib/encryption"
	"github.com/tfkhdyt/fiber-toolbox/hash"
	"github.com/tfkhdyt/fiber-toolbox/response"
)

type UserService interface {
	create(payload *UserRegisterDTO) (*response.MessageResponse, error)
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

	privateKey, publicKey, err := encryption.GenerateKeys()
	if err != nil {
		return nil, err
	}

	encryptedPrivateKey, err := encryption.EncryptPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	publicKeyStr, err := encryption.MarshalPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	if err := s.repo.create(&User{
		Username:   user.Username,
		Email:      user.Email,
		Password:   user.Password,
		PublicKey:  publicKeyStr,
		PrivateKey: encryptedPrivateKey,
	}); err != nil {
		return nil, err
	}

	return response.NewMessageResponse(true, "user created successfully"), nil
}
