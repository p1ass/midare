package usecase

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/mrjones/oauth"
	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/entity"
	"github.com/p1ass/midare/lib/logging"
	"go.uber.org/zap"
)

func (u *Usecase) UploadImage(periods []*entity.Period, shareID string, accessToken *oauth.AccessToken) string {
	logging.New().Info("uploadImage", zap.String("uuid", shareID))
	go u.uploadImageThroughCloudFunctions(shareID, periods, accessToken)

	return config.ReadAllowCORSOriginURL() + "/share/" + shareID
}

func (u *Usecase) uploadImageThroughCloudFunctions(uuid string, periods []*entity.Period, accessToken *oauth.AccessToken) {
	type request struct {
		Name    string           `json:"name"`
		IconURL string           `json:"iconUrl"`
		UUID    string           `json:"uuid"`
		Periods []*entity.Period `json:"periods"`
	}

	user, err := u.twiCli.AccountVerifyCredentials(accessToken)
	if err != nil {
		logging.New().Error("uploadImageThroughCloudFunctions: get account info" + err.Error())
		return
	}

	req := &request{
		Name:    user.Name,
		IconURL: user.ImageURL,
		UUID:    uuid,
		Periods: periods,
	}
	encoded, _ := json.Marshal(req)

	_, err = http.Post(config.ReadCloudFunctionsURL(), "application/json", bytes.NewBuffer(encoded))
	if err != nil {
		logging.New().Error("post period data to cloud functions" + err.Error())
	}
}
