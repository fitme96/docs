package main

import (
	"fmt"
)

const englishHelloPrefix = "hello, "
const spanish = "Spanish"
const spanishhelloprefix = "Hola, "
const frenchhelloprefix = "Bonjour, "
const french = "French"
const china = "中文"
const chinahelloprefix = "你好, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + name
}
func greetingPrefix(language string) (prefix string) {

	switch language {
	case french:
		prefix = frenchhelloprefix
	case spanish:
		prefix = spanishhelloprefix
	case china:
		prefix = chinahelloprefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
}
