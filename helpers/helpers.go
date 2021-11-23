package helpers

import (
	"encoding/json"
	"example-ci/models"
	"fmt"
	"log"
	"os"
)

func DecodePosts(filename string) (posts *[]models.Post, err error) {
	log.Print(filename)
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("opening file: %v", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&posts)
	if err != nil {
		fmt.Printf("decoding json: %v", err)
		return
	}

	return

}
