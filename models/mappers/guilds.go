package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
)

func MapGuild(guild entities.Guild) *amqp.ConfigurationGetAnswer {
	serverID := ""
	if guild.ServerID != nil {
		serverID = *guild.ServerID
	}

	return &amqp.ConfigurationGetAnswer{
		GuildId:         guild.ID,
		ServerId:        serverID,
		ChannelServers:  mapChannelServers(guild.ChannelServers),
		AlmanaxWebhooks: mapAlmanaxWebhooks(guild.AlmanaxWebhooks),
		RssWebhooks:     mapFeedWebhooks(guild.FeedWebhooks),
		TwitchWebhooks:  mapTwitchWebhooks(guild.TwitchWebhooks),
		TwitterWebhooks: mapTwitterWebhooks(guild.TwitterWebhooks),
		YoutubeWebhooks: mapYoutubeWebhooks(guild.YoutubeWebhooks),
	}
}

func mapChannelServers(channelServers []entities.ChannelServer) []*amqp.ConfigurationGetAnswer_ChannelServer {
	result := make([]*amqp.ConfigurationGetAnswer_ChannelServer, 0)

	for _, channelServer := range channelServers {
		result = append(result, &amqp.ConfigurationGetAnswer_ChannelServer{
			ChannelId: channelServer.ChannelID,
			ServerId:  channelServer.ServerID,
		})
	}

	return result
}

func mapAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) []*amqp.ConfigurationGetAnswer_AlmanaxWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_AlmanaxWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_AlmanaxWebhook{
			ChannelId: webhook.ChannelID,
			WebhookId: webhook.WebhookID,
			Language:  webhook.Locale,
		})
	}

	return result
}

func mapFeedWebhooks(webhooks []entities.WebhookFeed) []*amqp.ConfigurationGetAnswer_RssWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_RssWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_RssWebhook{
			ChannelId: webhook.ChannelID,
			WebhookId: webhook.WebhookID,
			FeedId:    webhook.FeedTypeID,
			Language:  webhook.Locale,
		})
	}

	return result
}

func mapTwitchWebhooks(webhooks []entities.WebhookTwitch) []*amqp.ConfigurationGetAnswer_TwitchWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_TwitchWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_TwitchWebhook{
			ChannelId:  webhook.ChannelID,
			WebhookId:  webhook.WebhookID,
			StreamerId: webhook.StreamerID,
		})
	}

	return result
}

func mapTwitterWebhooks(webhooks []entities.WebhookTwitter) []*amqp.ConfigurationGetAnswer_TwitterWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_TwitterWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_TwitterWebhook{
			ChannelId: webhook.ChannelID,
			WebhookId: webhook.WebhookID,
			Name:      webhook.TwitterAccount.Name,
			Language:  webhook.Locale,
		})
	}

	return result
}

func mapYoutubeWebhooks(webhooks []entities.WebhookYoutube) []*amqp.ConfigurationGetAnswer_YoutubeWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_YoutubeWebhook, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_YoutubeWebhook{
			ChannelId: webhook.ChannelID,
			WebhookId: webhook.WebhookID,
			VideastId: webhook.VideastID,
		})
	}

	return result
}
