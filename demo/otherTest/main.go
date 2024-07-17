package main

import (
	"encoding/json"
	"fmt"
	"github.com/Thenecromance/OurStories/application/models"
)

func main() {
	t := models.Travel{}

	t.TogetherWith.Set(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	marshal, err := json.Marshal(t)
	if err != nil {
		return
	}

	fmt.Println(string(marshal))

}
