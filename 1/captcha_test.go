package main

import "testing"
import "os"
import "io/ioutil"
import "log"
import "bytes"

func TestCaptchaPart1(t *testing.T) {
	Result := captcha([]byte("1122"))
	Want := 3
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha([]byte("1111"))
	Want = 4
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha([]byte("1234"))
	Want = 0
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha([]byte("91212129"))
	Want = 9
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}

}

func TestCaptchaPart2(t *testing.T) {
	Result := captcha_part2([]byte("1212"))
	Want := 6
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha_part2([]byte("1221"))
	Want = 0
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha_part2([]byte("123425"))
	Want = 4
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha_part2([]byte("123123"))
	Want = 12
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = captcha_part2([]byte("12131415"))
	Want = 4
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestCaptchaPart1WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	buf = bytes.Trim(buf, "\x0A")
	Result := captcha(buf)
	Want := 1031
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}

func TestCaptchaPart2WithInputFile(t *testing.T) {
	pwd, _ := os.Getwd()
	buf, err := ioutil.ReadFile(pwd + "/input")
	if err != nil {
		log.Fatal(err)
	}
	buf = bytes.Trim(buf, "\x0A")
	Result := captcha_part2(buf)
	Want := 1080
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}
