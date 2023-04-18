package main

func main() {
	// To begin, a variable can be seen in two ways.
	// the first way is just a container for storing data we hold onto.
	// the second way is that of a name for a spot in memory that we push and pull data from.

	// it generally follows the format of
	// var [name] [type] = [value];
	// like this
	var variable_name float32 = 1

	// the data in the variable can be of a multitude of data points... from
	// boolean values
	var boolean_true bool = true
	var boolean_false bool = false
	// integers
	var int_a int = 1
	var int_b int = 5
	var int_c int = 210698462510602195
	// floating point numbers
	var float_a float32 = 1.1
	var float_b float32 = 3.14159
	var float_c float32 = 21021.106
	// strings
	var string_a string = "this is a string"
	var string_b string = "this is also a string"
	var string_c string = "yo what the fu__"
	// they can also represent extremely complex data types that can be wildly different!
	var array_123 [3]int = [3]int{1, 2, 3}
	var map_KV map[string]string = map[string]string{"key1": "value1", "key2": "value2"}
	// Although for this, in the beginning, you do not have to understand the complex variables. They are taught in lesson 5.

	// sometimes GO can "infer" what variable type you are using when you declare a variable.
	var probably_a_string = "yeah, this is probably a string"
	var probably_an_integer = 2
	var probably_a_float = 2.5
	// this feature can be very useful for extremely complex datatypes like this.
	var function_mul_by_2 = func(x int) int { return x * 2 }
	// that way i do not have to type out func(int)int as my datatype

	// data in variables can be retrieved and used in function calls like print.
	println("int_a =", int_a)
	println("string_c =", string_c)
	// data in variables can be used in basic equations just as easily as numbers.
	println("1.1 + 3.14159 =", float_a+float_b)

	// and data in variables can change!
	var counter int = 1
	println("I have counted", counter, "times")
	counter = counter + 1
	println("I have counted", counter, "times")
	counter = counter + 1
	println("I have counted", counter, "times")
	counter = counter + 1
	println("I have counted", counter, "times")

	// there is another concept to variables called the scope of a variable
	// this is the area in which the variables remain alive and not freed from memory.
	// for example, here we declare a function. you do not need to know what a function is.
	i_have_my_own_scope := func() {
		var variable_in_functions_scope int = 1

		// getting rid of the unused variable error
		_ = variable_in_functions_scope
	}

	// if i try to reference the variable outside of it's scope, i end up with an "___ is not defined" error.
	// uncomment the line below and you will see it.
	// println(variable_in_functions_scope)
	// this is because a variable created inside a function belongs to the "local scope" of that function, and can only be used inside that function.
	// we will get into scope in detail on lesson 4 when we really get into what a function is and what it is used for.

	// ignore below.
	// just cleaning up the declared and not used errors in a garbage way.
	Use(variable_name, probably_a_string, probably_an_integer, probably_a_float, boolean_false, boolean_true, int_a, int_b, int_c, float_a, float_b, float_c, string_a, string_b, string_c, array_123[0], map_KV["key1"], function_mul_by_2, i_have_my_own_scope)
}
func Use(v ...interface{}) {}
