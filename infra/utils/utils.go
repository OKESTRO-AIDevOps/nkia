package utils

import "time"

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

func FindFromSlice[K comparable](slice []K, el K) int {

	slice_len := len(slice)

	return_stat := -1

	for i := 0; i < slice_len; i++ {

		if slice[i] == el {

			return_stat = i

			break

		}

	}

	return return_stat

}

func GetStringTimeNow() string {

	c_time := time.Now()

	str_time_now := c_time.Format("2006-01-02 15:04:05")

	return str_time_now
}
