package math_test

import (
	"testing"
	"time"

	"golang.source-fellows.com/samples/unittest/math"
)

func TestAdd(t *testing.T) {
	res := math.Add(1, 2)
	if res != 3 {
		t.Fatalf("Not the expected result %v", res)
	}

}

var testdata = []struct {
	name   string
	param1 int
	param2 int
	result int
}{
	{"positiv1", 1, 2, 3},
	{"positiv2", 2, 3, 5},
	{"negativ1", -2, 3, 5}, //stimmt nicht! - zur Demo von Testfehlern
	{"negativ2", 2, -3, 7}, //stimmt nicht! - zur Demo von Testfehlern
}

func TestTableAdd(t *testing.T) {
	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			t.Parallel()
			result := math.Add(tt.param1, tt.param2)
			time.Sleep(time.Second * 1)
			if result != tt.result {
				t.Errorf("%v not expected. Expected: %v", result, tt.result)
			}
		})
	}
}
