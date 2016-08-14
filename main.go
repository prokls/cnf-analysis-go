package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	input "github.com/prokls/cnf-analysis-go/input"
	output "github.com/prokls/cnf-analysis-go/output"
	"github.com/prokls/cnf-analysis-go/stats"
)

const USAGE = `usage: cnf-analysis-go [-h] [-f {xml,json}] [--ignore IGNORE] [-u UNITS] [-n]
                       [-p] [-s]
                       dimacsfiles [dimacsfiles ...]

CNF analysis

positional arguments:
  dimacsfiles           filepath of DIMACS file

optional arguments:
  -h, --help            show this help message and exit
  -f {xml,json}, --format {xml,json}
                        format to store feature data in
  --ignore IGNORE       a prefix for lines that shall be ignored (like "c")
  -u UNITS, --units UNITS
                        how many units (= processes) should run in
                        concurrently
  -n, --no-hashes       do not compute hashes for the CNF file considered
  -p, --fullpath        use full path instead of basename in featurefiles
  -s, --skip-existing   skip CNF file if file.stats.json exists
`

type work struct {
	input       string
	output      string
	format      int
	ignoreLines []string
	fullpath    bool
	hashes      bool
}

func worker(workDist chan work, w *sync.WaitGroup) {
	defer w.Done()

	for job := range workDist {
		var err error
		pconf := input.NewParsingConfig()
		fconf := stats.NewFeatureConfig()
		oconf := output.NewOutputConfig()

		oconf.Format = job.format
		pconf.IgnoreLines = job.ignoreLines
		fconf.FullPath = job.fullpath
		fconf.Hashes = job.hashes

		if len(pconf.IgnoreLines) == 0 {
			pconf.IgnoreLines = append(pconf.IgnoreLines, "c", "%")
		}

		log.Printf("considering %s", job.input)

		// read file
		fd, err := os.Open(job.input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "read file %s failed: %s", job.input, err.Error())
			return
		}

		// parse file
		cnf, err := input.ReadCNFFile(fd, pconf)
		fd.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while processing %s: %s\n", job.input, err.Error())
			return
		}

		// evaluate features
		stat := output.NewStats()
		stats.Metadata(stat, job.input, fconf)

		err = evaluate(cnf, &stat.Fts, fconf)
		if err != nil {
			fmt.Fprintf(os.Stderr, "evaluation failed: %s", err.Error())
			return
		}

		log.Printf("writing file %s", job.output)

		// write features
		out, err := os.Create(job.output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not write file %s: %s\n", job.output, err.Error())
			return
		}

		err = output.WriteFeatures(stat, out, oconf)
		out.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error while writing features: %s\n", err.Error())
			return
		}
	}

	return
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func deriveFilePath(base string, skipExisting bool) (string, error) {
	ext := filepath.Ext(base)
	woExt := base[0 : len(base)-len(ext)]
	newFile := woExt + ".stats.json"

	if !exists(newFile) {
		return newFile, nil
	}

	if skipExisting {
		return "", nil
	}

	// exists & !skip: move existing file
	timestamp := time.Now().UTC().Format("20060102150405")
	altFile := fmt.Sprintf("%s.backup%s.stats.json", woExt, timestamp)
	err := os.Rename(newFile, altFile)
	if err != nil {
		return newFile, err
	}
	log.Printf("existing %s has been renamed to %s", newFile, altFile)

	return newFile, nil
}

func main() {
	var files []string
	var ignoreLines []string
	format := output.JSONFormat
	units := 4
	skip_existing := false
	fullpath := false
	hashes := true

	skip := true
	for i, arg := range os.Args {
		if skip {
			skip = false
			continue
		}
		if arg == "-h" || arg == "--help" {
			fmt.Println(USAGE)
			os.Exit(0)
		} else if arg == "-f" || arg == "--format" {
			form := os.Args[i+1]
			if form != "json" && form != "xml" {
				fmt.Fprintf(os.Stderr, "invalid format supplied: '%s'", form)
				os.Exit(1)
			} else if form == "xml" {
				fmt.Fprint(os.Stderr, "XML not yet supported")
				os.Exit(1)
			}
			skip = true
		} else if arg == "--ignore" {
			ignoreLines = append(ignoreLines, os.Args[i+1])
			skip = true
		} else if arg == "-u" || arg == "--units" {
			u, err := strconv.Atoi(os.Args[i+1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "--units parameter invalid: %s", err.Error())
				os.Exit(1)
			} else if u <= 0 {
				fmt.Fprintf(os.Stderr, "--units must be positive")
				os.Exit(1)
			}
			units = u
			skip = true
		} else if arg == "-n" || arg == "--no-hashes" {
			hashes = false
		} else if arg == "-p" || arg == "--fullpath" {
			fullpath = true
		} else if arg == "-s" || arg == "--skip-existing" {
			skip_existing = true
		} else {
			files = append(files, arg)
		}
	}

	if len(files) < units {
		units = len(files)
	}

	var jobs []work
	for _, file := range files {
		out, err := deriveFilePath(file, skip_existing)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
		}
		if out == "" {
			fmt.Fprintf(os.Stderr, "%s was processed previously - skipping\n", file)
			continue
		}
		jobs = append(jobs, work{
			input:       file,
			output:      out,
			format:      format,
			ignoreLines: ignoreLines,
			fullpath:    fullpath,
			hashes:      hashes,
		})
	}

	var w sync.WaitGroup
	workDist := make(chan work, 1)
	for i := 0; i < units; i++ {
		w.Add(1)
		go worker(workDist, &w)
	}
	for _, j := range jobs {
		workDist <- j
	}
	close(workDist)
	w.Wait()
}
