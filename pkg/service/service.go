package service

import (
	"pishikot/pkg/models"
	"pishikot/pkg/service/eventservice"
	"pishikot/pkg/service/middlewares"
	"pishikot/pkg/store"
	"pishikot/pkg/store/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Service struct {
	User         *db.UserRepo
	Task         *db.TaskRepo
	EventService *eventservice.EventService
	Config       *models.Config
}

func NewService(store *store.Store) *Service {
	return &Service{
		User:         store.UserRepo,
		Task:         store.TaskRepo,
		Config:       store.Config,
		EventService: eventservice.NewEventService(),
	}
}

func (s *Service) Start() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"*"},
		AllowWebSockets:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
	}))
	r.POST("/users/login", s.Login)
	r.POST("/users/register", s.Register)
	authMiddleware, err := middlewares.GetJWTMiddleware(*s.Config)
	if err != nil {
		panic(err)
	}
	a := r.Group("/api", authMiddleware.MiddlewareFunc())
	a.GET("/auth/accessible", s.Accessible)
	a.GET("/user", s.GetUser)
	a.POST("/check_task", s.CheckTask)
	a.GET("/auth/refresh_token", authMiddleware.RefreshHandler)
	a.GET("/ws/event", s.UpgradeConnection)
	r.Run()
}
