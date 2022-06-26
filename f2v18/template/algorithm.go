//go:build gogensed

//go:generate gogensed gen/tint T=int
package template

// 返回 slice 中首个 >= value 的元素的索引，或若找不到则返回 len(slice)-1
func LowerBound(slice []T, value T) T {
	left, right := 0, len(slice)-1
	for left <= right {
		m := left + ((right - left + 1) >> 1)
		if slice[m] < value {
			left = m + 1
			continue
		}
		right = m
	}
	return right
}
