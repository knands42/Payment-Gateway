package process_transaction

import (
	"context"
	tracer_adapter "github.com/caiofernandes00/payment-gateway/adapter/trace"
	"log"

	"github.com/caiofernandes00/payment-gateway/adapter/broker"
	"github.com/caiofernandes00/payment-gateway/domain/entity"
	"github.com/caiofernandes00/payment-gateway/domain/repository"
)

type ProcessTransaction struct {
	transactionRepository repository.TransactionRepository
	Producer              broker.ProducerInterface
	Topic                 string
	otel                  tracer_adapter.TraceClosure
}

func NewProcessTransaction(transactionRepository repository.TransactionRepository, producer broker.ProducerInterface, topic string, otel tracer_adapter.TraceClosure) *ProcessTransaction {
	return &ProcessTransaction{
		transactionRepository: transactionRepository,
		Producer:              producer,
		Topic:                 topic,
		otel:                  otel,
	}
}

func (p *ProcessTransaction) Execute(ctx context.Context, input TransactionDTOInput) (out TransactionDTOOutput, err error) {
	log.Printf("Processing transaction %s", input.ID)

	p.otel(ctx, "process-transaction-validation", func(ctx context.Context) {
		transaction, err := entity.NewTransaction(input.ID, input.AccountId, input.Amount)
		if err != nil {
			out, err = p.handleRejectedTransaction(ctx, input, err.Error())
			return
		}

		cc, err := entity.NewCreditCard(input.CreditCardNumber, input.CreditCardName, input.CreditCardExpirationMonth, input.CreditCardExpirationYear, input.CreditCardCVV)
		if err != nil {
			out, err = p.handleRejectedTransaction(ctx, input, err.Error())
			return
		}

		transaction.SetCreditCard(*cc)
		out, err = p.handleApprovedTransaction(ctx, input)
	})

	return
}

func (p *ProcessTransaction) handleRejectedTransaction(ctx context.Context, input TransactionDTOInput, errorMessage string) (out TransactionDTOOutput, err error) {
	p.otel(ctx, "handle-rejected-transaction", func(ctx context.Context) {
		err = p.transactionRepository.Insert(ctx, input.ID, input.AccountId, entity.STATUS_REJECTED, errorMessage, input.Amount)
		if err != nil {
			log.Printf("Transaction %s is invalid with error: %s", input.ID, err.Error())
			out = TransactionDTOOutput{}
			return
		}

		out = TransactionDTOOutput{
			ID:           input.ID,
			Status:       entity.STATUS_REJECTED,
			ErrorMessage: errorMessage,
		}

		err = p.publish(ctx, out, []byte(input.ID))
		if err != nil {
			log.Printf("Transaction %s is invalid with error: %s", input.ID, err.Error())
			out = TransactionDTOOutput{}
			return
		}

		log.Printf("Failed Transaction %s with output: %s", input.ID, out)
	})

	return
}

func (p *ProcessTransaction) handleApprovedTransaction(ctx context.Context, input TransactionDTOInput) (out TransactionDTOOutput, err error) {
	p.otel(ctx, "usecase-handle-approved-transaction", func(ctx context.Context) {
		err = p.transactionRepository.Insert(ctx, input.ID, input.AccountId, entity.STATUS_APPROVED, "", input.Amount)
		if err != nil {
			log.Printf("Transaction %s is invalid with error: %s", input.ID, err.Error())
			out = TransactionDTOOutput{}
			return
		}

		out = TransactionDTOOutput{
			ID:           input.ID,
			Status:       entity.STATUS_APPROVED,
			ErrorMessage: "",
		}

		err = p.publish(ctx, out, []byte(input.ID))
		if err != nil {
			log.Printf("Transaction %s is invalid with error: %s", input.ID, err.Error())
			out = TransactionDTOOutput{}
			return
		}

		log.Printf("Approved Transaction %s with output value: %s", input.ID, out)
	})
	return
}

func (p *ProcessTransaction) publish(ctx context.Context, output TransactionDTOOutput, key []byte) (err error) {
	p.otel(ctx, "usecase-publish-transaction", func(ctx context.Context) {
		err = p.Producer.Publish(ctx, output, key, p.Topic)
	})

	return
}
