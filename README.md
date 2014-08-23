# go-romankana

Roman Alphabet <-> Japanese Hiragana/Katkakana Convert Library for Go.

(Go porting from https://github.com/ymrl/romankana)

## Example

```go
package main

import (
	"fmt"
	"github.com/koyachi/go-romankana"
)

func main() {
	fmt.Println(romankana.KanaRoman("ほげほげ"))
	// Output:
	// hogehoge

	fmt.Println(romankana.KanaRoman("ホゲホゲ"))
	// Output:
	// hogehoge

	fmt.Println(romankana.RomanKana("hogehoge"))
	// Output:
	// ホゲホゲ
}
```

## Links

- https://github.com/ymrl/romankana
