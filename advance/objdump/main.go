package main

func main() {
	var slice = make([]string, 100000)
	println(slice)

	var ch = make(chan int, 5)
	println(ch)

	var hash = make(map[int]string, 10)
	println(hash)
}
