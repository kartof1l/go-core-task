package main

func Intersection(a, b []int) (bool, []int) {
	m := make(map[int]bool)
	for _, v := range a {
		m[v] = true
	}

	res := []int{}
	seen := make(map[int]bool)
	for _, v := range b {
		if m[v] && !seen[v] {
			res = append(res, v)
			seen[v] = true
		}
	}

	return len(res) > 0, res
}