package main

import (
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb"
)

func main() {

	db, err := leveldb.OpenFile("test", nil)
	if err != nil {
		return
	}
	//err = db.Put([]byte("1"), []byte("value2224"), nil)
	//err = db.Put([]byte("2000"), []byte("value2226"), nil)
	//err = db.Put([]byte("20"), []byte("value2227"), nil)
	//err = db.Put([]byte("10"), []byte("value2228"), nil)
	defer db.Close()
	iter := db.NewIterator(nil, nil)
	iter.Last()
	key := iter.Key()
	value := iter.Value()
	fmt.Printf("key: %s | value: %s\n", key, value)
	for iter.Prev() {
		fmt.Printf("key: %s | value: %s\n", key, value)
	}
	fmt.Println("\n")
	iter.Last()
	key = iter.Key()
	value = iter.Value()
	fmt.Printf("key: %s | value: %s\n", key, value)

	iter.Release()
	err = iter.Error()
	fmt.Println("\n")
}
