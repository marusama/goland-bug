package mypackage_test

import (
	"fmt"
	"testing"

	"github.com/marusama/goland-bug/mypackage"
)

type Counters struct {
	Clickz int
}

func TestA(t *testing.T) {
	fmt.Println(Counters{Clickz: 1})
	fmt.Println(mypackage.Counters{Clicks: 1})
}
