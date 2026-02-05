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
	cinema.ID = uuid.NewString()
	return s.repo.Create(cinema)
}

func (s *CinemaService) Update(cinema *cinemaModel.Cinema) error {
	return s.repo.Update(cinema)
}

func (s *CinemaService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *CinemaService) Get(status string, keyword string, page int, limit int) ([]cinemaModel.Cinema, int, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	return s.repo.Get(status, keyword, page, limit)
}

