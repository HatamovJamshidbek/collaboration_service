package models

type Collaboration struct {
	Id             string
	Composition_id string
	User_Id        string
	Role           string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	Limit, Offset  int
}
