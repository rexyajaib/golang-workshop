package helper

func RemoveFromSlice[T any](s []T, index int) []T {
	values := append(s[:index], s[index+1:]...)
	return values
}

func AddToSliceIfNotExists[T comparable](s []T, value T) []T {
	for _, v := range s {
		if v == value {
			return s
		}
	}
	return append(s, value)
}

func GetFromSlice[T any](s []T, index int) T {
	return s[index]
}
