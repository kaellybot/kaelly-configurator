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

func New(channelServerRepo servers.Repository, almanaxRepo almanax.Repository,
	feedRepo feeds.Repository, twitchRepo twitch.Repository,
	twitterRepo twitter.Repository, youtubeRepo youtube.Repository) (*Impl, error) {
	return &Impl{
		channelServerRepo:  channelServerRepo,
		almanaxWebhookRepo: almanaxRepo,
		feedWebhookRepo:    feedRepo,
		twitchWebhookRepo:  twitchRepo,
		twitterWebhookRepo: twitterRepo,
		youtubeWebhookRepo: youtubeRepo,
	}, nil
}

func (service *Impl) GetAlmanaxWebhook(guildID, channelID string,
	game amqp.Game) (*entities.WebhookAlmanax, error) {
	return service.almanaxWebhookRepo.Get(guildID, channelID, game)
}

func (service *Impl) GetFeedWebhook(guildID, channelID, feedTypeID string,
	game amqp.Game) (*entities.WebhookFeed, error) {
	return service.feedWebhookRepo.Get(guildID, channelID, feedTypeID, game)
}

func (service *Impl) GetTwitchWebhook(guildID, channelID, streamerID string,
) (*entities.WebhookTwitch, error) {
	return service.twitchWebhookRepo.Get(guildID, channelID, streamerID)
}

func (service *Impl) GetTwitterWebhook(guildID, channelID, twitterID string,
) (*entities.WebhookTwitter, error) {
	return service.twitterWebhookRepo.Get(guildID, channelID, twitterID)
}

func (service *Impl) GetYoutubeWebhook(guildID, channelID, videastID string,
) (*entities.WebhookYoutube, error) {
	return service.youtubeWebhookRepo.Get(guildID, channelID, videastID)
}

func (service *Impl) SaveChannelServer(channelServer entities.ChannelServer) error {
	return service.channelServerRepo.Save(channelServer)
}

func (service *Impl) SaveAlmanaxWebhook(webhook entities.WebhookAlmanax) error {
	return service.almanaxWebhookRepo.Save(webhook)
}

func (service *Impl) SaveFeedWebhook(webhook entities.WebhookFeed) error {
	return service.feedWebhookRepo.Save(webhook)
}

func (service *Impl) SaveTwitchWebhook(webhook entities.WebhookTwitch) error {
	return service.twitchWebhookRepo.Save(webhook)
}

func (service *Impl) SaveTwitterWebhook(webhook entities.WebhookTwitter) error {
	return service.twitterWebhookRepo.Save(webhook)
}

func (service *Impl) SaveYoutubeWebhook(webhook entities.WebhookYoutube) error {
	return service.youtubeWebhookRepo.Save(webhook)
}

func (service *Impl) DeleteAlmanaxWebhook(webhook *entities.WebhookAlmanax) error {
	if webhook != nil {
		return service.almanaxWebhookRepo.Delete(*webhook)
	}
	return nil
}

func (service *Impl) DeleteFeedWebhook(webhook *entities.WebhookFeed) error {
	if webhook != nil {
		return service.feedWebhookRepo.Delete(*webhook)
	}
	return nil
}

func (service *Impl) DeleteTwitchWebhook(webhook *entities.WebhookTwitch) error {
	if webhook != nil {
		return service.twitchWebhookRepo.Delete(*webhook)
	}
	return nil
}

func (service *Impl) DeleteTwitterWebhook(webhook *entities.WebhookTwitter) error {
	if webhook != nil {
		return service.twitterWebhookRepo.Delete(*webhook)
	}
	return nil
}

func (service *Impl) DeleteYoutubeWebhook(webhook *entities.WebhookYoutube) error {
	if webhook != nil {
		return service.youtubeWebhookRepo.Delete(*webhook)
	}
	return nil
}
