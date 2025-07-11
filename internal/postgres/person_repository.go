package postgres

import (
	"database/sql"
	"errors"
	"person-crud/internal/models"
)

type personRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *personRepository {
	return &personRepository{db: db}
}

func (r *personRepository) Create(p *models.Person) error {
	return r.db.QueryRow(
		`INSERT INTO persons (email, phone, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`,
		p.Email, p.Phone, p.FirstName, p.LastName,
	).Scan(&p.ID)
}

func (r *personRepository) GetAll() ([]*models.Person, error) {
	rows, err := r.db.Query(`SELECT id, email, phone, first_name, last_name FROM persons ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []*models.Person
	for rows.Next() {
		p := new(models.Person)
		if err := rows.Scan(&p.ID, &p.Email, &p.Phone, &p.FirstName, &p.LastName); err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func (r *personRepository) GetByID(id int) (*models.Person, error) {
	p := new(models.Person)
	err := r.db.QueryRow(
		`SELECT id, email, phone, first_name, last_name FROM persons WHERE id = $1`, id,
	).Scan(&p.ID, &p.Email, &p.Phone, &p.FirstName, &p.LastName)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return p, err
}

func (r *personRepository) Update(id int, p *models.Person) error {
	res, err := r.db.Exec(
		`UPDATE persons SET email=$1, phone=$2, first_name=$3, last_name=$4 WHERE id=$5`,
		p.Email, p.Phone, p.FirstName, p.LastName, id,
	)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *personRepository) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func InitSchema(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		phone TEXT NOT NULL,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL
	)`)
	return err
}
