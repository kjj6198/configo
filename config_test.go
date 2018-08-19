package configo

import (
	"os"
	"testing"
)

const (
	testFilename = "./test.yml"
)

func TestLoad(t *testing.T) {
	ok := Load(testFilename)

	if ok {
		t.Log("config.Load success when valid yml file.")
	} else {
		t.Error("config.Load FAIL")
	}
}

func TestLoadFailed(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("should panic when filename is empty")
		} else {
			t.Log("panic when filename is empty")
		}
	}()

	Load("")
}

func TestGetVariable(t *testing.T) {
	if os.Getenv("TEST_ENV") == "helloworld" {
		t.Log("load yml variable into ENV variable.")
	} else {
		t.Errorf("should load variable into ENV, but received `%s`", os.Getenv("TEST_ENV"))
	}
}
