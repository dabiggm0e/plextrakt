package plex

// Plex webhook json structure
type Webhook struct {
	Event    string
	User     bool
	Owner    bool
	Account  Account
	Server   Server
	Player   Player
	Metadata Metadata
}

type Account struct {
	Id    int
	Thumb string
	Title string
}

type Server struct {
	Title string
	UUID  string
}

type Player struct {
	Local         bool
	PublicAddress string
	Title         string
	UUID          string
}

type Metadata struct {
	LibrarySectionType    string
	RatingKey             string
	Key                   string
	Guid                  string
	LibrarySectionTitle   string
	LibrarySectionID      int
	LibrarySectionKey     string
	Studio                string
	Mediatype             string `json:"type"`
	Title                 string
	ContentRating         string
	Summary               string
	Rating                float32
	AudienceRating        float32
	ViewCousnt            int
	LastViewedAt          int32
	Year                  int
	Tagline               string
	Thumb                 string
	Art                   string
	Duration              int
	OriginallyAvailableAt string
	AddedAt               int32
	UpdatedAt             int32
	AudienceRatingImage   string
	ChapterSource         string
	PrimaryExtraKey       string
	RatingImage           string
	Genre                 []Genre
	Director              []Director
	Writer                []Writer
	Producer              []Producer
	Country               []Country
	Collection            []Collection
	Role                  []Role
	Similar               []Similar
}

type Genre struct {
	Id     int
	Filter string
	Tag    string
	Count  int
}

type Director struct {
	Id     int
	Filter string
	Tag    string
}

type Writer struct {
	Id     int
	Filter string
	Tag    string
}

type Producer struct {
	Id     int
	Filter string
	Tag    string
}

type Country struct {
	Id     int
	Filter string
	Tag    string
	Count  int
}

type Collection struct {
	Id     int
	Filter string
	Tag    string
}

type Role struct {
	Id     int
	Filter string
	Tag    string
	Role   string
	Thumb  string
}

type Similar struct {
	Id     int
	Filter string
	Tag    string
}
