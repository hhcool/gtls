package log

import (
	"bufio"
	"io"
	"runtime"
)

// SafeWriter
// @Description: 获取一个io.Writer，方便集成
// @return *io.PipeWriter
func SafeWriter() *io.PipeWriter {
	reader, writer := io.Pipe()
	go scan(reader)
	runtime.SetFinalizer(writer, writerFinalizer)
	return writer
}
func scan(reader *io.PipeReader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		Info(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		Errorf("Error while reading from Writer: %s", err)
	}
	_ = reader.Close()
}
func writerFinalizer(writer *io.PipeWriter) {
	_ = writer.Close()
}
