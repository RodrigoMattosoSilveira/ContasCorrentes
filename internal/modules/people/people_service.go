package people

import (
	"log"

	"gorm.io/gorm"
)

type PeopleService struct {
	db *gorm.DB
}

func NewPeopleService(db *gorm.DB) *PeopleService {
	return &PeopleService{db: db}
}

func (s *PeopleService) GetAllPeople() ([]Person, error) {
	var People []Person
	result := s.db.Order("created_at desc").Find(&People)
	return People, result.Error
}

func (s *PeopleService) CreatePerson(Person *Person) error {
	log.Println("AddPerson - PeopleService: Adding Person")
	result := s.db.Create(Person)
	return result.Error
}

func (s *PeopleService) DeletePerson(id uint) error {
	result := s.db.Delete(&Person{}, id)
	return result.Error
}