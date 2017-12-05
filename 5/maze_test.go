package main

import "testing"
import "os"
import "io/ioutil"
import "log"

func TestMazePart1WithTestInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/test_input")
	if err != nil {
		log.Fatal(err)
	}
	Result := maze_part1(buf)
	Want := 5
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestMazePart1WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	Result := maze_part1(buf)
	Want := 355965
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestMazePart2WithTestInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/test_input")
	if err != nil {
		log.Fatal(err)
	}
	Result := maze_part2(buf)
	Want := 10
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestMazePart2WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	Result := maze_part2(buf)
	Want := 26948068
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}
