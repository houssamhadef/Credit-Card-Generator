package functions

func reverse(slice string) string {
	reversed := ""

	for i := len(slice) - 1; i >= 0; i-- {
		reversed += string(slice[i])
	}
	return reversed
}
