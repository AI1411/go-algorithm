package min

func Bitwise(base int, value int, values ...int) int {
	min := value
	for _, val := range values {
		min = min&((min-val)>>base) | val&(^(min-val)>>base)
	}
	return min
}
