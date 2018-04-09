package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func parseEnvLine(line string) (string, string) {
	parts := strings.Split(line, "=")
	key := parts[0]
	value := strings.Join(parts[1:], "=")
	return key, value
}

func getEnvMapByLines(lines []string) map[string]string {
	m := map[string]string{}
	for _, line := range lines {
		key, value := parseEnvLine(line)
		m[key] = value
	}
	return m
}

func main() {
	env := getEnvMapByLines(os.Environ())
	envJson, err := json.MarshalIndent(env, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(envJson))
}
