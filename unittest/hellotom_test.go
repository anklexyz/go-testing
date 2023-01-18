package unittest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestHelloTom(t *testing.T) {
//	output := HelloTom()
//	expectOutput := "Tom"
//	if output != expectOutput {
//		t.Errorf("Expected %s do not match actual %s", expectOutput, output)
//	}
//}

// 直接用assert
// 要以Test开头
func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
}
