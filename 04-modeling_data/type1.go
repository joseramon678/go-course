package main

import "time"

type page struct {
	number int
	header string
	text   string
}

type book struct {
	title       string
	autor       string
	releaseDate *time.Time
	pages       []page
}
