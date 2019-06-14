package strings

// decodeEncoding decodes spreadsheet column encoding.
func decodeEncoding(enc string) int {
	result := 0
	for _, letter := range enc {
		result *= 26
		result += int(letter-'A') + 1
	}

	return result
}
