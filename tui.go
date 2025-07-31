package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	slide "github.com/tmickleydoyle/moonwalk/slide"
)

type tuiModel struct {
	files          []FileInfo
	stats          Statistics
	cursor         int
	viewport       int
	height         int
	width          int
	loading        bool
	done           bool
	showDetails    bool
	searchMode     bool
	searchQuery    string
	filtered       []FileInfo
	showPath       bool
	showHelp       bool
	lastScanPath   string
	errorMsg       string
}

type scanCompleteMsg struct {
	files []FileInfo
	stats Statistics
}

type scanProgressMsg struct {
	path string
}

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1)

	statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1)

	fileStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575"))

	dirStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

	selectedStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#874BFD")).
			Foreground(lipgloss.Color("#FAFAFA"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262"))

	searchStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#FF5F87")).
			Padding(0, 1)
)

func initialTUIModel() tuiModel {
	return tuiModel{
		files:        []FileInfo{},
		stats:        Statistics{ExtensionCount: make(map[string]int)},
		loading:      true,
		filtered:     []FileInfo{},
		showPath:     true,
		showHelp:     false,
		lastScanPath: path,
	}
}

func (m tuiModel) Init() tea.Cmd {
	return tea.Batch(
		scanFiles(),
		tea.EnterAltScreen,
	)
}

func scanFiles() tea.Cmd {
	return func() tea.Msg {
		var files []FileInfo
		stats := Statistics{
			ExtensionCount: make(map[string]int),
		}

		err := slide.Slide(path, func(filePath string, info os.FileInfo, depth int, err error) error {
			if err != nil {
				return err
			}

			fileInfo := FileInfo{
				Path:    filePath,
				Depth:   depth,
				ModTime: info.ModTime(),
			}

			if info.IsDir() {
				fileInfo.Type = "directory"
				stats.TotalDirs++
			} else {
				fileInfo.Type = "file"
				fileInfo.Size = info.Size()
				stats.TotalFiles++
				stats.TotalSize += info.Size()

				if info.Size() > stats.LargestFileSize {
					stats.LargestFile = filePath
					stats.LargestFileSize = info.Size()
				}

				ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(filePath), "."))
				if ext != "" {
					stats.ExtensionCount[ext]++
				}

				if search != "" {
					content, err := os.ReadFile(filePath)
					if err == nil && strings.Contains(string(content), search) {
						fileInfo.SearchHit = true
						stats.SearchMatches++
					}
				}
			}

			files = append(files, fileInfo)
			return nil
		}, extension, maxDepth)

		if err != nil {
			return tea.Quit
		}

		if stats.TotalFiles > 0 {
			stats.AverageFileSize = stats.TotalSize / int64(stats.TotalFiles)
		}

		return scanCompleteMsg{files: files, stats: stats}
	}
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, nil

	case scanCompleteMsg:
		m.files = msg.files
		m.stats = msg.stats
		m.loading = false
		m.filtered = m.getFilteredFiles()
		return m, nil

	case tea.KeyMsg:
		if m.searchMode {
			switch msg.String() {
			case "enter":
				m.searchMode = false
				m.filtered = m.getFilteredFiles()
				m.cursor = 0
				m.viewport = 0
			case "esc":
				m.searchMode = false
				m.searchQuery = ""
				m.filtered = m.getFilteredFiles()
				m.cursor = 0
				m.viewport = 0
			case "backspace":
				if len(m.searchQuery) > 0 {
					m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
				}
			default:
				if len(msg.String()) == 1 {
					m.searchQuery += msg.String()
				}
			}
			return m, nil
		}

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				if m.cursor < m.viewport {
					m.viewport = m.cursor
				}
			}

		case "down", "j":
			maxItems := len(m.filtered) - 1
			if m.cursor < maxItems {
				m.cursor++
				visibleItems := m.height - 6
				if m.cursor >= m.viewport+visibleItems {
					m.viewport = m.cursor - visibleItems + 1
				}
			}

		case "home", "g":
			m.cursor = 0
			m.viewport = 0

		case "end", "G":
			m.cursor = len(m.filtered) - 1
			visibleItems := m.height - 6
			if len(m.filtered) > visibleItems {
				m.viewport = len(m.filtered) - visibleItems
			}

		case "d":
			m.showDetails = !m.showDetails

		case "p":
			m.showPath = !m.showPath

		case "h", "?":
			m.showHelp = !m.showHelp

		case "/":
			m.searchMode = true
			m.searchQuery = ""

		case "r":
			m.loading = true
			m.files = []FileInfo{}
			m.filtered = []FileInfo{}
			return m, scanFiles()
		}
	}

	return m, nil
}

func (m tuiModel) getFilteredFiles() []FileInfo {
	if m.searchQuery == "" && search == "" {
		return m.files
	}

	var filtered []FileInfo
	query := strings.ToLower(m.searchQuery)

	for _, file := range m.files {
		if search != "" && !file.SearchHit {
			continue
		}

		if m.searchQuery != "" {
			fileName := strings.ToLower(filepath.Base(file.Path))
			if !strings.Contains(fileName, query) {
				continue
			}
		}

		filtered = append(filtered, file)
	}

	return filtered
}

func (m tuiModel) View() string {
	if m.loading {
		loadingText := fmt.Sprintf("🌙 Scanning files from %s...", m.lastScanPath)
		return fmt.Sprintf("\n%s\n\n%s\n\n%s",
			titleStyle.Render("MoonWalk"),
			loadingText,
			helpStyle.Render("Press q to quit"))
	}

	if m.errorMsg != "" {
		return fmt.Sprintf("\n%s\n\n❌ Error: %s\n\n%s",
			titleStyle.Render("MoonWalk"),
			m.errorMsg,
			helpStyle.Render("Press r to retry or q to quit"))
	}

	var s strings.Builder

	// Header
	title := "MoonWalk - Directory Scanner"
	if search != "" {
		title += fmt.Sprintf(" (searching: %s)", search)
	}
	s.WriteString(titleStyle.Render(title))
	s.WriteString("\n")
	
	// Show current path
	if m.showPath {
		pathInfo := fmt.Sprintf("📁 %s", m.lastScanPath)
		s.WriteString(helpStyle.Render(pathInfo))
		s.WriteString("\n")
	}
	s.WriteString("\n")

	// Search bar
	if m.searchMode {
		searchBar := fmt.Sprintf("Search: %s█", m.searchQuery)
		s.WriteString(searchStyle.Render(searchBar))
		s.WriteString("\n\n")
	}

	// Stats
	if m.showDetails {
		s.WriteString(fmt.Sprintf("Files: %d | Dirs: %d | Size: %s",
			m.stats.TotalFiles,
			m.stats.TotalDirs,
			formatSize(m.stats.TotalSize)))
		if search != "" {
			s.WriteString(fmt.Sprintf(" | Matches: %d", m.stats.SearchMatches))
		}
		s.WriteString("\n\n")
	}

	// File list
	visibleItems := m.height - 6
	if m.showDetails {
		visibleItems -= 2
	}
	if m.searchMode {
		visibleItems -= 2
	}

	start := m.viewport
	end := start + visibleItems
	if end > len(m.filtered) {
		end = len(m.filtered)
	}

	for i := start; i < end; i++ {
		file := m.filtered[i]
		line := m.renderFileItem(file, i == m.cursor)
		s.WriteString(line)
		s.WriteString("\n")
	}

	// Footer/Help
	s.WriteString("\n")
	if m.showHelp {
		s.WriteString(helpStyle.Render("=== HELP ==="))
		s.WriteString("\n")
		s.WriteString(helpStyle.Render("Navigation: ↑/↓ or j/k | g/G: top/bottom"))
		s.WriteString("\n")
		s.WriteString(helpStyle.Render("Actions: /: search | d: details | p: path | r: refresh"))
		s.WriteString("\n")
		s.WriteString(helpStyle.Render("Other: h/?: toggle help | q: quit"))
		s.WriteString("\n")
	} else if m.searchMode {
		s.WriteString(helpStyle.Render("Enter: confirm search | Esc: cancel | Type to search"))
	} else {
		help := "↑/↓: navigate | d: details | p: path | /: search | h: help | r: refresh | q: quit"
		s.WriteString(helpStyle.Render(help))
	}

	return s.String()
}

func (m tuiModel) renderFileItem(file FileInfo, selected bool) string {
	var icon, name, details string

	// Better icons based on file type and extensions
	if file.Type == "directory" {
		icon = "📁"
		name = dirStyle.Render(filepath.Base(file.Path))
	} else {
		// Choose icon based on file extension
		ext := strings.ToLower(filepath.Ext(file.Path))
		switch ext {
		case ".go":
			icon = "🐹"
		case ".js", ".ts":
			icon = "🟨"
		case ".py":
			icon = "🐍"
		case ".md":
			icon = "📝"
		case ".json":
			icon = "🔧"
		case ".yml", ".yaml":
			icon = "⚙️"
		case ".txt":
			icon = "📄"
		case ".png", ".jpg", ".jpeg", ".gif":
			icon = "🖼️"
		default:
			icon = "📄"
		}
		
		name = fileStyle.Render(filepath.Base(file.Path))
		if showSize || m.showDetails {
			details = fmt.Sprintf(" (%s)", formatSize(file.Size))
		}
	}

	// Search hit indicator overrides icon
	if file.SearchHit {
		icon = "🔍"
	}

	// Depth indentation
	indent := strings.Repeat("  ", file.Depth)

	// Show full path if requested
	displayName := name
	if m.showPath && selected {
		displayName = fileStyle.Render(file.Path)
	}

	line := fmt.Sprintf("%s%s %s%s", indent, icon, displayName, details)

	if selected {
		return selectedStyle.Render(line)
	}
	return line
}

func runTUI() error {
	p := tea.NewProgram(
		initialTUIModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	_, err := p.Run()
	return err
}