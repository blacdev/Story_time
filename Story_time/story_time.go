package stry

// Takes in a json file in bytes and maps each object to a struct

// Read file
// map byte


type( 
	Story struct {
		Title string `json:title`
		Story_text string `json:story`
		Options struct {
			Text string `json:text`
			Arc string `json:arc`
		}

	}

)

var StoryWeb []Story


