package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Body struct {
	MaxdateStr string `json:"maxdate"`
}

const location = "Asia/Tokyo"

func init() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Panic(err)
	}
	time.Local = loc
}

func main() {
	url := os.Getenv("EBICA_URL")
	if url == "" {
		log.Panic("EBICA_URL is empty")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	io.Copy(buf, resp.Body)
	bytes := buf.Bytes()
	body := new(Body)
	json.Unmarshal(bytes, body)

	md, err := time.Parse("2006/01/02", body.MaxdateStr)
	if err != nil {
		log.Panic(err)
	}

	now := time.Now()
	log.Print(now)
	if now.AddDate(0, 1, 0).Day() != md.Day() {
		log.Panic("not found update")
	}

	log.Print("found update!")
}
