package service

import (
	"github.com/google/uuid"

	cinemaModel "github.com/NhatPixel/cinema-service/internal/model"
	cinemaRepo "github.com/NhatPixel/cinema-service/internal/repository"
)

type CinemaService struct {
	repo *cinemaRepo.CinemaRepository
}

func NewCinemaService(repo *cinemaRepo.CinemaRepository) *CinemaService {
	return &CinemaService{
		repo: repo,
	}
}

func (s *CinemaService) Create(cinema *cinemaModel.Cinema) error {
	if cinema.ID == "" {
		cinema.ID = uuid.NewString()
	}
	return s.repo.Create(cinema)
}