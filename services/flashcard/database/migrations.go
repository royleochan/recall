package database

import "github.com/royleochan/recall/services/flashcard/models"

func Migrate() {
	DBCon.AutoMigrate(models.Board{}, models.Flashcard{})
}
