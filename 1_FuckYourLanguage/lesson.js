const prompt = require('prompt-sync')({sigint: true});

// ask him for details
let loopTimes = parseInt(prompt("how many times should i say 'Hello world!: "));

// do the loop!
for (let i = 0; i < loopTimes; i++) {
    console.log('Hello world!');
}