package main

// https://github.com/charmbracelet/bubbletea/tree/master/examples/list-fancy

import (
	"fmt"
	"os"
	"strings"
    tea "github.com/charmbracelet/bubbletea"
)

type entry struct {
    title string
    author string
    year int16
}

type model struct {
    entries map[string]entry
    cursor string
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
    switch msg.String() {
        case "q":
        return m, tea.Quit
        }
    }
    return m, nil
}

func (m model) View() string{
    return "This is a view\n"
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

func main() {
    if _, err := tea.NewProgram(initializeModel(), tea.WithAltScreen()).Run(); err != nil {
        fmt.Println("Error running program:", err)
        os.Exit(1)
    }
}
