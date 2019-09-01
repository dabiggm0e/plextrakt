package plex

type Webhook struct {
	event    string
	user     bool
	owner    bool
	Account  Account  `json:"Account"`
	Server   Server   `json:"Server"`
	Player   Player   `json:"Player"`
	Metadata Metadata `json:"Metadata"`
}

type Account struct {
	id    int
	thumb string
	title string
}

type Server struct {
	title string
	uuid  string
}

type Player struct {
	local         bool
	publicAddress string
	title         string
	uuid          string
}

type Metadata struct {
	librarySectionType    string
	ratingKey             string
	key                   string
	guid                  string
	librarySectionTitle   string
	librarySectionID      int
	librarySectionKey     string
	studio                string
	Mediatype             string `json:"type"`
	title                 string
	contentRating         string
	summary               string
	rating                float32
	audienceRating        float32
	viewCousnt            int
	lastViewedAt          int32
	year                  int
	tagline               string
	thumb                 string
	art                   string
	duration              int
	originallyAvailableAt string
	addedAt               int32
	updatedAt             int32
	audienceRatingImage   string
	chapterSource         string
	primaryExtraKey       string
	ratingImage           string
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
	id     int
	filter string
	tag    string
	count  int
}

type Director struct {
	id     int
	filter string
	tag    string
}

type Writer struct {
	id     int
	filter string
	tag    string
}

type Producer struct {
	id     int
	filter string
	tag    string
}

type Country struct {
	id     int
	filter string
	tag    string
	count  int
}

type Collection struct {
	id     int
	filter string
	tag    string
}

type Role struct {
	id     int
	filter string
	tag    string
	role   string
	thumb  string
}

type Similar struct {
	id     int
	filter string
	tag    string
}
