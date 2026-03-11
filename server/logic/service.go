package logic

// Service defines the business operations available to the presentation layer.
type Service interface {
	CreatePlace(place Place) error
	GetPlaces(postcode, filter string, limit, offset int) ([]Place, error)
	SignUp(account Account) (Session, error)
	LogIn(account Account) error
}

// ServiceImpl implements Service using a Database.
type ServiceImpl struct {
	DB Database
}

func (s *ServiceImpl) CreatePlace(place Place) error {
	return CreatePlace(s.DB, place)
}

func (s *ServiceImpl) GetPlaces(postcode, filter string, limit, offset int) ([]Place, error) {
	return GetPlaces(s.DB, postcode, filter, limit, offset)
}

func (s *ServiceImpl) SignUp(account Account) (Session, error) {
	return SignUp(s.DB, account)
}

func (s *ServiceImpl) LogIn(account Account) error {
	return LogIn(s.DB, account)
}
