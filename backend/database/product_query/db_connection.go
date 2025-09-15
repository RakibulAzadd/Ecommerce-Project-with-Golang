package product_query

import (
	"database/sql"
	"ecommerce/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func GetConnection(cnf config.Config) {

	var err error

	// PostgreSQL connection string
	pslInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cnf.DbHost, cnf.DbPort, cnf.DbUser, cnf.DbPassword, cnf.DbName,
	)

	db, err = sql.Open("postgres", pslInfo)

	if err != nil {
		log.Fatal("❌ Error while opening DB:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("❌ Cannot connect to PostgreSQL:", err)
	}

	fmt.Println("✅ PostgreSQL Connected Successfully")
}

func CloseConnection() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Println("❌ Error closing DB:", err)
			return
		}
		fmt.Println("✅ PostgreSQL Connection Closed Successfully")
	}
}
