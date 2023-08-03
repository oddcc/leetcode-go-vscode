package problems

import "math/rand"

/*
 * @lc app=leetcode.cn id=380 lang=golang
 *
 * [380] Insert Delete GetRandom O(1)
 */

// @lc code=start
type RandomizedSet struct {
	values  []int
	hashSet map[int]int
}

func Constructor() RandomizedSet {
	r := RandomizedSet{}
	r.values = make([]int, 0)
	r.hashSet = make(map[int]int)
	return r
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exixts := this.hashSet[val]
	if exixts {
		return false
	}
	this.values = append(this.values, val)
	this.hashSet[val] = len(this.values) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	i, exixts := this.hashSet[val]
	if !exixts {
		return false
	}
	this.values[i] = this.values[len(this.values)-1]
	this.hashSet[this.values[i]] = i
	this.values = this.values[:len(this.values)-1]
	delete(this.hashSet, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.values))
	return this.values[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
