package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/s3blob"
)

var serverURL string

func readfile(c *gin.Context) {
	ctx := c.Request.Context()
	bucket, err := blob.OpenBucket(ctx, serverURL)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer bucket.Close()

	reader, err := bucket.NewReader(ctx, "gopher.png", nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer reader.Close()

	c.Header("Content-Type", "image/png")
	io.Copy(c.Writer, reader)

}

func writefile(c *gin.Context) {
	ctx := c.Request.Context()
	bucket, err := blob.OpenBucket(ctx, serverURL)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer bucket.Close()

	writer, err := bucket.NewWriter(ctx, "gopher.png", nil)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer writer.Close()
	_, err = io.Copy(writer, c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func main() {

	os.Setenv("AWS_ACCESS_KEY", "minioadmin")
	os.Setenv("AWS_SECRET_KEY", "minioadmin")

	log.Println("starting...")

	serverURL = "s3://gocloud?endpoint=localhost:9000&disableSSL=true&s3ForcePathStyle=true&region=DE"

	r := gin.Default()
	r.GET("/", readfile)
	r.POST("/", writefile)
	r.Run()
}
