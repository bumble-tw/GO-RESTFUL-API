package models

import (
	"time"

	"example.com/db"
)

type Event struct {
	ID int64
	Name string `binding:"required"`
	Desc string	`binding:"required"`
	Location string `binding:"required"`
	Date time.Time `binding:"required"`
	UserId int
}

var events = []Event{}

func(e *Event) Save() error{
	// SQL 查詢語句，用來插入事件的資料到 events 表中。問號（?）是佔位符，稍後會被具體的值替換。
	query := `
		INSERT INTO events (name, description, location, date, user_id) 
		VALUES (?,?,?,?,?)
	`

	// 使用 db 對象的 Prepare 方法準備 SQL 語句，這有助於防止 SQL 注入攻擊，並且可以提高性能。
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// 確保在 Save 方法結束時關閉 stmt 對象，釋放資源。
	defer stmt.Close()

	// 執行 SQL 語句，並將事件的名稱、描述、位置、日期和使用者 ID 傳遞進去，替換掉 SQL 語句中的佔位符。
	result, err := stmt.Exec(e.Name, e.Desc, e.Location, e.Date, e.UserId)
	if err != nil {
		return err
	}

	// 獲取剛剛插入的數據行的 ID。
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents () ([]Event,error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next(){
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Desc, &event.Location, &event.Date, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error){
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query,eventId) // for id case use queryRow

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Desc, &event.Location, &event.Date, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}