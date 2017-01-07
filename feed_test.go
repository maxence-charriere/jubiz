package jubiz

import "testing"

func TestGetFeed(t *testing.T) {
	f, err := GetFeed()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", f)
}
