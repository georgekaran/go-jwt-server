package service

type Service interface {
	FindOne(id int)
	FindAll()
	Save(t interface{})
	SaveAll(t interface{})
	Delete(t interface{})
	DeleteAll(t interface{})
}
