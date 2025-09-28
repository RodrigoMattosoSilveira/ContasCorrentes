package associates

import "gorm.io/gorm"

type AssociatesService struct {
	db *gorm.DB
}

func NewAssociatesService(db *gorm.DB) *AssociatesService {
	return &AssociatesService{db: db}
}

func (s *AssociatesService) GetAllAssociates() ([]Associate, error) {
	var associates []Associate
	result := s.db.Order("created_at desc").Find(&associates)
	return associates, result.Error
}

func (s *AssociatesService) CreateAssociate(associate *Associate) error {
	result := s.db.Create(associate)
	return result.Error
}

func (s *AssociatesService) DeleteAssociate(id uint) error {
	result := s.db.Delete(&Associate{}, id)
	return result.Error
}
