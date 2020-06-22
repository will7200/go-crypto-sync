package personalcapital

type SpHeader struct {
	SPHEADERVERSION  int      `json:"SP_HEADER_VERSION"`
	UserStage        string   `json:"userStage"`
	IsDelegate       bool     `json:"isDelegate"`
	UserGUID         string   `json:"userGuid"`
	DeviceName       string   `json:"deviceName"`
	BetaTester       bool     `json:"betaTester"`
	AccountsMetaData []string `json:"accountsMetaData"`
	Success          bool     `json:"success"`
	QualifiedLead    bool     `json:"qualifiedLead"`
	Developer        bool     `json:"developer"`
	PersonID         int      `json:"personId"`
	AuthLevel        string   `json:"authLevel"`
	Username         string   `json:"username"`
	Status           string   `json:"status"`
	Errors           []struct {
		Code    int `json:"code"`
		Details struct {
			FieldName string `json:"fieldName"`
		} `json:"details"`
		Message string `json:"message"`
	} `json:"errors"`
	Csrf          string `json:"csrf"`
	SPDATACHANGES []struct {
		ServerChangeID int `json:"serverChangeId"`
		Details        struct {
			ID int `json:"id"`
		} `json:"details"`
		EventType string `json:"eventType"`
	} `json:"SP_DATA_CHANGES"`
}
