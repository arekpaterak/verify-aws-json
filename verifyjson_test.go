package main

import (
	"testing"
)

func TestVerifyPolicy(t *testing.T) {
	testCases := []struct {
		name     string
		filePath string
		want	 bool
		wantError  bool
	}{
		{
			name:     "Example",
			filePath: "testdata/example.json",
			want:     false,
			wantError:  false,
		},
		{
			name:     "ValidPolicy",
			filePath: "testdata/true.json",
			want:     true,
			wantError:  false,
		},
		{
			name:     "NonExistentFile",
			filePath: "testdata/non_existent.json",
			want:     false,
			wantError:  true,
		},
		{
			name:     "EmptyResource",
			filePath: "testdata/empty_resource.json",
			want:     true,
			wantError:  false,
		},
		{
			name:     "InvalidJSON",
			filePath: "testdata/invalid.json",
			want:     false,
			wantError:  true,
		},
		{
			name:     "Empty",
			filePath: "testdata/empty.json",
			want:     false,
			wantError:  true,
		},
		{
			name:     "NoResource",
			filePath: "testdata/no_resource.json",
			want:     true,
			wantError:  false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := verifyPolicy(testCase.filePath)
			if (err != nil) != testCase.wantError {
				t.Fatalf("verifyPolicy() error = %v, wantError %v", err, testCase.wantError)
				return
			}
			if got != testCase.want {
				t.Fatalf("verifyPolicy() = %v, want %v", got, testCase.want)
			}
		})
	}
}