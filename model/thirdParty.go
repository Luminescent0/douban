package model

type URL struct {
	AuthURL  string
	TokenURL string
}
type Config struct {
	RedirectURL  string
	ClientID     string
	ClientSecret string
	Endpoint     URL
	Scopes       []string
}
