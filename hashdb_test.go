package hashdb

import (
	"fmt"
	"os"
	"testing"
)

func Test_db_open(t *testing.T) {
	db := &db{
		fileName: "test.map",
	}
	if err := db.Open(); err != nil {
		t.Errorf("db.open() error = %v", err)
	}

	db.Insert(&Entry{
		key:    2,
		offset: 0,
		length: 15,
	})
	db.Insert(&Entry{
		key:    3,
		offset: 16,
		length: 15,
	})
	db.Close()
	if err := db.Open(); err != nil {
		t.Errorf("db.open() error = %v", err)
	}
	i, ok := db.Find(3)
	if ok {
		fmt.Println("Found", i)
	} else {
		t.Errorf("db.Find() not Found %d", 3)
	}
	fmt.Println("GO")
	db.Dump()
	fmt.Println("DONE")

	for i := 0; i < 100000; i++ {
		db.Insert(&Entry{
			key:    uint64(i),
			offset: uint64(i),
			length: 15,
		})
	}
	fmt.Println("END POP")
	db.Dump()
	fmt.Println("Done")
	db.Close()
	os.Remove(db.FileName())
}
