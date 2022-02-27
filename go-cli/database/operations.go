// ./database/operations.go
package database

import (
	"fmt"
)

func (db Database) CreateMovie(title, description string, year int) error {
	var err error

	err = db.SqlDb.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Ping succeeded")

	queryStatement := `
	INSERT INTO Movies(Title, Description, Year) VALUES (?, ?, ?);
	SELECT ISNULL(SCOPE_IDENTITY(), -1);
	`

	query, err := db.SqlDb.Prepare(queryStatement)
	if err != nil {
		return err
	}
	fmt.Println("Prepare succeeded")

	defer query.Close()

	_, newRecord := query.Exec(title, description, year)
	fmt.Println("Record inserted: ", newRecord)

	// newRecord := query.QueryRow(
	// 	sql.Named("Title", title),
	// 	sql.Named("Description", description),
	// 	sql.Named("Year", year),
	// )
	// fmt.Println("QueryRow succeeded")

	// var newID int64
	// err = newRecord.Scan(&newID)
	// if err != nil {
	// 	return -1, err
	// }

	return nil
}
