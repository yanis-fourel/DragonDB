package store

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

var dbfile string = filepath.Join("_dragondata", "data.db")

type Store struct {
	data map[string]string
	file *os.File
}

func New() (*Store, error) {
	os.MkdirAll(filepath.Dir(dbfile), fs.ModePerm)
	file, err := os.OpenFile(dbfile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &Store{
		data: make(map[string]string),
		file: file,
	}, nil
}

func (s *Store) Close() error {
	return s.file.Close()
}

func (s *Store) Set(key string, value string) {
	s.data[key] = value
	_ = s.writeDisk(s)
}

func (s *Store) Get(key string) string {
	return s.data[key]
}

func (s *Store) writeDisk(mem *Store) error {
	s.file.Seek(0, 0)
	for k, v := range mem.data {
		_, _ = fmt.Fprintf(s.file, "%s=%s\n", k, v)
		// fuck err
	}

	return nil
}

func (s *Store) readDisk(key string) (string, error) {
	s.file.Seek(0, 0)
	var k, v string
	for {
		_, err := fmt.Fscanf(s.file, "%s=%s\n", &k, &v)
		if err == io.EOF {
			return "", nil
		}
		fmt.Println("Found:", k, v)

		if err != nil {
			return "", err
		}
		if k == key {
			return v, nil
		}
	}
}
