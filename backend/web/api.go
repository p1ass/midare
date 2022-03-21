package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/p1ass/midare/period"
)

// GetMe gets my profile.
func (h *Handler) GetMe(c *gin.Context) {
	_, token := h.getAccessToken(c)
	if token == nil {
		return
	}

	user, err := h.usecase.GetUser(c.Request.Context(), token)
	if err != nil {
		sendError(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAwakePeriods gets awake periods from tweets.
func (h *Handler) GetAwakePeriods(c *gin.Context) {
	type getAwakePeriodsRes struct {
		Periods  []*period.Period `json:"periods"`
		ShareURL string           `json:"shareUrl"`
	}

	userID, accessToken := h.getAccessToken(c)
	if accessToken == nil {
		return
	}

	periods, shareURL, err := h.usecase.GetAwakePeriods(c.Request.Context(), userID, accessToken)
	if err != nil {
		sendError(err, c)
		return
	}

	c.JSON(http.StatusOK, &getAwakePeriodsRes{
		Periods:  periods,
		ShareURL: shareURL,
	})
}
