package main

import (
	"encoding/json"
	"fmt"

	"github.com/dragonraid/gowit/wit"
)

func main() {
	witAPI, _ := wit.New()
	body, _, _ := witAPI.Message("How many people between tuesday and friday").Do()
	fmt.Println(body.Text)
	js, _ := json.MarshalIndent(body.Entities, "", "    ")
	fmt.Println(string(js))
}
