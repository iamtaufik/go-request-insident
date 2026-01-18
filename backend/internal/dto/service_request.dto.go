package dto

type DraftServiceRequestPayload struct {
	RequestType  string `json:"request_type" binding:"required"`
	ReportedFrom string `json:"reported_from" binding:"required"`
	SLAType      string `json:"sla_type" binding:"required"`
}

type UpdateServiceRequestPayload struct {
	Company      string `json:"company" binding:"required"`
	Organization string `json:"organization" binding:"required"`
	Department   string `json:"department" binding:"required"`
	Impact       string `json:"impact" binding:"required"`
	Urgency      string `json:"urgency" binding:"required"`
	Priority     string `json:"priority" binding:"required"`
	Summary      string `json:"summary" binding:"required"`
	Note         string `json:"note" binding:"required"`
}

type ServiceRequestAttachmentPayload struct {
	FileName string `json:"file_name" binding:"required"`
	FileURL  string `json:"file_url" binding:"required"`
	FileType string `json:"file_type" binding:"required"`
	FileSize int64  `json:"file_size" binding:"required"`
}