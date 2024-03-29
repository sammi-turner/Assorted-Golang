package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Runs shell commands and store the output as a string.
func Shell(command string) string {
	out, _ := exec.Command("bash", "-c", command).Output()
	return strings.TrimSpace(string(out))
}

// Takes the nth rune from a string, or returns an error if that index does not exist.
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

// Flushes stdout.
func FlushStdout() {
	writer := bufio.NewWriter(os.Stdout)
	writer.Flush()
}

// Prompts the user with a given string and returns their input.
func UserInput(s string) string {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSuffix(userInput, "\n")
}

// Checks if the OS is Unix-based.
func IsUnix() bool {
	goos := runtime.GOOS
	switch goos {
	case "darwin", "freebsd", "linux", "openbsd", "solaris":
		return true
	default:
		return false
	}
}

// Flushes stdout and clears the console window.
func ClearConsole() {
	FlushStdout()
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Returns true if the named file exists in the working directory.
func FileExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// Reads the contents of a file into a string.
func ReadFile(file string) (string, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

// Writes the contents of a string to a file.
func WriteFile(file string, text string) error {
	err := os.WriteFile(file, []byte(text), 0644)
	if err != nil {
		return err
	}
	return nil
}

// Append the contents of a string to a file.
func AppendFile(file string, text string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.WriteString(f, text)
	if err != nil {
		return err
	}

	return nil
}

// Deleted a file from the working directory.
func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

// NewRNG returns a new pseudo-random number generator seeded with the current time
func NewRNG() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Returns a slice of strings with the empty strings removed.
func RemoveEmptyStrings(sos []string) []string {
	var result []string
	for _, str := range sos {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

// Returns the number of instances of a rune that occur within a string.
func CountRuneInstances(s string, r rune) int {
	return strings.Count(s, string(r))
}

// Returns either the nth string from a slice of strings, or an error.
func GetNthString(strings []string, n int) (string, error) {
	if n < 0 || n >= len(strings) {
		return "", fmt.Errorf("index out of range. The length of the slice is %d", len(strings))
	}
	return strings[n], nil
}

// Returns the number of instances of a string in a slice of strings.
func CountStringInstances(sos []string, s string) int {
	count := 0
	for _, item := range sos {
		if item == s {
			count++
		}
	}
	return count
}

// Returns the nth portion of a string, as delimited by a rune, or an empty string.
func NthDelimitedString(s string, r rune, n int) string {
	split := strings.Split(s, string(r))
	if n < len(split) && n >= 0 {
		return split[n]
	}
	return ""
}

// Returns either an integer or an error message.
func ToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("'%s' is not a valid integer", s)
	}
	return i, nil
}

// Returns true if the string represents an integer, or false if it does not.
func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Returns either a float64 or an error message.
func ToFloat64(s string) (float64, error) {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting string to float64: %w", err)
	}
	return f, nil
}

// Returns true if the string represents a float64, or false if it does not.
func IsFloat64(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Returns the number of (possibly empty) substrings in a string, as delimited by a rune.
// A non-empty string without the delimiter will return 1.
// The empty string will return 0.
func SubstringCount(s string, r rune) int {
	if s == "" {
		return 0
	}
	return strings.Count(s, string(r)) + 1
}

// Prints each string in a slice on a new line.
func PrintStrings(sos []string) {
	for _, s := range sos {
		fmt.Println(s)
	}
}

// Creates a slice of strings from a parent string and a delimiter rune.
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
