package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testStore(t *testing.T) *Store {
	t.Helper()
	dir, err := os.MkdirTemp("C:/Users/Jejo/AppData/Local/Temp", "docxterminal-test-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	return &Store{Dir: dir}
}

// ponytail: tests use an explicit Windows temp root because the harness TMPDIR can be malformed.

func TestSanitize(t *testing.T) {
	ok := []string{"git", "pnpm", "a", "Node.js", "foo_bar", "x-1"}
	for _, n := range ok {
		if _, err := sanitize(n); err != nil {
			t.Fatalf("%s: %v", n, err)
		}
	}
	bad := []string{"", "../etc", "foo/bar", "has space", ".md"}
	for _, n := range bad {
		if _, err := sanitize(n); err == nil {
			t.Fatalf("%s: expected error", n)
		}
	}
}

func TestCRUD(t *testing.T) {
	s := testStore(t)

	if err := s.Create("git"); err != nil {
		t.Fatal(err)
	}
	if err := s.Create("git"); err == nil {
		t.Fatal("dup create should fail")
	}
	body, err := s.Read("git")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(body, "# git") {
		t.Fatalf("body: %q", body)
	}

	p, err := s.FilePath("git")
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Ext(p) != ".md" {
		t.Fatal(p)
	}
	if err := os.WriteFile(p, []byte("# Git\n\nstatus\n"), 0644); err != nil {
		t.Fatal(err)
	}
	body, _ = s.Read("git")
	if !strings.Contains(body, "status") {
		t.Fatal(body)
	}

	if err := s.Create("npm"); err != nil {
		t.Fatal(err)
	}
	docs, err := s.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(docs) != 2 || docs[0] != "git" || docs[1] != "npm" {
		t.Fatalf("list: %v", docs)
	}

	if err := s.Delete("git"); err != nil {
		t.Fatal(err)
	}
	if s.Exists("git") {
		t.Fatal("still exists")
	}
	if err := s.Delete("git"); err == nil {
		t.Fatal("delete missing should fail")
	}
}

func TestPathTraversal(t *testing.T) {
	s := testStore(t)
	if _, err := s.FilePath("../secret"); err == nil {
		t.Fatal("traversal allowed")
	}
}
