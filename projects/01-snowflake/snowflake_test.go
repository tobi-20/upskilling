package snowflake

import "testing"

func TestGenerateSnowflake(t *testing.T) {
	got := GenerateSnowflake(5, customEpoch+3, 2)
	want := int64(12603394)

	if got != want {
		t.Errorf("GenerateSnowflake(3, customEpoch+3, 2) = %d, want %d", got, want)
	}
}
func TestNoCollision(t *testing.T) {
	a := GenerateSnowflake(1, 1794508485000, 0)
	b := GenerateSnowflake(1, 1794508485000, 1)
	if a == b {
		t.Errorf("There's a collision as %d and %d are equal", a, b)
	}
}

func TestSameMillisecond_GivesDifferentIDs(t *testing.T) {
	g, err := NewGenerator(1)
	if err != nil {
		t.Fatal(err)
	}

	now := customEpoch + 500

	a, _ := g.NextID(now)
	b, _ := g.NextID(now)

	if b <= a {
		t.Errorf("second ID %d should be bigger than first %d", b, a)
	}
}

// clock going backwards, sequence resetting in a new millisecond, and NewGenerator(1024) returning an error.
func TestClockGoingBackwards_GivesAnError(t *testing.T) {
	x := &Generator{
		worker:   1,
		lastTime: 1794508485001,
		sequence: 0,
	}
	_, err := x.NextID(1794508485000)
	if err == nil {
		t.Fatal(err)
	}

}

func TestSequenceResetsInANewMs(t *testing.T) {
	x := &Generator{
		worker:   1,
		lastTime: 1794508485001,
		sequence: 1,
	}
	want := 0
	_, err := x.NextID(x.lastTime + 1)
	if err != nil {
		t.Fatal("clock moved backwards")
	}
	if x.sequence != 0 {
		t.Fatalf("Error: want %d got %d", want, x.sequence)
	}

}

func TestNewGenerator_RejectsWorker1024(t *testing.T) {
	_, err := NewGenerator(1024)
	if err == nil {
		t.Fatal("Worker is greater than 1023")
	}
}

func TestSequenceExhaustion(t *testing.T) {
	s := &Generator{}
	now := uint64(1794508485001)
	for i := 0; i <= 4095; i++ {
		_, err := s.NextID(now)
		if err != nil {
			t.Fatal(err)
		}

	}
	_, err := s.NextID(now)
	if err == nil {
		t.Fatal("this should fail with: Sequence limit exceeded")
	}
}
