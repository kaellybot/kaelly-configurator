package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/mappers"
	"github.com/kaellybot/kaelly-configurator/utils/replies"
	"github.com/rs/zerolog/log"
)

func (service *Impl) guildCreateRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationGuildCreateRequest
	if !isValidConfigurationGuildCreateRequest(request) {
		replies.FailedAnswer(ctx, service.broker, amqp.RabbitMQMessage_CONFIGURATION_GUILD_CREATE_ANSWER,
			message.Language)
		return
	}

	created, errCreate := service.guildService.Create(request.Id, message.Game)
	if errCreate != nil {
		log.Warn().Err(errCreate).Msg("Cannot create guild into DB, continuing...")
		replies.FailedAnswer(ctx, service.broker, amqp.RabbitMQMessage_CONFIGURATION_GUILD_CREATE_ANSWER,
			message.Language)
	} else {
		answer := mappers.MapGuildCreateAnswer(request, message.Game, created)
		replies.SucceededAnswer(ctx, service.broker, answer)
	}
}

func (service *Impl) guildDeleteRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationGuildDeleteRequest
	if !isValidConfigurationGuildDeleteRequest(request) {
		replies.FailedAnswer(ctx, service.broker, amqp.RabbitMQMessage_CONFIGURATION_GUILD_DELETE_ANSWER,
			message.Language)
		return
	}

	deleted, errDel := service.guildService.Delete(request.Id, message.Game)
	if errDel != nil {
		log.Warn().Err(errDel).Msg("Cannot delete guild from DB, continuing...")
		replies.FailedAnswer(ctx, service.broker, amqp.RabbitMQMessage_CONFIGURATION_GUILD_DELETE_ANSWER,
			message.Language)
	} else {
		answer := mappers.MapGuildDeleteAnswer(request, message.Game, deleted)
		replies.SucceededAnswer(ctx, service.broker, answer)
	}
}

func isValidConfigurationGuildCreateRequest(request *amqp.ConfigurationGuildCreateRequest) bool {
	return request != nil
}

func isValidConfigurationGuildDeleteRequest(request *amqp.ConfigurationGuildDeleteRequest) bool {
	return request != nil
}
