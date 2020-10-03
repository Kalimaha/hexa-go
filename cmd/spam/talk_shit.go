package spam

import "fmt"

func TalkShit() {
	fmt.Printf(SuggestShit())
}

func SuggestShit() string {
	return "Strunz!"
}
