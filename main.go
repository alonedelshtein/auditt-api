package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	helpers "auditt-api/helpers"
	model "auditt-api/model"
)

func main() {

	//db = utils.Connect("")

	gin.SetMode(gin.ReleaseMode)

	// Logging to a file.
	f, _ := os.Create("auditt-api.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")
	router.Static("/css", "./css")
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		header := param.Request.Header
		errorMessageLen := len(param.ErrorMessage)
		if errorMessageLen == 0 {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" \nHeader: %s",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				header,
			)
		} else {
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" Error: %s\"\nHeader: %s",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
				header,
			)
		}
	}))

	router.Use(cors.Default())

	router.POST("/service/v1/pullrequest", postPullRequest)
	router.GET("/view/v1/pullrequest", getPullRequest)

	router.Run("0.0.0.0" + ":" + "8082")
}
func getPullRequest(c *gin.Context) {
	var prs []model.PullRequestDB

	db := helpers.DBConnect()

	db.Find(&prs)

	c.HTML(http.StatusOK, "gridview", gin.H{
		"prs": prs,
	})
}

func postPullRequest(c *gin.Context) {
	var request map[string]interface{}
	buf, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, buf); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	}
	err = json.Unmarshal(buffer.Bytes(), &request)
	if err != nil {
		fmt.Println("error:", err)
	}
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf))
	pr := model.PullRequestDBFromGitHubEvent(request)
	reader := helpers.ScreenshotByURL(pr.URL)
	pr.Screenshot = helpers.UploadFileS3(reader, fmt.Sprintf("%s.jpeg", strconv.Itoa(pr.Id)))
	db := helpers.DBConnect()
	db.Create(&pr)

	c.IndentedJSON(http.StatusCreated, pr)
}
