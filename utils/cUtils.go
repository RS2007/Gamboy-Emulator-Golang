package utils

func MemsetInt8(arr []int8, value int8) {
	if len(arr) == 0 {
		return
	}
	arr[0] = value
	for bp := 1; bp < len(arr); bp *= 2 {
		copy(arr[bp:], arr[:bp])
	}

}
