package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
)

func MapGuild(guild entities.Guild, langage amqp.Language) *amqp.RabbitMQMessage {
	serverID := ""
	if guild.ServerID != nil {
		serverID = *guild.ServerID
	}

	notifiedChannels := make([]*amqp.ConfigurationGetAnswer_NotifiedChannel, 0)
	notifiedChannels = append(notifiedChannels, mapAlmanaxWebhooks(guild.AlmanaxWebhooks)...)
	notifiedChannels = append(notifiedChannels, mapFeedWebhooks(guild.FeedWebhooks)...)
	notifiedChannels = append(notifiedChannels, mapTwitterWebhooks(guild.TwitterWebhooks)...)

	return &amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: langage,
		ConfigurationGetAnswer: &amqp.ConfigurationGetAnswer{
			GuildId:          guild.ID,
			ServerId:         serverID,
			ServerChannels:   mapServerChannels(guild.ChannelServers),
			NotifiedChannels: notifiedChannels,
		},
	}
}

func mapServerChannels(channelServers []entities.ChannelServer) []*amqp.ConfigurationGetAnswer_ServerChannel {
	result := make([]*amqp.ConfigurationGetAnswer_ServerChannel, 0)

	for _, channelServer := range channelServers {
		result = append(result, &amqp.ConfigurationGetAnswer_ServerChannel{
			ChannelId: channelServer.ChannelID,
			ServerId:  channelServer.ServerID,
		})
	}

	return result
}

func mapAlmanaxWebhooks(webhooks []entities.WebhookAlmanax) []*amqp.ConfigurationGetAnswer_NotifiedChannel {
	result := make([]*amqp.ConfigurationGetAnswer_NotifiedChannel, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_NotifiedChannel{
			ChannelId:        webhook.ChannelID,
			WebhookId:        webhook.WebhookID,
			NotificationType: amqp.NotificationType_ALMANAX,
		})
	}

	return result
}

func mapFeedWebhooks(webhooks []entities.WebhookFeed) []*amqp.ConfigurationGetAnswer_NotifiedChannel {
	result := make([]*amqp.ConfigurationGetAnswer_NotifiedChannel, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_NotifiedChannel{
			ChannelId:        webhook.ChannelID,
			WebhookId:        webhook.WebhookID,
			Label:            webhook.FeedTypeID,
			NotificationType: amqp.NotificationType_RSS,
		})
	}

	return result
}

func mapTwitterWebhooks(webhooks []entities.WebhookTwitter) []*amqp.ConfigurationGetAnswer_NotifiedChannel {
	result := make([]*amqp.ConfigurationGetAnswer_NotifiedChannel, 0)
	for _, webhook := range webhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_NotifiedChannel{
			ChannelId:        webhook.ChannelID,
			WebhookId:        webhook.WebhookID,
			Label:            webhook.TwitterAccount.ID,
			NotificationType: amqp.NotificationType_TWITTER,
		})
	}

	return result
}
