package channels

import (
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
