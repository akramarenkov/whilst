package whilst_test

import (
	"fmt"
	"time"

	"github.com/akramarenkov/whilst"
)

func ExampleParse() {
	whl, err := whilst.Parse("2y3mo10d23h59m58s10ms30µs10ns")
	if err != nil {
		panic(err)
	}

	fmt.Println(whl)
	fmt.Println(whl.When(time.Date(2023, time.April, 1, 0, 0, 0, 0, time.UTC)))
	// Output:
	// 2y3mo10d23h59m58.01003001s
	// 2025-07-11 23:59:58.01003001 +0000 UTC
}
