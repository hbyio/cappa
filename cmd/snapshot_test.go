package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_ExecuteSnapshotNoConfig(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"snapshot"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	subString := "Not Found in"
	if !strings.Contains(string(out), subString) {
		t.Fatalf("expected output to contain \"%s\" in \"%s\"", subString, string(out))
	}
}

func Test_ExecuteSnapshotEmptyConfig(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"snapshot"})
	fakeconf := fakeConfig{}
	_ = fakeconf.create()
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fakeconf.remove()
	subString := "Some values are missing or are incorrect in your config file (run 'cappa init')"
	if !strings.Contains(string(out), subString) {
		t.Fatalf("expected output to contain \"%s\" in \"%s\"", subString, string(out))
	}
}

func Test_ExecuteSnapshot(t *testing.T) {
	cmd := NewRootCmd()
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{"snapshot"})
	fakeconf := fakeConfig{}
	_ = fakeconf.create()
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	fakeconf.remove()
	subString := "Some values are missing or are incorrect in your config file (run 'cappa init')"
	if !strings.Contains(string(out), subString) {
		t.Fatalf("expected output to contain \"%s\" in \"%s\"", subString, string(out))
	}
}
