package main

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/diegolikescode/adventofcode/y2015"
	"github.com/diegolikescode/adventofcode/y2016"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("plzz send dudes (I mean the date as argument like this: yyyy/dd[A-B] e.g. 2015/05B)")
		return
	}

	myRegex := regexp.MustCompile(`^[0-9]{4}/[0-9]{2}[A-B]{1}$`)
	didMatch := myRegex.MatchString(os.Args[1])
	if !didMatch {
		fmt.Println("plzz use yyyy/dd[A-B] format e.g. 2015/05B")
		return
	}
	// y2015 := y2015.SharedType{}

	date := strings.Split(os.Args[1], "/")
	callIt := "Day" + string(date[1][:len(date[1])-1]) + "_" + string(date[1][len(date[1])-1])
	fmt.Println("CALLING THIS", callIt)

	var y any
	switch date[0]{
	case "2015":
		y= y2015.SharedType{}

	case "2016":
		y = y2016.SharedType{}
	}

	inputTxtPath := "./y" + date[0] + "/" + "/input" + date[1][:len(date[1])-1] + ".txt"
	reflect.ValueOf(y).MethodByName(callIt).Call([]reflect.Value{reflect.ValueOf(inputTxtPath)})
}
