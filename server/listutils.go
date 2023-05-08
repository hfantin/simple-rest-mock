package server

func contains[T comparable](elems []T, v T) bool {
	// log.Printf("check if %v contains %v \n", elems, v)
	for _, s := range elems {
		if v == s {
			return true
		}
	}
	return false
}
