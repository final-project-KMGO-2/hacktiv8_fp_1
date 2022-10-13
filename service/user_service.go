package service

import (
	"context"
	"hacktiv8_fp_1/dto"
	"hacktiv8_fp_1/entity"
	"hacktiv8_fp_1/repository"

	"github.com/mashingan/smapping"
)

type UserService interface {
	InsertUser(ctx context.Context, userDTO dto.UserRegisterDTO) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepository: ur,
	}
}

func (s *userService) InsertUser(ctx context.Context, userDTO dto.UserRegisterDTO) (entity.User, error) {
	createdUser := entity.User{}
	err := smapping.FillStruct(&createdUser, smapping.MapFields(&userDTO))
	if err != nil {
		return createdUser, err
	}

	res, err := s.userRepository.InsertUser(ctx, createdUser)
	if err != nil {
		return createdUser, err
	}
	return res, nil
}

func (s *userService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.userRepository.GetUserByEmail(ctx, email)
}
