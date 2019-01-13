package db

import (
	"fmt"
	//"log"
	//"database/sql"
	"Product-Management-API/pkg/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	*gorm.DB
}

/* ConnectDB
Connect Initializes a New Database Connection returns a Gorm Database pointer
*/
func ConnectDB(directory string) (*Database, error) {
	fmt.Println("() Connecting To Database", directory)
	db, err := gorm.Open("sqlite3", directory+"PMA.db")
	if err != nil {
		fmt.Println("(-) Unable To Connect to Database:", err)

	}
	fmt.Println("(*) Connected to Database")
	SyncDatabases(db)

	return &Database{db}, nil

}

// SyncDatabases make sure tables are upto date
func SyncDatabases(db *gorm.DB) {
	db.AutoMigrate(&models.Project{})
	fmt.Println("Synced Databases")
}

// PopulateInitialData populates some initial databases
