package publish_subscriber

import (
	"fmt"
	"io"
	"os"
	"time"
)

type writerSubscriber struct {
	in     chan interface{}
	id     int
	Writer io.Writer
}

///
func (s *writerSubscriber) Notify(msg interface{}) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%#v", rec)
		}
	}()

	select {
	case s.in <- msg:
	case <-time.After(time.Second):
		err = fmt.Errorf("timeout\n ")
	}

	/// nothing in return cause variable err is supposed to return
	return

}
func (s *writerSubscriber) Close() {
	close(s.in)
}

/// NewWriterSubscriber accepts an ID and an io.Writer as the destination for writing the
func NewWriterSubscriber(id int, out io.Writer) Subscriber {
	if out == nil {
		out = os.Stdout
	}

	// creating a new pointer to writerSubscriber
	s := &writerSubscriber{
		id:     id,
		in:     make(chan interface{}),
		Writer: out, // out will be the input value or the value in line 25 above(out = os.Stdout)
	}

	go func() {
		// iterating over the in channel
		// if the s.in chan is closed then the for loop stop listening
		// and the current goRoutine will finish
		for msg := range s.in {
			fmt.Fprintf(s.Writer, "(W%d): %v\n", s.id, msg)
		}
	}()

	return s
}
