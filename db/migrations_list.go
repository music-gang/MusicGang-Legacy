package db

import (
	"github.com/IacopoMelani/Go-Starter-Project/pkg/manager/migration"
	"github.com/IacopoMelani/musicgang/db/migrations"
)

// Definisce la lista in sequenza di tutte le migrazioni
var migrationsList = []migration.Migrable{
	migrations.CreateMusicsTable{},
}

// InitMigrationsList - Si occupa di inizializzare la lista delle migrazioni
func InitMigrationsList() {
	migration.InitMigrationsList(migrationsList)
}
