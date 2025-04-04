package channels

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
)

func New(channelServerRepo servers.Repository, almanaxRepo almanax.Repository,
	feedRepo feeds.Repository, twitterRepo twitter.Repository) (*Impl, error) {
	return &Impl{
		channelServerRepo:  channelServerRepo,
		almanaxWebhookRepo: almanaxRepo,
		feedWebhookRepo:    feedRepo,
		twitterWebhookRepo: twitterRepo,
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

func (service *Impl) GetTwitterWebhook(guildID, channelID, twitterID string,
) (*entities.WebhookTwitter, error) {
	return service.twitterWebhookRepo.Get(guildID, channelID, twitterID)
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

func (service *Impl) SaveTwitterWebhook(webhook entities.WebhookTwitter) error {
	return service.twitterWebhookRepo.Save(webhook)
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

func (service *Impl) DeleteTwitterWebhook(webhook *entities.WebhookTwitter) error {
	if webhook != nil {
		return service.twitterWebhookRepo.Delete(*webhook)
	}
	return nil
}
