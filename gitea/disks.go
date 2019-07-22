package gitea

import (
	core "k8s.io/api/core/v1"
)

type Disk struct {
	Name                  string                 `json:"name,omitempty"`
	Id                    string                 `json:"id,omitempty"`
	Provider              string                 `json:"provider,omitempty"`
	SizeGb                int64                  `json:"size_gb,omitempty"`
	Type                  string                 `json:"type,omitempty"`
	Zone                  string                 `json:"zone,omitempty"`
	Status                string                 `json:"status,omitempty"`
	Endpoint              string                 `json:"endpoint,omitempty"`
	Iops                  int64                  `json:"iops,omitempty"`
	PersistentVolume      *PersistentVolume      `json:"persistent_volume,omitempty"`
	PersistentVolumeClaim *PersistentVolumeClaim `json:"persistent_volume_claim,omitempty"`
}

type PersistentVolumeClaim struct {
	Meta   *Meta                                              `json:"meta,omitempty"`
	Spec   *PersistentVolumeClaimSpec                         `json:"spec,omitempty"`
	Status *PersistentVolumeClaim_PersistentVolumeClaimStatus `json:"status,omitempty"`
}

type PersistentVolumeClaimSpec struct {
	StorageClassName string                                          `json:"storage_class_name,omitempty"`
	AccessModes      []string                                        `json:"access_modes,omitempty"`
	Resources        *PersistentVolumeClaimSpec_ResourceRequirements `json:"resources,omitempty"`
	VolumeName       string                                          `json:"volume_name,omitempty"`
}
type PersistentVolumeClaim_PersistentVolumeClaimStatus struct {
	Phase       string            `json:"phase,omitempty"`
	AccessModes []string          `json:"access_modes,omitempty"`
	Capacity    map[string]string `json:"capacity,omitempty"`
}

type PersistentVolumeClaimSpec_ResourceRequirements struct {
	Limits   map[string]string `json:"limits,omitempty"`
	Requests map[string]string `json:"requests,omitempty"`
}

type PersistentVolume struct {
	Meta   *Meta                                    `json:"meta,omitempty"`
	Spec   *PersistentVolume_PersistentVolumeSpec   `json:"spec,omitempty"`
	Status *PersistentVolume_PersistentVolumeStatus `json:"status,omitempty"`
}

type Meta struct {
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	SelfLink          string            `json:"self_link,omitempty"`
	ResourceVersion   string            `json:"resource_version,omitempty"`
	CreationTimestamp int64             `json:"creation_timestamp,omitempty"`
	Generation        int64             `json:"generation,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	Uid               string            `json:"uid,omitempty"`
}

type PersistentVolume_PersistentVolumeSpec struct {
	Capacity                      map[string]string            `json:"capacity,omitempty"`
	AccessModes                   []string                     `json:"access_modes,omitempty"`
	PersistentVolumeReclaimPolicy string                       `json:"persistent_volume_reclaim_policy,omitempty"`
	ClaimRef                      *ObjectReference             `json:"claim_ref,omitempty"`
	PersistentVolumeSource        *core.PersistentVolumeSource `json:"persistent_volume_source,omitempty"`
}

type ObjectReference struct {
	Kind            string `json:"kind,omitempty"`
	Namespace       string `json:"namespace,omitempty"`
	Name            string `json:"name,omitempty"`
	Uid             string `json:"uid,omitempty"`
	ApiVersion      string `json:"api_version,omitempty"`
	ResourceVersion string `json:"resource_version,omitempty"`
}

type PersistentVolume_PersistentVolumeStatus struct {
	Phase   string `json:"phase,omitempty"`
	Message string `json:"message,omitempty"`
	Reason  string `json:"reason,omitempty"`
}

type CreateDiskOption struct {
	Name     string `json:"name,omitempty"`
	Zone     string `json:"zone,omitempty"`
	DiskType string `json:"disk_type,omitempty"`
	SizeGb   int64  `json:"size_gb,omitempty"`
}

type DescribeDiskOption struct {
	Name     string `json:"name,omitempty"`
	Provider string `json:"provider,omitempty"`
}

type DescribeDiskResponse struct {
	Disk *Disk `json:"disk,omitempty"`
}

type DeleteDiskOption struct {
	Uid string `json:"uid,omitempty"`
}

type ListDiskRequest struct {
	Cluster string `json:"cluster,omitempty"`
}

type ListDiskResponse struct {
	Disks []*Disk `json:"disks,omitempty"`
}
