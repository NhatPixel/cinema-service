package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/NhatPixel/cinema-service/internal/model"
)

type CinemaRepository struct{
	db *sql.DB
}

func NewCinemaRepo(db *sql.DB) * CinemaRepository{
	return &CinemaRepository{db: db}
}

func (r *CinemaRepository) Create(n *model.Cinema) error {
	staffIDsJSON, err := json.Marshal(n.StaffIDs)
	if err != nil {
		return err
	}

	query := `
		INSERT INTO cinemas (id, name, address, location, status, manager_id, staff_ids)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err = r.db.Exec(
		query,
		n.ID,
		n.Name,
		n.Address,
		n.Location,
		n.Status,
		n.ManagerID,
		staffIDsJSON,
	)
	return err
}

func (r *CinemaRepository) Delete(id string) error {
	query := `DELETE FROM cinemas WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CinemaRepository) Update(n *model.Cinema) error {
	staffIDsJSON, err := json.Marshal(n.StaffIDs)
	if err != nil {
		return err
	}

	query := `
		UPDATE cinemas 
		SET name = ?, address = ?, location = ?, status = ?, manager_id = ?, staff_ids = ?
		WHERE id = ?
	`
	_, err = r.db.Exec(
		query,
		n.Name,
		n.Address,
		n.Location,
		n.Status,
		n.ManagerID,
		staffIDsJSON,
		n.ID,
	)
	return err
}

func (r *CinemaRepository) Get(status string, keyword string, page int, limit int) ([]model.Cinema, int, error) {
	where := ` WHERE 1=1`
	var args []any

	if status != "" {
		where += ` AND status = ?`
		args = append(args, status)
	}

	if keyword != "" {
		where += ` AND (name LIKE ? OR address LIKE ?)`
		like := "%" + keyword + "%"
		args = append(args, like, like)
	}

	countQuery := `SELECT COUNT(*) FROM cinemas` + where
	var total int
	if err := r.db.QueryRow(countQuery, args...).Scan(&total); err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit

	query := `
		SELECT id, name, address, location, status, manager_id, staff_ids
		FROM cinemas
	` + where + `
		LIMIT ? OFFSET ?
	`

	args = append(args, limit, offset)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var cinemas []model.Cinema
	for rows.Next() {
		var cinema model.Cinema
		var staffIDs sql.NullString

		if err := rows.Scan(
			&cinema.ID,
			&cinema.Name,
			&cinema.Address,
			&cinema.Location,
			&cinema.Status,
			&cinema.ManagerID,
			&staffIDs,
		); err != nil {
			return nil, 0, err
		}

		if staffIDs.Valid {
			_ = json.Unmarshal([]byte(staffIDs.String), &cinema.StaffIDs)
		}

		cinemas = append(cinemas, cinema)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, err
	}

	return cinemas, total, nil
}
