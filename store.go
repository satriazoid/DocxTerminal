package main

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

//go:embed seeds/*.md
var seedFS embed.FS

var nameRe = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9._-]*$`)

type Store struct {
	Dir string
}

func defaultDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".docxterminal"), nil
}

func OpenStore() (*Store, error) {
	dir, err := defaultDir()
	if err != nil {
		return nil, err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}
	s := &Store{Dir: dir}
	if err := s.seedIfEmpty(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Store) seedIfEmpty() error {
	docs, err := s.List()
	if err != nil {
		return err
	}
	if len(docs) > 0 {
		return nil
	}
	entries, err := seedFS.ReadDir("seeds")
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		b, err := seedFS.ReadFile("seeds/" + e.Name())
		if err != nil {
			return err
		}
		dst := filepath.Join(s.Dir, e.Name())
		if err := os.WriteFile(dst, b, 0644); err != nil {
			return err
		}
	}
	return nil
}

func sanitize(name string) (string, error) {
	name = strings.TrimSpace(name)
	name = strings.TrimSuffix(name, ".md")
	if !nameRe.MatchString(name) {
		return "", fmt.Errorf("invalid name %q (use letters, numbers, . _ -)", name)
	}
	return name, nil
}

func (s *Store) path(name string) (string, error) {
	n, err := sanitize(name)
	if err != nil {
		return "", err
	}
	return filepath.Join(s.Dir, n+".md"), nil
}

func (s *Store) List() ([]string, error) {
	entries, err := os.ReadDir(s.Dir)
	if err != nil {
		return nil, err
	}
	var out []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		out = append(out, strings.TrimSuffix(e.Name(), ".md"))
	}
	sort.Strings(out)
	return out, nil
}

func (s *Store) Read(name string) (string, error) {
	p, err := s.path(name)
	if err != nil {
		return "", err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return "", fmt.Errorf("%s: not found", name)
		}
		return "", err
	}
	return string(b), nil
}

func (s *Store) Exists(name string) bool {
	p, err := s.path(name)
	if err != nil {
		return false
	}
	_, err = os.Stat(p)
	return err == nil
}

func (s *Store) Create(name string) error {
	p, err := s.path(name)
	if err != nil {
		return err
	}
	if _, err := os.Stat(p); err == nil {
		return fmt.Errorf("%s: already exists", name)
	}
	n, _ := sanitize(name)
	body := "# " + n + "\n\n"
	return os.WriteFile(p, []byte(body), 0644)
}

func (s *Store) Delete(name string) error {
	p, err := s.path(name)
	if err != nil {
		return err
	}
	if err := os.Remove(p); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return fmt.Errorf("%s: not found", name)
		}
		return err
	}
	return nil
}

func (s *Store) FilePath(name string) (string, error) {
	return s.path(name)
}
