package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	path := filepath.Join("inputs", "day-02.txt")

	data, err := os.ReadFile(path)
	if err != nil {

		panic(err)
	}

	// Raw sequences of id ranges (e.g. "123-456")
	sequences := strings.Split(string(data), ",")
	// Sum of all invalid IDs as numeric values
	result := 0

	for _, raw := range sequences {
		// Get the first and last ID of the given range/sequence
		ids := strings.Split(raw, "-")
		firstId := ids[0]
		lastId := ids[1]

		// The intuitive but naive approach to solve the issue would be the following:
		// (1.) Extract the first and last ID of the range (here: `firstId` and `lastId`)
		// (2.) Check if the length of both `firstId` and `lastId` is equal and an odd number
		//		- Yes: All IDs in between must be valid. Check next range (1.)
		// (3.) Iterate over all IDs (as integers) in the range `[firstId, lastId]` (start/end inclusive)
		// (4.) Transform the given ID to a string
		// (5.) Check if the length of the string is even or odd
		//		- Odd: ID is valid, since it cant consist of two consecutive repeating patterns. Check next ID
		//		- Even: ID might be invalid, since the length allows for two consecutive repeating patterns. Go to (4.)
		// (6.) Split the ID into two halves and check if both parts are the same string
		//		- Yes: ID is invalid. Parse it as an integer and add it to the result
		//		- No: ID is valid. Do nothing
		// (7.) Check next ID
		//
		// The issue with this naive approach is that we end up checking all IDs within the given ranges, which can
		// end up with a lot of unnecessary calculations
		//
		// ------------------------------------------------------------------
		//
		// Instead, we can use the following approach:
		// (1.) Extract the first and last ID of the range (here: `firstId` and `lastId`)
		// (2.) Check if the length of both `firstId` and `lastId` is equal and an odd number
		//		- Yes: All IDs in between must be valid. Check next range (1.)
		// (3.) Extract the first half of the `firstId` and store (here: we store it in `firstIdHalf`)
		//		- If length is odd: Take the smaller half (e.g. `123` -> Take `1` instead of `12`)
		// (4.) Extract the first half of the `lastId` and store (here: we store it in `lastIdHalf`)
		//		- If length is odd: Take the greater half (e.g. `456` -> Take `45` instead of `4`)
		// (5.) Iterate over all IDs (as numbers) in the range `[firstIdHalf, lastIdHalf]` (start/end inclusive)
		// (6.) Convert the given ID to a string and concatenate it with itself
		// (7.) Since the previous step guarantees that we have an "invalid" ID (according to the challenge), we only
		//		need to check if the resulting ID (as number) lies within the given ID range `[firstId, lastId]`
		//		- Yes: Add the resulting ID (as number) to the `result`
		//		- No: Do nothing
		// (8.) Check next ID
		//
		// ------------------------------------------------------------------
		//
		// EXAMPLE:
		// Let the sequence be `123-2345`
		// (1.) `firstId` = `123`, lastId = `2345`
		// (2.) `firstId` and `lastId` don't have the same length and at least one ID has an even length
		// (3.) `firstIdHalf` = `1`
		// (4.) `lastIdHalf` = `23`
		// (5. - 8.) Iterate over all IDs in the range `[1, 23]` (start/end inclusive)
		// 		- All IDs in above range if we apply (6.): `11`, `22`, ..., `1010`, `1111`, ..., `2222`, `2323`
		// 		- Subset of IDs which are within ID range `[123, 2345]`: `1010`, `1111`, ..., `2222`, `2323`
		// 		- Add all IDs which are within the ID range to the result

		var firstIdHalf string
		var lastIdHalf string

		// Extract the first halves of `firstId` and `lastId`
		if len(firstId)/2 > 0 {
			firstIdHalf = firstId[:len(firstId)/2]
		} else {
			firstIdHalf = string(firstId[0])
		}
		if len(lastId)/2%2 == 0 {
			lastIdHalf = lastId[:len(lastId)/2]
		} else {
			lastIdHalf = lastId[:len(lastId)/2+1]
		}

		// Convert the halves to integers
		firstIdHalfInt, err := strconv.Atoi(firstIdHalf)
		if err != nil {
			panic(err)
		}
		secondIdHalfInt, err := strconv.Atoi(lastIdHalf)
		if err != nil {
			panic(err)
		}

		// Convert the full IDs to integers
		firstIdInt, err := strconv.Atoi(firstId)
		if err != nil {
			panic(err)
		}
		secondIdInt, err := strconv.Atoi(lastId)
		if err != nil {
			panic(err)
		}

		for i := firstIdHalfInt; i <= secondIdHalfInt; i++ {
			str := strconv.Itoa(i)
			// Create invalid "candidate" ID by concatenating the string with itself
			id := str + str
			// Parse the resulting ID as integer
			idInt, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			// If it's within the given ID range, add it to the result
			if idInt >= firstIdInt && idInt <= secondIdInt {
				result += idInt
			}
		}
	}

	fmt.Println("Solution (part 1):", result)
}
