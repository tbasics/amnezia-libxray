package libXray

import (
	"bufio"
	"io"
	"log"
	"os"
	"syscall"

	"github.com/amnezia-vpn/amnezia-libxray/nodep"
)

type Logger interface {
	io.Writer
	Warning(s string)
	Error(s string)
}

var (
	// Store the writer end of the redirected stderr and stdout
	// so that they are not garbage collected and closed.
	stderr, stdout *os.File
)

func writeLog(stream *os.File, log func(s string), logError func(s string)) {
	const logSize = 1024
	r := bufio.NewReaderSize(stream, logSize)
	for {
		line, _, err := r.ReadLine()
		str := string(line)
		if err != nil {
			str += " " + err.Error()
			logError(str)
			break
		} else {
			log(str)
		}
	}
}

func logStream(stream uintptr, log func(s string), logError func(s string)) (*os.File, error) {
	r, w, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	if err := syscall.Dup3(int(w.Fd()), int(stream), 0); err != nil {
		return nil, err
	}
	go writeLog(r, log, logError)
	return w, err
}

func InitLogger(logger Logger) string {
	log.SetOutput(logger)
	log.SetFlags(0)

	stream, err := logStream(os.Stderr.Fd(), logger.Error, logger.Error)
	if err != nil {
		return nodep.WrapError(err)
	}
	stderr = stream

	stream, err = logStream(os.Stdout.Fd(), logger.Warning, logger.Error)
	if err != nil {
		return nodep.WrapError(err)
	}
	stdout = stream

	return ""
}
