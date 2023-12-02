package store

import (
	"io/fs"
	"os"
	"path/filepath"
)

var dbfile string = filepath.Join("_dragon", "data.db")

type DiskStore struct {
	file *os.File
}

func NewDiskStore() (*DiskStore, error) {
	os.MkdirAll(filepath.Dir(dbfile), fs.ModePerm)
	file, err := os.OpenFile(dbfile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &DiskStore{file}, nil
}

func (d *DiskStore) Close() error {
	return d.file.Close()
}
