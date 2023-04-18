package main

func main() {
	// IF statements follow this basic format...
	// if (condition) {
	//     statements if true;
	//     statements if true;
	//     ...
	// }

	// if the condition evaluates to true, then the block of code indented below it gets run
	// if it is false, that block of code is skipped and it moves on.

	// you can directly use True and False for conditions, most likely stored in variables.
	var var_containing_true bool = true // boolean
	if var_containing_true {
		println("Of corse this is True!")
	}

	// or you can use comparisons and functions that return boolean values for your conditions.
	var number int = 5 // Integer

	// now here is a bunch of different ways to compare things!
	if number == 5 { // Notice the two "=" signs. "=" means assignment, "==" means comparison.
		println("the number equals 5!")
	}
	if number != 5 {
		println("the number does not equal 5!")
	}
	if number < 5 {
		println("the number is greater than 5!")
	}
	if number <= 5 {
		println("the number is greater than or equal to 5!")
	}
	if number > 5 {
		println("the number is less than 5!")
	}
	if number >= 5 {
		println("the number is less than or equal to 5!")
	}

	// you can even use "&&" for the and sign and "||" for the or sign to make more complex decisions!
	if number >= 3 && number <= 6 {
		println("the number lies between 3 and 6!")
	}
	if number == 5 || number == 8 {
		println("the number is either 5 or 8!")
	}
	// the next part of an IF statement is the ELSE block.
	// it is entirely optional and it looks like this.

	// if (condition) {
	//     statements if true;
	// } else {
	//     statements if false;
	// }
	var favorite_flavor string = "vanilla" // String

	if favorite_flavor == "vanilla" {
		println("I like vanilla too!")
	} else {
		println("You are objectively wrong. Vanilla is the greatest!")
	}

	// Our next item in the flow control adventure is the basic for loop.
	// It looks like this.

	// for condition {
	//     statements to repeat if true
	//     statements to repeat if true
	//     ...
	// }

	// here is an example.
	var num_times_printed int = 1 // Integer
	for num_times_printed <= 5 {
		println("I printed ", num_times_printed, " times!")
		num_times_printed = num_times_printed + 1
		// I can also use "num_times_printed += 1"if I want the line to be a bit shorter.
	}

	// look out for infinite loops in your code, as it can cause nasty issues for you
	// if you do not intend for it to be stuck somewhere forever.

	// for example, this one causes an infinite loop that can not be broken out of.
	for num_times_printed < 6 {
		// this happens because num_times_printed is ever incremented.
		// to halt execution of any program in the terminal, select the terminal and press
		//   CTRL + C
		// to cancel execution of code
	}
	println("Note how this is never, EVER printed!")
}
