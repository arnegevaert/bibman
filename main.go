package main

func main() {

}

func checkhealth() {
	// Get all BIB files

	// Get all PDF files

	// Get all MD files

	// Each PDF must have an associated BIB (ERR)

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
