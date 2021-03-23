package main

import (
	"bufio"
	"io"

	"github.com/pkg/errors"
)

func CountLine(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)
	for {
		_, err = br.ReadString('\n')
		if err != nil {
			break
		}
		lines++
	}
	if err != io.EOF {
		return 0, err
	}

	return lines, err
}

func test() error {
	err := errors.New("this is error")
	return errors.Wrap(err, "hahhahah")
}

func main() {

}
