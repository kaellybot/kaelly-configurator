package mappers

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/entities"
)

func MapGuildCreateAnswer(request *amqp.ConfigurationGuildCreateRequest,
	game amqp.Game, created bool) *amqp.RabbitMQMessage {
	return &amqp.RabbitMQMessage{
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Type:     amqp.RabbitMQMessage_CONFIGURATION_GUILD_CREATE_ANSWER,
		Language: amqp.Language_ANY,
		Game:     game,
		ConfigurationGuildCreateAnswer: &amqp.ConfigurationGuildCreateAnswer{
			Id:          request.Id,
			Name:        request.Name,
			MemberCount: request.MemberCount,
			Created:     created,
		},
	}
}

func MapGuildDeleteAnswer(request *amqp.ConfigurationGuildDeleteRequest,
	game amqp.Game, deleted bool) *amqp.RabbitMQMessage {
	return &amqp.RabbitMQMessage{
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Type:     amqp.RabbitMQMessage_CONFIGURATION_GUILD_DELETE_ANSWER,
		Language: amqp.Language_ANY,
		Game:     game,
		ConfigurationGuildDeleteAnswer: &amqp.ConfigurationGuildDeleteAnswer{
			Id:          request.Id,
			Name:        request.Name,
			MemberCount: request.MemberCount,
			Deleted:     deleted,
		},
	}
}

func MapGuild(guild entities.Guild, langage amqp.Language) *amqp.RabbitMQMessage {
	serverID := ""
	if guild.ServerID != nil {
		serverID = *guild.ServerID
	}

	return &amqp.RabbitMQMessage{
		Type:     amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
		Status:   amqp.RabbitMQMessage_SUCCESS,
		Language: langage,
		ConfigurationGetAnswer: &amqp.ConfigurationGetAnswer{
			GuildId:         guild.ID,
			ServerId:        serverID,
			ChannelServers:  mapChannelServers(guild.ChannelServers),
			AlmanaxWebhooks: mapAlmanaxWebhooks(guild.AlmanaxWebhooks),
			RssWebhooks:     mapFeedWebhooks(guild.FeedWebhooks),
			TwitchWebhooks:  mapTwitchWebhooks(guild.TwitchWebhooks),
			TwitterWebhooks: mapTwitterWebhooks(guild.TwitterWebhooks),
			YoutubeWebhooks: mapYoutubeWebhooks(guild.YoutubeWebhooks),
		},
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
			TwitterId: webhook.TwitterAccount.ID,
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
