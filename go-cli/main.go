// ./main.go
package main

import (
	"database/sql"
	"fmt"
	"go-cli/database"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", "localhost", "sa", "Password1234", "1467", "GolangSqlDB")
	condb, connectionError := sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}
	data := database.Database{
		SqlDb: condb,
	}

	// var (
	// 	sqlversion string
	// )
	// rows, err := condb.Query("select @@version")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for rows.Next() {
	// 	err := rows.Scan(&sqlversion)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(sqlversion)
	// }
	// defer condb.Close()
	fmt.Println("Inserting record")
	data.CreateMovie("Die Hard", "Christmas Movie", 1988)
}
