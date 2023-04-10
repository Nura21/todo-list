package models

import (
	"net/http"
	"todo-list-rest-go/db"
)

type Activity struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func FetchAllActivity() (Response, error) {
	var obj Activity
	var arrobj []Activity
	var res Response

	rows, err := db.CreateCon().Query("SELECT * FROM activities")

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Title, &obj.Email, &obj.CreatedAt)

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

func StoreActivity(title string, email string) (Response, error) {
	var res Response

	con := db.CreateCon()

	sqlStatement := "INSERT INTO activities (title, email) VALUES(?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(title, email)
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

func UpdateActivity(id int, title string, email string) (Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE activities SET title = ?, email = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err !=nil{
		return res, err
	}

	result, err := stmt.Exec(title, email, id)
	if err != nil{
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteActivity(id int) (Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "DELETE FROM activities WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil{
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil{
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
