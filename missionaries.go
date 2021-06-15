// Joshua Meranda
// Laurence Madsen
// Nathan Kruger
// Missionaries and Cannibals for CSI 380
// This program solves the classic Missionaries and Cannibals problem

package main

import (
	"fmt"
)

// A representation of the state of the game
type position struct {
	boatOnWestBank   bool // true is west bank, false is east bank
	westMissionaries int  // west bank missionaries
	westCannibals    int  // west bank cannibals
	eastMissionaries int  // east bank missionaries
	eastCannibals    int  // east bank cannibals
}

// Is this a legal position? In particular, does it have
// more cannibals than missionaries on either bank? Because that is illegal.
func valid(pos position) bool {
	// Ensure 0 <= n <= 3 where n is the amount of either missionaries or cannibals on either coast
	return pos.westMissionaries >= 0 && pos.westMissionaries <= 3 &&
		pos.westCannibals >= 0 && pos.westCannibals <= 3 &&
		pos.eastMissionaries >= 0 && pos.eastMissionaries <= 3 &&
		pos.eastCannibals >= 0 && pos.eastCannibals <= 3 &&

		// Ensure the proper amount of cannibals and missionaries
		pos.westMissionaries+pos.eastMissionaries <= 3 &&
		pos.westCannibals+pos.eastCannibals <= 3 &&

		// Ensure cannibals never outnumber missionaries where missionaries are present
		(pos.westMissionaries == 0 || pos.westMissionaries >= pos.westCannibals) &&
		(pos.eastMissionaries == 0 || pos.eastMissionaries >= pos.eastCannibals)
}

// What are all of the next positions we can go to legally from the current position
// Returns nil if there are no valid positions
func (pos position) successors() []position {
	//https://github.com/marianafranco/missionaries-and-cannibals/blob/master/java/src/State.java was inspiring
	var successors = make([]position, 0, 5) //array to hold valid moves

	factor := 1
	if !pos.boatOnWestBank {
		factor = -1
	}

	for i := 0; i < 3; i++ { // number of missionaries
		for j := 2 - i; j >= 0; j-- { //number of cannibals
			if i+j <= 2 && i+j > 0 {
				missionaries := i * factor
				cannibals := j * factor
				move := position{
					!pos.boatOnWestBank,
					pos.westMissionaries - missionaries,
					pos.westCannibals - cannibals,
					pos.eastMissionaries + missionaries,
					pos.eastCannibals + cannibals,
				}

				if valid(move) { // if move is valid add it to the list of successors
					successors = append(successors, move)
				}
			}
		}
	}

	if len(successors) > 0 {
		return successors //return all valid moves if there are any
	}

	return nil // returns nil if there are no valid moves
}

// A recursive depth-first search that goes through to find the goal and returns the path to get there
// Returns nil if no solution found
func dfs(start position, goal position, solution []position, visited map[position]bool) []position {
	if start == goal {
		return []position{start}
	} else if start == (position{true, 0, 2, 3, 1}) {
		return nil
	}

	// Mark start as visited
	visited[start] = true

	temp := start.successors()
	// Iterate over potential moves
	for _, child := range temp {
		if !visited[child] { // Only check unvisited nodes
			subSolution := dfs(child, goal, solution, visited)
			if subSolution != nil {
				// Prepend solution array with current start position
				return append([]position{start}, subSolution...)
			}
		}
	}

	// No potential moves rendered a solution
	return nil
}

func main() {
	start := position{boatOnWestBank: true, westMissionaries: 3, westCannibals: 3, eastMissionaries: 0, eastCannibals: 0}
	goal := position{boatOnWestBank: false, westMissionaries: 0, westCannibals: 0, eastMissionaries: 3, eastCannibals: 3}
	solution := dfs(start, goal, []position{start}, make(map[position]bool))
	fmt.Println(solution)
}
