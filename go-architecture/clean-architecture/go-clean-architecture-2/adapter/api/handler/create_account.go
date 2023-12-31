package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"go-transactions/adapter/api/response"
	"go-transactions/domain"
	"go-transactions/infrastructure/validation"
	"go-transactions/usecase"

	"github.com/go-playground/validator/v10"
)

// CreateAccountHandler defines the dependencies of the HTTP handler for the use case
type CreateAccountHandler struct {
	uc        usecase.CreateAccountUseCase
	log       *log.Logger
	validator *validator.Validate
}

// NewCreateAccountHandler creates new CreateAccountHandler with its dependencies
func NewCreateAccountHandler(
	uc usecase.CreateAccountUseCase,
	log *log.Logger,
	v *validator.Validate,
) CreateAccountHandler {
	return CreateAccountHandler{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

// Handle handles http request
func (c CreateAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateAccountInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		c.log.Println("failed to marshal message:", err)
		response.NewError([]string{err.Error()}, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if err := c.validator.Struct(input); err != nil {
		errs := validation.ErrMessages(err)
		c.log.Println("invalid input:", errs)
		response.NewError(errs, http.StatusBadRequest).Send(w)
		return
	}

	output, err := c.uc.Execute(r.Context(), input)
	if err != nil {
		c.log.Println("failed to creating account:", err)
		switch err {
		case domain.ErrAccountAlreadyExists:
			response.NewError([]string{err.Error()}, http.StatusUnprocessableEntity).Send(w)
			return
		default:
			response.NewError([]string{err.Error()}, http.StatusInternalServerError).Send(w)
			return
		}
	}

	c.log.Println("success to creating account")
	response.NewSuccess(output, http.StatusCreated).Send(w)
}
