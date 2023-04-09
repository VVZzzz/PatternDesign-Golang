package facade

import (
	"fmt"
	"testing"
)

func TestFacadeAPI(t *testing.T) {
	facadeApi := NewDoAPI()
	ret := facadeApi.DoFunc()
	fmt.Println(ret)
}
