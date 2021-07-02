package taigo

import (
	"strconv"
)

type HistoryService struct {
	client *Client
}

// ListCompleteHistory
func (s *HistoryService) ListCompleteHistory(objId int, endpoint Endpoint) ([]History, error) {
	url := s.client.MakeURL("history", mapEndpointName(endpoint), strconv.Itoa(objId))
	var h []History
	_, err := s.client.Request.Get(url, &h)
	if err != nil {
		return nil, err
	}
	return h, nil
}

// CreateHistoryEntry
func (s *HistoryService) CreateHistoryEntry(history *History, entity TaigaBaseObject, endpoint Endpoint) error {
	// only handle comments
	if history.Comment == "" {
		return nil
	}

	url := s.client.MakeURL(string(endpoint), strconv.Itoa(entity.GetID()))

	var baseObj map[string]interface{}
	_, err := s.client.Request.Get(url, &baseObj)
	if err != nil {
		return err
	}

	_, err = s.client.Request.Patch(url, map[string]interface{}{"version": baseObj["version"], "comment": history.Comment}, nil)
	return err
}

// mapEndpointName return the correct endpoint for the history requests (singular names)
func mapEndpointName(endpoint Endpoint) string {
	switch endpoint {
	case UserStoryEndpoint:
		return "userstory"
	case TaskEndpoint:
		return "task"
	case EpicEndpoint:
		return "epic"
	case IssueEndpoint:
		return "issue"
	default:
		return ""
	}
}
