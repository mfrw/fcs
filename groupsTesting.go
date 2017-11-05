// This program return a randomly distributed 5 groups per student
// Constraints: No student should test his/her own group

// Very very hacky script

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

var totalGroups int = 21

// utiltiy to test a value in slice
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// allocate per student, barring his own
func perStudent(mygroup int) {
	var testSlice []int
	j := 0
	for j < 5 {
		toTest := rand.Intn(21)
		toTest++
		// don't test own group
		if toTest == mygroup || groups[toTest] == 0 || toTest == 9 { // group 9 does not exist
			continue
		}

		// test uniqe groups
		if contains(testSlice, toTest) {
			continue
		}

		testSlice = append(testSlice, toTest)
		j++
		groups[toTest]--
	}
	sort.Ints(testSlice)
	fmt.Println(testSlice)
}

var groups map[int]int
var maxTests int = 27

// initialize groups to 26
func initGroups(totalGroups int) {
	groups = make(map[int]int)

	for i := 1; i <= totalGroups; i++ {
		if i == 9 {
			continue
		}
		groups[i] = maxTests
	}

}

// calculate a seemingly random distribution & ask the students to choose.
func main() {

	// 42 was giving very skewed results :(
	rand.Seed(41)
	initGroups(totalGroups)

	for i := 1; i <= totalGroups; i++ {
		// does not exist
		if i == 9 {
			continue
		}

		for j := 0; j < 5; j++ {
			perStudent(i)
		}
	}
	//	printRemaining(groups)
}

func printRemaining(g map[int]int) {
	var keys []int
	for k := range g {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Printf("Group: %2d has remaining %2d\n", k, g[k])
	}

}
