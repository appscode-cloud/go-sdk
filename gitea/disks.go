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
	Meta   *Meta                                              `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Spec   *PersistentVolumeClaimSpec                         `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status *PersistentVolumeClaim_PersistentVolumeClaimStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

type PersistentVolumeClaimSpec struct {
	StorageClassName string                                          `protobuf:"bytes,1,opt,name=storage_class_name,json=storageClassName" json:"storage_class_name,omitempty"`
	AccessModes      []string                                        `protobuf:"bytes,2,rep,name=access_modes,json=accessModes" json:"access_modes,omitempty"`
	Resources        *PersistentVolumeClaimSpec_ResourceRequirements `protobuf:"bytes,3,opt,name=resources" json:"resources,omitempty"`
	VolumeName       string                                          `protobuf:"bytes,4,opt,name=volume_name,json=volumeName" json:"volume_name,omitempty"`
}
type PersistentVolumeClaim_PersistentVolumeClaimStatus struct {
	Phase       string            `protobuf:"bytes,1,opt,name=phase" json:"phase,omitempty"`
	AccessModes []string          `protobuf:"bytes,2,rep,name=access_modes,json=accessModes" json:"access_modes,omitempty"`
	Capacity    map[string]string `protobuf:"bytes,3,rep,name=capacity" json:"capacity,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

type PersistentVolumeClaimSpec_ResourceRequirements struct {
	Limits   map[string]string `protobuf:"bytes,1,rep,name=limits" json:"limits,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Requests map[string]string `protobuf:"bytes,2,rep,name=requests" json:"requests,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

type PersistentVolume struct {
	Meta   *Meta                                    `protobuf:"bytes,1,opt,name=meta" json:"meta,omitempty"`
	Spec   *PersistentVolume_PersistentVolumeSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status *PersistentVolume_PersistentVolumeStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
}

type Meta struct {
	Name              string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace         string            `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	SelfLink          string            `protobuf:"bytes,3,opt,name=self_link,json=selfLink" json:"self_link,omitempty"`
	ResourceVersion   string            `protobuf:"bytes,4,opt,name=resource_version,json=resourceVersion" json:"resource_version,omitempty"`
	CreationTimestamp int64             `protobuf:"varint,5,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	Generation        int64             `protobuf:"varint,6,opt,name=generation" json:"generation,omitempty"`
	Labels            map[string]string `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Annotations       map[string]string `protobuf:"bytes,8,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Uid               string            `protobuf:"bytes,9,opt,name=uid" json:"uid,omitempty"`
}

type PersistentVolume_PersistentVolumeSpec struct {
	Capacity                      map[string]string            `protobuf:"bytes,1,rep,name=capacity" json:"capacity,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	AccessModes                   []string                     `protobuf:"bytes,2,rep,name=access_modes,json=accessModes" json:"access_modes,omitempty"`
	PersistentVolumeReclaimPolicy string                       `protobuf:"bytes,3,opt,name=persistent_volume_reclaim_policy,json=persistentVolumeReclaimPolicy" json:"persistent_volume_reclaim_policy,omitempty"`
	ClaimRef                      *ObjectReference             `protobuf:"bytes,4,opt,name=claim_ref,json=claimRef" json:"claim_ref,omitempty"`
	PersistentVolumeSource        *core.PersistentVolumeSource `protobuf:"bytes,5,opt,name=persistent_volume_source,json=persistentVolumeSource" json:"persistent_volume_source,omitempty"`
}

type ObjectReference struct {
	Kind            string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Namespace       string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Name            string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Uid             string `protobuf:"bytes,4,opt,name=uid" json:"uid,omitempty"`
	ApiVersion      string `protobuf:"bytes,5,opt,name=api_version,json=apiVersion" json:"api_version,omitempty"`
	ResourceVersion string `protobuf:"bytes,6,opt,name=resource_version,json=resourceVersion" json:"resource_version,omitempty"`
}

type PersistentVolume_PersistentVolumeStatus struct {
	Phase   string `protobuf:"bytes,1,opt,name=phase" json:"phase,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Reason  string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
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
	Disk *Disk `protobuf:"bytes,1,opt,name=disk" json:"disk,omitempty"`
}

type DeleteDiskOption struct {
	Uid string `json:"uid,omitempty"`
}

type ListDiskRequest struct {
	Cluster string `json:"cluster,omitempty"`
}

type ListDiskResponse struct {
	Disks []*Disk `protobuf:"bytes,1,rep,name=disks" json:"disks,omitempty"`
}
