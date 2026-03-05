package constants

import "testing"

func TestDayName(t *testing.T) {
	tests := []struct {
		day  Day
		want string
	}{
		{Sunday, "Sunday"},
		{Monday, "Monday"},
		{Tuesday, "Tuesday"},
		{Wednesday, "Wednesday"},
		{Thursday, "Thursday"},
		{Friday, "Friday"},
		{Saturday, "Saturday"},
		{Day(-1), "Unknown"},
		{Day(7), "Unknown"},
		{Day(100), "Unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := DayName(tt.day)
			if got != tt.want {
				t.Errorf("DayName(%d) = %q; want %q", tt.day, got, tt.want)
			}
		})
	}
}

func TestIotaValues(t *testing.T) {
	if Sunday != 0 {
		t.Errorf("Sunday = %d; want 0", Sunday)
	}
	if Saturday != 6 {
		t.Errorf("Saturday = %d; want 6", Saturday)
	}
}

func TestByteSize(t *testing.T) {
	if KB != 1024 {
		t.Errorf("KB = %d; want 1024", KB)
	}
	if MB != 1024*1024 {
		t.Errorf("MB = %d; want %d", MB, 1024*1024)
	}
	if GB != 1024*1024*1024 {
		t.Errorf("GB = %d; want %d", GB, 1024*1024*1024)
	}
}

func TestToBytes(t *testing.T) {
	tests := []struct {
		name string
		size uint64
		unit string
		want uint64
	}{
		{"5KB", 5, "KB", 5 * 1024},
		{"2MB", 2, "MB", 2 * 1024 * 1024},
		{"1GB", 1, "GB", 1 * 1024 * 1024 * 1024},
		{"3TB", 3, "TB", 3 * 1024 * 1024 * 1024 * 1024},
		{"0KB", 0, "KB", 0},
		{"invalid unit", 5, "PB", 0},
		{"empty unit", 5, "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToBytes(tt.size, tt.unit)
			if got != tt.want {
				t.Errorf("ToBytes(%d, %q) = %d; want %d", tt.size, tt.unit, got, tt.want)
			}
		})
	}
}

func TestIsWeekend(t *testing.T) {
	tests := []struct {
		day  Day
		want bool
	}{
		{Sunday, true},
		{Saturday, true},
		{Monday, false},
		{Wednesday, false},
		{Friday, false},
	}
	for _, tt := range tests {
		t.Run(DayName(tt.day), func(t *testing.T) {
			got := IsWeekend(tt.day)
			if got != tt.want {
				t.Errorf("IsWeekend(%d) = %v; want %v", tt.day, got, tt.want)
			}
		})
	}
}
