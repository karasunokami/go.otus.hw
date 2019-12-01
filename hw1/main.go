package main

import (
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("0.ru.pool.ntp.org")

	if err != nil {
		log.Fatalf("Error: %s", err)

		return
	}

	log.Printf("NTP   time: %s", ntpTime)
	log.Printf("Local time: %s", time.Now())
}
