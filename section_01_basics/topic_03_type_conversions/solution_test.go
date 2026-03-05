package typeconversions

import (
	"math"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct {
		name string
		c    Celsius
		want Fahrenheit
	}{
		{"boiling", 100, 212},
		{"freezing", 0, 32},
		{"body temp", 37, 98.6},
		{"negative", -40, -40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CelsiusToFahrenheit(tt.c)
			if math.Abs(float64(got-tt.want)) > 0.01 {
				t.Errorf("CelsiusToFahrenheit(%v) = %v; want %v", tt.c, got, tt.want)
			}
		})
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name string
		f    Fahrenheit
		want Celsius
	}{
		{"boiling", 212, 100},
		{"freezing", 32, 0},
		{"negative", -40, -40},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FahrenheitToCelsius(tt.f)
			if math.Abs(float64(got-tt.want)) > 0.01 {
				t.Errorf("FahrenheitToCelsius(%v) = %v; want %v", tt.f, got, tt.want)
			}
		})
	}
}

func TestStringToInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"valid", "42", 42, false},
		{"negative", "-7", -7, false},
		{"zero", "0", 0, false},
		{"invalid", "abc", 0, true},
		{"empty", "", 0, true},
		{"float string", "3.14", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToInt(%q) error = %v; wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StringToInt(%q) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestIntToString(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{42, "42"},
		{-7, "-7"},
		{0, "0"},
		{1000000, "1000000"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := IntToString(tt.input)
			if got != tt.want {
				t.Errorf("IntToString(%d) = %q; want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestByteCountToHuman(t *testing.T) {
	tests := []struct {
		name  string
		bytes uint64
		want  string
	}{
		{"bytes", 500, "500.00 B"},
		{"1KB", 1024, "1.00 KB"},
		{"1.5KB", 1536, "1.50 KB"},
		{"1MB", 1048576, "1.00 MB"},
		{"1GB", 1073741824, "1.00 GB"},
		{"0", 0, "0.00 B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ByteCountToHuman(tt.bytes)
			if got != tt.want {
				t.Errorf("ByteCountToHuman(%d) = %q; want %q", tt.bytes, got, tt.want)
			}
		})
	}
}
