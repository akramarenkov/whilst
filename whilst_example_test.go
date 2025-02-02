package whilst_test

import (
	"fmt"
	"time"

	"github.com/akramarenkov/whilst"
)

func ExampleParse() {
	from := time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)

	whl, err := whilst.Parse("2y")
	if err != nil {
		panic(err)
	}

	fmt.Println(whl)
	fmt.Println(whl.When(from))
	fmt.Println(whl.Duration(from))
	// Output:
	// 2y
	// 2025-04-01 00:00:00 +0000 UTC
	// 17544h0m0s
}

func ExampleParse_compound() {
	whl, err := whilst.Parse("2y 3mo 10d 23.5h 59.5m 58.01003001s 10ms 30µs 10ns")
	if err != nil {
		panic(err)
	}

	fmt.Println(whl)
	// Output:
	// 2y3mo10d24h30m28.02006002s
}
