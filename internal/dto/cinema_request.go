package dto

type CreateCinemaRequest struct {
	Name      string   `json:"name" binding:"required,min=1,max=100"`
	Address   string   `json:"address"`
	Location  string   `json:"location"`
	Status    string   `json:"status"`
	ManagerID string   `json:"manager_id"`
	ManagerIDs []string `json:"manager_ids"`
}

func (r *CreateCinemaRequest) ToModel() *model.Cinema {
	return &model.Cinema{
		Name:       r.Name,
		Address:    r.Address,
		Location:   r.Location,
		Status:     r.Status,
		ManagerID:  r.ManagerID,
		ManagerIDs: r.ManagerIDs,
	}
}
