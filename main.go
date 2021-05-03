package integers

import "fmt"

const english = "English"
const spanish = "Spanish"
const french = "French"
const helloPrefixEng = "Hello, "
const helloPrefixSpan = "Hola, "
const helloPrefixFren = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return GreetingPrefix(language) + name
}

func GreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpan
	case french:
		prefix = helloPrefixFren
	default:
		prefix = helloPrefixEng
	}
	return
}

func main() {
	fmt.Println(Hello("World", english))
}
