# bark
## Welcome to the Bark Programming Language!
```
   / \__
  (    @\___
  /         O
 /   (_____/
/_____/   U
```
This is the language that your doggo would code in cause he is such a goodboi (thats the true boolean in Bark btw)! Bark is written in Go and utilizes only the baked in Go functionality. Thanks to Thorsten Ball and his book "Writing an Interpreter in Go"! This has been a major help in getting this project off the ground!

Bark is a scripting language with its own REPL. It is built out with many standard programming language features and data types such as integer operations, boolean checks, string concat, array indexing, hash indexing, builtin functions, prefix operators, nested conditionals, return statements, functions, and closures.

Built in functions now available:
- len() - gets the length of a string or array
- first() - gets the first element of an array
- last() - gets the last element of an array
- rest() - returns a new array with all elements except the first
- push() - adds a new element to the end of the array (example: push(myArray, 7))
- speak() - this is Bark's "print" command. Anything put inbetween the parenthesis will be printed to the screen

Feel free to give it a clone and run by moving into the bark directory and then running: "go run main.go"

NOTE: you must have Go installed for this to work! Go here to install for your system if not already installed: https://go.dev/dl/

Basic Keywords:
 - toy: this defines our variables -- toy five = 5;
- trick: this is how we define our functions -- trick(x, y) {x + y;};
- goodboi: this is our true boolean value
- badboi: this is our false boolean value
- borkf: this is how we start an if statement -- borkf (10 > 5) { fetchit goodboit}
- woofwise: this is how we continue an if statment with other conditions (else in other languages) -- borkf (10 > 5) { fetchit goodboi;} woofwise {fetchit badboi;}
- fetchit: this is our return statement keyword


Example syntax:

    toy five = 5;
	toy ten = 10;

	toy add = trick(x, y) {
	x + y;
	};

	toy result = add(five, ten);

	borkf (5 < 10) {
		fetchit goodboi;
	} woofwise {
		fetchit badboi;
	}


## Future Ideas:
- eventually, I would like to allow users to create files ending in .bark to be run through the interpreter. For now, this language only works in the REPL. 
- I would also like to implement more builtin functions and features to more fully represent a production ready language.


## Parting Note:
This was a fun project that taught me a lot but is by no means meant to be a production level language. If you have any ideas or features you would like to add to Bark, please submit a PR and lets get Barking!
