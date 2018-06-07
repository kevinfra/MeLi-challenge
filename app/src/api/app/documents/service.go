package documents

import (
	"api/app/models"
	"database/sql"
	"strconv"

	"fmt"
	"io/ioutil"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"errors"
)

// DocumentService ...
type DocumentService struct {
	DB *sql.DB
	GDrive *drive.Service
	ClientConfig *oauth2.Config
}

// Document ...
func (ds *DocumentService) Document(driveId string) (*models.Document, error) {
	var d models.Document
	row := ds.DB.QueryRow(`SELECT driveId, title, description FROM documents WHERE driveId = ?`, driveId)
	if err := row.Scan(&d.ID, &d.Title, &d.Description); err != nil {
		if err == sql.ErrNoRows {
			if ds.GDrive == nil {
				return nil, errors.New("noGDrive")
			}
			body, driveErr := ds.GDrive.Files.Get(driveId).Do()
			if driveErr != nil {
				return nil, driveErr
			}
			d.DriveId = driveId
			d.Title = body.Name
			d.Description = body.Description
			ds.AddDocument(&d)
		} else {
			return nil, err
		}
	}
	return &d, nil
}

// Documents ...
func (ds *DocumentService) Documents() ([]*models.Document, error) {
	return nil, nil
}

// AddDocument ...
func (ds *DocumentService) AddDocument(d *models.Document) error {
	stmt, err := ds.DB.Prepare(`INSERT INTO documents(driveId,title,description) values(?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(d.DriveId, d.Title, d.Description)
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
	stmt, err := ds.DB.Prepare(`DELETE FROM items WHERE driveId = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(d.DriveId)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DocumentService) InitializeConfig() error {
	var err error
	if ds.ClientConfig != nil {
		return nil
	}
	ds.ClientConfig, err = driveConfig()
	if err != nil {
		panic(err.Error())
		return err
	}
	return nil
}

func (ds *DocumentService) Authorized() bool {
	return ds.GDrive != nil
}

func (ds *DocumentService) LoadURLForTokenAuth() string {
	return ds.ClientConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (ds *DocumentService) LoadFromDB() error {
	var err error
	tok, err := tokenFromDB(ds.DB)
    if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	ds.GDrive, err = getServiceWithToken(tok, ds.ClientConfig)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func (ds *DocumentService) LoadFromToken(token string) error {
	var err error
    saveToken(token, ds.DB)
	ds.GDrive, err = getServiceWithToken(token, ds.ClientConfig)
	return err
}

func getServiceWithToken(token string, config *oauth2.Config) (*drive.Service, error) {
	tok, err := config.Exchange(oauth2.NoContext, token)
    if err != nil {
		fmt.Printf(err.Error())
		return nil, err
    }
	client := config.Client(context.Background(), tok)
	service, err := drive.New(client)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// Retrieves a token from DB.
func tokenFromDB(db *sql.DB) (string, error) {
	var id int
	var token string
	row := db.QueryRow(`SELECT token FROM tokens WHERE id = (SELECT LAST_INSERT_ID())`)
	if err := row.Scan(&id, &token); err != nil {
		return "", err
	}
    return token, nil
}

// Saves a token.
func saveToken(token string, db *sql.DB) {
    stmt, err := db.Prepare(`INSERT INTO tokens(token) values(?)`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(token)
	if err != nil {
		panic(err)
	}
}

func driveConfig() (*oauth2.Config, error) {
	b, err := ioutil.ReadFile("/go/src/api/client_secret.json")
    if err != nil {
        panic("Unable to read client secret file: " + err.Error())
    }

    // If modifying these scopes, delete your previously saved client_secret.json.
    driveConfig, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)
    if err != nil {
        panic("Unable to parse client secret file to config: " + err.Error())
    }
    return driveConfig, nil
}
