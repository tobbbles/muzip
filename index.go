package main

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhowden/tag"

	"github.com/tobbbles/muzip/archive"
)

func Walk(dir string) ([]*archive.Archive, error) {
	var archives []*archive.Archive

	walkr := func(path string, info os.FileInfo, err error) error {
		if info.Name()[len(info.Name())-3:] != "zip" {
			return nil
		}

		a, err := Index(path)
		if err != nil {
			return err
		}

		archives = append(archives, a)
		return nil
	}

	if err := filepath.Walk(dir, walkr); err != nil {
		return nil, err
	}

	return archives, nil
}

func Index(file string) (*archive.Archive, error) {

	f, err := os.Open(file)
	if err != nil {
return nil, err
	}
	defer f.Close()

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer(nil)
	if _, err := io.Copy(b, f); err != nil {
		return nil, err
	}

	r, err := zip.NewReader(f, s.Size())
	if err != nil {
		return nil, err
	}

	sum, err := Sum(b.Bytes())
	if err != nil {
		return nil, err
	}

	a := &archive.Archive{
		Attr: &archive.Attributes{
			Name:     file,
			Hash: sum,
		},

		Contents: make([]*archive.Metadata, 0, len(r.File)),
	}

	// Iterate through the files in the archive,
	// printing some of their contents.
	c := 0
	for _, f := range r.File {
		ext := f.FileInfo().Name()[len(f.FileInfo().Name())-3:]

		switch strings.ToLower(ext) {
		case "mp3", "flac", "FLAC":
			m, err := Parse(f)
			if err != nil {
				panic(err)
			}

			a.Contents = append(a.Contents, m)
			c++

		default:
			continue
		}
	}

	a.TrackCount = c

	return a, nil
}

func Parse(f *zip.File) (*archive.Metadata, error) {
	rc, err := f.Open()
	if err != nil {
		log.Fatal(err)
	}

	buf, err := ioutil.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	rc.Close()

	sum, err := Sum(buf)
	if err != nil {
		return nil, err
	}

	m, err := tag.ReadFrom(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	trn, _ := m.Track()

	md := &archive.Metadata{
		File:           f.Name,
		Title:          m.Title(),
		Artist:         m.Artist(),
		Sum:            sum,
		Track:          trn,
		Size:           f.UncompressedSize64,
		CompressedSize: f.CompressedSize64,
		Type:           string(m.FileType()),
	}

	return md, nil
}

func Sum(buf []byte) (string, error) {
	h := sha256.New()

	if _, err := h.Write(buf); err != nil {
		return "", err
	}

	sum := fmt.Sprintf("%x", h.Sum(nil))

	return sum, nil
}
