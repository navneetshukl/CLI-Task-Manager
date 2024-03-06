package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
	"github.com/navneetshukl/task/models"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

func DBInit(dbpath string) error {
	var err error

	db, err = bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}

	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

// CreateTask function store the task in boltdb
func CreateTask(task string) (int, error) {
	var id int

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		id64, _ := b.NextSequence()
		id = int(id64)
		key := itob(id)
		return b.Put(key, []byte(task))
	})

	if err != nil {
		return -1, err
	}
	return id, nil
}

// AllTasks function return all  the data from BOLTDB
func AllTasks() ([]models.Task, error) {
	var tasks []models.Task

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			tasks = append(tasks, models.Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil

}

// DeleteTask function will delete data from BOLTDB
func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(taskBucket)
		return b.Delete(itob(key))
	})
}

// itob function convert integer to byte
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

// btoi converts bytes to integer
func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
