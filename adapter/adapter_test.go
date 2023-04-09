package adapter

import (
	"fmt"
	"testing"
)

func TestAdapter(t *testing.T) {

	adaptee := NewAdaptee()
	adapteeRes := adaptee.SpecificRequest()

	target := NewAdapter(adaptee)
	targetRes := target.Request()

	fmt.Printf("%s\n%s", adapteeRes, targetRes)
}
