package main

import (
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/google/licensecheck"
)

func run() error {
	fsys := os.DirFS("/")
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	glob := filepath.Join(pwd, "LICENSE*")[1:]
	log.Printf("search glob: %q", glob)
	files, err := fs.Glob(fsys, glob)
	if err != nil {
		return err
	}
	log.Printf("found %d files", len(files))
	for _, filename := range files {
		data, err := fs.ReadFile(fsys, filename)
		if err != nil {
			log.Printf("failed to read file %q: %s", filename, err.Error())
			continue
		}
		fmt.Println(filename)
		scan(data)
	}
	return nil
}

func scan(text []byte) {
	cov := licensecheck.Scan(text)
	fmt.Printf("%.1f%% of text covered by licenses:\n", cov.Percent)
	for _, m := range cov.Match {
		fmt.Printf("%s at [%d:%d] IsURL=%v\n", m.ID, m.Start, m.End, m.IsURL)
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	err := run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}