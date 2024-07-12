package main

import (
	"testing"
)

var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake  = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar    = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre      = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		filename string
		want     []Bookworm
		wantErr  bool
	}
	tests := map[string]testCase{
		"file exists": {
			filename: "testdata/bookworms_test.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			filename: "testdata/no_file.json",
			want:     nil,
			wantErr:  true,
		},
		"invalid JSON": {
			filename: "testdata/invalid.json",
			want:     nil,
			wantErr:  true,
		},
	}

	for testName, tc := range tests {
		t.Run(testName, func(t *testing.T) {
			got, err := loadBookworms(tc.filename)
			if tc.wantErr && err == nil {
				t.Fatalf("expected an error, got none")
			}

			if !tc.wantErr && err != nil {
				t.Fatalf("expected no error, got one %s", err.Error())
			}

			if !equalBookworms(got, tc.want) {
				t.Fatalf("arrays don't match: got %v, expected %v", got, tc.want)
			}
			// Using reflect is also an option (although not very performant)
			// if !reflect.DeepEqual(got, testCase.want) {
			// 	t.Fatalf("arrays don't match:  got %v, expected %v", got, testCase.want)
			// }
		})
	}
}

// equalBookworms is a helper to test the equality of two lists of Bookworms.
func equalBookworms(bookworms, target []Bookworm) bool {
	if len(bookworms) != len(target) {
		return false
	}

	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}

		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}

	return true
}

// equalBooks is a helper to test the equality of two lists of Books.
func equalBooks(books, target []Book) bool {
	if len(books) != len(target) {
		return false
	}

	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}

	return true
}
