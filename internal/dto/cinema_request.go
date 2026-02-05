package dto

import (
	"github.com/NhatPixel/cinema-service/internal/model"
)

type CreateRequest struct {
	Name      string   `json:"name" binding:"required,min=1,max=100"`
	Address   string   `json:"address"`
	Location  string   `json:"location"`
	Status    string   `json:"status"`
	ManagerID string   `json:"manager_id"`
	StaffIDs []string `json:"staff_ids"`
}

func (r *CreateRequest) ToModel() *model.Cinema {
	return &model.Cinema{
		Name:       r.Name,
		Address:    r.Address,
		Location:   r.Location,
		Status:     r.Status,
		ManagerID:  r.ManagerID,
		StaffIDs: r.StaffIDs,
	}
}

type GetRequest struct {
	Status  string `json:"status"`
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
	Limit   int    `json:"limit"`
}

type UpdateRequest struct {
	ID        string   `json:"id" binding:"required"`
	Name      string   `json:"name" binding:"required,min=1,max=100"`
	Address   string   `json:"address"`
	Location  string   `json:"location"`
	Status    string   `json:"status"`
	ManagerID string   `json:"manager_id"`
	StaffIDs []string `json:"staff_ids"`
}

func (r *UpdateRequest) ToModel() *model.Cinema {
	return &model.Cinema{
		ID:         r.ID,
		Name:       r.Name,
		Address:    r.Address,
		Location:   r.Location,
		Status:     r.Status,
		ManagerID:  r.ManagerID,
		StaffIDs: r.StaffIDs,
	}
}