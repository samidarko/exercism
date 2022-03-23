package grep

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	lineNumber      bool
	caseInsensitive bool
	filesMatching   bool
	allLineMatching bool
	multipleFiles   bool
	invertMatch     bool
}

func Search(pattern string, flags, files []string) []string {
	result := make([]string, 0)
	config := getConfig(flags, files)

	for _, file := range files {
		data, _ := os.ReadFile(file)
		for lineNumber, line := range strings.Split(string(data), "\n") {
			if line != "" && isMatching(config, pattern, line) {
				if config.filesMatching {
					result = append(result, file)
					break
				}
				result = append(result, line)
				position := len(result) - 1
				if config.lineNumber {
					result[position] = fmt.Sprint(lineNumber+1, ":", result[position])
				}
				if config.multipleFiles {
					result[position] = fmt.Sprint(file, ":", result[position])
				}
			}
		}
	}
	return result
}

func getConfig(flags, files []string) Config {
	config := Config{multipleFiles: len(files) > 1}
	for _, flag := range flags {
		switch flag {
		case "-n":
			config.lineNumber = true
		case "-i":
			config.caseInsensitive = true
		case "-l":
			config.filesMatching = true
		case "-x":
			config.allLineMatching = true
		case "-v":
			config.invertMatch = true
		}
	}
	return config
}

func linePreprocessing(config Config, line string) string {
	if config.caseInsensitive {
		return strings.ToLower(line)
	}
	return line
}

func isMatching(config Config, pattern, line string) bool {
	result := false
	pattern, line = linePreprocessing(config, pattern), linePreprocessing(config, line)
	if config.allLineMatching {
		result = pattern == line
	} else {
		result = strings.Contains(line, pattern)
	}

	if config.invertMatch {
		return !result
	}
	return result
}
