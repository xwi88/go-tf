// Package main
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/cobra"
	"github.com/xwi88/kit4go/datetime"
)

var (
	pid       int
	startTime time.Time
)

// startCMD real time service
var startCMD = &cobra.Command{
	Use:       "start",
	Short:     "Start the app service",
	Long:      ``,
	Example:   "app start\n  app start -c [file]\n  app start --config [file]",
	ValidArgs: []string{"start"},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("[app] PreRun load config file:%v", *confFile)
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pid = syscall.Getpid()
		log.Printf("[app] Run with pid: %v, at: %v", pid, datetime.GetNowWithZone(nil))

		app := fiber.New(fiber.Config{
			Prefork:               false,
			CaseSensitive:         true,
			StrictRouting:         true,
			DisableStartupMessage: true,
		})
		app.Use(recover.New())
		app.Use(compress.New())
		app.Use(cors.New(cors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		}))
		app.Use(requestid.New())
		app.Use(logger.New(logger.Config{
			Format:     "[${time}] pid(${pid}) request_id(${locals:requestid}) | ${status} | ${latency} | ${method} | ${path}​\n​",
			TimeFormat: "2006-01-02 15:04:05",
			TimeZone:   "Asia/Shanghai",
		}))

		// Match any route
		app.Use(func(c *fiber.Ctx) error {
			fmt.Println("🥇 First handler")
			return c.Next()
		})

		// Match all routes starting with /api
		// app.Use("/api", func(c *fiber.Ctx) error {
		// 	fmt.Println("🥈 Second handler")
		// 	return c.Next()
		// })

		// GET /api/register
		app.Get("/api/list", func(c *fiber.Ctx) error {
			fmt.Println("🥉 Last handler")
			return c.SendString("Hello, World 👋!")
		})

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGINT)
		go func() {
			startTime = time.Now()
			<-ch
			log.Printf("Recieve kill signal, terminating...")
			// TODO destroy resources
			elapse := int64(time.Since(startTime) / time.Second)
			log.Printf("closed on pid: %d, elapse: %ds", pid, elapse)
			signal.Stop(ch)
			os.Exit(0)
		}()

		log.Fatal(app.Listen(":6666"))
	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}
