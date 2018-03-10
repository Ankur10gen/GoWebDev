package main

import (
	"bytes"
	"fmt"
	"io"
)

type human struct {
	milesCompleted int
}

func (h *human) run() {
	h.milesCompleted++
}

type runner interface {
	run()
}

func runTwice(r runner) {
	r.run()
	r.run()
}

func main() {
	m := make(map[string]string)
	fmt.Println(m, len(m))

	arr := [4]string{"ank", "ur", "rai"}

	brr := arr[1:3]

	fmt.Println(arr, brr)

	brr[0] = "na"

	fmt.Println(arr, brr)

	veggies := [...]string{"carrot", "pea"}

	fmt.Println(veggies, len(veggies))

	var i int = 1
	var f float32
	fmt.Println(i == 1.0000000, 5%2, f, 3.14 == 3)

	ch := make(chan string, 2)
	fmt.Println(len(ch))
	var x string
	go func() {
		x = <-ch
	}()
	fmt.Println(len(ch), x)

	in := new(bytes.Buffer)
	in.WriteString("hello world")

	out := new(bytes.Buffer)
	out.ReadFrom(in)
	fmt.Println(out.String())

	out.Reset()
	fmt.Println(out.String())

	in = new(bytes.Buffer)
	in.WriteString("hello world")
	io.CopyN(out, in, 5)
	fmt.Println(out.String())

	bob := new(human)
	fmt.Println(runner(bob), bob.milesCompleted)
	runTwice(bob)
	fmt.Println(bob.milesCompleted)

	ages := map[string]int{"a": 1}
	age, ok := ages["b"]
	fmt.Println(age, ok, len(ages))

	{

		a := 1
		b := a
		b++
		fmt.Println(a, b)
	}
	{
		a := 1
		b := &a
		*b = *b + 2
		fmt.Println(a, *b, "abc"[0], "smith"[2:], "smith"[:4], "smith"[2:4], "smith"[:], "smith" == "smith", "smith" < "x")
	}

	bytes := []byte{'a', 'b', 'c'}

	fmt.Println(string(bytes))
	str := fmt.Sprintf("hello \"%s\"", "world")
	fmt.Println(str)
	str = fmt.Sprintf("hello %q", "world")
	fmt.Println(str)

	fmt.Printf("your balance: %d and %0.2f\n", 3, 4.5589)

	fmt.Println(5 ^ 2)

	fruits := [4]string{"apple", "orange", "mango"}
	tasty_fruits := fruits[1:3]
	fmt.Println(fruits == [4]string{}, tasty_fruits, cap(tasty_fruits))

	strings := []string{"hello", " world", "!"}
	var concatenated string
	var total int
	for i, v := range strings {
		total += i
		concatenated += v
	}
	fmt.Println(concatenated)
}
