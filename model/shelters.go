package model

import "time"

var shelters map[int]Shelter

func init() {
	shelters = map[int]Shelter{
		0: Shelter{
			Name: "Gateway 180",
			Beds: make([]int, 174),
			Location: Location{
				Address: "1000 N. 19th St.\nSt.Louis, MO 63135",
				Lat:     38.6376,
				Lng:     90.2047,
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:    0,
				AcceptsSingleWomen:  1,
				AcceptsChildren:     1,
				AcceptsDrunks:       1,
				MustHaveID:          0,
				MustHaveSSN:         0,
				AcceptsSexOffenders: 0,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		1: Shelter{
			Name: "Our Lady's Inn",
			Beds: make([]int, 174),
			Location: Location{
				Address: "4223 S. Compton    St. Louis, MO 6311",
				Lat:     38.5799,
				Lng:     90.2403,
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:    0,
				AcceptsSingleWomen:  1,
				MustBePregnant:      1,
				AcceptsChildren:     1,
				MaxChildren:         3,
				AcceptsDrunks:       1,
				MustHaveID:          0,
				MustHaveSSN:         0,
				AcceptsSexOffenders: 0,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		2: Shelter{
			Name: "Loaves & Fishes Inc.",
			Beds: make([]int, 21),
			Location: Location{
				Address: "2750 McKelvey Rd, Maryland Heights, MO 63043",
				Lat:     38.7333,
				Lng:     90.4457,
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:       0,
				AcceptsSingleWomen:     1,
				AcceptsChildren:        1,
				AcceptsDrunks:          1,
				MustHaveID:             0,
				MustHaveSSN:            0,
				AcceptsSexOffenders:    0,
				AcceptsViolentCriminal: 0,
				TransFriendly:          0,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
	}
	for i := range shelters {
		shelter := shelters[i]
		shelter.ShelterID = i
	}
}
