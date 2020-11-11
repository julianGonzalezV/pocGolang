package publish_subscriber

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

type mockWriter struct {
	/// the input string represent the bytes written to this struct
	testingFunc func(string)
}

/// Write function is required to implement io.Writer (param in NewWriterSubscriber in file writer_subsc_test.go )
func (m *mockWriter) Write(p []byte) (n int, err error) {
	fmt.Println("io.Writer mock implementation => ", p, " => ", string(p))
	m.testingFunc(string(p))
	return len(p), nil
}

func TestStdoutPrinter(t *testing.T) {
}

func TestWriter(t *testing.T) {
	sub := NewWriterSubscriber(0, nil)

	msg := "publishing message FROM  TestWriter"
	var wg sync.WaitGroup
	wg.Add(1)

	stdoutPrinter := sub.(*writerSubscriber) // recommended type assertion
	/// la linea anterior es para asegurar que la interface devuelta en la linea 30 (NewWriterSubscriber)
	/// está correctamente implementada , en este caso la impl es la writerSubscriber

	stdoutPrinter.Writer = &mockWriter{

		testingFunc: func(res string) {
			fmt.Println("Msg received in testingFunc", res, msg)
			if !strings.Contains(res, msg) {
				t.Fatal(fmt.Errorf("Bad string: %s", res))
			}
			wg.Done()
		},
	}

	err := sub.Notify(msg)
	if err != nil {
		wg.Done()
		fmt.Println("	t.Fatal(err)")
		t.Error(err)
	}
	wg.Wait()
	sub.Close()

}
