package gitea

import (
	"bytes"
	"encoding/json"
	"time"
)

// Cluster represents a repository
type Cluster struct {
	ID            int64       `json:"id"`
	Owner         *User       `json:"owner"`
	Name          string      `json:"name"`
	Private       bool        `json:"private"`
	Parent        *Cluster `json:"parent"`
	Archived      bool        `json:"archived"`
	Data          string    `json:"data"`
	OperationID   string `json:"operation_id"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at"`
	// swagger:strfmt date-time
	Updated     time.Time   `json:"updated_at"`
	// swagger:strfmt date-time
	Deleted     *time.Time   `json:"deleted_at"`

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
	Configuration string `json:"configuration" binding:"Required"`
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
