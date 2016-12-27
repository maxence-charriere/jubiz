package main

import "testing"

func TestGetFeed(t *testing.T) {
	f, err := getFeed()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", f)
}
