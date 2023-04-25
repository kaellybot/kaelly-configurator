package channels

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
	"github.com/kaellybot/kaelly-configurator/repositories/almanax"
	"github.com/kaellybot/kaelly-configurator/repositories/feeds"
	"github.com/kaellybot/kaelly-configurator/repositories/servers"
	"github.com/kaellybot/kaelly-configurator/repositories/twitter"
)

func New(channelServerRepo servers.ChannelServerRepository, almanaxRepo almanax.AlmanaxWebhookRepository,
	feedRepo feeds.FeedWebhookRepository, twitterRepo twitter.TwitterWebhookRepository) (*ChannelServiceImpl, error) {

	return &ChannelServiceImpl{
		channelServerRepo:  channelServerRepo,
		almanaxWebhookRepo: almanaxRepo,
		feedWebhookRepo:    feedRepo,
		twitterWebhookRepo: twitterRepo,
	}, nil
}

func (service *ChannelServiceImpl) GetAlmanaxWebhook(guildId, channelId string,
	locale amqp.Language) (*entities.WebhookAlmanax, error) {

	return service.almanaxWebhookRepo.Get(guildId, channelId, locale)
}

func (service *ChannelServiceImpl) GetFeedWebhook(guildId, channelId, feedTypeId string,
	locale amqp.Language) (*entities.WebhookFeed, error) {

	return service.feedWebhookRepo.Get(guildId, channelId, feedTypeId, locale)
}

func (service *ChannelServiceImpl) GetTwitterWebhook(guildId, channelId string,
	locale amqp.Language) (*entities.WebhookTwitter, error) {

	return service.twitterWebhookRepo.Get(guildId, channelId, locale)
}

func (service *ChannelServiceImpl) SaveChannelServer(channelServer entities.ChannelServer) error {
	return service.channelServerRepo.Save(channelServer)
}

func (service *ChannelServiceImpl) SaveAlmanaxWebhook(webhook entities.WebhookAlmanax) error {
	return service.almanaxWebhookRepo.Save(webhook)
}

func (service *ChannelServiceImpl) SaveFeedWebhook(webhook entities.WebhookFeed) error {
	return service.feedWebhookRepo.Save(webhook)
}

func (service *ChannelServiceImpl) SaveTwitterWebhook(webhook entities.WebhookTwitter) error {
	return service.twitterWebhookRepo.Save(webhook)
}

func (service *ChannelServiceImpl) DeleteAlmanaxWebhook(webhook *entities.WebhookAlmanax) error {
	if webhook != nil {
		return service.almanaxWebhookRepo.Delete(*webhook)
	}
	return nil
}

func (service *ChannelServiceImpl) DeleteFeedWebhook(webhook *entities.WebhookFeed) error {
	if webhook != nil {
		return service.feedWebhookRepo.Delete(*webhook)
	}
	return nil
}

func (service *ChannelServiceImpl) DeleteTwitterWebhook(webhook *entities.WebhookTwitter) error {
	if webhook != nil {
		return service.twitterWebhookRepo.Delete(*webhook)
	}
	return nil
}
