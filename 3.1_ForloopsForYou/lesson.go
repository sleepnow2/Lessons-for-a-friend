package main

func main() {
	// Today, we will be talking about the for loop.
	// Basically, a for loop is almost exactly a while loop, except that it wraps
	var i = 0    // the assignment
	for i < 10 { // the condition
		// do something
		i = i + 1 // and the iteration
	}
	// all into one statement.

	// the above statement would look like this!
	for i = 1; i < 10; i++ {
		// do something
	}

	println("==============SECT0===============")
	// but For loops can be even more useful than that.
	// for loops take in any "itterator" which is any object that implements the "__next__" function.
	// this can be anything from the range function above
	for i = 20; i < 100; i += 10 {
		println(i) // counts multiples of 10 from 20 to 100
	}

	println("==============SECT1===============")
	// to pulling values out of arrays of items
	var arr = []string{"print", "hello", "world!"}
	for _, str := range arr {
		println(str)
	}

	println("==============SECT2===============")
	// to pulling characters out of a string
	var str = "onebyone"
	for _, chr := range str {
		println(string(chr))
	}
}
