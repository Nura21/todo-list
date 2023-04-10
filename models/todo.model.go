package models

import (
	"net/http"
	"todo-list-rest-go/db"
)

type Todo struct {
	Id              int    `json:"id"`
	ActivityGroupId string `json:"activity_group_id"`
	Title           string `json:"title"`
	Priority        string `json:"priority"`
	CreatedAt       string `json:"created_at"`
}

func FetchAllTodo() (Response, error) {
	var obj Todo
	var arrobj []Todo
	var res Response

	rows, err := db.CreateCon().Query("SELECT * FROM todos")

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.ActivityGroupId, &obj.Title, &obj.Priority, &obj.CreatedAt)

		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func StoreTodo(activity_group_id int, title string, priority string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO todos (activity_group_id, title, priority) VALUES(?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(activity_group_id, title, priority)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateTodo(activity_group_id int, id int, title string, priority string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE todos SET activity_group_id = ?, title = ?, priority = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(activity_group_id, title, priority, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteTodo(id int) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM todos WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
