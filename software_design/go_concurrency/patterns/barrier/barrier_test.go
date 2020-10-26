package barrier

import (
	"strings"
	"testing"
)

/// barrierTest is a tests with subTests
// if we want to run test by test then run the command in the next way:
//	$go test -run=TestBarrier/All_correct -v .
// l funcion debe iniciar con la palabra Test  y con T mayus para que golang
// lo detecte
func TestBarrier(t *testing.T) {
	t.Run("All correct", func(t *testing.T) {
		endPoints := []string{"http://httpbin.org/headers", "http://httpbin.org/headers"}
		result := captureStdout(endPoints...)
		if !strings.Contains(result, "Accept-Encoding") || !strings.Contains(result, "User-Agent") {
			t.Log("ENTRA A FALLAR")
			t.Fail()
		}
		t.Log("No falla", result)
	})
	t.Run("One call  incorrect", func(t *testing.T) {
		endPoints := []string{"http://httpbin.org/headers", "http://malformed-url"}
		result := captureStdout(endPoints...)
		if !strings.Contains(result, "Error") {
			t.Fail()
		}
		t.Log("One call  incorrect ERROR => ", result)
	})
	t.Run("Time out", func(t *testing.T) {
		endPoints := []string{"http://httpbin.org/headers", "http://httpbin.org/User-Agent"}
		timeoutMilliseconds = 1 // this is a package variable
		result := captureStdout(endPoints...)
		if !strings.Contains(result, "Timeout") {
			t.Fail()
		}
		t.Log("Time out Ok => ", result)

	})
}
