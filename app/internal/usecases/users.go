package usecases

import (
	"database/sql"

	"github.com/vincent87720/daymood/app/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model *model.User
}

func NewUser(userModel *model.User) *User {
	return &User{
		Model: userModel,
	}
}

func (user *User) Read(db *sql.DB) ([]interface{}, *model.ModelError) {

	userXi, modelErr := model.Read(user.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return userXi, nil
}

func (user *User) ReadAll(db *sql.DB) ([]interface{}, *model.ModelError) {

	userXi, modelErr := model.ReadAll(user.Model, db)
	if modelErr != nil {
		return nil, modelErr
	}
	return userXi, nil
}

func (user *User) Create(db *sql.DB) *model.ModelError {

	hashedPassword, err := hashPassword(user.Model.Password)
	if err != nil {
		return &model.ModelError{Model: "userUsecase", Code: 1, Message: "PasswordHashError"}
	}
	user.Model.Password = hashedPassword

	modelErr := model.Create(user.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (user *User) CreateMultiple(db *sql.DB, userXi interface{}) *model.ModelError {
	return nil
}

func (user *User) Update(db *sql.DB) *model.ModelError {

	modelErr := model.Update(user.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}

func (user *User) Delete(db *sql.DB) *model.ModelError {

	modelErr := model.Delete(user.Model, db)
	if modelErr != nil {
		return modelErr
	}
	return nil
}
