package objects_test

import (
	"reflect"
	"testing"
)

// eq compares the two given interfaces with reflect.DeepEqual.
func eq(t *testing.T, e, g interface{}) {
	if !reflect.DeepEqual(e, g) {
		t.Errorf("expected\n%#v\ngot\n%#v", e, g)
	}
}

// eqBool compares the two given boolean values.
func eqBool(t *testing.T, e, g bool) {
	if e != g {
		t.Errorf("expected %t, got %t", e, g)
	}
}

// eqInt compares the two given integers.
func eqInt(t *testing.T, e, g int) {
	if e != g {
		t.Errorf("expected %d, got %d", e, g)
	}
}

// eqStr compares the two given strings.
func eqStr(t *testing.T, e, g string) {
	if e != g {
		t.Errorf("expected %s, got %s", e, g)
	}
}
