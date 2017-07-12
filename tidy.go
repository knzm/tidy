package tidy

import (
	"bufio"
	"fmt"
	"io"
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
	// ToDo
	return n
}
