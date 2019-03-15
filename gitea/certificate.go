package gitea

type Certificate struct {
	ID         int64  `json:"id"`
	Name       string `json:"name,omitempty"`
	CommonName string `json:"common_name,omitempty"`
	IssuedBy   string `json:"issued_by,omitempty"`
	ValidFrom  int64  `json:"valid_from,omitempty"`
	ExpireDate int64  `json:"expire_date,omitempty"`
	// those feilds will not included into list response.
	// only describe response will include the underlying
	// feilds.
	Sans         []string `json:"sans,omitempty"`
	Cert         string   `json:"cert,omitempty"`
	Key          string   `json:"key,omitempty"`
	Version      int32    `json:"version,omitempty"`
	SerialNumber string   `json:"serial_number,omitempty"`
}

type LoadCertificateOptions struct {
	// required: true
	// unique: true
	Name     string `json:"name" binding:"Required;AlphaDashDot;MaxSize(100)"`
	CertData string `json:"cert" binding:"Required"`
	KeyData  string `json:"key" binding:"Required"`
}

type DeployCertificateOptions struct {
	SecretName  string `json:"secret_name" binding:"Required"`
	ClusterName string `json:"cluster_name" binding:"Required"`
	Namespace   string `json:"namespace"`
}
