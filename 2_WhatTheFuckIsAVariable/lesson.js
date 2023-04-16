var a = 1;
let b = 1.2;
let c = "yo what the fu___";

let scope1 = 1;
if (true) {
    for(let i = 1;;){
        // does stuff
    }

    for(let i = 2;;){
        // does other stuff
    }
}
console.log(scope1);
console.log(scope2);





console.log(a);
a = 2;
console.log(a);

console.log(b);
console.log(c);