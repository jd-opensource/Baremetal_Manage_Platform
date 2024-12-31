package util 

import(
	"testing"
)

func TestInArray(t *testing.T){
	v := InArray("aaa", []string{"aaa", "bbb"})
	if !v{
		t.Fatal("not in array")
	}else{
		t.Log("in array")
	}
}
