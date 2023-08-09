package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/phincon-backend/laza/helper"
	"io"
	"net/url"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := helper.GetLogrusLogger()
		startTime := time.Now()

		// Capture the request details
		method := c.Request.Method
		path := c.Request.URL.Path
		pathParam := c.Params
		query, _ := url.QueryUnescape(c.Request.URL.RawQuery)
		ip := c.ClientIP()

		// Read and capture the request body
		var bodyBytes []byte
		if c.Request.Header.Get("Content-Type") != "multipart/form-data" {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// Continue handling the request
		c.Next()

		// Calculate duration
		duration := time.Since(startTime)

		// Get response status code
		statusCode := c.Writer.Status()

		// Get response size (content length)
		respSize := c.Writer.Size()

		// Get response error message if present
		errMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		// Log the information
		logger.Infof("Method: %s, Path: %s, PathParam: %s Query: %s, IP: %s, RequestBody: %s, StatusCode: %d, Duration: %s, ResponseSize: %d, ErrorMessage: %s",
			method, path, pathParam, query, ip, string(bodyBytes), statusCode, duration, respSize, errMessage)
	}
}
