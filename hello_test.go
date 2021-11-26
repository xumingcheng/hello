package hello

import "testing"

func HelloTest(t *testing.T)  {
	// Return a greeting that embeds the name in a message.
	want:="hello,world"
	if got:=Hello();got!=want{
		t.Errorf("Hello()=%q,want %q",got,want)
	}

}