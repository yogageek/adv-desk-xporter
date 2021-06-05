package config

import (
	"time"
)

var (
	IFPURL            string
	MongodbURL        string
	MongodbUsername   string
	MongodbPassword   string
	MongodbDatabase   string
	AdminUsername     string
	AdminPassword     string
	OutboundURL       string
	Token             string
	Datacenter        string
	Cluster           string
	Workspace         string
	Namespace         string
	SSOURL            string
	AppID             string
	IFPStatus         = "Down"
	TaipeiTimeZone, _ = time.LoadLocation("Asia/Taipei")
	UTCTimeZone, _    = time.LoadLocation("UTC")
	EnvPath           = "local.env"
)
