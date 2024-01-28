package storage

import (
	"encoding/json"
	"io"
	"os"

	"github.com/walterdl/gophercises3/story"
)

func StoryFromFile(path string) (story.Story, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var result story.Story
	err = json.Unmarshal(content, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
