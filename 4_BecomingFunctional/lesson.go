package main

// functions are the bread and butter of life.
// they can do some wondrous things.
// they generally look something like this

// function name(input, input, ...) {
//     things to do
//     return value
// }
// although, the return statement is purely optional, as not every function returns something!

// and they are called by using something like this
// name(input, input, ...)
// they can be used almost anywhere a number or a variable can be used!

// they are used to do things to inputs that can be useful later,
// or to package repeated sections of code into its own box.

func mul_by_two(number float32) float32 {
	return number * 2
}

// println("mul_by_two(5) =",mul_by_two(5)) called in main later

func mul_two_numbers(num1 float32, num2 float32) float32 {
	return num1 * num2
}

//println("mul_two_numbers(6, 3) =",mul_two_numbers(6, 3))  called in main later

func calc_factorial(inp int) int {
	var result int = 1
	for inp > 0 {
		result *= inp
		inp -= 1
	}
	return result
}

// println("calc_factorial(10) =",calc_factorial(10))  called in main later

// functions can even call themselves, and that is called recursion.
// here is an example of a function that calculates the nth element of the fibonacci sequence.
// (horrible performance. Never use this if you need the fibonacci sequence for anything. my computer locks up at values past 42)
func fibonacci(num int) int {
	if num <= 2 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

// println("fibonacci(5) =",fibonacci(5))  called in main later

// This, unfortunately, is where we have to really talk about scope,
// now that our programs are becoming more complicated.

// a variable that is declared in a function belongs to the "local scope" of that function.
// it can only be used inside that function.

func foo() {
	var x int = 300
	println(x)
}

// if we uncomment the corresponding line in main, it will give us an error
// println(x)

// functions can also be declared inside another function. this is called nesting functions!
// you can think of it like a russian nesting doll!
func find_y_from_two_points(x float32, x1 float32, y1 float32, x2 float32, y2 float32) float32 {
	var find_slope = func(x1 float32, y1 float32, x2 float32, y2 float32) float32 {
		return (y1 - y2) / (x1 - x2)
	}
	var find_y_intercept = func(x float32, y float32, m float32) float32 {
		return -x*m + y
	}
	var slope = find_slope(x1, y1, x2, y2)
	var intercept = find_y_intercept(x1, y1, slope)

	return x*slope - intercept
}

//println("find_y_from_two_points(3,1,2,2,4) =",find_y_from_two_points(3,1,2,2,4)) ran in main below

// better yet, inner functions are able to use variables in the outer functions.
// This is because those variables are still "inside that function"
func find_y_from_two_points_v2(x float32, x1 float32, y1 float32, x2 float32, y2 float32) float32 {
	var find_slope = func() float32 {
		return (y1 - y2) / (x1 - x2)
	}
	var find_y_intercept = func(m float32) float32 {
		return -x1*m + y1
	}
	var slope = find_slope()
	var intercept = find_y_intercept(slope)

	return x*slope - intercept
}

//println("find_y_from_two_points_v2(3,1,2,2,4) =",find_y_from_two_points_v2(3,1,2,2,4)) ran in main below

// you can even pass functions as inputs to other functions...
func apply(f func(float32) float32, val float32) float32 {
	return f(val)
}

// println("apply(mul_by_two, 4) =", apply(mul_by_two, 4)) ran in main below

// or get functions out of functions...
func gen_linear(slope float32, intercept float32) func(float32) float32 {
	var result = func(x float32) float32 {
		return slope*x + intercept
	}
	return result
}

//var line func(float32,float32)float32 = gen_linear(2, 2) ran in main below
//println("line(2) =", line(2)) ran in main below

// Go does not support the short hand lambda notations of other languages, but that is OK.

// regarding scope and passing functions out of other functions, I am not sure how it works,
// other than treating it as though the function you are calling never left its scope.

func main() {
	println("mul_by_two(5) =", mul_by_two(5))

	println("mul_two_numbers(6, 3) =", mul_two_numbers(6, 3))

	println("calc_factorial(10) =", calc_factorial(10))

	println("apply(mul_by_two, 4) =", apply(mul_by_two, 4))

	var line func(float32) float32 = gen_linear(2, 2)
	println("line(2) =", line(2))

	//println(x)
}
