package entity

import "github.com/p1ass/midare/twitter"

type Period struct {
	OkiTime *twitter.Tweet `json:"okiTime"`
	NeTime  *twitter.Tweet `json:"neTime"`
}
