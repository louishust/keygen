package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"math/rand"
	"os"
	"time"
)

const VERSION = "1.0.0"

var options struct {
	Version bool `short:"v" long:"version" description:"Print version"`
	Length  int  `short:"l" long:"length" description:"key length" default:"16"`
	Mode    int  `short:"m" long:"mode" description:"key mode 0: alphanum 1: alphanum+special char" default:"0"`
}

var num = [...]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var lowChars [26]rune
var upChars [26]rune
var specialChars = [...]rune{'(', ')', '#', '$', '!', '@'}
var myseed int64 = time.Now().Unix()

const ( // iota is reset to 0
	c0 = iota // c0 == 0
	c1
	c2
	c3
)

func exitWithMessage(message string) {
	fmt.Println("Error:", message)
	os.Exit(1)
}

func initOptions() {
	_, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		exitWithMessage(err.Error())
		os.Exit(1)
	}

	if options.Version {
		fmt.Printf("keygen v%s\n", VERSION)
		os.Exit(0)
	}
}

func initArray() {
	var i rune
	for i = 'a'; i <= 'z'; i++ {
		lowChars[i-'a'] = i
	}
	for i = 'A'; i <= 'Z'; i++ {
		upChars[i-'A'] = i
	}
}

func random(min, max int) int {
	rand.Seed(myseed)
	myseed = myseed + int64(max)
	return rand.Intn(max-min+1) + min
}

func getChar(c int) rune {
	switch c {
	case c0:
		return num[random(0, len(num)-1)]
	case c1:
		return lowChars[random(0, len(lowChars)-1)]
	case c2:
		return upChars[random(0, len(upChars)-1)]
	case c3:
		return specialChars[random(0, len(specialChars)-1)]
	default:
		fmt.Println("Error getChar")
		os.Exit(1)
	}

	return '.'
}

func keygen() {
	len := options.Length
	mode := options.Mode
	fmt.Printf("%c", getChar(random(c1, c2)))
	last := c2
	if mode == 1 {
		last = c3
	}
	for i := 1; i < len; i++ {
		fmt.Printf("%c", getChar(random(c0, last)))
	}
	fmt.Println()

}

func main() {
	initOptions()
	initArray()
	keygen()
}
