package model

type Pair struct {
	Name     string    `json:"name"`
	Elements []Element `json:"els"`
}
