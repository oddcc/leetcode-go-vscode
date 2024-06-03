package problems

/*
 * @lc app=leetcode.cn id=1103 lang=golang
 *
 * [1103] Distribute Candies to People
 */

// @lc code=start
func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)

	current := 1
	currentPeople := 0

	for current > 0 {
		if candies <= current {
			res[currentPeople] += candies
			break
		}
		res[currentPeople] += current
		candies -= current
		current += 1
		currentPeople += 1
		if currentPeople >= num_people {
			currentPeople = 0
		}
	}

	return res
}

// @lc code=end
