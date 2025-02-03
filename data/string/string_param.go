package main

func foo(str string) {
	println("foo => " + str)
	str = "hello-golang"
	println("foo => " + str)
}

func main() {
	var a = "hello"
	println("main => " + a)
	foo(a)
	println("main => " + a)
}
