package action

import (
	"net/http"

	"go-clean-architecture/adapter/api/logging"
	"go-clean-architecture/adapter/api/response"
	"go-clean-architecture/adapter/logger"
	"go-clean-architecture/domain"
	"go-clean-architecture/usecase"
)

type FindAccountBalanceAction struct {
	uc  usecase.FindAccountBalanceUseCase
	log logger.Logger
}

func NewFindAccountBalanceAction(uc usecase.FindAccountBalanceUseCase, log logger.Logger) FindAccountBalanceAction {
	return FindAccountBalanceAction{
		uc:  uc,
		log: log,
	}
}

func (a FindAccountBalanceAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_balance_account"

	var accountID = r.URL.Query().Get("account_id")
	if !domain.IsValidUUID(accountID) {
		var err = response.ErrParameterInvalid
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("invalid parameter")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), domain.AccountID(accountID))
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusBadRequest,
			).Log("error fetching account balance")

			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		default:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error when returning account balance")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning account balance")

	response.NewSuccess(output, http.StatusOK).Send(w)
}
