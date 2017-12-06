package main

import "testing"
import "os"
import "io/ioutil"
import "log"
import "bytes"
import "strconv"

func TestMemPart1WithTestInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/test_input")
	if err != nil {
		log.Fatal(err)
	}
	list := []int{}
	for _, row := range bytes.Split(buf, []byte("\t")) {
		i, _ := strconv.Atoi(string(row))
		list = append(list, i)
	}
	Result, _ := mem(list)
	Want := 5
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestMemPart1WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	list := []int{}
	for _, row := range bytes.Split(buf, []byte("\t")) {
		i, _ := strconv.Atoi(string(row))
		list = append(list, i)
	}
	Result, _ := mem(list)
	Want := 14029
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestMemPart2WithTestInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/test_input")
	if err != nil {
		log.Fatal(err)
	}
	list := []int{}
	for _, row := range bytes.Split(buf, []byte("\t")) {
		i, _ := strconv.Atoi(string(row))
		list = append(list, i)
	}
	_, tmp := mem(list)
	Result, _ := mem(tmp)
	Result--
	Want := 4
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestMemPart2WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	list := []int{}
	for _, row := range bytes.Split(buf, []byte("\t")) {
		i, _ := strconv.Atoi(string(row))
		list = append(list, i)
	}
	_, tmp := mem(list)
	Result, _ := mem(tmp)
	Result--
	Want := 2765
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}
