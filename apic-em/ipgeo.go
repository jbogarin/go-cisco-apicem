package apicem

import (
	"fmt"
	"strings"
)

const ipgeoBasePath = "v1"

// IPGeoService is an interface with the IPgeowanIP API
type IPGeoService service

// HashMapstringIPGeoModel is ...
type HashMapstringIPGeoModel struct {
}

// IPGeoModel is ...
type IPGeoModel struct {
	Latitude        string `json:"latitude,omitempty"`
	Longitude       string `json:"longitude,omitempty"`
	City            string `json:"city,omitempty"`
	Continent       string `json:"continent,omitempty"`
	SubDivision     string `json:"subDivision,omitempty"`
	SubDivisionCode string `json:"subDivisionCode,omitempty"`
	CountryCode     string `json:"countryCode,omitempty"`
	ContinentCode   string `json:"continentCode,omitempty"`
	Country         string `json:"country,omitempty"`
}

// IPGeoModelResult is ...
type IPGeoModelResult struct {
	Version  string                  `json:"version,omitempty"`
	Response HashMapstringIPGeoModel `json:"response,omitempty"`
}

// GetCityInfo is ...
// This method is used to get the location details for a list of space-separated WAN IP addresses.
//
//
//  * @param scope Authorization Scope for RBAC
//
// * @return *IPGeoModelResult
func (s *IPGeoService) GetCityInfo(wanIP string, scope string) (*IPGeoModelResult, *Response, error) {

	path := ipgeoBasePath + "/ipgeo/{wanIP}"
	path = strings.Replace(path, "{"+"wanIP"+"}", fmt.Sprintf("%v", wanIP), -1)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("scope", scope)

	root := new(IPGeoModelResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
