package logic

import (
	"fmt"
	"log/slog"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Session struct {
	Email string `json:"email,omitempty"`
	Token string `json:"token"`
}

func SignUp(db Database, account Account) (Session, error) {
	email := strings.ToLower(account.Email)
	role := strings.ToLower(account.Role)

	if !IsValidEmail(email) {
		return Session{}, fmt.Errorf("Invalid email")
	}

	if !IsValidRole(role) {
		return Session{}, fmt.Errorf("Invalid role")
	}

	err := db.InsertRow("Accounts", []string{"email", "password", "role"}, []interface{}{email, account.Password, role})
	if err != nil {
		if isDuplicateAccountEmailError(err) {
			return Session{}, fmt.Errorf("An account with this email already exists")
		}
		return Session{}, err
	}

	token := generateToken()
	err = db.UpsertRow("Sessions", []string{"email", "token"}, []interface{}{email, token})
	if err != nil {
		return Session{}, err
	}

	return Session{Token: token}, nil
}

func LogIn(db Database, account Account) (Session, error) {
	email := strings.ToLower(account.Email)

	rows, err := db.Query("Accounts", "email", email)
	if err != nil {
		return Session{}, err
	}

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			rows.Close()
			return Session{}, err
		}
		rows.Close()
		slog.Info("Account does not exist", "email", email)
		return Session{}, fmt.Errorf("Incorrect email or password")
	}

	var result Account
	if err := rows.Scan(&result.Email, &result.Password, &result.Role); err != nil {
		rows.Close()
		return Session{}, err
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		return Session{}, err
	}

	if err := rows.Close(); err != nil {
		return Session{}, err
	}

	if result.Password != account.Password {
		return Session{}, fmt.Errorf("Incorrect email or password")
	}

	token := generateToken()
	err = db.UpsertRow("Sessions", []string{"email", "token"}, []interface{}{email, token})
	if err != nil {
		return Session{}, err
	}

	return Session{Token: token}, nil
}

func LogOut(db Database, token string) error {
	if token == "" {
		return fmt.Errorf("Invalid session token")
	}

	return db.DeleteRows("Sessions", "token", token)
}

func ValidateSession(db Database, token string) (string, error) {
	if token == "" {
		return "", fmt.Errorf("Invalid session token")
	}

	rows, err := db.Query("Sessions", "token", token)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return "", err
		}
		return "", fmt.Errorf("Invalid session token")
	}

	var session Session
	if err := rows.Scan(&session.Email, &session.Token); err != nil {
		return "", err
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return session.Email, nil
}

func IsValidSession(db Database, email, token string) bool {
	email = strings.ToLower(email)

	rows, err := db.Query("Sessions", "email", email)
	if err != nil {
		return false
	}
	defer rows.Close()

	if !rows.Next() {
		return false
	}

	var sessionEmail string
	var sessionToken string
	if err := rows.Scan(&sessionEmail, &sessionToken); err != nil {
		return false
	}

	if err := rows.Err(); err != nil {
		return false
	}

	return sessionToken == token
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsValidRole(role string) bool {
	return role == "tourist" || role == "local"
}

func isDuplicateAccountEmailError(err error) bool {
	errMsg := strings.ToLower(err.Error())
	return strings.Contains(errMsg, "unique constraint failed: accounts.email")
}

func generateToken() string {
	return uuid.NewString()
}
