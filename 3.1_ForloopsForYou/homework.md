## Problem 1: Quadrant selection
A common problem in math is to select which quadrant a given point lies in. There are four quadrants, numbered from 1 to 4, as shown in the diagram below:
![](https://open.kattis.com/problems/quadrant/file/statement/en/img-0001.png)
For example, Point $A$ lies in quadrant 1 since both its $x$ and $y$ values are positive and point $B$ lies in quadrant 2 since its x value is negative and its y value is positive.

INPUT:
you will be given a line of text that contains two non zero integers. The first integer is the $x$ coordinate and the second integer is the $y$ coordinate of the point.

OUTPUT:
you will output the quadrant number (1,2,3,4)

EX:
10 25 -> 1
-2 -3 -> 3
-1 54 -> 2

_Hint: If statements_
**Extra Credit** Complete this assignment while only using 2 if statements.

## Problem 2: FizzBuzz
FizzBuzz is a common problem to torture young graduates in interviews. This is a very simple problem, but it can tell you a lot about someone by how they approach and optimize this problem.

Basically, you print the numbers from 1 to $n$, replacing any of them divisible by $x$ with "Fizz" and any divisible by $y$ with Buzz.
If a number is divisible by both $x$ and $y$, then you print FizzBuzz.

INPUT:
one line that contains integers $x$, $y$, and $n$. where $(1\le x<y<n<1000)$.

OUTPUT:
Print numbers from 1 to $n$ in order, each on their own line, replacing those that are supposed to be replaced following the rules of fizzbuzz.

EX:
2 3 10
1
fizz
buzz
fizz
5
fizzbuzz
7
fizz
buzz
fizz

_hint: modulo. look it up._