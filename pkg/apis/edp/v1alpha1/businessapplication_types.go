package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BusinessApplicationSpec defines the desired state of BusinessApplication
type BusinessApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	GitRepoUrl      string `json:"git_repo_url"`
	AddRepoStrategy string `json:"add_repo_strategy"`
	Language        string `json:"language"`
	BuildTool       string `json:"build_tool"`
	Framework       string `json:"framework"`
	Database        bool   `json:"database"`
	Route           struct {
		Site string `json:"site"`
		Path string `json:"path"`
	}
}

// BusinessApplicationStatus defines the observed state of BusinessApplication
type BusinessApplicationStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	Action   string `json:"action"`
	Message  string `json:"message"`
	Status   string `json:"status"`
	Database struct {
		Enabled bool `json:"enabled"`
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BusinessApplication is the Schema for the businessapplications API
// +k8s:openapi-gen=true
type BusinessApplication struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BusinessApplicationSpec   `json:"spec,omitempty"`
	Status BusinessApplicationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BusinessApplicationList contains a list of BusinessApplication
type BusinessApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BusinessApplication `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BusinessApplication{}, &BusinessApplicationList{})
}
