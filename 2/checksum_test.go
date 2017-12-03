package main

import "testing"
import "os"
import "io/ioutil"
import "log"

func TestChecksumPart1(t *testing.T) {
	Result := checksum([]byte("5 1 9 5\n7 5 3\n2 4 6 8"))
	Want := 18
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestChecksumPart2(t *testing.T) {
	Result := checksum_part2([]byte("5 9 2 8\n9 4 7 3\n3 8 6 5"))
	Want := 9
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestChecksumPart1WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	Result := checksum(buf)
	Want := 51139
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestChecksumPart2WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	Result := checksum_part2(buf)
	Want := 272
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}
