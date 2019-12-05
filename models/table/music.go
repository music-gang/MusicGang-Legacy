package table

import (
	"time"

	"github.com/IacopoMelani/Go-Starter-Project/pkg/db"

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
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	Title     string    `json:"title" db:"title"`
	Artist    string    `json:"artist" db:"artist"`
	Genre     *string   `json:"genre" db:"genre"`
	Notes     *string   `json:"notes" db:"notes"`
	Prompter  *string   `json:"prompter" db:"prompter"`
}

// LoadAllMusic - Restituisce tutte le istanze di Music
func LoadAllMusic() ([]*Music, error) {

	m := NewMusic()

	db := db.GetConnection()

	query := "SELECT " + record.AllField(m) + " FROM " + m.GetTableName()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*Music

	for rows.Next() {

		m = NewMusic()

		if err := record.LoadFromRow(rows, m); err != nil {
			return nil, err
		}

		result = append(result, m)
	}

	return result, nil
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
func (m Music) GetTableName() string {
	return MusicsTableName
}
