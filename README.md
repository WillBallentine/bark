# bark
## Welcome to the Bark Programming Language!

This is the language that your doggo would code in cause he is such a goodboi (thats the true boolean in Bark btw)! Bark is written in Go and utilizes only the baked in Go functionality. Thanks to Thorsten Ball and his book "Writing an Interpreter in Go"! This has been a major help in getting this project off the ground!

Bark is still a work in progress and as of this latest update to this repo, it is only generating tokens with the basic syntax and a basic repl. More is coming soon though!

Feel free to give it a clone and run by moving into the bark directory and then running: go run main.go
NOTE: you must have Go installed for this to work! Go here to install for your system if not already installed: https://go.dev/dl/

Basic Keywords:
 - toy: this defines our variables -- toy five = 5;
- trick: this is how we define our functions -- trick(x, y) {x + y;};
- goodboi: this is our true boolean value
- badboi: this is our false boolean value
- borkf: this is how we start an if statement -- borkf (10 > 5) { fetchit goodboit}
- woofwise: this is how we continue an if statment with other conditions (else in other languages) -- borkf (10 > 5) { fetchit goodboit;} woofwise {fetchit badboi;}
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
