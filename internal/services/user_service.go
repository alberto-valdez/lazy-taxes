package services

type UserService interface {
	GetUser() (string, error)
}

type userService struct{}

func NewItemService() UserService {
	return &userService{}
}

func (s *userService) GetUser() (string, error) {
	return "init", nil
}
