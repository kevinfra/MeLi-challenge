package documents

import (
	"api/app/models"
	"database/sql"
	"strconv"
    "google.golang.org/api/drive/v3"
)

// DocumentService ...
type DocumentService struct {
	DB *sql.DB
	GDrive *drive.Service 
}

// Document ...
func (ds *DocumentService) Document(id string) (*models.Document, error) {
	var i models.Document
	row := ds.DB.QueryRow(`SELECT id, name, description FROM items WHERE id = ?`, id)
	if err := row.Scan(&i.ID, &i.Title, &i.Description); err != nil {
		return nil, err
	}
	return &i, nil
}

// Documents ...
func (ds *DocumentService) Documents() ([]*models.Document, error) {
	return nil, nil
}

// AddDocument ...
func (ds *DocumentService) AddDocument(d *models.Document) error {
	stmt, err := ds.DB.Prepare(`INSERT INTO items(name,description) values(?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(d.Title, d.Description)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	d.ID = strconv.Itoa(int(id))
	return nil
}

// SearchInDoc ...
func (ds *DocumentService) SearchInDoc(d *models.Document, word string) error {
	stmt, err := ds.DB.Prepare(`DELETE FROM items WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(d)
	if err != nil {
		return err
	}
	return nil
}
