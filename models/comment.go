package models

type Comment struct {
	Id             string
	Composition_id string
	User_id        string
	Content        string
	CreatedAt      string
	UpdatedAt      string
	DeletedAt      string
	Limit, Offset  int
}
