package google

import (
	"log"

	"google.golang.org/api/sheets/v4"
)

func SheetsService() *sheets.Service {
	srv, err := sheets.New(Client())
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	return srv
}
