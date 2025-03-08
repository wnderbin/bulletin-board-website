package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBFields struct {
	ID          int
	Name        string
	Description string
	Price       string
	Contacts    string
}

func CreateTable(tablename string, fields []string, db_path string) error {
	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		return err
	}

	defer db.Close()

	command := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s(id INTEGER PRIMARY KEY AUTOINCREMENT, %s TEXT, %s TEXT, %s TEXT, %s TEXT);", tablename, fields[0], fields[1], fields[2], fields[3])
	result, err := db.Exec(command)

	if err != nil {
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	return nil
}

func GetDataFromDB(tablename string, db_path string) ([]DBFields, error) {
	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	command := fmt.Sprintf("select * from %s", tablename)
	rows, err := db.Query(command)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	DBData := []DBFields{}

	for rows.Next() {
		d := DBFields{}
		err = rows.Scan(&d.ID, &d.Name, &d.Description, &d.Price, &d.Contacts)
		if err != nil {
			return nil, err
		}
		DBData = append(DBData, d)
	}

	return DBData, nil
}

func AddToDB(tablename string, fields []string, add []string, db_path string) error {
	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		return err
	}

	defer db.Close()

	command := fmt.Sprintf("insert into %s (%s, %s, %s, %s) values ('%s', '%s', '%s', '%s')", tablename, fields[0], fields[1], fields[2], fields[3], add[0], add[1], add[2], add[3])
	result, err := db.Exec(command)
	if err != nil {
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	return nil
}

func UpdateDB(tablename string, new_info string, field string, id int, db_path string) error {
	db, err := sql.Open("sqlite3", db_path)
	if err != nil {
		return err
	}

	defer db.Close()

	command := fmt.Sprintf("update %s set %s = '%s' where id = %d", tablename, field, new_info, id)

	result, err := db.Exec(command)
	if err != nil {
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	return nil
}

func DeleteFromDB(tablename string, id int, db_path string) error {
	db, err := sql.Open("sqlite3", db_path)

	if err != nil {
		return err
	}

	defer db.Close()

	command := fmt.Sprintf("delete from %s where id = %d", tablename, id)
	result, err := db.Exec(command)

	if err != nil {
		return err
	}

	log.Println(result.LastInsertId())
	log.Println(result.RowsAffected())

	res, _ := result.RowsAffected()
	if res == 0 {
		log.Println("Most likely, a non-existent entry in the database has now been deleted")
	}

	return nil
}
