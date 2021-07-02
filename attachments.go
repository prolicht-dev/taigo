package taigo

type AttachmentService struct {
	client *Client
}

// GetAttachment retrieves a attachment by its ID from a specified endpoint => https://taigaio.github.io/taiga-doc/dist/api.html#tasks-get-attachment
func (s *AttachmentService) GetAttachment(attachmentID int, endpoint Endpoint) (*Attachment, error) {
	a, err := getAttachmentForEndpoint(s.client, attachmentID, string(endpoint))
	if err != nil {
		return nil, err
	}
	return a, nil
}

// ListAttachments returns a list of attachments from a specified endpoint => https://taigaio.github.io/taiga-doc/dist/api.html#tasks-list-attachments
func (s *AttachmentService) ListAttachments(entity TaigaBaseObject, endpoint Endpoint) ([]Attachment, error) {
	queryParams := attachmentsQueryParams{
		endpointURI: string(endpoint),
		ObjectID:    entity.GetID(),
		Project:     entity.GetProject(),
	}

	attachments, err := listAttachmentsForEndpoint(s.client, &queryParams)
	if err != nil {
		return nil, err
	}
	return attachments, nil
}

// CreateAttachment creates a new attachment on a specified endpoint => https://taigaio.github.io/taiga-doc/dist/api.html#tasks-create-attachment
func (s *AttachmentService) CreateAttachment(attachment *Attachment, entity TaigaBaseObject, endpoint Endpoint) (*Attachment, error) {
	url := s.client.MakeURL(string(endpoint), "attachments")
	return newfileUploadRequest(s.client, url, attachment, entity)
}
