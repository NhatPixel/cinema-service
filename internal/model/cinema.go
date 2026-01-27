package model

type Cinema struct {
	ID        string    `json:"id"`
	Name    string    `json:"name"`
	Address     string    `json:"address"`
	Location   string    `json:"location"`
	Status	string	`json:"status"`
	ManagerID    string   `json:"manager_id"`
	ManagerIDs    []string   `json:"manager_ids"`
}