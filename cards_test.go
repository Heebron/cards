package cards

import (
	"fmt"
	"testing"
)

func TestCard_String(t *testing.T) {
	d := NewDeck()
	d.Shuffle(0)

	for {
		c, last := d.Next()
		if last == false {
			fmt.Print(c)
		} else {
			fmt.Println()
			if !d.IsExhausted() {
				t.Fail()
			}
			break
		}
	}
}
