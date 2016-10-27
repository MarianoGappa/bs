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
	unit := flag.String("u", "", "output unit ([B|K|M|G|T]; default adapts to value)")
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

	suffix := ""
	switch {
	case (*unit == "" && total < 1000) || *unit == "B":
		if *unit == "" {
			suffix = "B"
		}
		fmt.Println(strconv.FormatInt(int64(round(total)), 10) + suffix)
	case (*unit == "" && total < 1000000) || *unit == "K":
		if *unit == "" {
			suffix = "K"
		}
		fmt.Println(strconv.FormatFloat(total/1000, 'f', 1, 64) + suffix)
	case (*unit == "" && total < 1000000000) || *unit == "M":
		if *unit == "" {
			suffix = "M"
		}
		fmt.Println(strconv.FormatFloat(total/1000000, 'f', 1, 64) + suffix)
	case (*unit == "" && total < 1000000000000) || *unit == "G":
		if *unit == "" {
			suffix = "G"
		}
		fmt.Println(strconv.FormatFloat(total/1000000000, 'f', 1, 64) + suffix)
	default:
		if *unit == "" {
			suffix = "T"
		}
		fmt.Println(strconv.FormatFloat(total/1000000000000, 'f', 1, 64) + suffix)
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
