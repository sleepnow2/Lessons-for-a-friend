# functions are the bread and butter of life. 
# they can do some wondrous things.
# they generally look something like this

# def name(input, input, ...):
#     things to do
#     return value
# although, the return statement is purely optional, as not every function returns something!

# and they are called by using something like this
# name(input, input, ...)
# they can be used almost anywhere a number or a variable can be used!

# they are used to do things to inputs that can be useful later, 
# or to package repeated sections of code into its own box.

def mul_by_two(number):
    return number*2
print("mul_by_two(5) =",mul_by_two(5))

def mul_two_numbers(num1, num2):
    return num1*num2
print("mul_two_numbers(6, 3) =",mul_two_numbers(6, 3))

def calc_factorial(inp):
    result = 1
    while inp > 0:
        result *= inp
        inp -= 1
    return result
print("calc_factorial(10) =",calc_factorial(10))

# functions can even call themselves, and that is called recursion.
# here is an example of a function that calculates the nth element of the fibonacci sequence. 
# (horrible performance. Never use this if you need the fibonacci sequence for anything. my computer locks up at values past 42)
def fibonacci(num):
    if num <= 2:
        return 1
    else:
        return fibonacci(num-1) + fibonacci(num-2)
print("fibonacci(5) =",fibonacci(5))

# This, unfortunately, is where we have to really talk about scope, 
# now that our programs are becoming more complicated.

# a variable that is declared in a function belongs to the "local scope" of that function.
# it can only be used inside that function.

def foo():
    x = 300
    print(x)
# if we uncomment the variable below, it will give us an error
# print(x)

# functions can also be declared inside another function. this is called nesting functions!
# you can think of it like a russian nesting doll!
def find_y_from_two_points(x, x1, y1, x2, y2):
    def find_slope(x1, y1, x2, y2):
        return (y1-y2)/(x1-x2)
    
    def find_y_intercept(x,y,m):
        return y-x*m
    
    slope = find_slope(x1, y1, x2, y2)
    intercept = find_y_intercept(x1, y1, slope)

    return x*slope+intercept
print("find_y_from_two_points(3,1,2,2,4) =",find_y_from_two_points(3,1,2,2,4))

# better yet, inner functions are able to use variables in the outer functions.
# This is because those variables are still "inside that function"
def find_y_from_two_points_v2(x, x1, y1, x2, y2):
    def find_slope():
        return (y1-y2)/(x1-x2)
    
    def find_y_intercept(m):
        return y1-x1*m
    
    slope = find_slope()
    intercept = find_y_intercept(slope)

    return x*slope+intercept
print("find_y_from_two_points_v2(3,1,2,2,4) =",find_y_from_two_points_v2(3,1,2,2,4))

# you can even pass functions as inputs to other functions...
def apply(f, val):
    return f(val)
print("apply(mul_by_two, 4) =", apply(mul_by_two, 4))

# or get functions out of functions...
def gen_linear(slope, intercept):
    def result(x):
        return slope*x + intercept
    return result
line = gen_linear(2, 2)

print("line(2) =", line(2))

# if you want to return functions with a really short hand notation, you can use lambda.
# technically it is called an "anonymous function" because it is not given a name.
# instead, it is assigned directly to a variable... or returned from a function!
def gen_linear_lambda(slope, intercept):
    return lambda x: slope*x + intercept
line_lambda = gen_linear_lambda(2, 2)

print("line_lambda(2) =", line_lambda(2))

# regarding scope and passing functions out of other functions, I am not sure how it works, 
# other than treating it as though the function you are calling never left its scope.
# there is some deeper reading regarding functions being regarded as "first class objects",
# but this is generally deeper than you will normally get anyways.