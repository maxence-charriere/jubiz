package main

import (
	"fmt"
	"time"
)

func timeSpanFromNow(t time.Time) (span string) {
	d := time.Since(t)

	days := d.Hours() / 24
	span = fmt.Sprintf("%v jours", int(days))
	if days < 2 {
		span = fmt.Sprintf("%v jour", 1)
	}

	if d.Hours() > 24 {
		return
	}

	span = fmt.Sprintf("%v heures", int(d.Hours()))
	if d.Hours() < 2 {
		span = fmt.Sprintf("%v heure", 1)
	}

	if d.Minutes() > 60 {
		return
	}

	span = fmt.Sprintf("%v minutes", int(d.Minutes()))
	if d.Minutes() < 2 {
		span = fmt.Sprintf("%v minute", 1)
	}
	return
}
