package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
)

// Creating metrics
var (
	MLatencyMS = stats.Float64("repl/latency", "Latency in Mili-second per loop", "ms")

	MLineLenght = stats.Int64("repl/line_length", "The distribution of line lengths", "BY")
)

// creating tage
var (
	KeyMethod, _ = tag.NewKey("method")
	KeyStatus, _ = tag.NewKey("status")
	KeyError, _  = tag.NewKey("error")
)

func main() {
	br := bufio.NewReader(os.Stdin)

	for {
		if err := evaluateProcess(br); err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal("Error generate", err)
		}
	}
}

func evaluateProcess(br *bufio.Reader) error {
	fmt.Printf("> ")
	line, _, err := br.ReadLine()
	if err != nil {
		return err
	}

	out, err := doSomething(line)
	if err != nil {
		return err
	}

	fmt.Printf("< %s\n\n", out)
	return nil
}

func doSomething(input []byte) (out []byte, err error) {
	return bytes.ToUpper(input), nil
}
