package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	deniedPenalty   = 20 // Минуты штрафа за неудачные попытки
	secondsInMinute = 60
)

// парс времени
func parseTime(timeStr string, startSeconds int64) int64 {
	parts := strings.Split(timeStr, ":")
	hh, _ := strconv.Atoi(parts[0])
	mm, _ := strconv.Atoi(parts[1])
	ss, _ := strconv.Atoi(parts[2])
	return int64(hh)*3600 + int64(mm)*60 + int64(ss) - startSeconds
}

// Проверка на валидность результата
func isValidResult(result string) bool {
	return result == "ACCESSED" || result == "DENIED" || result == "FORBIDEN" || result == "PONG"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// начало чтения времени
	scanner.Scan()
	startTime := scanner.Text()
	startParts := strings.Split(startTime, ":")
	startHours, _ := strconv.Atoi(startParts[0])
	startMinutes, _ := strconv.Atoi(startParts[1])
	startSeconds, _ := strconv.Atoi(startParts[2])
	startSecondsTotal := int64(startHours*3600 + startMinutes*60 + startSeconds)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	teams := make(map[string]map[string]int)
	penalties := make(map[string]int)
	serversAttempt := make(map[string]map[string]int)

	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()

		parts := strings.Split(line, " ")
		teamName := strings.Trim(parts[0], "\"")
		timeStr := parts[1]
		serverID := parts[2]
		result := parts[3]

		if !isValidResult(result) {
			continue
		}

		requestTime := parseTime(timeStr, startSecondsTotal)

		if _, exists := teams[teamName]; !exists {
			teams[teamName] = make(map[string]int)
			serversAttempt[teamName] = make(map[string]int)
			penalties[teamName] = 0
		}

		if result == "ACCESSED" {
			if _, ok := teams[teamName][serverID]; !ok {
				teams[teamName][serverID] = 1
				penalties[teamName] += int(requestTime) / secondsInMinute

				serversAttempt[teamName][serverID] = 0
			}
		} else if result == "DENIED" || result == "FORBIDEN" {
			if _, ok := teams[teamName][serverID]; !ok {
				serversAttempt[teamName][serverID]++
				if serversAttempt[teamName][serverID] == 1 {
					penalties[teamName] += deniedPenalty
				}
			}
		}
	}
	// Формирование результатов
	results := make([][3]interface{}, 0, len(teams))
	for teamName, servers := range teams {
		numServers := len(servers)
		penalty := penalties[teamName]
		results = append(results, [3]interface{}{numServers, penalty, teamName})
	}
	// Сортировка результатов
	sort.SliceStable(results, func(i, j int) bool {
		if results[i][0] != results[j][0] {
			return results[i][0].(int) > results[j][0].(int)
		}
		if results[i][1] != results[j][1] {
			return results[i][1].(int) < results[j][1].(int)
		}
		return results[i][2].(string) < results[j][2].(string)
	})


	for rank, result := range results {
		fmt.Printf("%d \"%s\" %d %d\n", rank+1, result[2], result[0], result[1])
	}
}
