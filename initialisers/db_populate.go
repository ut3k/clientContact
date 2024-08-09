package initializers

import (
	"encoding/csv"
	"fmt"
	// "log"
	"os"
	"strconv"

	"github.com/ut3k/clientContact/models"
)

func PopulateClients() {
	file, err := os.Open("test/Klienci_exported_1.csv")
	if err != nil {
		fmt.Println("Error Reading CSV File:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error reading headres:", err)
	}

	ConnectToDataBase()

	for _, record := range records {
		nip, _ := strconv.Atoi(record[1])
		telefon, _ := strconv.Atoi(record[2])

		client := models.Clients{
			Name:           record[0],
			Nip:            nip,
			Phone:          telefon,
			ContactPersone: record[12],
			// PersonePhone:    record[0],
			WWW:             record[3],
			FB:              record[4],
			Email:           record[5],
			Adres:           record[6],
			AdresGoogleCODE: record[7],
			GoogleMapURL:    record[13],
			// LastContact:     record[8],
			// NextContact:     record[11],
		}

		// fmt.Println(client)
		DB.Create(&client)
		// if result.Error != nil {
		// 	log.Printf("Failed to insert record: %v", result.Error)
		// }
	}

}
