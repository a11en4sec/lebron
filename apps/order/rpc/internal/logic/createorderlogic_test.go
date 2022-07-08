package logic

import (
	"fmt"
	"testing"
	"time"
)

func TestGenOrderID(t *testing.T) {
	oid := genOrderID(time.Now())
	fmt.Println("oid:",oid)
	if len(oid) != 24 {
		t.Failed()
	} else {
		t.Log(oid)
	}
}