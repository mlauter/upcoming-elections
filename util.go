package main

func stringInSlice(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}

	return false
}
