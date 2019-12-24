package entity

type Setting struct {
	SessionLifetimeInDays int      `bson:"sessionLifetimeInDays"`
	AuthAllowedDomains    []string `bson:"authAllowedDomains"`
}