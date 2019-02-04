package gitea

import (
	clusterapi "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"time"
)

type Machine struct {
	ID            int64       `json:"id"`
	Owner         *User       `json:"owner"`
	Name          string      `json:"name"`
	//Cluster       *Cluster `json:"cluster"`
	ClusterId      int64 `json:"cluster_id"`
	Archived      bool        `json:"archived"`
	Data          clusterapi.Machine   `json:"data"`
	OperationID   string `json:"operation_id"`
	// swagger:strfmt date-time
	Created time.Time `json:"created_at"`
	// swagger:strfmt date-time
	Updated     time.Time   `json:"updated_at"`
	// swagger:strfmt date-time
	Deleted     *time.Time   `json:"deleted_at"`
}


// CreateMachineOption options when creating kubernetes cluster machine
// swagger:model

type CreateMachineOption struct {
	// Machine to create
	Machine clusterapi.Machine `json:"machine" binding:"Required"`
}
