package handlers

import (
	"fmt"
	"net/http"

	"github.com/rarimo/geo-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateBalanceV2(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewCreateBalanceV2(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	nullifier := req.Data.ID
	if !auth.Authenticates(UserClaims(r), auth.UserGrant(nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(nullifier).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if balance != nil {
		ape.RenderErr(w, problems.Conflict())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		events := prepareEventsWithRef(nullifier, "", false, r)
		Log(r).Debugf("%d events will be added for nullifier=%s", len(events), nullifier)
		if err = EventsQ(r).Insert(events...); err != nil {
			return fmt.Errorf("add open events: %w", err)
		}

		if err = BalancesQ(r).Insert(data.Balance{Nullifier: nullifier}); err != nil {
			return fmt.Errorf("failed to insert balance: %w", err)
		}
		return nil
	})
	if err != nil {
		Log(r).WithError(err).Error("Failed to create balance")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
