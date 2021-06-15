// Missionaries and Cannibals Tests for CSI 380 by David Kopec
// This program tests the valid() function and the position.successors()
// method are returning valid results.

package main

import "testing"

// Test a valid position
// {false, 2, 2, 1, 1}
func TestValid1(t *testing.T) {
	p := position{false, 2, 2, 1, 1}
	expected := true
	actual := valid(p)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", p, expected)
	}
}

// Test a valid position
// {true, 3, 0, 0, 3}
func TestValid2(t *testing.T) {
	p := position{true, 3, 0, 0, 3}
	expected := true
	actual := valid(p)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", p, expected)
	}
}

// Test an invalid position
// {true, 4, 0, 0, 3}
func TestInvalid1(t *testing.T) {
	p := position{true, 4, 0, 0, 3}
	expected := false
	actual := valid(p)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", p, expected)
	}
}

// Test an invalid position
// {true, -1, 0, 4, 3}
func TestInvalid2(t *testing.T) {
	p := position{true, -1, 0, 4, 3}
	expected := false
	actual := valid(p)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", p, expected)
	}
}

// Test an invalid position
// {false, 2, 3, 1, 0}
func TestInvalid3(t *testing.T) {
	p := position{false, 2, 3, 1, 0}
	expected := false
	actual := valid(p)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", p, expected)
	}
}

// Test an invalid position
// {false, 2, 1, 1, 2}
func TestInvalid4(t *testing.T) {
	p := position{false, 2, 1, 1, 2}
	expected := false
	actual := valid(p)
	if actual != expected {
		t.Errorf("Test failed: expected %v to be %v", p, expected)
	}
}

// does a slice of positions contain a particular position?
func contains(ps []position, p position) bool {
	for _, item := range ps {
		if item == p {
			return true
		}
	}
	return false
}

// Check if two slices of positions are equivalent
func checkEquivalent(p1s, p2s []position) bool {
	if len(p1s) != len(p2s) { // same length
		return false
	}

	for _, p := range p1s { // same items
		if !contains(p2s, p) {
			return false
		}
	}

	return true
}

// Test successors
// {true, 3, 3, 0, 0}
func TestSuccessors1(t *testing.T) {
	p := position{true, 3, 3, 0, 0}
	expected := []position{{false, 2, 2, 1, 1}, {false, 3, 2, 0, 1}, {false, 3, 1, 0, 2}}
	actual := p.successors()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: got %v but expected %v", actual, expected)
	}
}

// Test successors
// {false, 3, 1, 0, 2}
func TestSuccessors2(t *testing.T) {
	p := position{false, 3, 1, 0, 2}
	expected := []position{{true, 3, 2, 0, 1}, {true, 3, 3, 0, 0}}
	actual := p.successors()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: got %v but expected %v", actual, expected)
	}
}

// Test successors
// {true, 0, 0, 3, 3}
func TestSuccessors3(t *testing.T) {
	p := position{true, 0, 0, 3, 3}
	var expected []position = nil
	actual := p.successors()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: got %v but expected %v", actual, expected)
	}
}

// Test successors
// {true, 2, 2, 1, 1}
func TestSuccessors4(t *testing.T) {
	p := position{true, 2, 2, 1, 1}
	expected := []position{{false, 1, 1, 2, 2}, {false, 0, 2, 3, 1}}
	actual := p.successors()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: got %v but expected %v", actual, expected)
	}
}

// Test successors
// {false, 2, 2, 1, 1}
func TestSuccessors5(t *testing.T) {
	p := position{false, 2, 2, 1, 1}
	expected := []position{{true, 3, 3, 0, 0}, {true, 3, 2, 0, 1}}
	actual := p.successors()
	if !checkEquivalent(expected, actual) {
		t.Errorf("Test failed: got %v but expected %v", actual, expected)
	}
}

// Test DFS
func TestDFS(t *testing.T) {
	start := position{boatOnWestBank: true, westMissionaries: 3, westCannibals: 3, eastMissionaries: 0, eastCannibals: 0}
	goal := position{boatOnWestBank: false, westMissionaries: 0, westCannibals: 0, eastMissionaries: 3, eastCannibals: 3}
	solution := dfs(start, goal, []position{start}, make(map[position]bool))

	if solution[0] != start {
		t.Errorf("Test failed: solution[0]: %v is not the start position: %v",
			solution[0], goal)
		return
	}

	if solution[len(solution)-1] != goal {
		t.Errorf("Test failed: solution[len(solution) - 1]: %v is not the goal position: %v",
			solution[len(solution)-1], goal)
		return
	}

	for i, pos := range solution {
		if !valid(pos) {
			t.Errorf("Test failed: solution[%d]: %v is not valid", i, solution[i])
			return
		}
	}

	for i := 0; i < len(solution)-1; i++ {
		if !contains(solution[i].successors(), solution[i+1]) {
			t.Errorf("Test failed: solution[%d]: %v is not a successor of solution[%d]: %v",
				i+1, solution[i+1], i, solution[i])
			return
		}
	}
}
