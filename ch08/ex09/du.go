package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var vFlag = flag.Bool("v", false, "show verbose progress messages")

type sizeEntry struct {
	rootIndex int
	size      int64
}

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	sizeEntries := make(chan sizeEntry)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(i, root, &n, sizeEntries)
	}
	go func() {
		n.Wait()
		close(sizeEntries)
	}()

	// Print the results periodically.
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))
loop:
	for {
		select {
		case info, ok := <-sizeEntries:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles[info.rootIndex]++
			nbytes[info.rootIndex] += info.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes) // final totals
}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, root := range roots {
		fmt.Printf("%s: %d files  %.1f GB\n", root, nfiles[i], float64(nbytes[i])/1e9)
	}
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(rootIndex int, dir string, n *sync.WaitGroup, fileSizes chan<- sizeEntry) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(rootIndex, subdir, n, fileSizes)
		} else {
			fileSizes <- sizeEntry{rootIndex, entry.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
