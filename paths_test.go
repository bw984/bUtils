package main

import "testing"

func TestParents(t *testing.T) {
	cases := []struct {
		Description string
		Got         string
		Want        string
		Levels      int
	}{
		{"Level 0 with file", "C:/UTools/Outer/Inner/UTools.xlsx", "C:/UTools/Outer/Inner", 0},
		{"Level 0 with directory trailing slash", "C:/UTools/Outer/Inner/", "C:/UTools/Outer", 0},
		{"Level 0 with directory no trailing slash", "C:/UTools/Outer/Inner", "C:/UTools/Outer", 0},
		{"Level 1 with file", "C:/UTools/Outer/Inner/UTools.xlsx", "C:/UTools/Outer", 1},
		{"Level 1 with directory trailing slash", "C:/UTools/Outer/Inner/", "C:/UTools", 1},
		{"Level 1 with directory no trailing slash", "C:/UTools/Outer/Inner", "C:/UTools", 1},
		{"Level 1 file with root leading /", "/Users/bradywalters/RiderProjects/BitStorkParse2022/Program.cs", "/Users/bradywalters/RiderProjects", 1},
		{"Windows Path Level 0", "C:\\UTools\\Outer\\", "C:/UTools", 0},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := Parents(test.Got, test.Levels)
			if got != test.Want {
				t.Errorf("got %q, want %q", got, test.Want)
			}
		})
	}
}

func TestTrimLeft(t *testing.T) {
	cases := []struct {
		Description string
		Got         string
		Want        string
		Levels      int
	}{
		{"Corporate_Root", "\\\\rbi.local\\corporate_root\\company\\Spec Sheets\\Steel\\10.625 270mm\\370S1072 - 10.625 - U513S.pdf", "Spec Sheets/Steel/10.625 270mm/370S1072 - 10.625 - U513S.pdf", 3},
		{"Corporate_Root Slash Crazy", "\\\\\\\\rbi.local\\\\corporate_root\\company\\\\Spec Sheets////\\////Steel\\10.625 270mm\\370S1072 - 10.625 - U513S.pdf", "Spec Sheets/Steel/10.625 270mm/370S1072 - 10.625 - U513S.pdf", 3},
		{"Local File", "C:\\UTools\\Outputs\\test.xls", "Outputs/test.xls", 2},
		{"Local Directory", "C:\\UTools\\Outputs\\TestDir\\", "Outputs/TestDir", 2},
		{"Local Directory No Trim", "C:\\UTools\\Outputs\\TestDir\\", "C:/UTools/Outputs/TestDir", 0},
		{"Local Directory", "C:\\UTools\\Outputs\\TestDir\\", "", 10},
		{"Local Directory Slash Crazy", "C:\\UTools\\/////Outputs\\\\\\TestDir\\///", "C:/UTools/Outputs/TestDir", 0},
		{"Single File", "Test.xls", "Test.xls", 0},
		{"Single File", "Test.xls", "", 1},
		{"Single File", "C:/Test.xls", "Test.xls", 1},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := TrimLeft(test.Got, test.Levels)
			if got != test.Want {
				t.Errorf("Desc: %q -> got %q, want %q", test.Description, got, test.Want)
			}
		})
	}
}
