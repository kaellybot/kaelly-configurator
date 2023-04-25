package channels

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
)

type ChannelService interface {
	GetAlmanaxWebhook(guildId, channelId string, locale amqp.Language) (*entities.WebhookAlmanax, error)
	GetFeedWebhook(guildId, channelId, feedTypeId string, locale amqp.Language) (*entities.WebhookFeed, error)
	GetTwitterWebhook(guildId, channelId string, locale amqp.Language) (*entities.WebhookTwitter, error)
	SaveChannelServer(channelServer entities.ChannelServer) error
	SaveFeedWebhook(webhook entities.WebhookFeed) error
	SaveTwitterWebhook(webhook entities.WebhookTwitter) error
	SaveAlmanaxWebhook(webhook entities.WebhookAlmanax) error
	DeleteAlmanaxWebhook(webhook *entities.WebhookAlmanax) error
	DeleteFeedWebhook(webhook *entities.WebhookFeed) error
	DeleteTwitterWebhook(webhook *entities.WebhookTwitter) error
}

type ChannelServiceImpl struct {
	channelServerRepo  servers.ChannelServerRepository
	almanaxWebhookRepo almanax.AlmanaxWebhookRepository
	feedWebhookRepo    feeds.FeedWebhookRepository
	twitterWebhookRepo twitter.TwitterWebhookRepository
}
