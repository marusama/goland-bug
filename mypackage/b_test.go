package mypackage_test

import (
	"fmt"
	"testing"

	"github.com/marusama/goland-bug/mypackage"
)

func TestB(t *testing.T) {
	fmt.Println(mypackage.Counters{Clicks: 2})
}
