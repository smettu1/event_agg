package parse


import (
	"encoding/json"
	models "github.com/events/models"
	"log"
)

func ParseData(data string) (models.Input, error) {
	var element models.Input
	err := json.Unmarshal([]byte(data), &element)
	if err != nil || element.Id == 0 || element.Market == 0{
		log.Println("Invalid data...moving on")
		return element,err
	}
	return element, nil
}

