package sqrt

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestFromCSV(t *testing.T) {

	csvfile, err := os.Open("sqrtfile.csv")

	if err != nil {
		t.Fatalf("Error while opening file: - %s", err)
	}

	defer csvfile.Close()

	readFile := csv.NewReader(csvfile)
	for {
		record, err := readFile.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			t.Fatalf("ERROR reading test cases frp, file: -%s", err)
		}

		val, err := strconv.ParseFloat(record[0], 64)

		if err != nil {
			t.Fatalf("bad value found, cannot be converted, error: -%s", err)
		}

		expected, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			t.Fatalf("bad value found, cannot be converted, error: -%s", err)
		}

		t.Run(
			fmt.Sprintf("%f", val), func(t *testing.T) {

				res, err := Sqrt(val)
				if err != nil {
					t.Fatalf("Error: %s", err)
				}

				if !almostEqual(res, expected) {
					t.Fatalf("%f != %f", res, expected)
				}

			})

	}

}

func BenchmarkFromFileCases(bcm *testing.B) {
	for i := 0; i < bcm.N; i++ {
		_, err := Sqrt(float64(i))
		if err != nil {
			bcm.Fatalf("Error Benchmark: -%s", err)
		}
	}
}
