package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	rotationLeft  = "L"
	rotationRight = "R"
	initialDial   = 50
)

func main() {
	path := filepath.Join("inputs", "day-01.txt")

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Replace CRLF with LF (windows)
	lines := strings.ReplaceAll(string(data), "\r\n", "\n")

	// Current dial position
	currDial := initialDial
	// Number of times the dial ends up at 0
	hits := 0
	// Number of times the dial passes 0
	passes := 0

	for _, line := range strings.Split(lines, "\n") {
		// Extract direction ("L" or "R")
		direction := string(line[0])

		// Extract rotation amount
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if amount == 0 {
			continue
		}

		// Calculate the total number of full rotations and (remaining) effective rotation.
		// ---
		// Example: Assuming a rotation value of `312` is given.
		// Each `100` units of rotation causes a "full" rotation, i.e. the
		// dial will end up in the same position as before the full rotation occurred.
		// Therefore, `312` will cause `3` full rotations, which have no impact on the dial position.
		// Besides the full rotations, only an "effective" rotation of `12` units is left, which
		// will actually cause the dial position to change (unless 0).
		// ---

		fullRotations := amount / 100
		effectiveRotation := amount - (fullRotations * 100)

		// When taking the full rotations into account, we have to differentiate between 2 cases:
		// 1. The dial starts at 0, and we only do full rotations
		// 2. The dial does not start at 0, and we do any kind of rotations
		if currDial == 0 && effectiveRotation == 0 {
			hits++
			passes += fullRotations - 1
		} else {
			passes += fullRotations
		}

		if effectiveRotation == 0 {
			continue
		}

		// Since we already took the full rotations into account in the previous steps,
		// we only need to handle the remaining effective rotation.
		// We also need to take care of lower and upper bound wrapping.
		if direction == rotationLeft {
			newDial := ((currDial - effectiveRotation) + 100) % 100

			if newDial == 0 {
				// Case: We landed on 0 after the rotation.
				hits++
			} else if newDial > currDial && currDial != 0 {
				// Case: If the new dial position is greater than the current position, a wrap occurred.
				// But the rotation only really passed `0` if the current dial position was not already on 0.
				passes++
			}
			currDial = newDial
		} else if direction == rotationRight {
			newDial := (currDial + effectiveRotation) % 100

			if newDial == 0 {
				// Case: We landed on 0 after the rotation.
				hits++
			} else if newDial < currDial && currDial != 0 {
				// Case: If the new dial position is less than the current position, a wrap occurred.
				// But the rotation only really passed `0` if the current dial position was not already on 0.
				passes++
			}
			currDial = newDial
		}
	}

	fmt.Println("Solution (part 1):", hits)
	fmt.Println("Solution (part 2):", hits+passes)

}
