package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/traildb/traildb-go"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Ev struct {
	Timestamp int    `tdb:"timestamp"`
	action    string `tdb:"action"`
}

func main() {
	initTrail()
}

func initTrail() {
	cons, err := tdb.NewTrailDBConstructor("forum", "action")
	if err != nil {
		panic(err.Error())
	}

	rand.Seed(time.Now().Unix())
	actionList := []string{
		"index",
		"feature",
		"whoweare",
		"mentions",
		"pricing",
		"extra_features",
		"registration_form",
	}
	totalActionCount := len(actionList)

	tt := 0
	for i := 0; i <= 1000; i++ {
		userIdentifier := uuid.Must(uuid.NewV4()).String()
		userIdentifier = strings.Replace(userIdentifier, "-", "", -1)

		for j := 0; j <= rand.Intn(10); j++ {
			n := rand.Int() % totalActionCount
			cons.Add(userIdentifier, time.Now().Unix(), []string{actionList[n]})

			tt = tt + 1
		}

		fmt.Println(i)
	}
	cons.Finalize()

	cons.Close()
	log.Println("index %d", tt)
}
