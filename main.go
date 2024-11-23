package main

import (
	"log"
	"os"
	"path"
	"slices"
	"strings"
)

func main() {
	checkHealth()
}

func getEntries(dir string) ([]string){
    filenames, err := os.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Found", len(filenames), "files in", dir)
    result := make([]string, len(filenames))
    for i := range(len(filenames)) {
        result[i] = strings.Split(string(filenames[i].Name()), ".")[0]
    }
    return result
}

func checkHealth() {
	const BASE_PATH = "/home/arne/Work/Buckets/ZK Zettelkasten/Bibman/"
    log.Println("Base path:", BASE_PATH)

    BIB_PATH := path.Join(BASE_PATH, "BIB")
	var PDF_PATH = path.Join(BASE_PATH, "PDF")
	// var MD_PATH = path.Join(BASE_PATH, "MD")

	// Get all BIB, PDF and MD files
    bibEntries := getEntries(BIB_PATH)
    // mdEntries := getEntries(MD_PATH)
    pdfEntries := getEntries(PDF_PATH)

	// Each PDF must have an associated BIB (ERR)
    for i := range pdfEntries {
        pdfEntry := pdfEntries[i]
        if !slices.Contains(bibEntries, pdfEntry) {
            log.Println("ERROR:", pdfEntry, "not found in BIB entries.")
        }
    }

	// Each BIB may have an associated PDF (WARN)

	// Each MD file must have an associated BIB (ERR)

	// Each MD file may have an associated PDF (WARN)

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
