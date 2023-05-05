package channels

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
)

type Service interface {
	GetAlmanaxWebhook(guildID, channelID string, locale amqp.Language) (*entities.WebhookAlmanax, error)
	GetFeedWebhook(guildID, channelID, feedTypeID string, locale amqp.Language) (*entities.WebhookFeed, error)
	GetTwitterWebhook(guildID, channelID string, locale amqp.Language) (*entities.WebhookTwitter, error)
	SaveChannelServer(channelServer entities.ChannelServer) error
	SaveFeedWebhook(webhook entities.WebhookFeed) error
	SaveTwitterWebhook(webhook entities.WebhookTwitter) error
	SaveAlmanaxWebhook(webhook entities.WebhookAlmanax) error
	DeleteAlmanaxWebhook(webhook *entities.WebhookAlmanax) error
	DeleteFeedWebhook(webhook *entities.WebhookFeed) error
	DeleteTwitterWebhook(webhook *entities.WebhookTwitter) error
}

type Impl struct {
	channelServerRepo  servers.Repository
	almanaxWebhookRepo almanax.Repository
	feedWebhookRepo    feeds.Repository
	twitterWebhookRepo twitter.Repository
}
