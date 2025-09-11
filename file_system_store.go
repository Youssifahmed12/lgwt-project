package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	// rewinds the reader to the start
	f.database.Seek(0, io.SeekStart)
	league, _ := NewLeague(f.database)
	return league
}
