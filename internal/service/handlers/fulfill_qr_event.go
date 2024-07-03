package handlers

import (
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/rarimo/decentralized-auth-svc/pkg/auth"
	"github.com/rarimo/geo-points-svc/internal/data"
	"github.com/rarimo/geo-points-svc/internal/data/evtypes"
	"github.com/rarimo/geo-points-svc/internal/service/hmacsig"
	"github.com/rarimo/geo-points-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func FulfillQREvent(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewFulfillQREvent(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	event, err := EventsQ(r).FilterByID(req.Data.ID).FilterByStatus(data.EventOpen).Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get event by ID")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if event == nil {
		Log(r).Debugf("Event not found for id=%s status=%s", req.Data.ID, data.EventOpen)
		ape.RenderErr(w, problems.NotFound())
		return
	}

	if !auth.Authenticates(UserClaims(r), auth.UserGrant(event.Nullifier)) {
		ape.RenderErr(w, problems.Unauthorized())
		return
	}

	gotSig := r.Header.Get("Signature")
	wantSig, err := hmacsig.CalculateQREventSignature(SigVerifier(r), event.Nullifier, event.ID, req.Data.Attributes.QrCode)
	if err != nil { // must never happen due to preceding validation
		Log(r).WithError(err).Error("Failed to calculate HMAC signature")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	if gotSig != wantSig {
		log.Warnf("QR event fulfillment unauthorized access: HMAC signature mismatch: got %s, want %s", gotSig, wantSig)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	evType := EventTypes(r).Get(event.Type, evtypes.FilterInactive)
	if evType == nil {
		Log(r).Infof("Event type %s is inactive", event.Type)
		ape.RenderErr(w, problems.Forbidden())
		return
	}
	if evType.QRCodeValue != req.Data.Attributes.QrCode {
		Log(r).Debugf("QR code for event %s doesn't match: got %s, want %s", event.Type, req.Data.Attributes.QrCode, evType.QRCodeValue)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	balance, err := BalancesQ(r).FilterByNullifier(event.Nullifier).FilterDisabled().Get()
	if err != nil {
		Log(r).WithError(err).Error("Failed to get balance by nullifier")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if balance == nil {
		Log(r).Infof("Balance nullifier=%s is disabled", event.Nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	if !evType.AutoClaim {
		_, err = EventsQ(r).FilterByID(event.ID).Update(data.EventFulfilled, nil, nil)
		if err != nil {
			Log(r).WithError(err).Error("Failed to update event status")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, false))
		return
	}

	if !balance.IsVerified {
		Log(r).Infof("Balance nullifier=%s is not verified, fulfill or claim not allowed", event.Nullifier)
		ape.RenderErr(w, problems.Forbidden())
		return
	}

	err = EventsQ(r).Transaction(func() error {
		event, err = claimEvent(r, event, balance)
		return err
	})
	if err != nil {
		Log(r).WithError(err).Errorf("Failed to claim event %s and accrue %d points to the balance %s",
			event.ID, evType.Reward, event.Nullifier)
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, newEventClaimingStateResponse(balance.Nullifier, true))
}
