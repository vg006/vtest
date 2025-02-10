package asset

import "github.com/charmbracelet/lipgloss"

// Colours
var (
	LightBlue  = lipgloss.AdaptiveColor{Light: "#35FCDC", Dark: "#00d0ff"}
	DarkBlue   = lipgloss.AdaptiveColor{Light: "#04FAD3", Dark: "#4BCBE7"}
	Black      = lipgloss.AdaptiveColor{Light: "#000", Dark: "#000"}
	Indigo     = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	Green      = lipgloss.AdaptiveColor{Light: "#0dff00", Dark: "#0dff00"}
	Orange     = lipgloss.AdaptiveColor{Light: "#FFA500", Dark: "#FFA500"}
	Red        = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}
	WhiteSmoke = lipgloss.AdaptiveColor{Light: "#F5F5F5", Dark: "#F5F5F5"}
)

// Text Styles
var (
	Text = lipgloss.NewStyle().
		PaddingLeft(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(LightBlue).
		Foreground(LightBlue)

	Msg = lipgloss.NewStyle().
		PaddingLeft(1).
		Foreground(WhiteSmoke)

	Info = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(WhiteSmoke).
		Foreground(Black).
		Background(WhiteSmoke).
		Bold(true).
		Render("INFO")

	Done = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(Green).
		Foreground(Black).
		Background(Green).
		Bold(true).
		Render("DONE")

	Error = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(Orange).
		Foreground(Black).
		Background(Orange).
		Bold(true).
		Render("ERROR")

	Exit = lipgloss.NewStyle().
		PaddingLeft(1).
		PaddingRight(1).
		BorderStyle(lipgloss.ThickBorder()).
		BorderLeft(true).
		BorderForeground(Red).
		Foreground(Black).
		Background(Red).
		Bold(true).
		Render("EXIT")
)

// Logo
var VtestLogo = lipgloss.NewStyle().
	Foreground(LightBlue).
	PaddingLeft(1).
	Bold(true).
	BorderStyle(lipgloss.ThickBorder()).
	BorderLeft(true).
	BorderForeground(DarkBlue).
	Render(`
     ___  ___  __  ___
\  /  |  |__  /__   |
 \/   |  |___  __/  |
`)
