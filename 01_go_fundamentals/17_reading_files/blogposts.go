package _7_reading_files

import (
	"io/fs"
	"log"
)

// NewPostsFromFS returns a collection of blog posts parsed from a file system.
// It reads all files in the current directory of the provided file system (fs.FS),
// attempts to parse each file as a blog post, and returns a slice of Post instances.
// If any file fails to parse, the function will return an error immediately.
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

// getPost retrieves a single Post from a file system based on the provided directory entry (fs.DirEntry).
// It opens the file, parses it into a Post using newPost, and returns the Post instance.
func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}
	defer func(postFile fs.File) {
		err := postFile.Close()
		if err != nil {
			log.Fatalf("Error closing post file: %v", err)
		}
	}(postFile)

	return newPost(postFile)
}
