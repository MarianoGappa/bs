package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	insensitive := flag.Bool("i", false, "case insensitive matching")
	flag.Parse()

	iModif := ""
	if *insensitive {
		iModif = "(?i)"
	}

	scanner := bufio.NewScanner(os.Stdin)
	total := float64(0)
	for scanner.Scan() {
		for _, v := range strings.Fields(scanner.Text()) {
			re := iModif + "^[\\d\\.]+[B|K|M|G|T]$"
			if matched, err := regexp.MatchString(re, v); err == nil && matched {
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
		fmt.Println(strconv.FormatFloat(total/1000, 'f', 1, 64) + "K")
	case total < 1000000000:
		fmt.Println(strconv.FormatFloat(total/1000000, 'f', 1, 64) + "M")
	case total < 1000000000000:
		fmt.Println(strconv.FormatFloat(total/1000000000, 'f', 1, 64) + "G")
	default:
		fmt.Println(strconv.FormatFloat(total/1000000000000, 'f', 1, 64) + "T")
	}
}

func multiplier(l rune) float64 {
	switch l {
	case 'B', 'b':
		return float64(1)
	case 'K', 'k':
		return float64(1000)
	case 'M', 'm':
		return float64(1000000)
	case 'G', 'g':
		return float64(1000000000)
	case 'T', 't':
		return float64(1000000000000)
	}
	return float64(1)
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
