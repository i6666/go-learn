package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v , %s", e.When, e.What)
}

func run() error {
	e := MyError{
		time.Now(),
		"it didn't work",
	}
	return &e
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	r := strings.NewReader("Hello strong")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v \n", n, err, b)
		fmt.Printf("b[:n] = %q \n", b[:n])
		if err == io.EOF {
			break
		}

	}

}
