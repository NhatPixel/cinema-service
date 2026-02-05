package dto

import (
	"github.com/NhatPixel/cinema-service/internal/model"
)

type GetResponse struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Address   string   `json:"address"`
	Location  string   `json:"location"`
	Status    string   `json:"status"`
	ManagerID string   `json:"manager_id"`
	StaffIDs  []string `json:"staff_ids"`
}

func (c *GetResponse) FromModel(m *model.Cinema) {
	c.ID = m.ID
	c.Name = m.Name
	c.Address = m.Address
	c.Location = m.Location
	c.Status = m.Status
	c.ManagerID = m.ManagerID
	c.StaffIDs = m.StaffIDs
}
