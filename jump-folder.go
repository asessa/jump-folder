package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

var addFlag string
var listFlag bool
var pathFlag string
var removeFlag string
var bashFlag bool

var bookmarks []string
var filename string

// loadBookmarks loads bookmarked folders as TSV.
func loadBookmarks() {
	bookmarks = nil

	fh, err := os.Open(filename)
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		bookmarks = append(bookmarks, scanner.Text())
	}
}

// saveBookmarks writes back the bookmarks to the config file.
func saveBookmarks() {
	fh, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	w := bufio.NewWriter(fh)
	for _, bookmark := range bookmarks {
		fmt.Fprintln(w, bookmark)
	}
	w.Flush()
}

// remove an element from bookmarks.
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	flag.StringVar(&addFlag, "a", "", "Add the current folder in bookmarks")
	flag.BoolVar(&listFlag, "l", false, "List saved bookmarks")
	flag.StringVar(&pathFlag, "p", "", "Prints the path of a bookmark")
	flag.StringVar(&removeFlag, "d", "", "Deletes bookmark")
	flag.BoolVar(&bashFlag, "bash", false, "Bash/ZSH integration")

	flag.Parse()

	user, err := user.Current()
	if err != nil {
		panic("Unable to get current user")
	}

	filename = filepath.Join(user.HomeDir, ".jump-folder")

	loadBookmarks()

	// List bookmarks
	if listFlag {
		for _, line := range bookmarks {
			fmt.Printf("%v\n", line)
		}
		return
	}

	// Add bookmark
	if addFlag != "" {
		cwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		bookmarks = append(bookmarks, addFlag+"\t"+cwd)
		saveBookmarks()
		fmt.Println("Bookmark added.")
		return
	}

	// Display bookmark path
	if pathFlag != "" {
		for _, line := range bookmarks {
			bookmark := strings.Split(line, "\t")
			if bookmark[0] == pathFlag {
				fmt.Printf("%s\n", bookmark[1])
				break
			}
		}
		return
	}

	// Remove bookmark
	if removeFlag != "" {
		for idx, line := range bookmarks {
			bookmark := strings.Split(line, "\t")
			if bookmark[0] == removeFlag {
				bookmarks = remove(bookmarks, idx)
				fmt.Println("Bookmark deleted.")
				break
			}
		}
		saveBookmarks()
		return
	}

	// Jump to selected bookmark
	if len(os.Args) > 1 {
		for _, line := range bookmarks {
			bookmark := strings.Split(line, "\t")
			if bookmark[0] == os.Args[1] {
				fmt.Printf("%s\n", bookmark[1])
				return
			}
		}
	}

	if bashFlag {
		fmt.Println(`function jump {
if [ $# -lt 1 ]; then
	jump-folder
elif [ ${1:0:1} == "-" ]; then
	jump-folder $*
else
	cd $(jump-folder $*)
fi
}`)
		return
	}

	// Show help
	flag.Usage()
}
