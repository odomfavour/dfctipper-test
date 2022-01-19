package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/sessions"
)

type Server struct {
	SessionStore  *sessions.CookieStore
	webMux        *chi.Mux
	cfg           Config
	Templates     *Templates
	AuthTemplates *Templates
	SiteTemplates *Templates
	routes        map[string]route
	routeGroups   []routeGroup
	common        CommonPageData
}

// Links to be passed with common page data.
type Links struct {
	CoinbaseComment string
	POSExplanation  string
	APIDocs         string
	InsightAPIDocs  string
	Github          string
	License         string
	NetParams       string
	DownloadLink    string
	// Testnet and below are set via pdanalytics config.
	Testnet       string
	Mainnet       string
	TestnetSearch string
	MainnetSearch string
	OnionURL      string
}

type MenuItem struct {
	Href       string
	HyperText  string
	Info       string
	Attributes map[string]string
}

const (
	MainNavGroup = iota
	HistoricNavGroup
)

type NavGroup struct {
	Label     string
	MenuItems []MenuItem
}

type BreadcrumbItem struct {
	Href      string
	HyperText string
	Active    bool
}

// Cookies contains information from the request cookies.
type Cookies struct {
	DarkMode bool
}

// TimeDef is time.Time wrapper that formats time by default as a string without
// a timezone. The time Stringer interface formats the time into a string with a
// timezone.
type TimeDef struct {
	T time.Time
}

const (
	timeDefFmtHuman        = "2006-01-02 15:04:05 (MST)"
	timeDefFmtDateTimeNoTZ = "2006-01-02 15:04:05"
	timeDefFmtJS           = time.RFC3339
)

// String formats the time in a human-friendly layout. This ends up on the
// explorer web pages.
func (t TimeDef) String() string {
	return t.T.Format(timeDefFmtHuman)
}

// RFC3339 formats the time in a machine-friendly layout.
func (t TimeDef) RFC3339() string {
	return t.T.Format(timeDefFmtJS)
}

// UNIX returns the UNIX epoch time stamp.
func (t TimeDef) UNIX() int64 {
	return t.T.Unix()
}

func (t TimeDef) Format(layout string) string {
	return t.T.Format(layout)
}

// MarshalJSON implements json.Marshaler.
func (t *TimeDef) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.RFC3339())
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *TimeDef) UnmarshalJSON(data []byte) error {
	if t == nil {
		return fmt.Errorf("TimeDef: UnmarshalJSON on nil pointer")
	}
	tStr := string(data)
	tStr = strings.Trim(tStr, `"`)
	T, err := time.Parse(timeDefFmtJS, tStr)
	if err != nil {
		return err
	}
	t.T = T
	return nil
}

// CommonPageData is the basis for data structs used for HTML templates.
// explorerUI.commonData returns an initialized instance or CommonPageData,
// which itself should be used to initialize page data template structs.
type CommonPageData struct {
	Version             string
	BlockTimeUnix       int64
	DevAddress          string
	Links               *Links
	NavGroups           []NavGroup
	NetName             string
	Cookies             Cookies
	RequestURI          string
	RequestID           string
	CurrentPage         string
	Authenticated       bool
	ErrorAlertMessage   string
	SuccessAlertMessage string
	CSRFToken           template.HTML
}
