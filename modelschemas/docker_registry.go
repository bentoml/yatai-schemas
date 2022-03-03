package modelschemas

type DockerRegistrySchema struct {
	BentosRepositoryURI          string `json:"bentosRepositoryURI"`
	ModelsRepositoryURI          string `json:"modelsRepositoryURI"`
	BentosRepositoryURIInCluster string `json:"bentosRepositoryURIInCluster"`
	ModelsRepositoryURIInCluster string `json:"modelsRepositoryURIInCluster"`
	Server                       string `json:"server"`
	Username                     string `json:"username"`
	Password                     string `json:"password"`
	Secure                       bool   `json:"secure"`
}
