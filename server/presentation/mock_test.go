package presentation_test

import "server/logic"

type mockService struct {
	places []logic.Place
	err    error
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

func (m *mockService) SignUp(place logic.Account) error {
	return m.err
}
