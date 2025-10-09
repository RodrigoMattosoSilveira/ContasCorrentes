package persons

import "gorm.io/gorm"

type PersonsService struct {
	db *gorm.DB
}

func NewPersonsService(db *gorm.DB) *PersonsService {
	return &PersonsService{db: db}
}

func (s *PersonsService) GetAllPersons() ([]Person, error) {
	var Persons []Person
	result := s.db.Order("created_at desc").Find(&Persons)
	return Persons, result.Error
}

func (s *PersonsService) CreatePerson(Person *Person) error {
	result := s.db.Create(Person)
	return result.Error
}

func (s *PersonsService) DeletePerson(id uint) error {
	result := s.db.Delete(&Person{}, id)
	return result.Error
}