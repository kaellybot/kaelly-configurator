package channels

import (
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
)

type ChannelService interface {
	SaveChannelServer(channelServer entities.ChannelServer) error
	SaveAlmanaxWebhook(webhook entities.WebhookAlmanax) error
	SaveFeedWebhook(webhook entities.WebhookFeed) error
	SaveTwitterWebhook(webhook entities.WebhookTwitter) error
}

type ChannelServiceImpl struct {
	channelServerRepo  servers.ChannelServerRepository
	almanaxWebhookRepo almanax.AlmanaxWebhookRepository
	feedWebhookRepo    feeds.FeedWebhookRepository
	twitterWebhookRepo twitter.TwitterWebhookRepository
}
