package main

import "testing"
import "os"
import "io/ioutil"
import "log"

func TestPassphrasePart1WithTestInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/test_input")
	if err != nil {
		log.Fatal(err)
	}
	Result := passphrase(buf)
	Want := 2
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestPassphrasePart1WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	Result := passphrase(buf)
	Want := 383
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestPassphrasePart2WithTestInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/test_input_part2")
	if err != nil {
		log.Fatal(err)
	}
	Result := passphrase_anagram(buf)
	Want := 3
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestPassphrasePart2WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	Result := passphrase_anagram(buf)
	Want := 265
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}
