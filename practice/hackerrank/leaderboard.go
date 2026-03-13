package hackerrank

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	results := make([]int32, len(player))

	uniqueRanked := []int32{}
	if len(ranked) > 0 {
		uniqueRanked = append(uniqueRanked, ranked[0])
		for i := 1; i < len(ranked); i++ {
			if ranked[i] != ranked[i-1] {
				uniqueRanked = append(uniqueRanked, ranked[i])
			}
		}
	}

	// 2. Two-Pointer Approach
	// player scores are ascending, uniqueRanked is descending.
	// We start from the bottom of the leaderboard (the end of the slice).
	currRankIdx := len(uniqueRanked) - 1

	for i, score := range player {
		// Move the pointer up the leaderboard as long as the player's score
		// is better than or equal to the score at the current rank.
		for currRankIdx >= 0 && score >= uniqueRanked[currRankIdx] {
			currRankIdx--
		}

		// If currRankIdx is -1, they are #1.
		// Otherwise, their rank is currRankIdx + 2 (1-based index + 1 for being below the next rank)
		results[i] = int32(currRankIdx + 2)
	}

	return results
}
