// Today, we will be talking about the for loop.
// Basically, a for loop is almost exactly a while loop, except that it wraps
let i = 0;      // the assignment
while (i < 10){ // the condition
    // do something
    i++;        // and the iteration
}
// all into one statement.

// the above statement would look like this!
for (let i = 0; i < 10; i++){
    // do something
}

console.log("==============SECT0===============") 
// but For loops can be even more useful than that.
// for loops take in any "itterator" which is any object that implements the "__next__" function.
// this can be anything from the range function above
for (let i = 20; i < 100; i+=10){
    console.log(i); // counts multiples of 10 from 20 to 100
}


console.log("==============SECT1===============") 
// to pulling values out of arrays of items
let arr = ["print", "hello", "world!"];
for (let index in arr){
    console.log(arr[index]);
}

console.log("==============SECT2===============")
// to pulling characters out of a string
let str = "onebyone";
for (let char_pos in str){
    console.log(str[char_pos]);
}