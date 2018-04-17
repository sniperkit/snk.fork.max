package config

import (
	"io/ioutil"
	"testing"
)

func TestReadContent(t *testing.T) {
	buf, err := ioutil.ReadFile("./test.yml")

	if err != nil {
		t.Errorf("Expected: nil, got: %v", err)
	}

	c, err := ReadContent(string(buf))

	if err != nil {
		t.Errorf("Expected: nil, got: %v", err)
	}

	if c == nil {
		t.Errorf("Expected: struct, got: %v", c)
	}

	if c.Tasks["hello"].Summary != "Hello task" {
		t.Errorf("Expected: 'Hello task', got: %s", c.Tasks["Hello"])
	}
}

func TestReadFile(t *testing.T) {
	c, err := ReadFile("")

	if err == nil {
		t.Errorf("Expected: error, got: %v", err)
	}

	if c != nil {
		t.Errorf("Expected: nil, got: %v", c)
	}

	c, err = ReadFile("./test.yml")

	if err != nil {
		t.Errorf("Expected: nil, got: %v", err)
	}

	if c == nil {
		t.Errorf("Expected: struct, got: %v", c)
	}

	if c.Tasks["hello"].Summary != "Hello task" {
		t.Errorf("Expected: 'Hello task', got: %s", c.Tasks["Hello"])
	}
}
