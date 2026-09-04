package main

import (
	"strings"
	"testing"
)

func TestRenderMarkdown(t *testing.T) {
	src := "# Title\n\n- item\n"
	out := renderMarkdown(src, 60)
	if out == src {
		t.Fatal("expected glamour transform")
	}
	if !strings.Contains(out, "Title") {
		t.Fatal("missing Title")
	}
	if !strings.Contains(out, "item") {
		t.Fatal("missing item")
	}
	_ = renderMarkdown("", 60)
}
