package actioninfo

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type failingParser struct {
	current string
	infos   int
}

func (p *failingParser) Parse(data string) error {
	p.current = data
	if data == "bad parse" {
		return errors.New("invalid record")
	}
	return nil
}
func (p *failingParser) ActionInfo() (string, error) {
	p.infos++
	if p.current == "bad info" {
		return "discard this", errors.New("invalid activity")
	}
	return p.current, nil
}

func TestInfoContinuesAfterEachError(t *testing.T) {
	oldOut, oldLog := os.Stdout, log.Writer()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	defer func() { os.Stdout = oldOut; log.SetOutput(oldLog) }()
	os.Stdout = writer
	var messages strings.Builder
	log.SetOutput(&messages)
	parser := &failingParser{}
	Info([]string{"bad parse", "first", "bad info", "last"}, parser)
	writer.Close()
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	if string(output) != "first\nlast\n" {
		t.Fatalf("unexpected output: %q", output)
	}
	if parser.infos != 3 {
		t.Fatalf("ActionInfo called %d times, want 3", parser.infos)
	}
	if !strings.Contains(messages.String(), "invalid record") || !strings.Contains(messages.String(), "invalid activity") {
		t.Fatalf("missing logged error: %s", &messages)
	}
}
