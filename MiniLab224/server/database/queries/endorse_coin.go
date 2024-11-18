package queries

import (
	"server/database"
	"server/database/models"
	"time"
)

func Endorse(id string) error {
	result := database.Database.Create(&models.EndorsedCoinInfo{
		ID:          id,
		FeatureDate: time.Now(),
	})
	return result.Error
}

func RemoveEndorsement(id string) error {
	result := database.Database.Delete(&models.EndorsedCoinInfo{}, id)
	return result.Error
}
