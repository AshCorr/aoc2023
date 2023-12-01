package days

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var DICTIONARY = [][]byte{
	{'0', 0},
	{'1', 1},
	{'2', 2},
	{'3', 3},
	{'4', 4},
	{'5', 5},
	{'6', 6},
	{'7', 7},
	{'8', 8},
	{'9', 9},
	{'z', 'e', 'r', 'o', 0},
	{'o', 'n', 'e', 1},
	{'t', 'w', 'o', 2},
	{'t', 'h', 'r', 'e', 'e', 3},
	{'f', 'o', 'u', 'r', 4},
	{'f', 'i', 'v', 'e', 5},
	{'s', 'i', 'x', 6},
	{'s', 'e', 'v', 'e', 'n', 7},
	{'e', 'i', 'g', 'h', 't', 8},
	{'n', 'i', 'n', 'e', 9},
}

func match_pattern(values string, possible_numbers [][]byte, depth int) int {
	if len(possible_numbers) == 0 || len(values) == 0 {
		return -1
	}

	NEW_POSSIBLE_NUMBERS := [][]byte{}

	for _, possible_number := range possible_numbers {
		if possible_number[depth] == values[0] {
			if possible_number[depth+1] < 10 {
				return int(possible_number[depth+1])
			}

			NEW_POSSIBLE_NUMBERS = append(NEW_POSSIBLE_NUMBERS, possible_number)
		}
	}

	return match_pattern(values[1:], NEW_POSSIBLE_NUMBERS, depth+1)
}

func One(data []string, part int) int {
	sum := 0

	for _, line := range data {
		numbers := []int{}

		for i := 0; i < len(line); i++ {
			value := match_pattern(line[i:], DICTIONARY, 0)

			if value != -1 {
				numbers = append(numbers, value)
			}
		}

		sum += (numbers[0] * 10) + numbers[len(numbers)-1]
	}

	return sum
}
