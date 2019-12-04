package table

import (
	"time"

	record "github.com/IacopoMelani/Go-Starter-Project/pkg/models/table_record"
)

const (
	MusicsColRecoordID = "record_id"
	MusicsColCreatedAt = "created_at"
	MusicsColUpdatedAt = "updated_at"
	MusicsColTitle     = "title"
	MusicsColArtist    = "artist"
	MusicsColGenre     = "genre"
	MusicsColNotes     = "Notes"
	MusicsColPrompter  = "prompter"

	MusicsTableName = "musics"
)

// Music - Definisce la struct relativa alla tabella musics
type Music struct {
	tr        *record.TableRecord
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Title     string    `db:"title"`
	Artist    string    `db:"artist"`
	Genre     *string   `db:"genre"`
	Notes     *string   `db:"notes"`
	Prompter  *string   `db:"prompter"`
}

// NewMusic - Costruttore della struct Music
func NewMusic() *Music {

	m := new(Music)
	m.tr = record.NewTableRecord(true, false)

	return m
}

// GetTableRecord - Restituisce l'istanza di TableRecord
func (m Music) GetTableRecord() *record.TableRecord {
	return m.tr
}

// GetPrimaryKeyName - Restituisce il nome della chiave primaria
func (m Music) GetPrimaryKeyName() string {
	return MusicsColRecoordID
}

// GetTableName - Restituisce il nome della tabella
func GetTableName() string {
	return MusicsTableName
}
