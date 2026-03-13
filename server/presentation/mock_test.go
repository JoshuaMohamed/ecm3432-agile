package presentation_test

import "server/logic"

type mockService struct {
	places []logic.Place
	err    error
	email  string
}

func (m *mockService) CreatePlace(place logic.Place) error {
	return m.err
}

func (m *mockService) GetPlaces(postcode, filter string, limit, offset int) ([]logic.Place, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.places, nil
}

func (m *mockService) SignUp(place logic.Account) (logic.Session, error) {
	return logic.Session{}, m.err
}

func (m *mockService) LogIn(place logic.Account) (logic.Session, error) {
	return logic.Session{}, m.err
}

func (m *mockService) LogOut(token string) error {
	return m.err
}

func (m *mockService) ValidateSession(token string) (string, error) {
	if m.err != nil {
		return "", m.err
	}
	return m.email, nil
}
