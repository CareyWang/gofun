package handles

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
	"net/url"
)

func GenerateSurge(context *gin.Context) {
	urlInput := context.Param("url")
	_subscriptionUrl, err := base64.URLEncoding.DecodeString(urlInput)
	if err != nil {
		context.JSON(500, err)
		return
	}
	subscriptionUrl := string(_subscriptionUrl)

	path := "resources/yoyu.conf"

	b, err := ioutil.ReadFile(path)
	if err != nil {
		context.JSON(500, err)
		return
	}

	conf := string(b)

	conf = strings.ReplaceAll(conf, ":url", subscriptionUrl)
	conf = strings.ReplaceAll(conf, ":eUrl", url.QueryEscape(subscriptionUrl))
	context.String(200, conf)
}
