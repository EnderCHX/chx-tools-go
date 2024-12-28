package encrypt

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	testString := "Hello, world!"

	fmt.Println(Sha1(testString))
	fmt.Println(Sha256(testString))
	fmt.Println(Sha512(testString))
}
