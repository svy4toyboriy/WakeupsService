package migrations

import (
	"WebService/db"
	"log"

	"github.com/go-pg/migrations/v7"
)

func Run(db *db.PgDB) error {
	oldVersion, newVersion, err := migrations.Run(db, "init")
	if err != nil {
		return err
	}
	oldVersion, newVersion, err = migrations.Run(db, "up")
	if err != nil {
		return err
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d", oldVersion, newVersion)
	} else {
		log.Printf("version is %d", oldVersion)
	}
	return nil
}
