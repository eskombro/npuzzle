package checker

import (
	"reflect"

	"github.com/boringteam/npuzzle/src/utils"
)

func CheckResult(current []int16, result []int16) bool {
	if reflect.DeepEqual(current, result) {
		return (true)
	}
	return (false)
}

func countInversions(tab []int16) int16 {
	inversions := 0
	for i, _ := range tab {
		for a := i + 1; a < len(tab); a++ {
			if tab[a] != 0 && tab[i] != 0 && tab[i] > tab[a] {
				inversions++
			}
		}
	}
	return int16(inversions)
}

func CheckSolvable(givenTab []int16, result []int16, size int16) bool {
	/*
		If the grid width is odd, then the number of inversions in a solvable situation is even.
		If the grid width is even, and the blank is on an even row counting from the bottom (seco nd-last, fourth-last etc)
		then the number of inversions in a solvable situation is odd.
		If the grid width is even, and the blank is on an odd row counting from the bottom (last, third-last, fifth-last etc)
		then the number of inversions in a solvable situation is even.
	*/
	if size%2 == 1 {
		if countInversions(givenTab)%2 == countInversions(result)%2 {
			return true
		}
	} else {
		rowFromBottom := size - (utils.GetEmptyTile(givenTab) / size)
		if (size % 4) == 2 {
			rowFromBottom -= 1
		}
		if rowFromBottom%2 == 1 {

			if countInversions(givenTab)%2 != countInversions(result)%2 {
				return true
			}
		} else {

			if countInversions(givenTab)%2 == countInversions(result)%2 {
				return true
			}
		}
	}
	return false
}

func BuildCorrectResult(size int16) []int16 {
	tab := make([]int16, size*size)
	utils.InitUtils(tab)

	var start int16 = 0
	step_len := size
	for step_len > 0 {
		start = buildCrown(tab, start, step_len, size)
		step_len -= 2
	}
	i := 0
	for tab[i] < size*size {
		i++
	}
	tab[i] = 0
	return (tab)
}

func buildCrown(tab []int16, step_start int16, step_len int16, total_len int16) int16 {
	// Fill top line
	offset := (total_len - step_len) / 2
	block_start := offset*total_len + offset
	val := step_start + 1
	for i := block_start; i < block_start+step_len; i++ {
		tab[i] = val
		val++
	}
	// Fill right column
	val--
	step_start = block_start + step_len - 1
	for i := step_start; i < step_start+total_len*step_len; i += total_len {
		tab[i] = val
		val++
	}
	// Fill bottom line
	val--
	step_start = step_start + total_len*(step_len-1)
	for i := step_start; i > step_start-step_len; i-- {
		tab[i] = val
		val++
	}
	// Fill left line
	val--
	step_start = step_start - (step_len - 1)
	for i := step_start; i > block_start; i -= total_len {
		tab[i] = val
		val++
	}
	val--
	return (val)
}
