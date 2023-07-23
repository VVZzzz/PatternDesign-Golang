package builder

import (
	"fmt"
	"testing"
)

func TestBuilderX(t *testing.T) {
	builderx := &BuilderX{}
	director := NewDirector(builderx)
	director.Construct()
	fmt.Println(director.GetResult())
}

func TestBuilderY(t *testing.T) {
	buildery := &BuilderY{}
	director := NewDirector(buildery)
	director.Construct()
	fmt.Println(director.GetResult())
}
