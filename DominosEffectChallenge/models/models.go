package models

type ServiceGraph struct {
	Name      string   `json:"name"`
	Health    float32  `json:"health"`
	DependsOn []string `json:"depends_on"`
}


