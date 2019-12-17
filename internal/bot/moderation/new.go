package moderation

import (
    "github.com/arturoguerra/d2arena/internal/router"
//    "github.com/arturoguerra/d2arena/internal/config"
    "net/http"
    "fmt"
)

type (
    Requests struct {
       Client *http.Client
       Internal *http.Client
    }

    AddHeaderTransport struct {
        T http.RoundTripper
        HeaderMap map[string]string
    }
)

func (adt *AddHeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    for k, v := range adt.HeaderMap {
        fmt.Println("Adding header: " + k + ":" + v)
        req.Header.Add(k, v)
    }

    req.Header.Add("Content-Type", "application/json")

    return adt.T.RoundTrip(req)
}

func NewAddHeaderTransport(headers map[string]string) *AddHeaderTransport {
    T := http.DefaultTransport
	return &AddHeaderTransport{T, headers}
}

var requests Requests
var StaffChannelID string

func init() {
    StaffChannelID = "655803628644335661"
    //faceit := config.LoadFaceit()
    client := &http.Client{}
    internal := &http.Client{}

    requests = Requests{
        client,
        internal,
    }
}


func New(r *router.Router) {
    r.On("ban", Ban)
    r.On("remove", Delete)
    r.On("profile", getProfile)
}
