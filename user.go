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
			Email: "",
		},
		&user{
			Name:  "Andjaradji Rooseno",
			Email: "",
		},
		&user{
			Name:  "Muhammad Ilham",
			Email: "muhammad.ilham@klikacc.com",
		},
		&user{
			Name:  "Meigara Juma",
			Email: "",
		},
		&user{
			Name:  "Kalys Khairy Lasoma",
			Email: "",
		},
	}
}

func randomIndex(max int) []int {
	candidate := []int{}
	for len(candidate) < max {
		isUnique := true
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(max)

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
