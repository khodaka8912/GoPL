package main

import (
	"time"
)

var m *multiSort

func main() {
	m = newMultiSort(tracks)

}

type multiSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (m *multiSort) Less(i, j int) bool {
	return m.less(m.t[i], m.t[j])
}

func (m *multiSort) Swap(i, j int) {
	m.t[i], m.t[j] = m.t[j], m.t[i]
}

func (m *multiSort) Len() int {
	return len(m.t)
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, mustParseDuration("3m38s")},
	{"Go", "Moby", "Moby", 1992, mustParseDuration("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, mustParseDuration("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, mustParseDuration("4m24s")},
}

func (m *multiSort) byTitle() {
	m.less = func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		return m.less(x, y)
	}
}

func (m *multiSort) byArtist(x, y *Track) {
	m.less = func(x, y *Track) bool {
		if x.Artist != y.Artist {
			return x.Artist < y.Artist
		}
		return m.less(x, y)
	}
}

func (m *multiSort) byAlbum() {
	m.less = func(x, y *Track) bool {
		if x.Album != y.Album {
			return x.Album < y.Album
		}
		return m.less(x, y)
	}
}

func (m *multiSort) byYear() {
	m.less = func(x, y *Track) bool {
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		return m.less(x, y)
	}
}

func (m *multiSort) byLength() {
	m.less = func(x, y *Track) bool {
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return m.less(x, y)
	}
}

func newMultiSort(t []*Track) *multiSort {
	return &multiSort{t, func(_, _ *Track) bool { return false }}
}

func mustParseDuration(s string) time.Duration {
	t, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}
	return t
}
