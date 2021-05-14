package main

func toLowerCase(str string) string {
	res := make([]byte, len(str))

	// go thru string
	for i := 0; i < len(str); i++ {
		// check if uppercase
		if 65 <= str[i] && str[i] <= 90 {
			// write this byte + 32
			res[i] = str[i] + 32
		} else {
			// just write this byte
			res[i] = str[i]
		}
	}

	// return byte slice as string
	return string(res)
}
