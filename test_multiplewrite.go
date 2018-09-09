package main

import (
	"fmt"
	"encoding/json"
	"github.com/btcsuite/goleveldb/leveldb"
)

type AlertData map[string]interface{}

func main() {
	db, err := leveldb.OpenFile("test", nil)
	if err != nil {
		return
	}

	testversion := 5
	jsondata := AlertData{
		"version": testversion,
		"relayunti": 2,
		"expiration": 3,
		"cancel": 4,
		"setcancel": 5,
		"minver": 6,
		"maxver": 7,
		"setsubver": "8",
		"priority": 9,
		"comment": "10",
		"statusbar": "11",
		"reserved": "12"}

	bJsondata, err := json.MarshalIndent(jsondata, "", "    ")

	err = db.Put([]byte(fmt.Sprintf("%020d", 1)), bJsondata, nil)
	data, err := db.Get([]byte(fmt.Sprintf("%020d", 1)), nil)

	fmt.Printf(string(data))
	defer db.Close()
}
