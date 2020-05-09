package socialmedia

import "time"

type MoodState int

const (
	MoodStateNeutral MoodState = iota
	MoodStateHappy
	MoodStateSad
	MoodStateAngry
	MoodStateHopeful
	MoodStateThrilled
	MoodStateBored
	MoodStateShy
	MoodStateComical
	MoodStateOnCloudNine
)

type AuditableContent struct {
	TimeCreated  time.Time
	TimeModified time.Time
	CreatedBy    string
	ModifiedBy   string
}

type Post struct {
	AuditableContent
	Caption      string
	MessageBody  string
	URL          string
	ImageURI     string
	ThumbnailUri string
	keywords     []string
	Likers       []string
	AuthorMood   MoodState
}

var Moods map[string]MoodState

func init() {
	Moods = map[string]MoodState{
		"neutral":   MoodStateNeutral,
		"happy":     MoodStateHappy,
		"sad":       MoodStateSad,
		"angry":     MoodStateAngry,
		"hopeful":   MoodStateHopeful,
		"thrilled":  MoodStateThrilled,
		"bored":     MoodStateBored,
		"shy":       MoodStateShy,
		"comical":   MoodStateComical,
		"cloudnine": MoodStateOnCloudNine,
	}
}

func NewPost(username string, mood MoodState, caption string, messageBody string, url string, imageURI string, thumbnailURI string, keywords []string) *Post {
	auditableContent := AuditableContent{
		CreatedBy:   username,
		TimeCreated: time.Now(),
	}
	return &Post{
		Caption:          caption,
		MessageBody:      messageBody,
		URL:              url,
		ImageURI:         imageURI,
		ThumbnailUri:     thumbnailURI,
		AuthorMood:       mood,
		keywords:         keywords,
		AuditableContent: auditableContent,
	}
}
