package logic

import (
	"fmt"
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

	token := generateToken()
	err := db.UpsertRow("Sessions", []string{"email", "token"}, []interface{}{email, token})
	if err != nil {
		return Session{}, err
	}

	return Session{Token: token}, nil
}

func LogIn(db Database, account Account) error {
	email := strings.ToLower(account.Email)

	if !IsValidEmail(email) {
		return fmt.Errorf("Invalid email")
	}

	// Log in

	return nil
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsValidRole(role string) bool {
	return role == "tourist" || role == "local"
}

func generateToken() string {
	return uuid.NewString()
}
