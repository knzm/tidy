package tidy

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
)

type Number int

func (n Number) Int() int {
	return int(n)
}

func (n Number) Digits() []byte {
	b := []byte(strconv.Itoa(int(n)))
	for i := 0; i < len(b); i++ {
		b[i] -= '0'
	}
	return b
}

func NumberFromDigits(b []byte) (Number, error) {
	x := 0
	for _, d := range b {
		if d < 0 || d > 9 {
			err := fmt.Errorf("Not a digit: %d", d)
			return Number(0), err
		}
		x = x*10 + int(d)
	}
	return Number(x), nil
}

func ParseInput(r io.Reader) ([]Number, error) {
	sc := bufio.NewScanner(r)

	if !sc.Scan() {
		return nil, sc.Err()
	}
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		return nil, err
	}

	var ns []Number

	for i := 0; i < n; i++ {
		if !sc.Scan() {
			break
		}
		d, err := strconv.Atoi(sc.Text())
		if err != nil {
			return nil, err
		}
		ns = append(ns, Number(d))
	}

	err = sc.Err()
	if err != nil {
		return nil, err
	}

	return ns, nil
}

func PrintOutput(w io.Writer, i int, input, output Number) {
	fmt.Fprintf(w, "Case #%d: %d\n", i+1, output.Int())
}

func Solve(n Number) Number {
	digits := n.Digits()

	// upper bound of each digit
	bounds := make([]byte, len(digits))

	var bound byte = 9
	for i := len(digits) - 1; i >= 0; i-- {
		bound = func(x, y byte) byte {
			if x < y {
				return x
			}
			return y
		}(digits[i], bound)
		bounds[i] = bound
	}

	result := make([]byte, len(digits))
	for i := 0; i < len(result); i++ {
		result[i] = 9
	}
	for i := 0; i < len(result); i++ {
		if digits[i] > bounds[i] {
			result[i] = digits[i] - 1
			break
		}
		result[i] = digits[i]
	}

	n, err := NumberFromDigits(result)
	if err != nil {
		log.Panic(err)
	}

	return n
}
