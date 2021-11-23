package processor

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"
	"github.com/ozonmp/bss-workplace-api/internal/config"
	"github.com/ozonmp/bss-workplace-api/internal/infra/kafka"
	"github.com/ozonmp/bss-workplace-api/internal/model"
	"github.com/ozonmp/bss-workplace-api/internal/repo"
	bss_workplace_facade "github.com/ozonmp/bss-workplace-facade/pkg/bss-workplace-facade"
	"github.com/rs/zerolog/log"
)

type EventProcessor struct {
	EventsRepo repo.WorkplaceEventRepo
	Consumer   *kafka.Consumer
	topic      string
}

const BATCH_SIZE = 10

func NewEventsProcessor(cfg config.Config, db *sqlx.DB) (*EventProcessor, error) {
	eventsRepo := repo.NewWorkplaceEventRepo(db, BATCH_SIZE)

	consumer, err := kafka.NewKafkaConsumer(cfg.Kafka.Brokers, cfg.Kafka.GroupID,
		func(ctx context.Context, message *sarama.ConsumerMessage) error {
			var workplaceEvent bss_workplace_facade.WorkplaceEvent

			err := proto.Unmarshal(message.Value, &workplaceEvent)
			if err != nil {
				log.Error().Msgf("Message unmarshall error %v", err)
				return err
			}

			eventsRepo.Add(ctx, *model.CreateEventFromProtoEvent(workplaceEvent))

			log.Debug().Msgf("New kafka message from topic %v", message.Topic)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return &EventProcessor{
		EventsRepo: eventsRepo,
		Consumer:   consumer,
		topic:      cfg.Kafka.Topic,
	}, nil

}

func (e *EventProcessor) StartProcessor(ctx context.Context) {
	e.Consumer.StartConsuming(ctx, e.topic)
}
