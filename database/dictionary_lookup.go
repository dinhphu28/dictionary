package database

import (
	"database/sql"
	"log"
	"time"
)

func LookupInDB(db *sql.DB, word string) ([]Entry, error) {
	// NOTE: for performance measurement
	t := time.Now()
	defer func() {
		log.Println("single db lookup:", time.Since(t))
	}()

	rows, err := db.Query(`
		SELECT headword, html
		FROM entries
		WHERE lower(headword) = lower(?)
	`, word)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Entry

	for rows.Next() {
		var e Entry
		if err := rows.Scan(&e.Headword, &e.HTML); err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	return result, nil
}

func FindByHeadwordStartsWith(db *sql.DB, prefix string, limit int) ([]string, error) {
	rows, err := db.Query(`
		SELECT headword
		FROM entries
		WHERE lower(headword) LIKE lower(?) || '%'
	`, prefix)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []string

	for rows.Next() {
		var e string
		if err := rows.Scan(&e); err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	return result, nil
}
