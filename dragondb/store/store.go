package store

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

	res := Store{
		data: make(map[string]string),
		file: file,
	}
	res.readDisk()
	return &res, nil
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
	s.file.Truncate(0)
	s.file.Seek(0, 0)
	for k, v := range mem.data {
		fmt.Println("Writing to disk:", k, v)
		_, err := fmt.Fprintf(s.file, "%s=%s\n", k, v)
		_ = err // fuck err
	}

	return nil
}

func (s *Store) readDisk() error {
	s.file.Seek(0, 0)
	scann := bufio.NewScanner(s.file)

	for scann.Scan() {
		line := scann.Text()
		k, v, _ := strings.Cut(line, "=")
		// fuck errors

		fmt.Printf("Found: '%s' -> '%s'\n", k, v)

		s.data[k] = v
	}

	return nil
}
