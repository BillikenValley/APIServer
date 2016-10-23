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
				Lng:     -90.2047,
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
				Lng:     -90.2403,
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
				Address: "",
				Lat:     38.607488,
				Lng:     -90.203525
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
				TransFriendly:          1,

			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		3: Shelter{
			Name: "Peter & Paul",
			Beds: make([]int, 60),
			Location: Location{
				Address: "711 Allen Ave, St. Louis, MO 63104",
				Lat:     38.607303,
				Lng:    -90.203450,
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:       1,
				AcceptsSingleWomen:     0,
				AcceptsChildren:        0,
				AcceptsDrunks:          1,
				MustHaveID:             0,
				MustHaveSSN:            0,
				AcceptsSexOffenders:    1,
				AcceptsViolentCriminal: 0,
				TransFriendly:          1,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		4: Shelter{
			Name: "Room at The Inn",
			Beds: make([]int, 20),
			Location: Location{
				Address: "3415 Bridgeland Dr, Bridgeton, MO 63044",
				Lat:     38.746497,
				Lng:     -90.429620
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
				TransFriendly:          1,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		5: Shelter{
			Name: "Saint Patrick Center - Women's night",
			Beds: make([]int, 20),
			Location: Location{
				Address: "800 N Tucker Blvd, St. Louis, MO 63101",
				Lat:     38.633424,
				Lng:     -90.195435
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:       0,
				AcceptsSingleWomen:     1,
				AcceptsChildren:        0,
				AcceptsDrunks:          1,
				MustHaveID:             0,
				MustHaveSSN:            0,
				AcceptsSexOffenders:    1,
				AcceptsViolentCriminal: 0,
				TransFriendly:          1,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		6: Shelter{
			Name: "Salvation Army Family Haven",
			Beds: make([]int, 64),
			Location: Location{
				Address: "10740 Page Ave, St. Louis, MO 63132",
				Lat:     38.692118,
				Lng:     -90.398464
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:       1,
				AcceptsSingleWomen:     1,
				AcceptsChildren:        1,
				AcceptsDrunks:          1,
				MustHaveID:             0,
				MustHaveSSN:            0,
				AcceptsSexOffenders:    0,
				AcceptsViolentCriminal: 0,
				TransFriendly:          1,
			},
			OpenTime:    time.Now(),
			ClosingTime: time.Now().Add(2 * time.Hour),
			ShelterCredentials: ShelterCredentials{
				Username: "david",
				Password: "password",
			},
		},
		7: Shelter{
			Name: "12th and Park Shelter",
			Beds: make([]int, 125),
			Location: Location{
				Address: "1410 S Tucker Blvd, St. Louis, MO 63104",
				Lat:     38.614828,
				Lng:     -90.203185
			},
			ShelterConstraints: map[string]int{
				AcceptsSingleMen:       1,
				AcceptsSingleWomen:     0,
				AcceptsChildren:        0,
				AcceptsDrunks:          1,
				MustHaveID:             0,
				MustHaveSSN:            0,
				AcceptsSexOffenders:    0,
				AcceptsViolentCriminal: 0,
				TransFriendly:          1,
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
