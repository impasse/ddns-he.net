package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func ddns() {
	domain := os.Getenv("DOMAIN")
	password := os.Getenv("KEY")

	if domain == "" {
		log.Fatalln("DOMAIN not set")
	}

	if password == "" {
		log.Fatalln("KEY not set")
	}

	updateURL := fmt.Sprintf("https://dyn.dns.he.net/nic/update?hostname=%s", domain)

	client := &http.Client{}
	req, err := http.NewRequest("GET", updateURL, nil)
	if err != nil {
		log.Println(err)
		return
	}

	req.SetBasicAuth(domain, password)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Result: %s\n", string(body))
}

func main() {
	interval, err := time.ParseDuration(os.Getenv("INTERVAL"))

	if err != nil {
		interval = 5 * time.Minute
	}

	tick := time.NewTicker(interval)

	for {
		select {
		case <-tick.C:
			go ddns()
		}
	}
}
