package dayseven

import (
	"strings"
	"testing"
)

func TestBuildOperatingSystem(t *testing.T) {
	data := strings.NewReader("$ cd /\n$ ls\ndir a\n100 f\ndir foo\n200 g")

	root, _ := buildOperatingSystem(data)

	if root.directories[0].Name() != "a" {
		t.Errorf("Dir 'a' was not created properly")
	}

	if root.directories[1].Name() != "foo" {
		t.Errorf("Dir 'foo' was not created properly")
	}

	if root.files[0].Name() != "f" || root.files[0].Size() != 100 {
		t.Errorf("File 'f' was not created properly")
	}

	if root.files[1].Name() != "g" || root.files[1].Size() != 200 {
		t.Errorf("File 'g' was not created properly")
	}
}

func TestBuildOperatingSystemAndTraverseToFoo(t *testing.T) {
	data := strings.NewReader("$ cd /\n$ ls\ndir a\n100 f\ndir foo\n200 g\n$ cd foo")
	_, curr := buildOperatingSystem(data)

	if curr.Name() != "foo" {
		t.Errorf("Did not traverse correctly. In %v, should be in %v", curr.Name(), "foo")
	}

	data2 := strings.NewReader("$ cd /\n$ ls\ndir a\ndir foo\n$ cd foo\n$ ls\ndir bar\n$ cd bar")
	_, curr2 := buildOperatingSystem(data2)

	if curr2.Name() != "bar" {
		t.Errorf("Did not traverse correctly. In %v, should be in %v", curr2.Name(), "bar")
	}

	data3 := strings.NewReader("$ cd /\n$ ls\ndir a\ndir foo\n$ cd foo\n$ ls\ndir bar\n$ cd bar\n$ cd ..\n$ cd ..")
	_, curr3 := buildOperatingSystem(data3)

	if curr3.Name() != "/" {
		t.Errorf("Did not traverse correctly. In %v, should be in %v", curr3.Name(), "/")
	}
}
