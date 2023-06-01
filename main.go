package main

import (
	"fmt"
	"math/rand"
)

/*
A 30×30 grid of squares contains 900 fleas, initially one flea per square.
When a bell is rung, each flea jumps to an adjacent square at random
(usually 4 possibilities, except for fleas on the edge of the grid or at the corners).

What is the expected number of unoccupied squares after 50 rings of the bell?
Give your answer rounded to six decimal places.
330.71
*/

const (
	gridSizeX = 30
	gridSizeY = 30

	jumps = 50
)

type Flea struct {
	X, Y int
}

func main() {
	total := 0.0
	for j := 0; j < 1000; j++ {
		fleas := make([]Flea, 0, gridSizeX*gridSizeY)
		for x := 0; x < gridSizeX; x++ {
			for y := 0; y < gridSizeY; y++ {
				fleas = append(fleas, Flea{X: x, Y: y})
			}
		}

		for jump := 0; jump < jumps; jump++ {
			m := make(map[string]int)
			for i := range fleas {
				fleas[i].jump()
				cellPosition := fmt.Sprintf("%d,%d", fleas[i].X, fleas[i].Y)
				m[cellPosition]++
			}
			if jump == 49 {
				total += float64(gridSizeX*gridSizeY - len(m))
			}
		}
	}
	average := total / 1000
	fmt.Printf("Average: %.6f", average)
}

func (c Flea) possibleJumps() []string {
	possibleJumps := []string{"up", "down", "left", "right"}
	if c.X == gridSizeX-1 {
		possibleJumps = removeJump(possibleJumps, 3)
	}
	if c.X == 0 {
		possibleJumps = removeJump(possibleJumps, 2)
	}
	if c.Y == gridSizeY-1 {
		possibleJumps = removeJump(possibleJumps, 1)
	}
	if c.Y == 0 {
		possibleJumps = removeJump(possibleJumps, 0)
	}
	return possibleJumps
}

func (c *Flea) jump() {
	possibleJumps := c.possibleJumps()
	switch possibleJumps[rand.Intn(len(possibleJumps))] {
	case "up":
		c.Y -= 1
	case "down":
		c.Y += 1
	case "left":
		c.X -= 1
	case "right":
		c.X += 1
	default:
		panic("no")
	}
}

func removeJump(jumps []string, i int) []string {
	jumps[i] = jumps[len(jumps)-1]
	return jumps[:len(jumps)-1]
}
