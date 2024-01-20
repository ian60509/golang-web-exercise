package math

import (
	"errors"
	"testing"
)

func TestIsNegative(t *testing.T)  {
	err := errors.New("This is an error")
	if isNegative(-1) == true {
		t.Log("OK")
	} else{
		t.Error(err)
	}
	return 
}