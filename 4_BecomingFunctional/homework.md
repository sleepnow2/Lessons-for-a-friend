You are someone who is tasked with finding zeros for any arbitrary function, so in order to make your life easier, you split your problem into many smaller sub problems. Some to help shake the rust off, and others to get going towards your goal.

While doing your preliminary research, you remember someone talking about Newtons method for finding zeros and you decide that it would be a pretty decent idea.

# Problem 1: Rust-be-gone!
In order to remember how to make a function again, you have decided that you need to make a function that evaluates one particular quadratic function at any point x.

You remember that quadratic functions are of the form 
$$a*x^2 + b*x + c.$$ or of the form 
$$(x+root1)*(x+root2)$$a, b, and c or root1 and root2 can be hard coded, but your function needs to take in one number and output one number.

_hint: Don't over think it. You can get this done in 2-3 lines_
_hint 2: Use the second form. It will let you check your work later_

**Extra Credit**: Make a function that takes in the values a, b, and c or root1 and root2. This function should return functions that satisfy problem 1.

# Problem 2: Deriving Derivatives
In your studies, you remember that newtons method required the derivative, or the slope, of a line at any one point.
You also remember from your calculus class how to estimate derivatives.

$$f'(x) = \frac{f(x + h) - f(x)}h$$ (for small values of h.)
f'(x) just stands for the derivative of the function f at point x. Think of it like the slope of the line at any one point.

![image to display what I mean](https://koki0702.github.io/dezero-book/images/1-5.png)

Your task is to write a function that takes two parameters.
* The first parameter is another function that has one input number and returns one output number.
* The second parameter is x, the point at which you are evaluating.
* The third parameter is a number that will represent h.
* This function is then to return the estimated derivative of your function at point x

_hint: Don't over think it._

**Extra credit:** Have the function take in just the func and h and return a function that evaluates the derivative.

This returned function should have take in one number and return one number.

# Problem 3: Finding Zeros
With these two solved, you have all the tools to find out newtons method.

You remember that Newtons method is an iterative function, so you figure that 50 or so iterations are good enough for most applications, but you want to keep it general for future use.

The formua is as follows:$$x_{n+1} = x_n - \frac{f(x_n)}{f'(x_n)}$$where $f'$ is the derivative of the equasion $f$.
![Picture to describe what I mean](https://tutorial.math.lamar.edu/classes/calci/NewtonsMethod_Files/image001.png)
You realize that $x_0$ can be any guess you wish.

* Make a function that takes in a function $f$, an initial guess $x_0$, and a small number h, and an integer itts for the number of times you need to iterate.
* This function should return a single number, which is where it ends up after taking newtons method.

_hint: We built all of the pieces in the previous two problems, now it is just linking them together._

**Extra Credit:** See if you can make h and itts optional.
This is not possible in GO.