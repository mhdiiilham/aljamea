package main

import (
	"math/rand"
	"time"
)

type user struct {
	Name  string `json:"string"`
	Email string `json:"email"`
}

func newJamaah() []*user {
	return []*user{
		&user{
			Name:  "Ardi Gunawan",
			Email: "muhammad.ilham@klikacc.com",
		},
		&user{
			Name:  "Andjaradji Rooseno",
			Email: "muhammad.ilham@klikacc.com",
		},
		&user{
			Name:  "Muhammad Ilham",
			Email: "muhammad.ilham@klikacc.com",
		},
		&user{
			Name:  "Meigara Juma",
			Email: "muhammad.ilham@klikacc.com",
		},
		&user{
			Name:  "Kalys Khairy Lasoma",
			Email: "muhammad.ilham@klikacc.com",
		},
	}
}

func randomIndex(max int) []int {
	candidate := []int{}
	for len(candidate) < 3 {
		isUnique := true
		randomIndex := rand.Intn(max)
		rand.Seed(time.Now().UnixNano())

		for _, value := range candidate {
			if randomIndex == value {
				isUnique = false
			}
		}
		if isUnique {
			candidate = append(candidate, randomIndex)
		}
	}
	return candidate
}
