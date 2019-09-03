package plex

// https://support.plex.tv/articles/115002267687-webhooks/#toc-2

// global constants
const (
	PlexMovieType = "movie"
	PlexShowType  = "show"
)

// Webhook describes the plex webhook json structure
type Webhook struct {
	Event    string
	User     bool
	Owner    bool
	Account  Account
	Server   Server
	Player   Player
	Metadata Metadata
}

//Account describes the structure of the Account in Webhook
type Account struct {
	ID    int
	Thumb string
	Title string
}

//Server describes the structure of the Server in Webhook
type Server struct {
	Title string
	UUID  string
}

//Player describes the structure of the Player in Webhook
type Player struct {
	Local         bool
	PublicAddress string
	Title         string
	UUID          string
}

//Metadata describes the structure of the Metadata in Webhook
type Metadata struct {
	LibrarySectionType   string
	RatingKey            string
	Key                  string
	ParentRatingKey      string
	GrandParentRatingKey string
	GuID                 string
	ParentGuID           string
	GrandParentGuID      string
	GrandParentKey       string
	ParentKey            string
	LibrarySectionTitle  string
	LibrarySectionID     int
	LibrarySectionKey    string
	Studio               string
	Mediatype            string `json:"type"`

	Title            string
	ParentTitle      string
	GrandParentTitle string

	Index                 int
	ParentIndex           int
	ContentRating         string
	Summary               string
	Rating                float32
	AudienceRating        float32
	ViewCount             int
	LastViewedAt          int32
	Year                  int
	Tagline               string
	Thumb                 string
	GrandParentThumb      string
	Art                   string
	GrandParentArt        string
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

//Genre describes the structure of the Account in Metadata
type Genre struct {
	ID     int
	Filter string
	Tag    string
	Count  int
}

//Director describes the structure of the Director in Metadata
type Director struct {
	ID     int
	Filter string
	Tag    string
}

//Writer describes the structure of the Writer in Metadata
type Writer struct {
	ID     int
	Filter string
	Tag    string
}

//Producer describes the structure of the Producer in Metadata
type Producer struct {
	ID     int
	Filter string
	Tag    string
}

//Country describes the structure of the Country in Metadata
type Country struct {
	ID     int
	Filter string
	Tag    string
	Count  int
}

//Collection describes the structure of the Collection in Metadata
type Collection struct {
	ID     int
	Filter string
	Tag    string
}

//Role describes the structure of the Role in Metadata
type Role struct {
	ID     int
	Filter string
	Tag    string
	Role   string
	Thumb  string
}

//Similar describes the structure of the Similar in Metadata
type Similar struct {
	ID     int
	Filter string
	Tag    string
}

//GetMediaType returns the media type
func (w *Webhook) GetMediaType() string {
	return w.Metadata.Mediatype
}

//IsMovie returns whether the media is a movie or not
func (w *Webhook) IsMovie() bool {
	return w.Metadata.Mediatype == PlexMovieType
}

//IsShow returns whether the media is a show or not
func (w *Webhook) IsShow() bool {
	return w.Metadata.Mediatype == PlexShowType
}

//GetSeason returns the show's season number
func (w *Webhook) GetSeason() int {
	return w.Metadata.ParentIndex
}

//GetEpisodeNo returns the show's episode number
func (w *Webhook) GetEpisodeNo() int {
	return w.Metadata.Index
}

//GetShowTitle returns the show's title
func (w *Webhook) GetShowTitle() string {
	return w.Metadata.GrandParentTitle
}

//GetEpisodeTitle returns the show's episode title
func (w *Webhook) GetEpisodeTitle() string {
	return w.Metadata.Title
}
