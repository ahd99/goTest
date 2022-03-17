package main

func main() {

}

func test1() {
	i, j, x, y := f1(1,2,"a","b")
	f1(1,2,"a","b")	// ignore all returns
	// i, j := f1(1,2,"a","b")  // Error: ignore all or get all returns.
	i, j, x, _ = f1(1,2,"a","b")	// igonre last return value.
	f2(f1(1,2,"a", "b"))	// because f1 returns type and order are equal to f2 parameters



	_,_,_,_ = i, j, x, y

}


//multiple ways to declare a function : f1 to f6
/*
- i, j, x, y are function parameters. (values send to this parameters when calleing function, are arguments)
- r, p, s are function return values. returns can be named or unnamed. 
- parameters and anme returns are local variables. returns initialize by zero-value.
- in function with return values, using return keyword is mandatory. 

- signature: type of function. function methods includes order and type of parameters and returns.
- So two functions are same type if they have the same sequence of parameter types and same sequence of return types. 
- method name, parameters names and returns names are not part of method type.

- go has not default parameter value conceept nor any way to specify parameters by name when calling fumc. 

- arguments are passed by value, so function receive a copy of them.

- go use dynamic stack size that start from small size and groes to giga-bytes if needed. 
	so dont worry about stack overflow in method calls specially in recursive calls.
*/

func f1(i int, j int, x string, y string) (int, int, string, string) {	
	return i+j, i-j, x+y, y+x	// when using no-name returns, return must specify a value for each return
}

func f2(i, j int, x, y string) (int, int, string) {
	return f3(i, j, x, y)	//because types and order of f3 returns is equal to f2 returns
}

func f3(i, j int, x, y string) (r int, p int, s string) {
	r = i + j
	p = i - j
	s = x + y
	return 		//bare return. it is possible because returns are named
}

func f4(i, j int, x, y string) (r, p int, s string) {
	return //i+j, i-j, x+y		// when using name returns, return can be empty (bare return)
}

func f5(i, j int, x, _ string) (r, p int, s string) {  // _ emphesize that parameter it nnot used
	return i+j, i-j, x  // second string parameter can not be used because _
}

func f6(int, int, string, string) (int, int, string) {  // no name parameters can not be used
	return 0,0,""
}

func f7() (r int) {	
	return 0
}

func f8() int {	//when function return 1 no-name value, paranthesis can be ommitted
	return 0
}

func f9() {	  //function with no return value. return keyword not needed in body
	return	// it is optional to use return keyword in function without return value 
}