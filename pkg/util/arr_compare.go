package util

type T interface {
}

type name struct {
}

func ArrCompare[T int | float32 | float64 | string | struct{}](arr []T) bool {

	// 通过排序法
	// a := make([]int, len(arr))
	// copy(a, arr)
	// sort.Ints(a)
	// return a[0] == a[len(a)-1]

	// 通过map
	var empty struct{}
	m := make(map[T]struct{})
	for _, t := range arr {
		m[t] = empty
	}
	return len(m) <= 1
}
