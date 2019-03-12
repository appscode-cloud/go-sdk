package gitea

import (
	"bytes"
	"encoding/json"
	pharmerApi "github.com/pharmer/pharmer/apis/v1beta1"
	"time"
)

// Cluster represents a kubernetes cluster
type Cluster struct {
	ID            int64       `json:"id"`
	Owner         *User       `json:"owner"`
	Name          string      `json:"name"`
	Private       bool        `json:"private"`
	Parent        *Cluster `json:"parent"`
	Archived      bool        `json:"archived"`
	Data          pharmerApi.Cluster    `json:"data"`
	OperationID   string `json:"operation_id,omitempty"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at"`
	// swagger:strfmt date-time
	Updated     time.Time   `json:"updated_at"`
	// swagger:strfmt date-time
	Deleted     *time.Time   `json:"deleted_at"`

}


// ClusterMetadata represents metadata of a cluster
type ClusterMetadata struct {
	Config                    *ClusterMetadataResponse_KubedConfig                `json:"config,omitempty"`
	CustomResourceDefinitions []string                                            `json:"customResourceDefinitions,omitempty"`
	Upgrades                  []*pharmerApi.Upgrade `json:"upgrades,omitempty"`

}

type SSHConfigGetResponse struct {
	Config *pharmerApi.SSHConfig `json:"config,omitempty"`
}

type ClusterMetadataResponse_KubedConfig struct {
	//Version             *appscode_version.Version `json:"version,omitempty"`
	SearchEnabled       bool                      `json:"searchEnabled,omitempty"`
	ReverseIndexEnabled bool                      `json:"reverseIndexEnabled,omitempty"`
	AnalyticsEnabled    bool                      `json:"analyticsEnabled,omitempty"`
}


type ClusterOperation struct {
	OperationID string `json:"operation_id"`
}

// CreateClusterOption options when creating kubernetes cluster
// swagger:model

type CreateClusterOption struct {
	// Name of the cluster to create
	//
	// required: true
	// unique: true
	Name string `json:"name" binding:"Required;AlphaDashDot;MaxSize(100)"`
	// Configuration of the cluster to create
	Configuration pharmerApi.Cluster `json:"configuration" binding:"Required"`
	// Whether the cluster is private
	Private bool `json:"private"`
}

// CreateCluster creates a cluster for authenticated user.
func (c *Client) CreateCluster(opt CreateClusterOption) (*Cluster, error) {
	body, err := json.Marshal(&opt)
	if err != nil {
		return nil, err
	}
	cluster := new(Cluster)
	return cluster, c.getParsedResponse("POST", "/user/clusters", jsonHeader, bytes.NewReader(body), cluster)
}

type ImportClusterOption struct {
	Provider string `json:"provider" binding:"Required"`
	KubeConfig string `json:"kube_config" binding:"Required"`
	Region string `json:"region" binding:"Required"`
	Zone string `json:"zone" binding:"Required"`
	SSHKey string `json:"ssh_key"`
	CredentialSecret string `json:"credential_secret" binding:"Requried"`
	CredentialSecretNamespace string `json:"credential_secret_namespace" binding:"Requried"`

}
