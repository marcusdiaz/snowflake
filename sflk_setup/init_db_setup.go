// You can edit this code!
// Click here and start typing.
package main

import "fmt"

import (
	"database/sql"

	_ "github.com/snowflakedb/gosnowflake"
)


func main() {
	fmt.Println("Hello, 世界")

db, err := sql.Open("snowflake", "marcusdiaz:%t0qFnl99MDWVNem@XLZYSKJ-TLA68665/voterdata")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}




