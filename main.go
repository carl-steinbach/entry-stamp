package main

import (
	"embed"
	"flag"
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
)

//go:embed fonts/arial_narrow.txt
var fontFile embed.FS

//go:embed config.yaml
var configFile []byte

const fontPath string = "fonts/arial_narrow.txt"

type StampConfig struct {
	Header string
	Footer string
}

func main() {
	var config StampConfig
	var fontSize int
	var dateString string
	var err error
	var selectedTime time.Time

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Could not read config file. %s", err)
	}

	// Command line arguments.
	flag.IntVar(&fontSize, "Größe", 32, "Die verwendete Schriftgröße.")
	flag.StringVar(&dateString, "Datum", time.Now().Format(time.DateOnly), "Das verwendete Datum im Format 'YYYY-MM-DD'; Ist automatisch als heutiges Datum konfiguriert.")
	flag.Parse()

	selectedTime, err = time.Parse(time.DateOnly, dateString)
	if err != nil {
		fmt.Printf("Das angegebene Datum konnte nicht gelesen werden, bitte das '--Datum' Argument im Format 'YYYY-MM-DD' angeben: %s\n", err)
	}

	err = CreateStamp(fontSize, fontFile, fontPath, selectedTime, config.Header, config.Footer)
	if err != nil {
		fmt.Printf("Fehler beim generieren des Stempels, bitte Eingabewerte & Zugriffsrechte kontrollieren: %s\n", err)
	}
}
