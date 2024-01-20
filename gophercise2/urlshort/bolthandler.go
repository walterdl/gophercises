package urlshort

import (
	"net/http"
	"strings"

	"github.com/boltdb/bolt"
)

const (
	dbBucket = "URLS"
)

func BoltHandler(fallback http.Handler) (http.HandlerFunc, error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}

	handler := func(w http.ResponseWriter, r *http.Request) {
		err := db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(dbBucket))
			path := strings.ToLower((*r).URL.Path)
			url := b.Get([]byte(path))

			if url == nil {
				fallback.ServeHTTP(w, r)
				return nil
			}

			redirect(string(url), w)
			return nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	return handler, nil
}

func connectDB() (*bolt.DB, error) {
	db, err := bolt.Open(
		"urls.db",
		// Read-write for owner but not execute. Deny read/write/execute for group and others.
		0600,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(dbBucket))
		return err
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
