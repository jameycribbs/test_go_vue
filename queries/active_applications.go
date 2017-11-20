package queries

import (
	"github.com/jameycribbs/phoenix_models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type appsData struct {
	ControlData
	Data []phoenix_models.ClientApplication `json:"data"`
}

func ActiveApps(db *gorm.DB, userID int) *appsData {
	var activeAppStateIDs = []int{1, 2, 3, 5}
	var apps appsData

	db.Where("application_state_id IN (?) AND assigned_worker_id = ?", activeAppStateIDs, userID).
		Preload("Program").
		Preload("Client").
		Preload("Client.Area").
		Preload("Client.MedicaidEligibility").
		Preload("Client.MedicaidEligibility.EligibilityCategory").
		Preload("Client.MedicaidRsps").
		Find(&apps.Data)

	return &apps
}
