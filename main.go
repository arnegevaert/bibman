package main

// https://github.com/charmbracelet/bubbletea/tree/master/examples/list-fancy

import (
	"fmt"
	"os"
	"path"
	"slices"
	"strings"
    tea "github.com/charmbracelet/bubbletea"
)

func main() {
	checkHealth()
}

type entry struct {
    title string
    author string
    year int16
}

type model struct {
    entries map[string]entry
    cursor string
}

func initializeModel() model {
    // TODO:
    // read filenames from directories
    // parse bibfiles
    return model{
        entries: map[string]entry{},
        cursor: "",
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func parseBibFile(filename string) (entry) {
    // TODO: parse the bib file
    return entry{
        title: filename,
        author: filename,
        year: 2000,
    }
}

func getEntries(dir string) ([]string){
    filenames, err := os.ReadDir(dir)
    if err != nil {
        logError(fmt.Sprint(err))
        os.Exit(1)
    }
    log("Found", fmt.Sprintf("%d", len(filenames)), "files in", dir)
    result := make([]string, len(filenames))
    for i := range(len(filenames)) {
        fn := filenames[i].Name()
        if strings.Contains(fn, "@") {
            fn = fn[1:]
        }
        result[i] = strings.Split(fn, ".")[0]
    }
    return result
}

func checkHealth() {
	const BASE_PATH = "/home/arne/Work/Buckets/ZK Zettelkasten/Bibman/"
    log("Base path:", BASE_PATH)

    BIB_PATH := path.Join(BASE_PATH, "BIB")
	var PDF_PATH = path.Join(BASE_PATH, "PDF")
	var MD_PATH = path.Join(BASE_PATH, "MD")

	// Get all BIB, PDF and MD files
    bibEntries := getEntries(BIB_PATH)
    mdEntries := getEntries(MD_PATH)
    pdfEntries := getEntries(PDF_PATH)

	// Each PDF must have an associated BIB (ERR)
    for i := range pdfEntries {
        pdfEntry := pdfEntries[i]
        if !slices.Contains(bibEntries, pdfEntry) {
            logError("PDF entry", fmt.Sprintf("\"%s\"", pdfEntry), "not found in BIB entries.")
        }
    }

	// Each BIB may have an associated PDF (WARN)
    for i := range bibEntries {
        bibEntry := bibEntries[i]
        if !slices.Contains(pdfEntries, bibEntry) {
            logWarning("BIB entry", fmt.Sprintf("\"%s\"", bibEntry), "not found in PDF entries.")
        }
    }

	// Each BIB may have an associated MD (WARN)
    for i := range bibEntries {
        bibEntry := bibEntries[i]
        if !slices.Contains(mdEntries, bibEntry) {
            logWarning("BIB entry", fmt.Sprintf("\"%s\"", bibEntry), "not found in MD entries.")
        }
    }

	// Each MD file must have an associated BIB (ERR)
    for i := range mdEntries {
        mdEntry := mdEntries[i]
        if !slices.Contains(bibEntries, mdEntry) {
            logError("MD entry", fmt.Sprintf("\"%s\"", mdEntry), "not found in BIB entries.")
        }
    }

	// Each MD file must have the correct structure (checkMarkdownFile)
}

func checkMarkdownFile() {
	// MD file should have frontmatter containing:
	// - title: must correspond to BIB title
	// - authors: must correspond to BIB authors
	// - year: must correspond to BIB year
	// - aliases: must contain BIB title

	// Below frontmatter,
	// the first non-empty line should be
	// a top-level title corresponding to the BIB title

}
