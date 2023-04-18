// IF statements follow this basic format...
// if (condition) {
//     statements if true;
//     statements if true;
//     ...
// }

// if the condition evaluates to true, then the block of code indented below it gets run
// if it is false, that block of code is skipped and it moves on.

// you can directly use True and False for conditions, most likely stored in variables.
let var_containing_true = true;
if (var_containing_true){
    console.log("Of corse this is True!");
}

// or you can use comparisons and functions that return boolean values for your conditions.
let number = 5; // Integer

// now here is a bunch of different ways to compare things!
if (number == 5) { // Notice the two "=" signs. "=" means assignment, "==" means comparison.
    console.log("the number equals 5!");
}
if (number != 5) {
    console.log("the number does not equal 5!");
}
if (number < 5) {
    console.log("the number is greater than 5!");
}
if (number <= 5) {
    console.log("the number is greater than or equal to 5!");
}
if (number > 5) {
    console.log("the number is less than 5!");
}
if (number >= 5) {
    console.log("the number is less than or equal to 5!");
}

// you can even use "&&" for the and sign and "||" for the or sign to make more complex decisions!
if (number >= 3 && number <= 6) {
    console.log("the number lies between 3 and 6!");
}
if (number == 5 || number == 8) {
    console.log("the number is either 5 or 8!");
}
// the next part of an IF statement is the ELSE block.
// it is entirely optional and it looks like this.

// if (condition) {
//     statements if true;
// } else {
//     statements if false;
// }
let favorite_flavor = "vanilla"; // String

if (favorite_flavor == "vanilla"){
    console.log("I like vanilla too!");
} else {
    console.log("You are objectively wrong. Vanilla is the greatest!")
}

// Our next item in the flow control adventure is the While loop.
// It looks like this.

// while (condition) {
//     statements to repeat if true
//     statements to repeat if true
//     ...
// }

// here is an example.
let num_times_printed = 1; // Integer
while (num_times_printed <= 5) {
    console.log("I printed " + num_times_printed + " times!");
    num_times_printed = num_times_printed + 1;
    // I can also use "num_times_printed += 1"if I want the line to be a bit shorter.
}

    
// look out for infinite loops in your code, as it can cause nasty issues for you
// if you do not intend for it to be stuck somewhere forever.

// for example, this one causes an infinite loop that can not be broken out of.
while (num_times_printed < 6){
    // this happens because num_times_printed is ever incremented.
    // to halt execution of any program in the terminal, select the terminal and press
    //   CTRL + C
    // to cancel execution of code
}
console.log("Note how this is never, EVER printed!");