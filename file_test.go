package jubiz

import "testing"
import "time"
import "os"

func TestFileMarshal(t *testing.T) {
	name := "test.json"
	a := Article{
		Title:   "Hello",
		PubDate: time.Now(),
	}

	if err := FileMarshal(name, a); err != nil {
		t.Error(err)
	}
	defer os.Remove(name)

	if err := FileMarshal("test.json", a); err != nil {
		t.Error(err)
	}

}

func TestFileUnmarshal(t *testing.T) {
	var a Article
	name := "test.json"

	if err := FileUnmarshal(name, &a); err == nil {
		t.Fatal("err should not be nil")
	}

	if err := FileMarshal(name, a); err != nil {
		t.Error(err)
	}
	defer os.Remove(name)

	if err := FileUnmarshal(name, &a); err != nil {
		t.Error(err)
	}
}
