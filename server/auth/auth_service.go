package auth

import (
	"github.com/tfkhdyt/chat-app/server/lib/encryption"
	"github.com/tfkhdyt/chat-app/server/user"
	"github.com/tfkhdyt/fiber-toolbox/hash"
	"github.com/tfkhdyt/fiber-toolbox/jwt"
)

type AuthService interface {
	login(payload *LoginRequestDTO) (*LoginResponseDTO, error)
}

type authService struct {
	repo user.UserRepo
}

func NewAuthService(repo user.UserRepo) AuthService {
	return &authService{repo: repo}
}

func (s *authService) login(payload *LoginRequestDTO) (*LoginResponseDTO, error) {
	user, err := s.repo.FindByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if err := hash.VerifyPassword(payload.Password, user.Password); err != nil {
		return nil, err
	}

	accessToken, err := jwt.GenerateJWT(&JwtClaims{ID: user.ID}, jwt.Access)
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.GenerateJWT(&JwtClaims{ID: user.ID}, jwt.Refresh)
	if err != nil {
		return nil, err
	}

	privateKey, err := encryption.DecryptPrivateKey(user.PrivateKey)
	if err != nil {
		return nil, err
	}

	privateKeyStr, err := encryption.MarshalPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	resp := &LoginResponseDTO{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		PrivateKey:   privateKeyStr,
	}

	return resp, nil
}
