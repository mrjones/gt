package gt

import (
	"testing"
)

func AssertTrueM(t *testing.T, cond bool, msg string) {
	if !cond { t.Fatal(msg) }
}
func AssertTrue(t *testing.T, cond bool) {
	AssertTrueM(t, cond, "")
}


func AssertFalseM(t *testing.T, cond bool, msg string) {
	if cond { t.Fatal(msg) }
}
func AssertFalse(t *testing.T, cond bool) {
	AssertTrueM(t, cond, "")
}

func AssertEqualM(t *testing.T, expected interface{}, actual interface{}, msg string) {
	if expected != actual {
		t.Fatalf("%s [expected: %s, actual: %s]", msg, expected, actual)
	}
}

func AssertNil(t *testing.T, o interface{}) {
	if o != nil { t.Fatalf("o is not Nil: %s", o) }
}

func AssertNotNil(t *testing.T, o interface{}) {
	if o == nil { t.Fatal("o is Nil") }
}
