package main

func FindDifference(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	for _, v := range slice2 {
		m[v] = true
	}

	res := []string{}
	for _, v := range slice1 {
		if !m[v] {
			res = append(res, v)
		}
	}
	return res
}