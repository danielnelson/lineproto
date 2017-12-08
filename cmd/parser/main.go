package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/danielnelson/lineproto"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Must have 1 argument")
		return
	}
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// stats := lineproto.Stats{}
	// _, err = lineproto.Parse("lineproto", input, lineproto.Statistics(&stats, "no match"))
	res, err := lineproto.Parse("lineproto", input)
	if err != nil {
		fmt.Println(err)
		return
	}
	// b, err := json.MarshalIndent(stats.ChoiceAltCnt, "", "  ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("stats: ", string(b))

	fmt.Printf("%T %q\n", res, res)
}
