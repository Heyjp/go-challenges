package greet

import (
	"fmt"
)

func Hi(name string) {
	fmt.Printf("Hello, %s, nice to meet you!\n", name)
}

func Salut(name string) {
	fmt.Printf("Salut, %v, ravis de faire votre connaissance!\n", name)
}

/*
func OldGreet() {
	m := make(map[string]string)

	m["en_Us"] = "Hello"
	m["fr"] = "Salut"

	lang := os.Getenv("LANG")
	log.Println(lang)
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Greetings! What is your name?")
	for reader.Scan() {
		text := strings.Replace(reader.Text(), " ", "", -1)
		text = strings.Replace(text, "	", "", -1)
		if len(reader.Text()) == 0 || len(text) == 0 {
			fmt.Println("Sorry, didn't catch that")
			continue
		}

		greeting := m[lang]
		fmt.Printf("%s, %s! It's good to meet you\n", greeting, reader.Text())
		break
	}
}
*/
