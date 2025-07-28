package repository

import (
	"level_6/database/models"

	"gorm.io/gorm"
)

// StudentRepository implements StudentsRepository interface
// SOLID: Dependency Inversion - depends on gorm.DB interface, not concrete DB
// SOLID: Single Responsibility - only handles data persistence
type StudentRepository struct {
	db *gorm.DB
}

// NewStudentRepository creates a new repository instance
// SOLID: Dependency Inversion - returns interface, not concrete type
func NewStudentRepository(db *gorm.DB) StudentsRepository {
	return &StudentRepository{db: db}
}

// SOLID: Interface Segregation - each method has single, focused responsibility
func (r *StudentRepository) Create(student *models.Students) error {
	return r.db.Create(student).Error
}

func (r *StudentRepository) GetByID(id uint) (*models.Students, error) {
	var student models.Students
	err := r.db.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) GetAll() ([]models.Students, error) {
	var students []models.Students
	err := r.db.Find(&students).Error
	return students, err
}

func (r *StudentRepository) Update(student *models.Students) error {
	return r.db.Save(student).Error
}

func (r *StudentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Students{}, id).Error
}
