package pairsum

import "log"

// given collection of number with length of n in ascending order
// each value possible value is 0 to n
// given one value that is sum of two number,
// find the pair number that sum == given sum

// example :
// nums := []int{1,2,3,9}
// sum := 8
// expected result is false, because no pair of number in array nums could sum to 8
// while give nums := []int{1,2,4,4}
// the result is true, since nums[2] + nums [3] == sum
func hasPairWithSum() {

	nums := []int{1, 2, 3, 4, 5, 6}
	sum := 7

	result := false

	for i := 0; i < len(nums)-1; i++ {

		for j := i + 1; j < len(nums); j++ {

			r := nums[i] + nums[j]

			log.Printf("test index %d with index %d, with result : %d", i, j, r)

			if r == sum {
				log.Printf("found at index : %d, number %d, with index : %d, number %d", i, nums[i], j, nums[j])
				result = true
				break
			}
		}

		if result == true {
			break
		}

	}
	log.Println(result)

}

func HasPairWithSum2(nums []int, sum int) bool {
	var (
		result bool
	)

	low, high := 0, len(nums)-1

	counter := 1

	for low < high {

		r := nums[low] + nums[high]

		if r == sum {
			result = true
			break
		} else if r > sum {
			log.Println("result bigger, move high to the left")
			high--
		} else if r < sum {
			log.Println("result smaller, move low to the right")
			low++
		}

		counter++

	}

	log.Println("done in : ", counter, "loop")

	return result

}
