package main

import (
	"sort"
)

func findMovies(durations []int, n int) ([][]int, []int, []int) {
	sort.Ints(durations)
	var moviePairs [][]int
	var singleEqual []int
	var singleLess []int

	// Check for pairs of movies with a total duration exactly equal to the flight duration
	for i, duration1 := range durations {
		for j := i + 1; j < len(durations); j++ {
			duration2 := durations[j]
			totalDuration := duration1 + duration2
			if totalDuration == n {
				moviePairs = append(moviePairs, []int{duration1, duration2})
			}
		}
	}

	// If no exact pairs found, check for pairs of movies with a total duration less than the flight duration
	if len(moviePairs) == 0 {
		for i, duration1 := range durations {
			for j := i + 1; j < len(durations); j++ {
				duration2 := durations[j]
				totalDuration := duration1 + duration2
				if totalDuration < n {
					moviePairs = append(moviePairs, []int{duration1, duration2})
				}
			}
		}
	}

	// Check for a single movie with a duration equal to the flight duration
	for _, duration := range durations {
		if duration == n {
			singleEqual = append(singleEqual, duration)
		} else if duration < n {
			singleLess = append(singleLess, duration)
		}
	}

	return moviePairs, singleEqual, singleLess
}
