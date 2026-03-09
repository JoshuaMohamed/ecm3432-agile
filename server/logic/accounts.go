package logic

import (
	"fmt"
	"regexp"
	"strings"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func SignUp(db Database, account Account) error {
	email := strings.ToLower(account.Email)
	role := strings.ToLower(account.Role)

	if !IsValidEmail(email) {
		return fmt.Errorf("Invalid email")
	}

	if !IsValidRole(role) {
		return fmt.Errorf("Invalid role")
	}

	err := db.CreateRow("Accounts", []string{"email", "password", "role"}, []interface{}{email, account.Password, role})
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return fmt.Errorf("An account with this email already exists")
		}
		return err
	}

	return nil
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func IsValidRole(role string) bool {
	return role == "tourist" || role == "local"
}
