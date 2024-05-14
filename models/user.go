package models

import (
	"errors"

	"example.com/db"
	"example.com/utils"
)

type User struct {
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
		INSERT INTO users (email, password) VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword) 

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, email, password FROM users WHERE email =?"
	row := db.DB.QueryRow(query, u.Email)


	var retrievedEmail, retrievedPassword string
	err := row.Scan(&u.ID, &retrievedEmail, &retrievedPassword)  //將資料庫查詢的結果掃描並將值分配給指定的變數

	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	return nil
}