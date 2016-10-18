package apicem

const relevanceBasePath = "v1"

// RelevanceService is an interface with the Relevance API
type RelevanceService service

// RelevanceListResult is ...
type RelevanceListResult struct {
	Response []string `json:"response,omitempty"`
	Version  string   `json:"version,omitempty"`
}

// GetAllRelevanceLevels is ...
// Get all Relevance(s)
//
//
//
//
// * @return *RelevanceListResult
func (s *RelevanceService) GetAllRelevanceLevels() (*RelevanceListResult, *Response, error) {

	path := relevanceBasePath + "/relevance"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(RelevanceListResult)
	resp, err := s.client.Do(req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
