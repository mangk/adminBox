package util

func SlicesDistinct[T comparable](slice []T) []T {
	// 使用 map 来记录已经出现的元素
	seen := make(map[T]bool)
	result := []T{}

	for _, v := range slice {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}

	return result
}
