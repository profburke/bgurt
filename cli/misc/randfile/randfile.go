package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func contains(set map[string]bool, candidate string) bool {
	_, found := set[candidate]
	return found
}

func visit(files *[]string, extensions map[string]bool) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		ext := strings.ToLower(filepath.Ext(path))
		if len(extensions) == 0 || contains(extensions, ext) {
			//			if !info.IsDir() { // Justin Case
			*files = append(*files, path)
			//			}
		}

		return nil
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: randfile <directory> [ext1] [ext2] ... [extn]")
		os.Exit(1)
	}
	rand.Seed(time.Now().Unix())

	dirpath := os.Args[1]
	var extensions map[string]bool

	if len(os.Args) > 2 {
		extensions = make(map[string]bool)
		for i := 2; i < len(os.Args); i++ {
			extension := os.Args[i]
			// filepath.Ext() includes the '.'
			if extension[0] != '.' {
				extension = "." + extension
			}
			extensions[extension] = true
		}
	}

	var files []string
	err := filepath.Walk(dirpath, visit(&files, extensions))
	if err != nil {
		fmt.Fprintf(os.Stderr, "", err)
		os.Exit(1)
	} else if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "randfile: no matching files found in '%s'\n", dirpath)
		os.Exit(1)
	}

	fmt.Println(files[rand.Intn(len(files))])
}

// Local Variables:
// compile-command: "go build"
// End:
