package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/sachinnagesh/log-middleware/config"
	"github.com/sachinnagesh/log-middleware/internal/model"
)

func PostLog(logs []model.LogPayload) {
	defer config.Wg.Done()
	posturl := os.Getenv("POST_ENDPOINT")
	buf := &bytes.Buffer{}
	err := json.NewEncoder(buf).Encode(logs)
	if err != nil {
		log.Error("Error while encoding logs", err.Error())
	}
	req, err := http.NewRequest("POST", posturl, buf)
	if err != nil {
		log.Error("Error while creating request to post endpoint!!!", err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	startTime := time.Now()

	for i := 0; i < 3; i++ {

		resp, err := client.Do(req)
		if err != nil {
			log.Error("Error while sending post request : ", err.Error())
			if i == 2 {
				log.Error("ERROR sending data to post endpoint even after 3 retries")
				os.Exit(1)
			}
			time.Sleep(2 * time.Second)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {

			log.Info("Data posted to ", posturl, " successfully at : ", time.Now(), ", batchSize : ", len(logs), ", TimeTaken : ", time.Since(startTime).Seconds(), "s")
			break
		} else {
			time.Sleep(2 * time.Second)
			if i == 2 {
				log.Error("ERROR sending data to post endpoint even after 3 retries")
				os.Exit(1)
			}

		}
	}

}
