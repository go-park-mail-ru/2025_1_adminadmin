package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOptions(t *testing.T) {
	params := NewOptions()
	require.Equal(t, 10, params.GetCount(), "Ожидалось count = 10")
	require.Equal(t, 0, params.GetOffset(), "Ожидалось offset = 0")
}

func TestWithCustomCount(t *testing.T) {
	tests := []struct {
		name     string
		count    int
		total    int
		expected int
	}{
		{"Count more than 0", -5, 100, 0},
		{"Count more than total", 150, 100, 100},
		{"Count in range", 50, 100, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := NewOptions(WithCustomCount(test.count, test.total))
			require.Equal(t, test.expected, params.GetCount(), "Неверное значение count")
		})
	}
}

func TestWithCustomOffset(t *testing.T) {
	tests := []struct {
		name     string
		offset   int
		total    int
		expected int
	}{
		{"Offset less than 0", -5, 100, 0},
		{"Offset more than total", 150, 100, 0},
		{"Offset in range", 50, 100, 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			params := NewOptions(WithCustomOffset(test.offset, test.total))
			require.Equal(t, test.expected, params.GetOffset(), "Неверное значение offset")
		})
	}
}

func TestCombinedOptions(t *testing.T) {
	params := NewOptions(
		WithCustomCount(20, 100),
		WithCustomOffset(10, 100),
	)

	require.Equal(t, 20, params.GetCount(), "Wrong count value")
	require.Equal(t, 10, params.GetOffset(), "Wrong offset value")
}
