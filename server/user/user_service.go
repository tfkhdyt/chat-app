package user

import (
	"crypto/x509"
	"encoding/base64"

	"github.com/tfkhdyt/chat-app/server/lib/encryption"
	"github.com/tfkhdyt/fiber-toolbox/exception"
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

	privateKey, publicKey, err := encryption.GenerateKeys()
	if err != nil {
		return nil, exception.NewInternalServerError("failed to generate keys", err)
	}

	encryptedPrivateKey, err := encryption.EncryptPrivateKey(privateKey)
	if err != nil {
		return nil, exception.NewInternalServerError("failed to encrypt private key", err)
	}

	publicKeyBytes := x509.MarshalPKCS1PublicKey(publicKey)
	publicKeyStr := base64.StdEncoding.EncodeToString(publicKeyBytes)

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
