package main

import (
	"fmt"
)
//Don’t use floating point numbers for money.
/*
So you’ve written some absurdly simple code, say for example:

	0.1 + 0.2

and got a really unexpected result:

	0.30000000000000004

Maybe you asked for help on some forum and got pointed to a long article with lots of formulas that didn’t seem to help with your problem.

Well, this site is here to:

    Explain concisely why you get that unexpected result
    Tell you how to deal with this problem
    If you’re interested, provide in-depth explanations of why floating-point numbers have to work like that and what other problems can arise

You should look at the Basic Answers first - but don’t stop there!
*/
func main() {
	var x float32
	x = 80917.65
	fmt.Println(int(x * 100))

	var y float64
	y = 80917.65
	fmt.Println(int(y * 100))
}
