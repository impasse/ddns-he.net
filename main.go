package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func ddns(client *http.Client, domain string, password string) {
	log.Println("Start ddns")
	updateURL := fmt.Sprintf("https://dyn.dns.he.net/nic/update?hostname=%s", domain)

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
	domain := os.Getenv("DOMAIN")
	password := os.Getenv("KEY")

	if domain == "" {
		log.Fatalln("DOMAIN not set")
	}

	if password == "" {
		log.Fatalln("KEY not set")
	}

	interval, err := time.ParseDuration(os.Getenv("INTERVAL"))

	if err != nil {
		interval = 5 * time.Minute
	}

	client := &http.Client{}

	tick := time.NewTicker(interval)

	for {
		select {
		case <-tick.C:
			go ddns(client, domain, password)
		}
	}
}
