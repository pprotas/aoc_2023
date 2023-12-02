package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FetchInput(day int) string {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		log.Fatal(fmt.Errorf("AOC_SESSION environment variable is not set"))
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: session,
	})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		log.Fatal(fmt.Errorf("bad request, check your session cookie"))
	}

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Fatal(fmt.Errorf("something went oopsie woopsie: %s", body))
	}

	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
