package romankana_test

import (
	"fmt"
	"github.com/koyachi/go-romankana"
)

func ExampleKanaRoman_hiragana() {
	fmt.Println(romankana.KanaRoman("ほげほげ"))
	// Output:
	// hogehoge
}

func ExampleKanaRoman_katakana() {
	fmt.Println(romankana.KanaRoman("ホゲホゲ"))
	// Output:
	// hogehoge
}

func ExampleRomanKana() {
	fmt.Println(romankana.RomanKana("hogehoge"))
	// Output:
	// ホゲホゲ
}
