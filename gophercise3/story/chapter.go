package story

import "errors"

type Story map[string]Arc

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

var story Story

const defaultArc = "intro"

func ChooseArc(arc string) (Arc, error) {
	if arc == "" {
		arc = defaultArc
	}

	if newArc, ok := story[arc]; ok {
		return newArc, nil
	}

	return Arc{}, errors.New("Arc not found")
}

func SetStory(s Story) {
	story = s
}
