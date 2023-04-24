package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
)

func MapGuild(guild entities.Guild) *amqp.ConfigurationGetAnswer {
	serverId := ""
	if guild.ServerId != nil {
		serverId = *guild.ServerId
	}

	return &amqp.ConfigurationGetAnswer{
		GuildId:         guild.Id,
		ServerId:        serverId,
		ChannelServers:  mapChannelServers(guild.ChannelServers),
		AlmanaxWebhooks: mapAlmanaxWebhooks(guild.AlmanaxWebhooks),
		RssWebhooks:     mapRssWebhooks(guild.RssWebhooks),
		TwitterWebhooks: mapTwitterWebhooks(guild.TwitterWebhooks),
	}
}

func mapChannelServers(channelServers []entities.ChannelServer) []*amqp.ConfigurationGetAnswer_ChannelServer {
	result := make([]*amqp.ConfigurationGetAnswer_ChannelServer, 0)

	for _, channelServer := range channelServers {
		result = append(result, &amqp.ConfigurationGetAnswer_ChannelServer{
			ChannelId: channelServer.ChannelId,
			ServerId:  channelServer.ServerId,
		})
	}

	return result
}

func mapAlmanaxWebhooks(webhooks []entities.AlmanaxWebhook) []*amqp.ConfigurationGetAnswer_AlmanaxWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_AlmanaxWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_AlmanaxWebhook{
			ChannelId: webhook.ChannelId,
			Language:  webhook.Language,
		})
	}

	return result
}

func mapRssWebhooks(webhooks []entities.RssWebhook) []*amqp.ConfigurationGetAnswer_RssWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_RssWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_RssWebhook{
			ChannelId: webhook.ChannelId,
			Language:  webhook.Language,
			FeedId:    webhook.FeedId,
		})
	}

	return result
}

func mapTwitterWebhooks(webhooks []entities.TwitterWebhook) []*amqp.ConfigurationGetAnswer_TwitterWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_TwitterWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_TwitterWebhook{
			ChannelId: webhook.ChannelId,
			Language:  webhook.Language,
		})
	}

	return result
}
