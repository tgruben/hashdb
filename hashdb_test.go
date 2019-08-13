package hashdb

import (
	"fmt"
	"os"
	"testing"
)

func Test_db_open(t *testing.T) {
	db := &Db{
		fileName: "test.map",
	}
	if err := db.Open(); err != nil {
		t.Errorf("db.open() error = %v", err)
	}

	db.Upsert(&Entry{
		Key:    2,
		Offset: 0,
		Length: 15,
	})
	db.Upsert(&Entry{
		Key:    3,
		Offset: 16,
		Length: 15,
	})
	db.Close()
	if err := db.Open(); err != nil {
		t.Errorf("db.open() error = %v", err)
	}
	fmt.Println("GO")
	db.Dump()
	fmt.Println("DONE")

	for i := 0; i < 100000; i++ {
		db.Upsert(&Entry{
			Key:    uint64(i),
			Offset: uint64(i),
			Length: 15,
		})
	}
	fmt.Println("END POP")
	db.Dump()
	fmt.Println("Done")
	db.Close()
	os.Remove(db.FileName())
}

func Test_db_scale(t *testing.T) {
	db := &Db{
		fileName: "test.map",
	}
	if err := db.Open(); err != nil {
		t.Errorf("db.open() error = %v", err)
	}

	fmt.Println("HELLO", db.totalSlots)
	db.Upsert(&Entry{
		Key:    2,
		Offset: 0,
		Length: 15,
	})
	db.Upsert(&Entry{
		Key:    3,
		Offset: 16,
		Length: 15,
	})
	db.scale(56)
	db.Close()
	if err := db.Open(); err != nil {
		t.Errorf("db.open() error = %v", err)
	}
	os.Remove(db.FileName())
}
