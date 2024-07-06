package _7_reading_files

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

// Post represents a post on a blog
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

// newPost parses an input stream (postBody) and returns a new Post instance and any encountered errors.
// We're using an io.Reader instead of fs.File to make it loosely coupled so that it can read from anywhere.
func newPost(postBody io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postBody)

	// Helper function to read a line and strip the prefix (tagName)
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	// Construct a new Post instance using parsed metadata and body content
	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

// readBody reads the remaining lines from scanner and concatenates them into a single string, excluding the first line.
func readBody(scanner *bufio.Scanner) string {
	// Skip the first line (already read in newPost)
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		_, _ = fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
