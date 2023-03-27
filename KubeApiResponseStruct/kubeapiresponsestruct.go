package kubeApiResponseStruct

type PodsStatus struct {
	PendingNumber   int `json:"nodenumber"`
	RunningNumber   int `json:"runningnumber"`
	SucceededNumber int `json:"succeedednumber"`
	FailedNumber    int `json:"failednumber"`
}

type NodeRespose struct {
	NodeNumber        int    `json:"nodenumber"`
	ContolPlaneNumber int    `json:"contolplanenumber"`
	WorkerNumber      int    `json:"workernumber"`
	KubernetesVersion string `json:"kubernetesversion"`
}

type NamespaceRespose struct {
	NamespaceNumber int `json:"NamespaceNumber"`
}

type PodsResponse struct {
	PodsNumber         int        `json:"podsnumber"`
	PodsNumberByStatus PodsStatus `json:"podsnumberbystatus"`
}

type IngressResponse struct {
	IngressNumber int `json:"ingressnumber"`
}

type ResourceResponce struct {
	Nodes      NodeRespose      `json:"nodes"`
	Namespaces NamespaceRespose `json:"namespaces"`
	Pods       PodsResponse     `json:"pods"`
	Ingresses  IngressResponse  `json:"ingresses"`
}
