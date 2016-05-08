package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"regexp"
	"strconv"
	"math"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := float64(0)
	for scanner.Scan() {
		for _, v := range strings.Fields(scanner.Text()) {
			if matched, err := regexp.MatchString("^[\\d\\.]+[B|K|M|G|T]$", v); err == nil && matched {
				i, err := strconv.ParseFloat(v[:len(v)-1], 64)
				if err == nil {
					total += i * multiplier(rune(v[len(v)-1]))
					break
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	switch {
	case total < 1000:
		fmt.Println(strconv.FormatInt(int64(round(total)), 10) + "B")
	case total < 1000000:
		fmt.Println(strconv.FormatFloat(total / 1000, 'f', 1, 64) + "K")
	case total < 1000000000:
		fmt.Println(strconv.FormatFloat(total / 1000000, 'f', 1, 64) + "M")
	case total < 1000000000000:
		fmt.Println(strconv.FormatFloat(total / 1000000000, 'f', 1, 64) + "G")
	default:
		fmt.Println(strconv.FormatFloat(total / 1000000000000, 'f', 1, 64) + "T")
	}
}

func multiplier(l rune) float64 {
	switch l {
	case 'B': return float64(1)
	case 'K': return float64(1000)
	case 'M': return float64(1000000)
	case 'G': return float64(1000000000)
	case 'T': return float64(1000000000000)
	}
	return float64(1)
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
