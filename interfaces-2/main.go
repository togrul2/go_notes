package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test"))
	wc.Close()

	bwc := wc.(*BufferedWriterCloser) // Type conversion
	// bwc := wc.(io.Reader)          // Also can be done with interfaces

	fmt.Printf("%v %T", bwc, bwc)

	r, ok := wc.(io.Reader) // Type conversion also supports error handling

	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	var myObj interface{} = NewBufferWriterCloser() // Empty interface, used to handle unknown types
	fmt.Println(myObj)
}

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

// Here we combined two interfaces into signle one
type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, nil
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, printErr := fmt.Println(string(v))
		if printErr != nil {
			return 0, printErr
		}
	}
	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
