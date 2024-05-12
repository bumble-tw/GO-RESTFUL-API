package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// 開啟或創建一個名為 "api.db" 的 SQLite3 數據庫文件，並初始化連接。
	// "sqlite3" 指定使用 SQLite3 驅動，由 go-sqlite3 包提供並在此應用中已匿名導入。
	// 函數返回的 DB 對象用於後續的數據庫操作。
	// 如果連接或文件創建過程中出現錯誤，err 變量將捕捉這個錯誤。
	err := error(nil)
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)  //最大連接數
	DB.SetMaxIdleConns(5)   //最大空閒連接數（閒置時期，保持開啟5個連接，以利於有req時，能快速連接。）

	createTables()
}


func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	
	}


	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

}