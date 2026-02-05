package model

type Cinema struct {
	ID        string    `json:"id"`
	Name    string    `json:"name"`
	Address     string    `json:"address"`
	Location   string    `json:"location"`
	Status	string	`json:"status"`
	ManagerID    string   `json:"manager_id"`
	StaffIDs    []string   `json:"staff_ids"`
}