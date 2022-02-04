package Ex_11_Control_Flow

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
 Defer
 - Used to delay execution of a statement until function exsits
 - Runs in LIFO Order (Last In First Out)
 - Arguments evaluated at time defer is executed, not at time of function execution

 Panic
 - Panic are used when your application goes into the condition where we cannot recover
 - Occur when program cannot continue at all
 - Don't use when file can't be opened, unless it is critical
 - Use for unrecoverable events - cannot obtain TCP Port for web server
 - If nothing handles panic, program will exit

Recover
- Used to recover from panics
- Only useful in deferred functions
- Current function will not attempt to continue, but higher function in call stack will
*/

// Order of Execution
// Code, Defer, Panic, Recover

func ControlFlow() {
	// We use the defer to close the resources
	deferFunc()
	deferOrderExample1()
	deferOrderExample2()
	deferUseCase()

	panicExample()
	panicAndDeferCombined()
	panicUseCase()
	callPanicer()

}

// Defer will execute after the function exit or before the function returns
func deferFunc() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End")
}

// Last In First Out
func deferOrderExample1() {
	defer fmt.Println("Start")
	defer fmt.Println("Middle")
	defer fmt.Println("End")
}

func deferOrderExample2() {
	a := "start"
	defer fmt.Println(a) // It will print "start"
	a = "end"
}

//Need to import libraries first to execute this one
func deferUseCase() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() // res is not closed here. it will be closed after the exection this line only. So we can access res in next line
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%s", robots)
}

func panicExample() {
	fmt.Println("start")
	panic("Something bad happened")
	fmt.Println("end")
}

func panicUseCase() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Go!"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
}

// Whenever some panic happened in our code, first it will check for all defer piece of code, and then execute them all and then finally it will execute the panic
func panicAndDeferCombined() {
	fmt.Println("Start")
	defer fmt.Println("this was deferred")
	panic("Something bad happened")
	fmt.Println("end")
}

func callPanicer() {
	fmt.Println("Start")
	panicker()
	fmt.Println("end")
}

// Recover
func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error: ", err)

			// While recovering if you feel like it is not possible, then we can rethrow that panic by uncommenting belw line
			// panic(err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("done panicking")
}
