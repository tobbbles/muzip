package state

import (
	"errors"
	"fmt"

	"github.com/tobbbles/muzip/archive"
)

type State struct {
	Archives map[string]*archive.Archive
}

func New() (*State, error) {
	archives := map[string]*archive.Archive{}

	s := &State{
		Archives: archives,
	}

	return s, nil
}

func (s *State) Exists(attr *archive.Attributes) bool {
	_, ok := s.Archives[fmt.Sprintf("%s-%s", attr.Name, attr.Hash)]
	return ok
}

func (s *State) Archive(attr *archive.Attributes) (*archive.Archive, error) {
	if s == nil {
		panic("State.Archive on nil *State")
	}

	if attr == nil {
		return nil, ErrMissingAttributes
	}

	a, found := s.Archives[fmt.Sprintf("%s-%s", attr.Name, attr.Hash)]
	if !found {
		return nil, ErrNoArchive
	}

	return a, nil
}

func (s *State) Save(path string) error {
	return errors.New("not implemented")
}

func Load(path string) (*State, error) {
	return nil, errors.New("Not implemented")
}

var (
	ErrNoArchive = errors.New("no archive in state matching name")
	ErrMissingAttributes = errors.New("missing archive attributes")
)