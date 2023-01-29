package model

import (
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type User struct {
	ID       int64  //流水號
	Username string //帳號
	Password string //密碼
	Name     string //名稱
	Email    string //信箱
	CreateAt string //建立時間
	UpdateAt string //最後編輯時間
}

func NewUser() *User {
	return &User{}
}

func (user *User) ReadAll(db *sql.DB) (userXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM users ORDER BY id DESC;")
	if err != nil {
		return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var userSchema User
	for row.Next() {
		err := row.Scan(
			&userSchema.ID, &userSchema.Username, &userSchema.Password,
			&userSchema.Name, &userSchema.Email, &userSchema.CreateAt,
			&userSchema.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
		}

		userXi = append(userXi, userSchema)
	}

	return userXi, nil
}

func (user *User) Read(db *sql.DB) (userXi []interface{}, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM users WHERE id = $1 ORDER BY id DESC;", user.ID)
	if err != nil {
		return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var userSchema User
	for row.Next() {
		err := row.Scan(
			&userSchema.ID, &userSchema.Username, &userSchema.Password,
			&userSchema.Name, &userSchema.Email, &userSchema.CreateAt,
			&userSchema.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
		}

		userXi = append(userXi, userSchema)
	}

	return userXi, nil
}

func (user *User) ReadByUsername(db *sql.DB) (userXi []User, modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	row, err := db.Query("SELECT * FROM users WHERE username = $1 ORDER BY id DESC;", user.Username)
	if err != nil {
		return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}
	defer row.Close()

	var userSchema User
	for row.Next() {
		err := row.Scan(
			&userSchema.ID, &userSchema.Username, &userSchema.Password,
			&userSchema.Name, &userSchema.Email, &userSchema.CreateAt,
			&userSchema.UpdateAt,
		)
		if err != nil {
			return nil, &ModelError{Model: "users", Code: 0, Message: err.Error()}
		}

		userXi = append(userXi, userSchema)
	}

	return userXi, nil
}

func (user *User) Create(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	stmt, err := db.Prepare("INSERT INTO users(username, password, name, email) VALUES($1,$2,$3,$4);")
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Password, user.Name, user.Email)
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "users", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}

func (user *User) Update(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	_, err = db.Exec("CALL updateUsers($1,$2,$3,$4)", user.ID, user.Username, user.Name, user.Email)
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	return nil
}

func (user *User) UpdatePassword(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	_, err = db.Exec("CALL updateUserPasswords($1,$2)", user.ID, user.Password)
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	return nil
}

func (user *User) Delete(db *sql.DB) (modelErr *ModelError) {
	err := db.Ping()
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = $1;")
	if err != nil {
		return &ModelError{Model: "users", Code: 0, Message: err.Error()}
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.ID)
	if err, ok := err.(*pq.Error); ok {
		fmt.Println(err.Code.Name())
		return &ModelError{Model: "users", Code: 3, Message: "Database error"}
	}

	rowsAff, err := res.RowsAffected()
	if rowsAff != 1 {
		return &ModelError{Model: "users", Code: 2, Message: "RowsAffected incorrect."}
	}
	return nil
}
