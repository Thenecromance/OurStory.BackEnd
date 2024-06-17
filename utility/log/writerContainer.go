package log

import "io"

type writerContainer struct {
	containers []io.Writer
}

func (wc *writerContainer) addWriter(w io.Writer) {
	wc.containers = append(wc.containers, w)
}

func (wc *writerContainer) Write(p []byte) (n int, err error) {
	for _, w := range wc.containers {
		n, err = w.Write(p)
		if err != nil {
			return
		}
	}
	return
}

func newWriterContainer() *writerContainer {
	return &writerContainer{
		containers: make([]io.Writer, 0),
	}
}
