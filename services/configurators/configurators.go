package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/kaellybot/kaelly-configurator/services/channels"
	"github.com/kaellybot/kaelly-configurator/services/guilds"
	"github.com/kaellybot/kaelly-configurator/utils/replies"
	"github.com/rs/zerolog/log"
)

func New(broker amqp.MessageBroker, guildService guilds.Service, channelService channels.Service) (*Impl, error) {
	return &Impl{
		guildService:   guildService,
		channelService: channelService,
		broker:         broker,
	}, nil
}

func GetBindings() []amqp.Binding {
	return []amqp.Binding{
		{
			Exchange:   amqp.ExchangeRequest,
			RoutingKey: requestsRoutingkey,
			Queue:      requestQueueName,
		},
		{
			Exchange:   amqp.ExchangeNews,
			RoutingKey: newsRoutingkey,
			Queue:      newsQueueName,
		},
	}
}

func (service *Impl) Consume() {
	log.Info().Msgf("Consuming configurator news and requests...")
	service.broker.Consume(requestQueueName, service.consumeRequests)
	service.broker.Consume(newsQueueName, service.consumeNews)
}

func (service *Impl) consumeRequests(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.GetType() {
	case amqp.RabbitMQMessage_CONFIGURATION_GET_REQUEST:
		service.handle(ctx, message, amqp.RabbitMQMessage_CONFIGURATION_GET_ANSWER,
			service.getRequest)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_REQUEST:
		service.handle(ctx, message, amqp.RabbitMQMessage_CONFIGURATION_SET_SERVER_ANSWER,
			service.serverRequest)
	case amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_REQUEST:
		service.handle(ctx, message, amqp.RabbitMQMessage_CONFIGURATION_SET_NOTIFICATION_ANSWER,
			service.notificationRequest)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Type not recognized, request ignored")
	}
}

// handle validates the request game before running the handler, so that a request
// never reads or writes another game's configuration.
func (service *Impl) handle(ctx amqp.Context, message *amqp.RabbitMQMessage,
	answerType amqp.RabbitMQMessage_Type, handler requestHandler) {
	// The broker guards already refuse these, so reaching this point means the guard
	// was bypassed. Answering is impossible: the reply would carry ANY_GAME too.
	if !message.IsGameSet() {
		log.Error().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Request without game received, request ignored")
		return
	}

	if !isGameSupported(message.GetGame()) {
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Str(constants.LogGame, message.GetGame().String()).
			Msgf("Game not supported, request refused")
		replies.FailedAnswer(ctx, service.broker, message, answerType)
		return
	}

	handler(ctx, message)
}

// isGameSupported reports whether the configurator can store configuration for this
// game. It is game-generic, so a game is supported as soon as its bot exists.
func isGameSupported(game amqp.Game) bool {
	switch game {
	case amqp.Game_DOFUS_GAME, amqp.Game_DOFUS_TOUCH:
		return true
	case amqp.Game_ANY_GAME, amqp.Game_DOFUS_RETRO:
		return false
	default:
		return false
	}
}

func (service *Impl) notificationRequest(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	request := message.ConfigurationSetNotificationRequest
	if !isValidNotificationRequest(request) {
		service.publishFailedSetNotificationAnswer(ctx, message, "")
		return
	}

	switch request.NotificationType {
	case amqp.NotificationType_ALMANAX:
		service.almanaxRequest(ctx, message)
	case amqp.NotificationType_RSS:
		service.rssRequest(ctx, message)
	case amqp.NotificationType_TWITTER:
		service.twitterRequest(ctx, message)
	case amqp.NotificationType_UNKNOWN:
		fallthrough
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Notification type not recognized, request ignored")
	}
}

func (service *Impl) consumeNews(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	//exhaustive:ignore Don't need to be exhaustive here since they will be handled by default case
	switch message.GetType() {
	case amqp.RabbitMQMessage_NEWS_GUILD:
		// News gets no answer, so an unusable game is only logged and dropped: acting
		// on it would create or delete a guild row under the wrong game.
		if !message.IsGameSet() || !isGameSupported(message.GetGame()) {
			log.Error().
				Str(constants.LogCorrelationID, ctx.CorrelationID).
				Str(constants.LogGame, message.GetGame().String()).
				Msgf("Guild news with an unusable game received, news ignored")
			return
		}

		service.guildNews(message)
	default:
		log.Warn().
			Str(constants.LogCorrelationID, ctx.CorrelationID).
			Msgf("Type not recognized, news ignored")
	}
}

func isValidNotificationRequest(request *amqp.ConfigurationSetNotificationRequest) bool {
	return request != nil
}
