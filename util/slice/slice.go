package slice

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func Map[S any, T any](slice []S, mapper func(S) T) []T {
	r := make([]T, len(slice))
	for i, v := range slice {
		r[i] = mapper(v)
	}
	return r
}
