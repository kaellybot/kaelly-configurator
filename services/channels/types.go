package channels

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitch"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
	"github.com/kaellybot/kaelly-configurator/repositories/youtube"
)

type Service interface {
	GetAlmanaxWebhook(guildID, channelID string, game amqp.Game) (*entities.WebhookAlmanax, error)
	GetFeedWebhook(guildID, channelID, feedTypeID string, game amqp.Game) (*entities.WebhookFeed, error)
	GetTwitchWebhook(guildID, channelID, streamerID string) (*entities.WebhookTwitch, error)
	GetTwitterWebhook(guildID, channelID, twitterID string) (*entities.WebhookTwitter, error)
	GetYoutubeWebhook(guildID, channelID, videastID string) (*entities.WebhookYoutube, error)
	SaveAlmanaxWebhook(webhook entities.WebhookAlmanax) error
	SaveChannelServer(channelServer entities.ChannelServer) error
	SaveFeedWebhook(webhook entities.WebhookFeed) error
	SaveTwitchWebhook(webhook entities.WebhookTwitch) error
	SaveTwitterWebhook(webhook entities.WebhookTwitter) error
	SaveYoutubeWebhook(webhook entities.WebhookYoutube) error
	DeleteAlmanaxWebhook(webhook *entities.WebhookAlmanax) error
	DeleteFeedWebhook(webhook *entities.WebhookFeed) error
	DeleteTwitchWebhook(webhook *entities.WebhookTwitch) error
	DeleteTwitterWebhook(webhook *entities.WebhookTwitter) error
	DeleteYoutubeWebhook(webhook *entities.WebhookYoutube) error
}

type Impl struct {
	channelServerRepo  servers.Repository
	almanaxWebhookRepo almanax.Repository
	feedWebhookRepo    feeds.Repository
	twitchWebhookRepo  twitch.Repository
	twitterWebhookRepo twitter.Repository
	youtubeWebhookRepo youtube.Repository
}
