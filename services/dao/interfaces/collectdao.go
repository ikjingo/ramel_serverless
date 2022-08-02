package interfaces

import "ramel-collectpage/models"

type UserDao interface {
	Create(u *models.Collect) error
	/*Update(u *models.User) error
	Delete(i int) error
	GetById(i int) (models.User, error)*/
	GetAll() ([]models.Collect, error)
}
