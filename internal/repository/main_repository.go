package repository

type MainRepository interface {
	FindAll() ([]interface{}, error)
	FindByID(id int) (interface{}, error)
	Update(id int, data interface{}) error
	Delete(id int) error
	Create(data interface{}) (interface{}, error)
}
