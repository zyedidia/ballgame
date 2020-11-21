package main

import "math/rand"

type AI struct {
}

func NewAI() *AI {
	return &AI{}
}

func (ai *AI) Get(a Action) float64 {
	if rand.Intn(10) > 8 {
		return 1.0
	}
	return 0.0
}

func (ai *AI) String() string {
	return "AI Controller"
}
