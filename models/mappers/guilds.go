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
		ChannelWebhooks: mapChannelWebhooks(guild.ChannelWebhooks),
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

func mapChannelWebhooks(channelWebhooks []entities.ChannelWebhook) []*amqp.ConfigurationGetAnswer_ChannelWebhook {
	result := make([]*amqp.ConfigurationGetAnswer_ChannelWebhook, 0)

	for _, channelWebhook := range channelWebhooks {
		result = append(result, &amqp.ConfigurationGetAnswer_ChannelWebhook{
			ChannelId: channelWebhook.ChannelId,
			Provider:  amqp.ConfigurationGetAnswer_ChannelWebhook_RSS, // TODO
			Language:  amqp.Language_ANY,                              // TODO
		})
	}

	return result
}
