DAY=$1
echo """
package days

import (
	\"bufio\"
)

type Day$DAY struct{}

func (_ Day$DAY) Part1(scanner *bufio.Scanner) {
}

func (_ Day$DAY) Part2(scanner *bufio.Scanner) {
}

""" > days/day$DAY.go
