package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var cfg *Config

func init() {
	loadConfig()
	if err := cfg.InitConfig(); err != nil {
		logrus.Fatalf("配置文件错误,%v\n", err)
	}
}

func AlertMarshal(alert *Alert) string {
	data, err := json.Marshal(&alert)
	if err != nil {
		fmt.Printf("AlertMarshal json marshal error ,%v", err)
		return ""
	}
	return string(data)
}

func RequestWebHook(data, webhook string) {
	url := fmt.Sprintf("%s/%s", cfg.Srv.WebHookAddress, webhook)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("webhook响应数据: %v\n", string(body))
}

func HandlerWebhook(c *gin.Context) {
	var notification Notification
	webhook := c.Param("webhook")
	if err := c.ShouldBind(&notification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	for _, hook := range cfg.Hooks {
		fmt.Println("Hook:", hook)
		if webhook == hook {
			if notification.Status == "firing" {
				for _, alert := range notification.Alerts {
					fmt.Printf("alert:%v\n", alert)
					if alert.Status == "firing" {
						RequestWebHook(AlertMarshal(&alert), webhook)
					}
				}
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "success",
			})
		} else {
			fmt.Println("webhook is not define.")
		}

	}

}

func main() {
	gin.SetMode(gin.ReleaseMode)
	gin.DisableConsoleColor()
	f, _ := os.OpenFile("gin.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/alert/:webhook", HandlerWebhook)

	srv := &http.Server{
		Addr:    cfg.ListenAddress,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
