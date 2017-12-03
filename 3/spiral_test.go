package main

import "testing"

func TestSpiralPart1(t *testing.T) {
	Result := spiral_part1(1)
	Want := 0
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = spiral_part1(12)
	Want = 3
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = spiral_part1(23)
	Want = 2
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = spiral_part1(1024)
	Want = 31
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
	Result = spiral_part1(265149)
	Want = 438
	if Result != Want {
		t.Errorf("Result was wrong, got: %d, want: %d.", Result, Want)
	}
}
