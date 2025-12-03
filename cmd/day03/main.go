package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

// Processes one bank and finds the largest joltage.
func findLargestJoltage(bank string, result *atomic.Int32) {
	totalBatteries := len(bank)

	// Keep track of the largest leading and trailing battery joltage
	var maxLeading rune
	var maxTrailing rune

	for idx, battery := range bank {
		if battery > maxLeading && idx < totalBatteries-1 {
			// If the current battery has a larger joltage than `maxLeading` then we set it as the new `maxLeading` value.
			// We also have to reset the `maxTrailing` since an update to `maxLeading` invalidates the `maxTrailing` value.
			// NOTE: The last battery in the bank can only overwrite the `maxTrailing` value, even if it has a larger joltage
			// than `maxLeading`.
			maxLeading = battery
			maxTrailing = '0'
		} else if battery > maxTrailing {
			// If the current battery has a larger joltage than `maxTrailing` then we set it as the new `maxTrailing` value.
			maxTrailing = battery
		}
	}

	// Combine runes to form a string which represents a joltage in range `[10, 99]` (start/end inclusive)
	composition := string(maxLeading) + string(maxTrailing)
	joltage, err := strconv.Atoi(composition)
	if err != nil {
		panic(err)
	}

	// Add the largest joltage to the atomic counter/result
	result.Add(int32(joltage))
}

func main() {
	path := filepath.Join("inputs", "day-03.txt")

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Replace CRLF with LF (windows)
	lines := strings.ReplaceAll(string(data), "\r\n", "\n")

	banks := strings.Split(lines, "\n")

	// Atomic counter which all goroutines will add their largest joltage to
	var result atomic.Int32

	var wg sync.WaitGroup

	// Process all banks concurrently
	for _, bank := range banks {
		wg.Add(1)
		go func() {
			defer wg.Done()
			findLargestJoltage(bank, &result)
		}()
	}

	wg.Wait()

	fmt.Println("Solution (part 1):", result.Load())
}
