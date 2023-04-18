# IF statements follow this basic format...
# if condition:
#     statements if true
#     statements if true
#     ...

# if the condition evaluates to true, then the block of code indented below it gets run
# if it is false, that block of code is skipped and it moves on.

# you can directly use True and False for conditions, most likely stored in variables.
var_containing_true = True
if var_containing_true:
    print("Of corse this is True!")

# or you can use comparisons and functions that return boolean values for your conditions.
number = 5 # Integer

# now here is a bunch of different ways to compare things!
if number == 5: # Notice the two "=" signs. "=" means assignment, "==" means comparison.
    print("the number equals 5!")
if number != 5:
    print("the number does not equal 5!")
if number < 5:
    print("the number is greater than 5!")
if number <= 5:
    print("the number is greater than or equal to 5!")
if number > 5:
    print("the number is less than 5!")
if number >= 5:
    print("the number is less than or equal to 5!")

# you can even use "and" and "or" to make more complex decisions!
if number >= 3 and number <= 6:
    print("the number lies between 3 and 6!")
if number == 5 or number == 8:
    print("the number is either 5 or 8!")

# the next part of an IF statement is the ELSE block.
# it is entirely optional and it looks like this.

# if condition:
#     statements if true
# else:
#     statements if false

favorite_flavor = "vanilla" # String

if favorite_flavor == "vanilla":
    print("I like vanilla too!")
else:
    print("You are objectively wrong. Vanilla is the greatest!")

# Our next item in the flow control adventure is the While loop.
# It looks like this.

# while condition:
#     statements to repeat if true
#     statements to repeat if true
#     ...

# here is an example.
number_of_times_printed = 1 # Integer
while number_of_times_printed <= 5:
    print(f"I printed {number_of_times_printed} times!")
    number_of_times_printed = number_of_times_printed + 1
    # I can also use "number_of_times_printed += 1" if I want the line to be a bit shorter.
    
# look out for infinite loops in your code, as it can cause nasty issues for you
# if you do not intend for it to be stuck somewhere forever.

# for example, this one causes an infinite loop that can not be broken out of.
while number_of_times_printed <= 6:
    # this happens because number_of_times_printed is ever incremented.
    # to halt execution of any program in the terminal, select the terminal and press
    #   CTRL + C
    # to cancel execution of code
    pass

print("Note how this is never, EVER printed!")