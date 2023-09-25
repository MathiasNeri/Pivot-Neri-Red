package red

import (
	"fmt"
	"time"
)

func Defil(s string) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}
}

func DefilAS(s string, c *Character) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}
	for _, letter := range c.classe {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}
}

func DefilLeft(s string, m *Monstre, s2 string) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}

	fmt.Print(m.curpv)

	for _, letter := range s2 {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}
}

func DefilIntro(s string, c *Character, s2 string) {
	for _, letter := range s {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}

	fmt.Print(c.nickname)

	for _, letter := range s2 {
		fmt.Print(string(letter))
		time.Sleep(30 * time.Millisecond)
	}
}
