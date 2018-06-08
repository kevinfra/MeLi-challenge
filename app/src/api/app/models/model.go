package models

// Item ...
type Item struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemServiceInterface ...
type ItemServiceInterface interface {
	Item(id string) (*Item, error)
	Items() ([]*Item, error)
	CreateItem(i *Item) error
	DeleteItem(id string) error
}

// Doc
type Document struct {
	ID 			string `json:"id"`
	DriveId		string `json:"driveId"`
	Title 		string `json:"title"`
	Description string `json:"description"`
}

// DocumentServiceInterface
type DocumentServiceInterface interface {
	Document(id string) (*Document, error)
	Documents() ([]*Document, error)
	AddDocument(d *Document) error
	SearchInDoc(id string, word string) (bool, error)
	InitializeConfig() error
	Authorized() bool
	LoadFromDB() error
	LoadURLForTokenAuth() string
	LoadFromToken(token string) error
}