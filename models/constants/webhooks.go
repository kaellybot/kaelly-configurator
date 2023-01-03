package constants

type WebhookType int

const (
	Almanax WebhookType = iota
	Rss
	Twitter
)
