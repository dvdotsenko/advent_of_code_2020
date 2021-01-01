package fio

import (
	"bufio"
	"bytes"
	"io"
	"os"
)


func FileAsLines(fn string) (lines chan string) {

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(f)

	lines = make(chan string, 1)

	go func () {
		for err == nil {
			var buffer bytes.Buffer
			var fragment []byte
			var isPartial bool
			for {
				fragment, isPartial, err = reader.ReadLine()
				buffer.Write(fragment)

				if err != nil {
					break
				}
				// Partial == false means end of line
				if !isPartial {
					break
				}
			}
			lines <- buffer.String()
		}
		f.Close()
		close(lines)
		if err != nil && err != io.EOF {
			panic(err)
		}
	}()

	return lines
}
