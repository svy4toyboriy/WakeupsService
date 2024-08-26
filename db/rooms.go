package db

import (
	"WebService/model"
)

func (db *PgDB) AddRooms(rooms []*model.Room) error {
	_, err := db.Model(&rooms).
		OnConflict("(room_number) DO UPDATE").Insert()
	return err
}
