package main

import (
	"encoding/json"
	"fmt"

	"github.com/BillikenValley/APIServer/model"
)

func main() {
	req := model.RangeRequirement{
		RequirmentMeta: model.RequirmentMeta{
			Name: model.AgeReq,
		},
		Expected: [][2]int64{
			[2]int64{0, 100},
		},
	}
	user := model.User{}
	shelter := model.Shelter{Requirements: []model.ShelterRequirement{req}}
	thing, _ := json.MarshalIndent(user, "", " ")
	thing2, _ := json.MarshalIndent(shelter, "", " ")
	fmt.Println(string(thing))
	fmt.Println(string(thing2))

}
