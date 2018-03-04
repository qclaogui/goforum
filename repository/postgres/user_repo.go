package postgres

import (
	"database/sql"

	"github.com/qclaogui/goforum/model"

	//use PostgreSQL
	_ "github.com/lib/pq"
)

// UserRepository represents a PostgreSQL UserRepository
type UserRepository struct {
	DB *sql.DB
}

// FindByID returns a user for a given id.
func (us *UserRepository) FindByID(id string) (*model.User, error) {

	var u model.User

	row := us.DB.QueryRow(`SELECT id, name FROM users WHERE id = $1`, id)

	if err := row.Scan(&u.ID, &u.Name); err != nil {
		return nil, err
	}

	return &u, nil
}
