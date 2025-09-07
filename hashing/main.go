package main

import "fmt"

var hash = [256]int{}

func main() {
	// arr := []rune{'A', 'r', 'c', 'e', 'b', 'C', 'a', 'f', 'a', 'r', 'b', 'c'}
	// Hash(arr)
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Printf("> ")
	// 	input, err := reader.ReadString(byte('\n'))
	// 	if err != nil {
	// 		fmt.Printf("error reading input: %+v", err)
	// 		fmt.Printf("exiting...")
	// 		return
	// 	}
	// 	input = strings.TrimSuffix(input, "\n")
	// 	input = strings.TrimPrefix(input, "> ")
	// 	if input == "exit" {
	// 		return
	// 	}

	// 	char := []rune(input)[0]

	// 	if char > 255 {

	// 		fmt.Printf("invalid character must be between %c - %c\n", 0, 255)
	// 		continue
	// 	}
	// 	fmt.Printf("Occurrences: %d\n", FetchOccurrences(char))
	// }
	arr := []int{1, 2, 34, 32, 432, 543, 25, 3, 42, 1, 531, 352, 565, 243, 45, 12, 43}
	high, low := FindHighestLowestFrequencyElement(arr)
	fmt.Printf("High: %d, Low: %d", high, low)
}

func Hash(arr []rune) {
	for _, elem := range arr {
		hash[elem] += 1
	}
}

func FetchOccurrences(ch rune) int {
	return hash[ch]
}

func FindHighestLowestFrequencyElement(arr []int) (int, int) {
	occurrenceMap := map[int]int{}
	for _, elem := range arr {
		occurrenceMap[elem] += 1
	}
	high := 0
	low := len(arr)
	for _, count := range occurrenceMap {
		if high < count {
			high = count
		}
		if low > count {
			low = count
		}
	}
	return high, low
}
