package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	slide "github.com/tmickleydoyle/moonwalk/slide"
	"github.com/fatih/color"
)

var (
	dir        string
	path       string
	extension  string
	search     string
	format     string
	showSize   bool
	summary    bool
	maxDepth   int
)

// FileInfo holds information about a file for output formatting
type FileInfo struct {
	Path      string    `json:"path"`
	Type      string    `json:"type"`
	Size      int64     `json:"size,omitempty"`
	ModTime   time.Time `json:"modified_time,omitempty"`
	Depth     int       `json:"depth"`
	SearchHit bool      `json:"search_hit,omitempty"`
	Matches   []string  `json:"matches,omitempty"`
}

// Statistics holds summary information about the walk
type Statistics struct {
	TotalFiles      int     `json:"total_files"`
	TotalDirs       int     `json:"total_directories"`
	TotalSize       int64   `json:"total_size_bytes"`
	AverageFileSize int64   `json:"average_file_size_bytes"`
	LargestFile     string  `json:"largest_file"`
	LargestFileSize int64   `json:"largest_file_size_bytes"`
	ExtensionCount  map[string]int `json:"extension_count"`
	SearchMatches   int     `json:"search_matches"`
}

func main() {
	flag.StringVar(&dir, "path", "", "starting point directory")
	flag.StringVar(&extension, "ext", "", "filter files by extension (e.g., .go, .txt)")
	flag.StringVar(&search, "search", "", "search for text within files")
	flag.StringVar(&format, "format", "text", "output format (text, json, csv)")
	flag.BoolVar(&showSize, "size", false, "show file sizes")
	flag.BoolVar(&summary, "summary", false, "show summary statistics")
	flag.IntVar(&maxDepth, "depth", -1, "maximum directory depth (-1 for unlimited)")
	flag.Parse()

	if dir != "" {
		path = dir
	} else {
		path, _ = os.Getwd()
	}

	// Add a dot to extension if needed and not empty
	if extension != "" && !strings.HasPrefix(extension, ".") {
		extension = "." + extension
	}

	// Create slices to store files and directories
	var files []FileInfo
	stats := Statistics{
		ExtensionCount: make(map[string]int),
	}

	err := slide.Slide(path, func(path string, info os.FileInfo, depth int, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		// Create file info object
		fileInfo := FileInfo{
			Path:    path,
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

			// Update largest file stats
			if info.Size() > stats.LargestFileSize {
				stats.LargestFile = path
				stats.LargestFileSize = info.Size()
			}

			// Count file extensions
			ext := strings.ToLower(strings.TrimPrefix(filepath.Ext(path), "."))
			if ext != "" {
				stats.ExtensionCount[ext]++
			}

			// Search within files if requested
			if search != "" && !info.IsDir() {
				content, err := ioutil.ReadFile(path)
				if err == nil {
					if strings.Contains(string(content), search) {
						fileInfo.SearchHit = true
						stats.SearchMatches++

						// Find matching lines for context
						matches := []string{}
						scanner := bufio.NewScanner(strings.NewReader(string(content)))
						lineNum := 0
						for scanner.Scan() {
							lineNum++
							line := scanner.Text()
							if strings.Contains(line, search) {
								matches = append(matches, fmt.Sprintf("Line %d: %s", lineNum, line))
								if len(matches) >= 5 { // Limit to 5 matches per file
									break
								}
							}
						}
						fileInfo.Matches = matches
					}
				}
			}
		}

		// Add to our results array
		files = append(files, fileInfo)

		return nil
	}, extension, maxDepth)

	if err != nil {
		log.Println(err)
		return
	}

	// Calculate average file size if there are files
	if stats.TotalFiles > 0 {
		stats.AverageFileSize = stats.TotalSize / int64(stats.TotalFiles)
	}

	// Output based on format
	switch format {
	case "json":
		outputJSON(files, stats)
	case "csv":
		outputCSV(files)
	default:
		outputText(files, stats)
	}
}

func outputJSON(files []FileInfo, stats Statistics) {
	var result interface{}

	if summary {
		result = stats
	} else {
		result = files
	}

	jsonData, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func outputCSV(files []FileInfo) {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// Write header
	header := []string{"Path", "Type", "Depth"}
	if showSize {
		header = append(header, "Size")
	}
	if search != "" {
		header = append(header, "SearchMatch")
	}
	writer.Write(header)

	// Write data
	for _, file := range files {
		// Skip directories if we're only showing search results
		if search != "" && file.Type == "directory" {
			continue
		}
		
		// Skip non-matches if we're searching
		if search != "" && !file.SearchHit {
			continue
		}

		row := []string{file.Path, file.Type, fmt.Sprintf("%d", file.Depth)}
		if showSize {
			row = append(row, fmt.Sprintf("%d", file.Size))
		}
		if search != "" {
			if file.SearchHit {
				row = append(row, "true")
			} else {
				row = append(row, "false")
			}
		}
		writer.Write(row)
	}
}

func outputText(files []FileInfo, stats Statistics) {
	// Print summary if requested
	if summary {
		fmt.Printf("\nSummary:\n")
		fmt.Printf("Total Files: %d\n", stats.TotalFiles)
		fmt.Printf("Total Directories: %d\n", stats.TotalDirs)
		fmt.Printf("Total Size: %s\n", formatSize(stats.TotalSize))
		fmt.Printf("Average File Size: %s\n", formatSize(stats.AverageFileSize))
		fmt.Printf("Largest File: %s (%s)\n", stats.LargestFile, formatSize(stats.LargestFileSize))
		
		// Print extension counts
		fmt.Println("\nFile Extensions:")
		for ext, count := range stats.ExtensionCount {
			fmt.Printf(".%s: %d files\n", ext, count)
		}
		
		if search != "" {
			fmt.Printf("\nSearch Results:\n")
			fmt.Printf("Found \"%s\" in %d files\n", search, stats.SearchMatches)
		}
		return
	}

	// Print files
	for _, file := range files {
		// Skip directories if we're only showing search results
		if search != "" && file.Type == "directory" {
			continue
		}
		
		// Skip non-matches if we're searching
		if search != "" && !file.SearchHit {
			continue
		}

		// Display with colors based on type
		if file.Type == "directory" {
			color.Cyan("[Directory] %s", file.Path)
		} else {
			output := fmt.Sprintf("[File] %s", file.Path)
			if showSize {
				output += fmt.Sprintf(" (%s)", formatSize(file.Size))
			}
			color.Green(output)

			// Show matches if this is a search hit
			if file.SearchHit {
				color.Yellow("   Matches for \"%s\":", search)
				for _, match := range file.Matches {
					color.Yellow("   - %s", match)
				}
			}
		}
	}
}

// formatSize formats a size in bytes to a human-readable string
func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}
