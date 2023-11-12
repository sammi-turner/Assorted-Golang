package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const dq = "\""

func Shell(command string) string {
	out, _ := exec.Command("bash", "-c", command).Output()
	return strings.TrimSpace(string(out))
}

func NthRune(s string, n int) (rune, error) {
	count := 0
	for _, runeValue := range s {
		if count == n {
			return runeValue, nil
		}
		count++
	}
	return 0, errors.New("index out of bounds")
}

func IsSQLiteInstalled() (bool, error) {
	s := Shell("sqlite3 --version")
	c, _ := NthRune(s, 0)
	if c == '3' {
		return true, nil
	}
	return false, nil
}

func DbExists(name string) bool {
	dbName := name + ".db"
	_, err := os.Stat(dbName)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func CreateKeyValueTable(t string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sCREATE TABLE %s (key TEXT, value TEXT);%s", t, dq, t, dq)
	s := Shell(command)
	return s
}

func InsertKeyValuePair(t, k, v string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sINSERT INTO %s (key, value) VALUES ('%s', '%s');%s", t, dq, t, k, v, dq)
	s := Shell(command)
	return s
}

func DeleteKeyValuePair(t, k string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sDELETE FROM %s WHERE key='%s';%s", t, dq, t, k, dq)
	s := Shell(command)
	return s
}

func SelectAllPairs(t string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sSELECT * FROM %s;%s", t, dq, t, dq)
	s := Shell(command)
	return s
}

func SelectAllKeys(t string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sSELECT key FROM %s;%s", t, dq, t, dq)
	s := Shell(command)
	return s
}

func SelectAllValues(t string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sSELECT value FROM %s;%s", t, dq, t, dq)
	s := Shell(command)
	return s
}

func SelectRowFromKey(t, k string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sSELECT * FROM %s WHERE key='%s';%s", t, dq, t, k, dq)
	s := Shell(command)
	return s
}

func SelectValueFromKey(t, k string) string {
	command := fmt.Sprintf("sqlite3 %s.db %sSELECT value FROM %s WHERE key='%s';%s", t, dq, t, k, dq)
	s := Shell(command)
	return s
}

func CountStringInstances(sos []string, s string) int {
	count := 0
	for _, item := range sos {
		if item == s {
			count++
		}
	}
	return count
}

func SplitString(s string, d rune) []string {
	var result []string
	start := 0

	for i, r := range s {
		if r == d {
			result = append(result, s[start:i])
			start = i + 1
		}
	}

	result = append(result, s[start:])
	return result
}

func TableContainsKey(t, k string) bool {
	s := SelectAllPairs(t)
	split := SplitString(s, '\n')
	n := CountStringInstances(split, k)
	return n == 1
}
