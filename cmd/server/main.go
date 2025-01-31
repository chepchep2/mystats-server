package main

import (
    "context"
    "log"
    "mystats-server/internal/config"
    "mystats-server/internal/handler"
    "mystats-server/internal/router"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    // Database 설정
    dbConfig := config.NewDatabaseConfig()
    client := config.NewEntClient(dbConfig)
    defer client.Close()

    // 자동 마이그레이션 실행
    if err := client.Schema.Create(context.Background()); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }

    // 핸들러 초기화
    h := handler.NewHandler(client)

    // Fiber 앱 생성
    app := fiber.New()

    // 미들웨어 설정
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowHeaders: "Origin, Content-Type, Accept, Authorization",
        AllowMethods: "GET, POST, PUT, DELETE",
    }))

    // 기본 라우트
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("MyStats API Server")
    })

    // API 라우터 설정
    router.SetupRouter(h, app)

    // 서버 시작
    log.Fatal(app.Listen(":8080"))
}
