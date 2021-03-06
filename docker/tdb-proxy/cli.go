package main

import (
	"encoding/csv"
	"fmt"
	"github.com/traildb/traildb-go"
	"log"
	"os"
)

var SESSION_LIMIT = uint64(30 * 60)

func main() {
	db, err := tdb.Open("forum.tdb")
	if err != nil {
		panic(err)
	}
	trail, err := tdb.NewCursor(db)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("result.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fmt.Println("Number of trails: ", db.NumTrails)

	count := 0
	for i := uint64(0); i < db.NumTrails; i++ {
		err := tdb.GetTrail(trail, i)

		if err != nil {
			panic(err)
		}

		line := []string{}
		transformed := "0"

		for {
			evt := trail.NextEvent()
			if evt == nil {
				break
			}

			evtMap := evt.ToMap()
			if evtMap["action"] != "" {
				fmt.Println("Action done: ", evtMap["action"])
				line = append(line, evtMap["action"])
				count += 1
			}
			if evtMap["action"] == "registration_ok" {
				transformed = "1"
			}
		}

		line = append([]string{transformed}, line...)
		err = writer.Write(line)
		if err != nil {
			panic(err)
		}
		fmt.Println("--------")
	}
}
