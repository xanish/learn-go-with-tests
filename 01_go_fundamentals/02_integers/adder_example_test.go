package _2_integers

import "fmt"

// ExampleAdd is a Testable Example. They are compiled whenever tests are executed.
// Because such examples are validated by the Go compiler, you can be confident your
// documentation's examples always reflect current code behavior. Adding this code will
// cause the example to appear in your godoc documentation, making your code even more
// accessible. If ever your code changes so that the example is no longer valid, your
// build will fail.
func ExampleAdd() {
	// Notice the special format of the comment, // Output: 6. While the example will always
	// be compiled, adding this comment means the example will also be executed. If we remove it
	// and then run go test, and you will see ExampleAdd is no longer executed.
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
