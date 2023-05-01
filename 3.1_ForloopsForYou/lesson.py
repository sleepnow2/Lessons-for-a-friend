# Today, we will be talking about the for loop.
# Basically, a for loop is almost exactly a while loop, except that it wraps
i = 0           # the assignment
while i < 10:   # the condition
    # do something
    i = i + 1   # and the iteration
# all into one statement.

# the above statement would look like this!
for i in range(10):
    pass # do something

print("==============SECT0===============") 
# but For loops can be even more useful than that.
# for loops take in any "itterator" which is any object that implements the "__next__" function.
# this can be anything from the range function above
for i in range(20,100,10):
    print(i) # counts multiples of 10 from 20 to 100

print("==============SECT1===============") 
# to pulling values out of arrays of items
for string in ["print", "hello", "world!"]:
    print(string)

print("==============SECT2===============")
# to pulling characters out of a string
for char in "onebyone":
    print(char)