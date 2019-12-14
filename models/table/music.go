package table

import (
	"time"

	"github.com/IacopoMelani/musicgang/models/dto"

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
	ID        int64     `json:"id"`
	Title     string    `json:"title" db:"title"`
	Artist    string    `json:"artist" db:"artist"`
	Genre     *string   `json:"genre" db:"genre"`
	Notes     *string   `json:"notes" db:"notes"`
	Prompter  *string   `json:"prompter" db:"prompter"`
}

// InsertNewMusic - Si occupa di inserire una nuova canzone
func InsertNewMusic(m dto.MusicDTO) error {

	newMusic := NewMusic()

	newMusic.Title = m.Title
	newMusic.Artist = m.Artist
	newMusic.SetGenre(m.Genre)
	newMusic.SetNotes(m.Notes)
	newMusic.SetPrompter(m.Prompter)

	newMusic.CreatedAt = time.Now().UTC()
	newMusic.UpdatedAt = time.Now().UTC()

	if err := record.Save(newMusic); err != nil {
		return err
	}

	return nil
}

// LoadAllMusic - Restituisce tutte le istanze di Music
func LoadAllMusic() ([]*Music, error) {

	m := NewMusic()

	db := db.GetConnection()

	query := "SELECT " + record.AllField(m) + " FROM " + m.GetTableName() + " ORDER BY " + MusicsColUpdatedAt + " DESC "

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := []*Music{}

	for rows.Next() {

		m = NewMusic()

		if err := record.LoadFromRow(rows, m); err != nil {
			return nil, err
		}

		m.ID = m.tr.GetID()

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

// SetGenre - Imposta il genere della canzone
func (m *Music) SetGenre(genre string) {
	m.Genre = &genre
}

// SetNotes - Imposta le note per la canzone
func (m *Music) SetNotes(notes string) {
	m.Notes = &notes
}

// SetPrompter - Imposta il genere della canzone
func (m *Music) SetPrompter(prompter string) {
	m.Prompter = &prompter
}
