package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {

	dat, err := ioutil.ReadFile("myfile.txt")
	check(err)
	fmt.Println(string(dat))

	fmt.Println("After replacements")

	// better way of replacement (all at once)
	// Note : use Replacer.WriteString() to write this to file (see docs)
	r := strings.NewReplacer("[3]", "[0]", "[4]", "[1]", "[1]", "[2]", "[2]", "[3]", "[0]", "[4]")
	str := r.Replace(string(dat))
	fmt.Print(str)


}
