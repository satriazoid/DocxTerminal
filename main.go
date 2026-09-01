package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func usage() {
	fmt.Fprint(os.Stderr, `dt — DocxTerminal

  dt                 TUI
  dt list            list docs
  dt show  NAME      print markdown
  dt add   NAME      create + open editor
  dt edit  NAME      open editor
  dt delete NAME     delete

Storage: ~/.docxterminal/*.md
Editor:  $EDITOR, else vim/nano
`)
}

func main() {
	s, err := OpenStore()
	if err != nil {
		die(err)
	}

	args := os.Args[1:]
	if len(args) == 0 {
		if err := runTUI(s); err != nil {
			die(err)
		}
		return
	}

	switch args[0] {
	case "list", "ls":
		docs, err := s.List()
		if err != nil {
			die(err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		for _, d := range docs {
			fmt.Fprintln(w, d+"\t"+d+".md")
		}
		w.Flush()
	case "show", "cat":
		need(args, 1)
		body, err := s.Read(args[1])
		if err != nil {
			die(err)
		}
		fmt.Print(body)
	case "add":
		need(args, 1)
		if err := s.Create(args[1]); err != nil {
			die(err)
		}
		p, _ := s.FilePath(args[1])
		if err := runEditor(p); err != nil {
			die(err)
		}
	case "edit":
		need(args, 1)
		if !s.Exists(args[1]) {
			die(fmt.Errorf("%s: not found", args[1]))
		}
		p, _ := s.FilePath(args[1])
		if err := runEditor(p); err != nil {
			die(err)
		}
	case "delete", "rm":
		need(args, 1)
		if err := s.Delete(args[1]); err != nil {
			die(err)
		}
		fmt.Println("deleted", args[1])
	case "-h", "--help", "help":
		usage()
	default:
		usage()
		os.Exit(2)
	}
}

func need(args []string, n int) {
	if len(args) < n+1 {
		usage()
		os.Exit(2)
	}
}
