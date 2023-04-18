# To begin, a variable can be seen in two ways.
# the first way is just a container for storing data we hold onto.
# the second way is that of a name for a spot in memory that we push and pull data from. 

# it generally follows the format of 
# [name] = [value]
# like this
variable_name = 1

# the data in the variable can be of a multitude of data points... from 
# boolean values
boolean_true = True
boolean_false = False
# integers
int_a = 1
int_b = 5
int_c = 210698462510602195
# floating point numbers
float_a = 1.1
float_b = 3.14159
float_c = 21021.106
# strings
string_a = "this is a string"
string_b = "this is also a string"
string_c = "yo what the fu__"
# they can also represent extremely complex data types that can be wildly different!
list_123 = [1, 2, 3]
dictionary_KV = {"key1":"value1", "key2":"value2"}
lambda_mul_by_2 = lambda x: x*2
complex_number = 1+2j
# Although for this, in the beginning, you do not have to understand the complex variables. They are taught in lesson 5.


# data in variables can be retrieved and used in function calls like print.
print("int_a =",int_a)
print("string_c =",string_c)
# data in variables can be used in basic equations just as easily as numbers.
print("1.1 + 3.14159 =",float_a + float_b)

# and data in variables can change!
counter = 1
print(f"I have counted {counter} times")
counter = counter + 1
print(f"I have counted {counter} times")
counter = counter + 1
print(f"I have counted {counter} times")
counter = counter + 1
print(f"I have counted {counter} times")

# there is another concept to variables called the scope of a variable
# this is the area in which the variables remain alive and not freed from memory.
# for example, here we declare a function. you do not need to know what a function is.
def i_have_my_own_scope():
    variable_in_functions_scope = 1

# if i try to reference the variable outside of it's scope, i end up with an "___ is not defined" error.
# uncomment the line below and you will see it.
# print(variable_in_functions_scope)
# this is because a variable created inside a function belongs to the "local scope" of that function, and can only be used inside that function.
# we will get into scope in detail on lesson 4 when we really get into what a function is and what it is used for.