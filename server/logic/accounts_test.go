package logic_test

import (
	"errors"
	"strings"
	"testing"

	"server/logic"
)

type mockAccountRows struct {
	rows    []logic.Account
	idx     int
	err     error
	scanErr error
}

func (m *mockAccountRows) Next() bool {
	if m.err != nil {
		return false
	}
	if m.idx >= len(m.rows) {
		return false
	}
	m.idx++
	return true
}

func (m *mockAccountRows) Scan(dest ...interface{}) error {
	if m.scanErr != nil {
		return m.scanErr
	}
	if m.idx == 0 || m.idx > len(m.rows) {
		return errors.New("scan out of bounds")
	}
	if len(dest) != 3 {
		return errors.New("unexpected scan arg count")
	}
	emailPtr, ok := dest[0].(*string)
	if !ok {
		return errors.New("invalid email dest")
	}
	passwordPtr, ok := dest[1].(*string)
	if !ok {
		return errors.New("invalid password dest")
	}
	rolePtr, ok := dest[2].(*string)
	if !ok {
		return errors.New("invalid role dest")
	}

	row := m.rows[m.idx-1]
	*emailPtr = row.Email
	*passwordPtr = row.Password
	*rolePtr = row.Role
	return nil
}

func (m *mockAccountRows) Err() error {
	return m.err
}

func (m *mockAccountRows) Close() error {
	return nil
}

type mockSessionRows struct {
	rows    []logic.Session
	idx     int
	err     error
	scanErr error
}

func (m *mockSessionRows) Next() bool {
	if m.err != nil {
		return false
	}
	if m.idx >= len(m.rows) {
		return false
	}
	m.idx++
	return true
}

func (m *mockSessionRows) Scan(dest ...interface{}) error {
	if m.scanErr != nil {
		return m.scanErr
	}
	if m.idx == 0 || m.idx > len(m.rows) {
		return errors.New("scan out of bounds")
	}
	if len(dest) != 2 {
		return errors.New("unexpected scan arg count")
	}
	emailPtr, ok := dest[0].(*string)
	if !ok {
		return errors.New("invalid email dest")
	}
	tokenPtr, ok := dest[1].(*string)
	if !ok {
		return errors.New("invalid token dest")
	}

	row := m.rows[m.idx-1]
	*emailPtr = row.Email
	*tokenPtr = row.Token
	return nil
}

func (m *mockSessionRows) Err() error {
	return m.err
}

func (m *mockSessionRows) Close() error {
	return nil
}

func TestSignUp(t *testing.T) {
	tests := []struct {
		name    string
		account logic.Account
		db      *mockDB
		wantErr string
	}{
		{
			name:    "success",
			account: logic.Account{Email: "User@Example.com", Password: "secret", Role: "Tourist"},
			db:      &mockDB{},
		},
		{
			name:    "invalid email",
			account: logic.Account{Email: "invalid", Password: "secret", Role: "tourist"},
			db:      &mockDB{},
			wantErr: "Invalid email",
		},
		{
			name:    "invalid role",
			account: logic.Account{Email: "user@example.com", Password: "secret", Role: "admin"},
			db:      &mockDB{},
			wantErr: "Invalid role",
		},
		{
			name:    "duplicate email",
			account: logic.Account{Email: "user@example.com", Password: "secret", Role: "tourist"},
			db:      &mockDB{insertErr: errors.New("constraint failed: UNIQUE constraint failed: Accounts.email (1555)")},
			wantErr: "An account with this email already exists",
		},
		{
			name:    "db error",
			account: logic.Account{Email: "user@example.com", Password: "secret", Role: "tourist"},
			db:      &mockDB{insertErr: errors.New("db down")},
			wantErr: "db down",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := logic.SignUp(tt.db, tt.account)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("SignUp() error = %v, want contains %q", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("SignUp() unexpected error: %v", err)
			}
			if session.Token == "" {
				t.Fatal("SignUp() expected non-empty token")
			}
		})
	}
}

func TestLogIn(t *testing.T) {
	tests := []struct {
		name    string
		account logic.Account
		db      *mockDB
		wantErr string
	}{
		{
			name:    "success",
			account: logic.Account{Email: "user@example.com", Password: "secret"},
			db: &mockDB{rows: &mockAccountRows{rows: []logic.Account{{
				Email:    "user@example.com",
				Password: "secret",
				Role:     "tourist",
			}}}},
		},
		{
			name:    "account missing",
			account: logic.Account{Email: "user@example.com", Password: "secret"},
			db:      &mockDB{rows: &mockAccountRows{}},
			wantErr: "Incorrect email or password",
		},
		{
			name:    "password mismatch",
			account: logic.Account{Email: "user@example.com", Password: "wrong"},
			db: &mockDB{rows: &mockAccountRows{rows: []logic.Account{{
				Email:    "user@example.com",
				Password: "secret",
				Role:     "tourist",
			}}}},
			wantErr: "Incorrect email or password",
		},
		{
			name:    "query error",
			account: logic.Account{Email: "user@example.com", Password: "secret"},
			db:      &mockDB{queryErr: errors.New("db down")},
			wantErr: "db down",
		},
		{
			name:    "upsert error",
			account: logic.Account{Email: "user@example.com", Password: "secret"},
			db: &mockDB{
				rows:      &mockAccountRows{rows: []logic.Account{{Email: "user@example.com", Password: "secret", Role: "tourist"}}},
				upsertErr: errors.New("db down"),
			},
			wantErr: "db down",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			session, err := logic.LogIn(tt.db, tt.account)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("LogIn() error = %v, want contains %q", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("LogIn() unexpected error: %v", err)
			}
			if session.Token == "" {
				t.Fatal("LogIn() expected non-empty token")
			}
		})
	}
}

func TestLogOut(t *testing.T) {
	tests := []struct {
		name       string
		token      string
		db         *mockDB
		wantErr    string
		wantDelete bool
	}{
		{
			name:       "success",
			token:      "abc123",
			db:         &mockDB{},
			wantDelete: true,
		},
		{
			name:    "empty token",
			token:   "",
			db:      &mockDB{},
			wantErr: "Invalid session token",
		},
		{
			name:  "delete error",
			token: "abc123",
			db: &mockDB{
				deleteErr: errors.New("db down"),
			},
			wantErr:    "db down",
			wantDelete: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := logic.LogOut(tt.db, tt.token)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("LogOut() error = %v, want contains %q", err, tt.wantErr)
				}
			} else if err != nil {
				t.Fatalf("LogOut() unexpected error: %v", err)
			}

			if tt.wantDelete && len(tt.db.deleteCalls) == 0 {
				t.Fatal("LogOut() expected DeleteRows to be called")
			}
			if tt.wantDelete && tt.token != "" {
				last := tt.db.deleteCalls[len(tt.db.deleteCalls)-1]
				if last.key != "token" || last.value != tt.token {
					t.Fatalf("DeleteRows() called with %s=%s, want token=%s", last.key, last.value, tt.token)
				}
			}
		})
	}
}

func TestIsValidSession(t *testing.T) {
	tests := []struct {
		name  string
		email string
		token string
		db    *mockDB
		want  bool
	}{
		{
			name:  "valid session",
			email: "user@example.com",
			token: "abc123",
			db: &mockDB{rows: &mockSessionRows{rows: []logic.Session{{
				Email: "user@example.com",
				Token: "abc123",
			}}}},
			want: true,
		},
		{
			name:  "missing session",
			email: "user@example.com",
			token: "abc123",
			db:    &mockDB{rows: &mockSessionRows{}},
			want:  false,
		},
		{
			name:  "wrong token",
			email: "user@example.com",
			token: "wrong",
			db: &mockDB{rows: &mockSessionRows{rows: []logic.Session{{
				Email: "user@example.com",
				Token: "abc123",
			}}}},
			want: false,
		},
		{
			name:  "query error",
			email: "user@example.com",
			token: "abc123",
			db:    &mockDB{queryErr: errors.New("db down")},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := logic.IsValidSession(tt.db, tt.email, tt.token)
			if got != tt.want {
				t.Fatalf("IsValidSession() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateSession(t *testing.T) {
	tests := []struct {
		name      string
		token     string
		db        *mockDB
		wantEmail string
		wantErr   string
	}{
		{
			name:      "valid token",
			token:     "abc123",
			db:        &mockDB{rows: &mockSessionRows{rows: []logic.Session{{Email: "user@example.com", Token: "abc123"}}}},
			wantEmail: "user@example.com",
		},
		{
			name:    "empty token",
			token:   "",
			db:      &mockDB{},
			wantErr: "Invalid session token",
		},
		{
			name:    "missing token",
			token:   "abc123",
			db:      &mockDB{rows: &mockSessionRows{}},
			wantErr: "Invalid session token",
		},
		{
			name:    "query error",
			token:   "abc123",
			db:      &mockDB{queryErr: errors.New("db down")},
			wantErr: "db down",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			email, err := logic.ValidateSession(tt.db, tt.token)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("ValidateSession() error = %v, want contains %q", err, tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("ValidateSession() unexpected error: %v", err)
			}
			if email != tt.wantEmail {
				t.Fatalf("ValidateSession() email = %q, want %q", email, tt.wantEmail)
			}
		})
	}
}

func TestEmailAndRoleValidation(t *testing.T) {
	if !logic.IsValidEmail("user@example.com") {
		t.Fatal("expected valid email")
	}
	if logic.IsValidEmail("bad") {
		t.Fatal("expected invalid email")
	}
	if !logic.IsValidRole("tourist") || !logic.IsValidRole("local") {
		t.Fatal("expected valid built-in roles")
	}
	if logic.IsValidRole("admin") {
		t.Fatal("expected invalid role")
	}
}
