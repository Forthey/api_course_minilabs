package models

import (
	"gorm.io/gorm"
	"time"
)

type EndorsedCoinInfo struct {
	gorm.Model
	ID          string
	FeatureDate time.Time
}
