package aoc2024

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

func dayTwo(input []byte) (int, error) {
	reports, err := readDayTwoInput(input)
	if err != nil {
		return 0, fmt.Errorf("could not read reports: %w", err)
	}

	result := 0
	for _, report := range reports {
		if err := validateDayTwoReport(report); err != nil {
			slog.Info("Report is not safe", slog.Any("error", err), slog.Any("report", report))
			continue
		}

		slog.Info("Report is safe", slog.Any("report", report))
		result++
	}

	return result, nil
}

func readDayTwoInput(input []byte) ([][]int, error) {
	result := [][]int{}

	scanner := bufio.NewScanner(bytes.NewReader(input))

	for scanner.Scan() {
		report := []int{}
		for _, part := range strings.Split(scanner.Text(), " ") {
			if part == "" {
				continue
			}

			level, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("could not convert part to int: %w", err)
			}

			report = append(report, level)
		}

		result = append(result, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not read input: %w", err)
	}

	return result, nil
}

func validateDayTwoReport(report []int) error {
	isDecreasing, isIncreasing := false, false

	for i := 0; i < len(report)-1; i++ {
		currentLevel, nextLevel := report[i], report[i+1]

		switch {
		case isDecreasing && currentLevel < nextLevel:
			return fmt.Errorf("start increasing between '%d' and '%d'", currentLevel, nextLevel)
		case isIncreasing && currentLevel > nextLevel:
			return fmt.Errorf("start decreasing between '%d' and '%d'", currentLevel, nextLevel)
		case currentLevel > nextLevel:
			isDecreasing = true
		case currentLevel < nextLevel:
			isIncreasing = true
		}

		if distance := int(math.Abs(float64(currentLevel - nextLevel))); distance < 1 || distance > 3 {
			return fmt.Errorf("invalid distance of '%d' between '%d' and '%d'", distance, currentLevel, nextLevel)
		}
	}

	return nil
}
