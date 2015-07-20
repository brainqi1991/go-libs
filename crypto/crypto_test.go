package crypto

import (
	"testing"
)

func Test_Md5(t *testing.T) {
	str := "brain2015.."
	result := Md5(str)
	t.Log("Md5: " + result)
}

func Test_Sha1(t *testing.T) {
	str := "brain2015.."
	result := Sha1(str)
	t.Log("Sha1: " + result)
}

var base64_encode string

func Test_Base64Encode(t *testing.T) {
	str := "brain2015.."
	result := Base64EncodeString(str)
	t.Log("Base64Encode: " + result)
	base64_encode = result
}

func Test_Base64Decode(t *testing.T) {
	result, err := Base64DecodeString(base64_encode)

	if err != nil {
		t.Error("解码出错了!")
	}

	if result != "brain2015.." {
		t.Error("解码Base64不对!" + base64_encode)
	} else {
		t.Logf("解码正确: %s --> %s\n", base64_encode, result)
	}
}