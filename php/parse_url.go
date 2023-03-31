package php

import (
	"net/url"
)

type urlInfo struct {
	Scheme   string                 `json:"scheme"`
	Host     string                 `json:"host"`
	Port     string                 `json:"port,omitempty"`
	User     string                 `json:"user,omitempty"`
	Password string                 `json:"pass,omitempty"`
	Path     string                 `json:"path"`
	Query    map[string]interface{} `json:"query"`
	Fragment string                 `json:"fragment,omitempty"`
}

func ParseUrl(urlStr string) *urlInfo {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil
	}

	queryValues, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil
	}

	args := make(map[string]interface{})
	for key, values := range queryValues {
		if len(values) == 1 {
			args[key] = values[0]
		} else {
			args[key] = values
		}
	}

	username := ""
	password := ""
	if u.User != nil {
		username = u.User.Username()
		password, _ = u.User.Password()
	}

	port := ""
	if u.Port() != "" {
		port = u.Port()
	}

	return &urlInfo{
		Scheme:   u.Scheme,
		Host:     u.Hostname(),
		Port:     port,
		User:     username,
		Password: password,
		Path:     u.Path,
		Query:    args,
		Fragment: u.Fragment,
	}
}
