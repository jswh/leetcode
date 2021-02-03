package leetcode


/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	const MaxUint = ^uint32(0)
	const MinUint = 0
	const MaxInt = int32(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	var result int32
	xx := int32(x)
	for xx != 0 {
		// 负数的余数也是负数
		n := xx % 10
		if result > MaxInt/10 || (result == MaxInt / 10 && n > 7) {
			return 0
		}
		if result < MinInt/10 || (result == MinInt / 10 && n < -8) {
			return 0
		}
		result = result * 10 + n
		xx = (xx - n) / 10
	}

	return int(result)

}
// @lc code=end

