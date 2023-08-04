package problems

/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	for i := 0; i < l; {
		start := i
		current := i
		gasWhenArrive := 0
		moves := 0
		for {
			tmp := gasWhenArrive + gas[current] - cost[current]
			if tmp >= 0 {
				current++
				current %= l
				if current == start {
					return current
				}
				gasWhenArrive = tmp
				moves++
			} else {
				break
			}
		}
		if moves == 0 {
			i++
		} else {
			i += moves
		}
	}

	return -1
}

// @lc code=end
