package utils

func InsertIntoSlice[K comparable](a []K, index int, value K) []K {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func PushBackSlice[K comparable](a []K, value K) []K {

	a = append(a, value)
	return a
}

func DeleteFromSlice[K comparable](slice []K, s int) []K {
	return append(slice[:s], slice[s+1:]...)
}
