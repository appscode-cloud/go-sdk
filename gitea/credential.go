package gitea

import (
	"time"
	pharmerApi "github.com/pharmer/pharmer/apis/v1beta1"
)

type Credential struct {
	ID            int64       `json:"id"`
	Owner         *User       `json:"owner"`
	Name          string      `json:"name"`
	Parent        *Credential `json:"parent"`
	Archived      bool        `json:"archived"`
	Data          pharmerApi.Credential    `json:"data"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at"`
	// swagger:strfmt date-time
	Updated     time.Time   `json:"updated_at"`
	// swagger:strfmt date-time
	Deleted     *time.Time   `json:"deleted_at"`
}

// CreateCredentialOption options when creating kubernetes cluster
// swagger:model

type CreateCredentialOption struct {
	// Name of the credential to create
	//
	// required: true
	// unique: true
	Name string `json:"name" binding:"Required;AlphaDashDot;MaxSize(100)"`
	// Credential data
	Data pharmerApi.Credential `json:"data" binding:"Required"`
}
