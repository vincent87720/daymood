package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/settings"
	"golang.org/x/crypto/bcrypt"
)

func (user *User) Login(db *sql.DB, s settings.Settings) (valid bool, err error) {
	userXi, modelErr := user.Model.ReadByUsername(db)
	if modelErr != nil {
		return false, modelErr
	}

	if len(userXi) < 1 {
		return false, &UsecaseError{Usecase: "authentication", Code: 1, Message: "UserNotFount"}
	}

	checkResult := checkPasswordHash(user.Model.Password, userXi[0].Password)
	if checkResult == true {
		user.Model = &userXi[0]
		return true, nil
	}
	return false, &UsecaseError{Usecase: "authentication", Code: 2, Message: "Login fail"}
}

func (user *User) Logout(db *sql.DB) (err error) {
	return nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
