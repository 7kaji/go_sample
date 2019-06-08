package entity

// User is user models property
type User struct {
	ID        uint   `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"firstname"`
	LastName  string `db:"last_name" json:"lastname"`
}
