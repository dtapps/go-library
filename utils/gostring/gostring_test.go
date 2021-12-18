package gostring

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	str := "iPhone 11 Pro Max<iPhone12,5>"
	fmt.Printf("%d\n", strings.LastIndex(str, "<"))
	fmt.Printf("%d\n", strings.LastIndex(str, "("))
	fmt.Printf("%d\n", strings.LastIndex("iPad (6th generation, WiFi)<iPad7,5>", "<"))
	fmt.Printf("%d\n", strings.LastIndex("iPad (6th generation, WiFi)<iPad7,5>", "("))
	s := str[0:17]
	fmt.Printf("%s\n", s)
	str = "iPad (6th generation, WiFi)<iPad7,5>"
	s = str[0:5]
	fmt.Printf("%s\n", s)
	fmt.Printf(strings.TrimSpace(s))
}

func TestToInt64(t *testing.T) {
	log.Println(ToInt64("120"))
	log.Println(ToInt64("120.9"))
}

func TestString(t *testing.T) {
	str := "wx6566ef69e8738ad9"
	fmt.Println(strings.Contains(str, "wx"))
	myString := "www.5lmh.com"
	if strings.HasPrefix(myString, "www") {
		fmt.Println("Hello to you too")
	} else {
		fmt.Println("Goodbye")
	}
}
