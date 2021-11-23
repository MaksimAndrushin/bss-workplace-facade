package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type ConsumeFunction func(ctx context.Context, message *sarama.ConsumerMessage) error

type Consumer struct {
	fn            ConsumeFunction
	ConsumerGroup sarama.ConsumerGroup
}

func NewKafkaConsumer(brokers []string, group string, consumeFunction ConsumeFunction) (*Consumer, error) {

	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		fn:            consumeFunction,
		ConsumerGroup: consumerGroup,
	}, nil
}

func (c *Consumer) StartConsuming(ctx context.Context, topic string) error {
	go func() {
		for {
			if err := c.ConsumerGroup.Consume(ctx, []string{topic}, c); err != nil {
				log.Error().Msgf("Error from Consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
		}
	}()

	return nil
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		ctx := context.TODO()

		err := c.fn(ctx, message)
		if err != nil {
			log.Error().Msgf("Kafka message processing error - %v", err)
		} else {
			session.MarkMessage(message, "")
		}
	}

	return nil
}
