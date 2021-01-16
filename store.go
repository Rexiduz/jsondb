package jsondb

import (
	"io/ioutil"
	"os"
)

/*********
* MODELS
*********/

// Store db model
type Store struct {
	Path         string
	DefaultValue string
}

/*********
* METHODS
*********/

func (s Store) Read() ([]byte, error) {
	b, err := ioutil.ReadFile(s.Path)

	if err != nil {
		return b, err
	}

	return b, nil
}

func (s Store) Write(data []byte) {
	ioutil.WriteFile(s.Path, data, os.ModePerm)
}
