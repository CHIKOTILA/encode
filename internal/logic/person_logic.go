package logic

import "person-crud/internal/models"

type Repository interface {
	Create(p *models.Person) error
	GetAll() ([]*models.Person, error)
	GetByID(id int) (*models.Person, error)
	Update(id int, p *models.Person) error
	Delete(id int) error
}

type personLogic struct {
	repo Repository
}

func NewPersonLogic(r Repository) *personLogic {
	return &personLogic{repo: r}
}

func (l *personLogic) Create(p *models.Person) error {
	return l.repo.Create(p)
}

func (l *personLogic) GetAll() ([]*models.Person, error) {
	return l.repo.GetAll()
}

func (l *personLogic) GetByID(id int) (*models.Person, error) {
	return l.repo.GetByID(id)
}

func (l *personLogic) Update(id int, p *models.Person) error {
	return l.repo.Update(id, p)
}

func (l *personLogic) Delete(id int) error {
	return l.repo.Delete(id)
}
