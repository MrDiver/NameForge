package main

import (
	"flag"
	"fmt"
	"strings"

	forge "CodingGhost.de/nameforge"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	showHelp := flag.Bool("h", false, "")
	sentence := flag.String("s", "", "")
	listUniverse := flag.Bool("lu", false, "")
	universe := flag.String("u", "STARWARS", "")
	listLanguage := flag.Bool("ll", false, "")
	language := flag.String("l", "EN", "")
	listAll := flag.Bool("la", false, "")

	split := flag.String("split", " ", "")
	join := flag.String("join", "-", "")
	number := flag.Int("n", 1, "")

	flag.Parse()
	flag.Usage = func() {
		border := lipgloss.NewStyle().Border(lipgloss.ThickBorder()).Width(60)
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA")).PaddingTop(1).PaddingLeft(2).PaddingBottom(1)
		title := style.Align(lipgloss.Center).Width(60)

		style = style.UnsetPadding().PaddingTop(1).PaddingLeft(3)
		options := lipgloss.JoinVertical(lipgloss.Left, style.Render("-h --help"), style.Render("-s"), style.Render("\n\n-lu"), style.Render("-u"), style.Render("\n-ll"), style.Render("-l"), style.Render("\n-la"), style.Render("-split"), style.Render("\n-join"), style.Render("\n-n"))
		style = style.Width(50).Faint(true)
		usages := lipgloss.JoinVertical(lipgloss.Left, style.Render("Displays this help message"), style.Render("The format used for forging your sentence containing a comma separated list of tags N,ADJ,V,ADV,NAME"), style.Render("Displays the available universes"), style.Render("Sets the universe to be used \n\tdefault: "+*universe), style.Render("Displays the available languages"), style.Render("Sets the language to be used\n\tdefault: EN"), style.Render("List all available of everything"), style.Render("String used for splitting words\nand names which are then rejoined with join"), style.Render("String used for joining all split words together"), style.Render("Number of names to be generated"))
		helps := lipgloss.JoinHorizontal(lipgloss.Left, options, usages)

		titlestring := border.Render(title.Bold(true).Render("🧰 Welcome to NameForge 🔨"), title.Bold(false).Faint(true).Render("Your everyday tool for producing random names"), helps)

		fmt.Println(titlestring)
	}

	if flag.NFlag() < 1 || *showHelp {
		flag.Usage()
		return
	}
	if *listUniverse || *listAll {
		keys := make([]string, 0, len(forge.AvailableUniverses()))
		for k := range forge.AvailableUniverses() {
			keys = append(keys, k)
		}
		style := lipgloss.NewStyle()
		fmt.Println(style.Bold(true).PaddingLeft(2).Render("Universes:"))
		fmt.Println(style.PaddingLeft(4).Render(lipgloss.JoinHorizontal(0.2, keys...)))
	}

	if *listLanguage || *listAll {
		keys := make([]string, 0, len(forge.AvailableLanguages()))
		for k := range forge.AvailableLanguages() {
			keys = append(keys, k)
		}
		style := lipgloss.NewStyle()
		fmt.Println(style.Bold(true).PaddingLeft(2).Render("Languages:"))
		fmt.Println(style.PaddingLeft(4).Render(lipgloss.JoinHorizontal(0.2, keys...)))
	}

	if val, ok := forge.AvailableUniverses()[*universe]; ok {
		forge.SetUniverse(val)
	}

	if val, ok := forge.AvailableLanguages()[*language]; ok {
		forge.SetLanguage(val)
	}

	splitfunc := func(s string) []string { return strings.Split(s, *split) }
	forge.SetSplit(splitfunc)

	joinfunc := func(s []string) string { return strings.Join(s, *join) }
	forge.SetJoin(joinfunc)

	if *sentence != "" {
		for _ = range *number {
			fmt.Println(forge.SentenceC(*sentence))
		}
	}
}
