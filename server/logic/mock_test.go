package logic_test

import "server/logic"

type deleteCall struct {
	table string
	key   string
	value string
}

type mockDB struct {
	rows         logic.DBRows
	createErr    error
	insertErr    error
	upsertErr    error
	deleteErr    error
	getPlacesErr error
	queryErr     error
	deleteCalls  []deleteCall
}

func (m *mockDB) CreateTable(details logic.TableDetails) error {
	return m.createErr
}

func (m *mockDB) InsertRow(table string, fields []string, values []interface{}) error {
	return m.insertErr
}

func (m *mockDB) UpsertRow(table string, fields []string, values []interface{}) error {
	return m.upsertErr
}

func (m *mockDB) DeleteRows(table, key, value string) error {
	m.deleteCalls = append(m.deleteCalls, deleteCall{table: table, key: key, value: value})
	return m.deleteErr
}

func (m *mockDB) GetPlaces(searchPrefix string, limit, offset int) (logic.DBRows, error) {
	if m.getPlacesErr != nil {
		return nil, m.getPlacesErr
	}
	return m.rows, nil
}

func (m *mockDB) Query(table, key, value string) (logic.DBRows, error) {
	if m.queryErr != nil {
		return nil, m.queryErr
	}
	return m.rows, nil
}

func (m *mockDB) Close() error {
	return nil
}
