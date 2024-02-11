package gault

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/mohieey/gault/encrypt"
)

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func Memory(encodingKey string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		keyValues:   map[string]string{},
	}
}

func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = map[string]string{}
		return nil
	}

	var sb strings.Builder
	_, err = io.Copy(&sb, f)
	if err != nil {
		return err
	}
	decryptedJSON, err := encrypt.Decrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}

	r := strings.NewReader(decryptedJSON)

	jsonDecoder := json.NewDecoder(r)
	err = jsonDecoder.Decode(&v.keyValues)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vault) save() error {
	var sb strings.Builder
	enc := json.NewEncoder(&sb)
	err := enc.Encode(v.keyValues)
	if err != nil {
		return err
	}

	encryptedJSON, err := encrypt.Encrypt(v.encodingKey, sb.String())
	if err != nil {
		return err
	}

	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = fmt.Fprint(f, encryptedJSON)
	if err != nil {
		return err
	}

	return nil
}

func (v *Vault) Get(key string) (string, error) {
	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("value not found")
	}

	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value
	err = v.save()
	if err != nil {
		return err
	}

	return nil
}
