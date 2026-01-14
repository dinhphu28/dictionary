package database

type Manifest struct {
	ID        string `json:"id"`
	ShortName string `json:"short_name"`
	FullName  string `json:"full_name"`
	Database  string `json:"database"`
	Version   string `json:"version"`
}
