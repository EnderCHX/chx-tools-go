package codec

import (
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	data := "hello world"
	encoded := Base64Encode(data)
	decoded, err := Base64Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	if decoded != data {
		t.Error("decoded data is not equal to original data")
	}
	fmt.Println(encoded)
	fmt.Println(decoded)
}
