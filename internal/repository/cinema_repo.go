package repository

import (
	"database/sql"

	"github.com/NhatPixel/cinema-service/internal/model"
)

type CinemaRepository struct{
	db *sql.DB
}

func NewCinemaRepo(db *sql.DB) * CinemaRepository{
	return &CinemaRepository{db: db}
}

func (r *CinemaRepository) Create(n *model.Cinema) error {
	query := `
		INSERT INTO cinemas (id, name, address, location, status, manager_id)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, n.ID, n.Name, n.Address, n.Location, n.Status, n.ManagerID)
	return err
}