package repository

import "level_6/database/models"

type StudentsRepository interface {
	Create(student *models.Students) error
	GetByID(id uint) (*models.Students, error)
	GetAll() ([]models.Students, error)
	Update(student *models.Students) error
	Delete(id uint) error
	GetByUsername(username string) (*models.Students, error)
}
