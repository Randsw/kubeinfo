package kubeApiResponseStruct

// PodsStatus is the representation of kubernetes pod phase.
type PodsStatus struct {
	PendingNumber   int `json:"nodenumber"`
	RunningNumber   int `json:"runningnumber"`
	SucceededNumber int `json:"succeedednumber"`
	FailedNumber    int `json:"failednumber"`
}

// NodeResponse is represenatation of kubernetes nodes information include number of nodes by type
// kubernetes verison and OS type
type NodeRespose struct {
	NodeNumber        int    `json:"nodenumber"`
	ContolPlaneNumber int    `json:"contolplanenumber"`
	WorkerNumber      int    `json:"workernumber"`
	KubernetesVersion string `json:"kubernetesversion"`
	OsImage           string `json:"osimage"`
}

// NamespaceRespose is represenatation of kubernetes namespace count
type NamespaceResponse struct {
	NamespaceNumber int `json:"NamespaceNumber"`
}

type PodsResponse struct {
	PodsNumber         int        `json:"podsnumber"`
	PodsNumberByStatus PodsStatus `json:"podsnumberbystatus"`
}

type IngressResponse struct {
	IngressNumber int `json:"ingressnumber"`
}

type FluxKustomizationsStatus struct {
	Ready    int `json:"ready"`
	NotReady int `json:"notready"`
	Unknown  int `json:"unknown"`
}

type FluxKustomizationsResponse struct {
	FluxKustomizationsNumber         int                      `json:"fluxkustomizationsnumber"`
	FluxKustomizationsNumberbyStatus FluxKustomizationsStatus `json:"fluxkustomizationsnumberbystatus"`
}

type FluxHelmreleasesStatus struct {
	Ready    int `json:"ready"`
	NotReady int `json:"notready"`
	Unknown  int `json:"unknown"`
}

type FluxHelmreleasesResponse struct {
	FluxHelmreleasesNumber         int                      `json:"fluxhelmreleasesnumber"`
	FluxHelmreleasesNumberbyStatus FluxKustomizationsStatus `json:"fluxhelmreleasesnumberbystatus"`
}

type ResourceResponce struct {
	Nodes              NodeRespose                `json:"nodes"`
	Namespaces         NamespaceResponse          `json:"namespaces"`
	Pods               PodsResponse               `json:"pods"`
	Ingresses          IngressResponse            `json:"ingresses"`
	FluxKustomizations FluxKustomizationsResponse `json:"fluxkustomizations"`
	FluxHelmreleases   FluxHelmreleasesResponse   `json:"fluxhelmreleases"`
}

type HealthzResponce struct {
	App_name string `json:"app_name"`
	Status   string `json:"status"`
	Tag      string `json:"tag"`
	Hash     string `json:"hash"`
	Date     string `json:"date"`
}
