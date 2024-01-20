package urlshort

import (
	"github.com/boltdb/bolt"
)

// SeedBoltDB seeds the BoltDB with some data.
func SeedBoltDB() error {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(dbBucket))
		if err != nil {
			return err
		}

		b.Put([]byte("/urlshort"), []byte("https://github.com/gophercises/urlshort"))
		b.Put([]byte("/urlshort-final"), []byte("https://github.com/gophercises/urlshort/tree/solution"))

		return nil
	})

	return err
}
