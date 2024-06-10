package main

import "testing"

func TestDistance1(t *testing.T) {

	rid := RoomId{}
	rid.Parse(`aaaaa-bbb-z-y-x-123[abxyz]`)
	if have, want := rid.GetSectorId(), 123; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := rid.IsValid(), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	rid.Parse(`a-b-c-d-e-f-g-h-987[abcde]`)
	if have, want := rid.GetSectorId(), 987; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := rid.IsValid(), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	rid.Parse(`not-a-real-room-404[oarel]`)
	if have, want := rid.GetSectorId(), 404; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := rid.IsValid(), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	rid.Parse(`totally-real-room-200[decoy]`)
	if have, want := rid.GetSectorId(), 200; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := rid.IsValid(), false; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}

func TestDistance2(t *testing.T) {

	rid := RoomId{}
	rid.Parse(`qzmt-zixmtkozy-ivhz-343[aaaaa]`)
	rid.Checksum = rid.ComputeChecksum()
	if have, want := rid.GetSectorId(), 343; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
	if have, want := rid.IsValid(), true; have != want {
		t.Errorf("have %v, want %v", have, want)
	}

	if have, want := rid.GetRealName(), "very encrypted name"; have != want {
		t.Errorf("have %v, want %v", have, want)
	}
}
