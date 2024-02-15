package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/JMjirapat/qrthrough-api/config"
	"github.com/JMjirapat/qrthrough-api/infrastructure"
	"github.com/JMjirapat/qrthrough-api/internal/adaptor/repo"
	"github.com/JMjirapat/qrthrough-api/internal/core/model"
	"gorm.io/gorm"
)

func init() {
	config.InitConfig()
	infrastructure.InitDB()
}

func main() {
	records, err := readData("cmd/alumni_csv/alumni.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records {
		alumni_id, err := strconv.Atoi(record[0])
		if err != nil {
			fmt.Println("Error during conversion")
			continue
		}

		alumni := model.Alumni{
			ID:        alumni_id,
			Firstname: record[1],
			Lastname:  record[2],
			Tel:       strings.TrimSpace(record[3]),
		}

		repo := repo.NewAlumniRepo(infrastructure.DB)
		if _, err := repo.GetById(alumni_id); err != nil {
			if err == gorm.ErrRecordNotFound {
				// Create a new record if it doesn't exist
				repo.Create(&alumni)
			}
			continue
		}

		repo.UpdateById(alumni_id, alumni)

		fmt.Printf("Student Code: %v | Name: %s %s | Tel: %s\n", alumni.ID, alumni.Firstname, alumni.Lastname,
			alumni.Tel)
	}
	fmt.Printf("Imported")
}

func readData(fileName string) ([][]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip first line
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return records, nil
}
