package sqrt

import (
	"fmt"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return Abs(v1-v2) <= 0.001
}

func TestSimple(t *testing.T) {
	val, err := Sqrt(2)

	if err != nil {
		t.Fatalf("error in calculation - %s", err)
	}

	if !almostEqual(val, 1.414214) {
		t.Fatalf("bad value -%f", val)
	}
}

type TestCase struct {
	value    float64
	expected float64
}

func TestMultiCases(t *testing.T) {
	testCases := []TestCase{
		{0.0, 0.0},
		{1, 1.00001},
		{1.0, 1.0},
		{2.0, 1.414214},
		{4.0, 2.0},
		{25.0, 5.0},
	}

	for _, testCase := range testCases {
		t.Run(
			fmt.Sprintf("for Value V=%f", testCase.value), func(t *testing.T) {

				res, err := Sqrt(testCase.value)
				if err != nil {
					t.Fatalf("Error: %s", err)
				}

				if !almostEqual(res, testCase.expected) {
					t.Fatalf("%f != %f", res, testCase.expected)
				}
				//fmt.Printf("Value %f and expected %f are equal", res, testCase.expected)

			})
	}

}

func BenchmarkSqrt(bcm *testing.B) {
	for i := 0; i < bcm.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			bcm.Fatal(err)
		}
	}
}
