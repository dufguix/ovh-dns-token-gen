package main

import "github.com/ovh/go-ovh/ovh"

type application struct {
	ID          int    `json:"applicationId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Key         string `json:"applicationKey"`
}

type credential struct {
	ID         int              `json:"credentialId"`
	AppID      int              `json:"applicationId"`
	Status     string           `json:"status"`
	LastUse    string           `json:"lastUse"`
	Expiration string           `json:"expiration"`
	Creation   string           `json:"creation"`
	OvhSupport bool             `json:"ovhSupport"`
	Rules      []ovh.AccessRule `json:"rules"`
	AllowedIPs []string         `json:"allowedIPs"`
}
