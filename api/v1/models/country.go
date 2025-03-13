package v1models

// Country represents a country from the REST Countries API
type Country struct {
	Name        Name                `json:"name"`
	TLD         []string            `json:"tld"`
	Independent bool                `json:"independent"`
	Status      string              `json:"status"`
	UNMember    bool                `json:"unMember"`
	Currencies  map[string]Currency `json:"currencies"`
	Capital     []string            `json:"capital"`
	Region      string              `json:"region"`
	Languages   map[string]string   `json:"languages"`
	Population  int                 `json:"population"`
	LatLng      []float64           `json:"latlng"`
	Continents  []string            `json:"continents"`
}

// Name represents the name information of a country
type Name struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

// NativeName represents a native name of a country
type NativeName struct {
	Official string `json:"official"`
	Common   string `json:"common"`
}

// Currency represents a currency of a country
type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type CountryRequest struct {
	Name string `form:"name"`
}

type CountryData struct {
	Name        Name                `json:"name"`
	Independent bool                `json:"independent"`
	UNMember    bool                `json:"unMember"`
	Currencies  map[string]Currency `json:"currencies"`
	Capital     []string            `json:"capital"`
	Region      string              `json:"region"`
	Languages   map[string]string   `json:"languages"`
	Population  string              `json:"population"`
	LatLng      []float64           `json:"latlng"`
	Continents  []string            `json:"continents"`
}
type CountryResponse struct {
	Data CountryData `json:"data"`
}
