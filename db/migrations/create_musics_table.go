package migrations

// CreateMusicsTable - Definisce la struct per la tabella musics
type CreateMusicsTable struct{}

// GetMigrationName - Restituisce il nome della migrazione
func (c CreateMusicsTable) GetMigrationName() string {
	return "create_musics_table"
}

// Down - Definisce la regola di rollback
func (c CreateMusicsTable) Down() string {
	return "DROP TABLE IF EXISTS "
}

// Up - Definisce la regola di migrate
func (c CreateMusicsTable) Up() string {
	return `
		CREATE TABLE musics(
			record_id INT AUTO_INCREMENT,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, 
			title TEXT NOT NULL,
			artist TEXT NOT NULL,
			genre TEXT NULL,
			notes TEXT NULL,
			prompter TEXT NULL,
			PRIMARY KEY(record_id)
		)
	`
}
