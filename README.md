# bark
## Welcome to the Bark Programming Language!

This is the language that your doggo would code in cause he is such a goodboi (thats the true boolean in Bark btw)!

Bark is still a work in progress and as of this latest update to this repo, it is only a parser of the basic syntax and a basic repl. More is coming soon though!

Basic Syntax:
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

	!-/*5;
	5 < 10 > 5;
	borkf (5 < 10) {
		fetchit goodboi;
	} woofwise {
		fetchit badboi;
	}
