package configurators

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-configurator/models/constants"
	"github.com/rs/zerolog/log"
)

func (service *Impl) guildNews(message *amqp.RabbitMQMessage) {
	newsGuild := message.NewsGuildMessage

	switch newsGuild.Event {
	case amqp.NewsGuildMessage_CREATE:
		service.guildCreateRequest(newsGuild.Id, message.Game)
	case amqp.NewsGuildMessage_DELETE:
		service.guildDeleteRequest(newsGuild.Id, message.Game)
	case amqp.NewsGuildMessage_UNKNOWN:
		fallthrough
	default:
		log.Warn().
			Str(constants.LogEvent, newsGuild.Event.String()).
			Msg("Guild event not handled, ignoring it")
		return
	}
}

func (service *Impl) guildCreateRequest(guildID string, game amqp.Game) {
	errCreate := service.guildService.Create(guildID, game)
	if errCreate != nil {
		log.Warn().Err(errCreate).
			Msg("Cannot create guild into DB, continuing...")
	}
}

func (service *Impl) guildDeleteRequest(guildID string, game amqp.Game) {
	errDel := service.guildService.Delete(guildID, game)
	if errDel != nil {
		log.Warn().Err(errDel).
			Msg("Cannot delete guild from DB, continuing...")
	}
}
