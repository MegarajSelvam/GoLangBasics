package Ex_14_Interfaces

import (
	"bytes"
	"fmt"
	"io"
)

func Interface() {
	basicInterface()
	typeInterface()
	compositeInterface()
	emptyInterface()
	advancedInterface()
}

// Interface Declaration
// Interfaces describes behavior
type Writer interface {
	Write([]byte) (int, error)
}

// Implementing Interface
// This type is implementing this interface
type ConsoleWriter struct{}

/* In Other programming langages we use this kind of approach to implement interface

public class ConsoleWriter : Writer
{
	public (int, error) Write(byte[] input)
	{

	}
}

*/

// But in Go, we implicity implement the interface
// Below code snippets indicates, "ConsoleWriter" struct implemented "Writer" Interface.
// All we did is created a method specific to type which follows the same signature of interface
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	// Write your own logics
	fmt.Println("I am basic console writer")
	n, err := fmt.Println((string(data)))
	return n, err
}

func basicInterface1() {
	// Creating a Writer Interface
	var writer Writer

	// Assigning the struct which implements Writer interface
	writer = ConsoleWriter{}

	// Calling the interface method
	writer.Write([]byte("Hello Basic Console Writer!"))
}

// Creating another Struct which also implement this interface

type AdvancedConsoleWriter struct{}

func (advancedConsoleWriter AdvancedConsoleWriter) Write(data []byte) (int, error) {
	// Write you own logics
	fmt.Println("I am advanced console writer")
	n, err := fmt.Println((string(data)))
	return n, err
}

func basicInterface2() {
	var writerV2 Writer = AdvancedConsoleWriter{}
	writerV2.Write([]byte("Hello Advanced Console Writer"))
}

func basicInterface() {
	basicInterface1()
	basicInterface2()
}

// Any type that have a method associated with it can implement interface. We don't need struct. In below example, IntCounter is int type. Not a struct
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (intCoutner *IntCounter) Increment() int {
	*intCoutner++
	return int(*intCoutner)
}

// Composing Interface - Making Interface for Interface
type WriterV2 interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

func typeInterface() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

// Composed Interface Example.
// Embedding Interfaces within other interfaces
type WriterCloser interface {
	WriterV2
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}

		_, err = fmt.Println(string(v))

		if err != nil {
			return 0, err
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

// This is just for intializing the BufferWriterCloser struct
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func compositeInterface() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Writer"))
	wc.Close()
}

func TypeConversionInInterface() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Writer"))
	wc.Close()

	// Not Safe
	bwc := wc.(*BufferedWriterCloser)
	fmt.Println(bwc)

	// Safe way
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

// Emtpy Interface - Interfaces that has no methods
type EmptyInterface interface{}

func emptyInterface() {
	emptyInterface1()
	emptyInterface2()
}

func emptyInterface1() {
	// Creating empty interface and assigning value
	var myObj interface{} = NewBufferedWriterCloser()
	if wc, ok := myObj.(WriterCloser); ok {
		wc.Write([]byte("Hello Listeners, this is a test"))
		wc.Close()
	}

	r, ok := myObj.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}
}

func emptyInterface2() {
	var i interface{} = true
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is")
	}
}

////////////////////////////////

func advancedInterface() {
	var wc WriterCloser = &myWriterCloserV3{}
	fmt.Println(wc)
}

type WriterV3 interface {
	Write([]byte) (int, error)
}

type CloserV3 interface {
	Close() error
}

type WriterCloserV3 interface {
	WriterV3
	CloserV3
}

type myWriterCloserV3 struct{}

func (mwc *myWriterCloserV3) Write(data []byte) (int, error) {
	return 0, nil
}

func (mvc myWriterCloserV3) Close() error {
	return nil
}

/*
// Best Practices
- Use many small interfaces
	-> Single method interfaces are some of the most powerful and flexible
	-> io.Writer, io.Reader, interface{}

- Don't export interfaces for types that will be consumed

- Do export interfaces for types that will be used by package

- Design functions and methods to receive interfaces whenever possible

*/

/*
// Summary

Basics:

	type Writer interface{
		Write([]byte)(int, error)
	}

	type ConsoleWriter struct {}

	func (cw ConsoleWriter)Write(data []byte)(int, error){
		n, err := fmt.Println(string(data))
		return n, err
	}

Composing Interfaces:

	type Writer interface{
		Write() string
	}

	type Closer interface{
		Close() error
	}

	type WriterCloser interface {
		Writer
		Closer
	}


Type Conversion:

var wc WriterCloser = NewBufferWriterCloser()
bwc := wc.(*BufferedWriterCloser)


Empty Interface and Type Switches:
	var i interface{} = 0
	switch i.(type) {
		case int:
			fmt.Println("i is an integer")
		case string:
			fmt.Println("i is a string")
		default:
			fmt.Println("I don't know what i is")
		}


Implementing with Values vs. Pointers
	- Method set of value is all methods with value receivers
	- Method set of pointer is all methods, regardless of receiver type
*/
