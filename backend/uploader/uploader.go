package uploader

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/p1ass/midare/config"
	"github.com/p1ass/midare/logging"
	"github.com/p1ass/midare/period"
	"github.com/p1ass/midare/twitter"
	"go.uber.org/zap"
)

type ImageUploader struct {
}

func NewImageUploader() *ImageUploader {
	return &ImageUploader{}
}

// Upload uploads image to cloud storage via Cloud Functions and returns share URL.
func (u *ImageUploader) Upload(periods []*period.Period, shareID string, twiCli twitter.Client) *url.URL {
	logging.New().Info(fmt.Sprintf("uploadImage: %s", shareID), zap.String("uuid", shareID))
	go u.uploadImageThroughCloudFunctions(shareID, periods, twiCli)

	parsed, err := url.Parse(config.ReadAllowCORSOriginURL())
	if err != nil {
		panic(err)
	}
	parsed.Path = path.Join(parsed.Path, "share", shareID)

	return parsed
}

func (u *ImageUploader) uploadImageThroughCloudFunctions(uuid string, periods []*period.Period, twiCli twitter.Client) {
	type request struct {
		Name    string           `json:"name"`
		IconURL string           `json:"iconUrl"`
		UUID    string           `json:"uuid"`
		Periods []*period.Period `json:"periods"`
	}

	// 本当はここでAPIを叩きたくないが、レイテンシの削減のために非同期でAPIを叩きたいため、ここで叩いている
	user, err := twiCli.GetMe(context.Background())
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
