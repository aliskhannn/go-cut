package cut

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Config struct {
	Fields    []int
	Delimiter string
	Separated bool
}

// parseFields converts 1-based fields slice to map[int]bool for quick lookup
func parseFields(fields []int) map[int]bool {
	m := make(map[int]bool)

	for _, f := range fields {
		if f > 0 {
			m[f-1] = true // internally 0-based
		}
	}

	return m
}

// Process reads from r line by line, splits by delimiter and writes selected fields to w.
func Process(r io.Reader, w io.Writer, cfg Config) error {
	fieldMap := parseFields(cfg.Fields)
	scanner := bufio.NewScanner(r)
	bw := bufio.NewWriter(w)
	defer func() {
		_ = bw.Flush()
	}()

	for scanner.Scan() {
		line := scanner.Text()

		if cfg.Separated && !strings.Contains(line, cfg.Delimiter) {
			continue
		}

		columns := strings.Split(line, cfg.Delimiter)
		var selected []string

		for i, col := range columns {
			if fieldMap[i] {
				selected = append(selected, col)
			}
		}

		if len(selected) > 0 {
			_, _ = bw.WriteString(strings.Join(selected, cfg.Delimiter) + "\n")
		}
	}

	return scanner.Err()
}

// ParseFieldArg parses a string like "1,3-5,7" into []int{1,3,4,5,7}
func ParseFieldArg(arg string) ([]int, error) {
	var result []int
	parts := strings.Split(arg, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			bounds := strings.SplitN(part, "-", 2)
			if len(bounds) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := strconv.Atoi(bounds[0])
			if err != nil {
				return nil, err
			}

			end, err := strconv.Atoi(bounds[1])
			if err != nil {
				return nil, err
			}

			for i := start; i <= end; i++ {
				result = append(result, i)
			}
		} else {
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}

			result = append(result, n)
		}
	}

	return result, nil
}
