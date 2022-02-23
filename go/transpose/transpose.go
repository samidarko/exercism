package transpose

func Transpose(input []string) []string {
	temp := make([][]rune, 0)

	for i, s := range input {
		for j, r := range s {
			if j >= len(temp) {
				arr := make([]rune, 0)
				for k := 0; k < i; k++ {
					arr = append(arr, ' ')
				}
				arr = append(arr, r)
				temp = append(temp, arr)
			} else {
				for k := len(temp[j]); k < i; k++ {
					temp[j] = append(temp[j], ' ')
				}
				temp[j] = append(temp[j], r)
			}
		}
	}

	result := make([]string, 0)

	for _, r := range temp {
		result = append(result, string(r))
	}

	return result
}
