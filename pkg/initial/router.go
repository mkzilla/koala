package initial

import (
	"github.com/mkzilla/koala/pkg/api/auth"
	cfv2 "github.com/mkzilla/koala/pkg/api/config"
	"github.com/mkzilla/koala/pkg/api/group"
	"github.com/mkzilla/koala/pkg/api/middleware"
	"github.com/mkzilla/koala/pkg/api/notification"
	"github.com/mkzilla/koala/pkg/api/statistics"
	"github.com/mkzilla/koala/pkg/api/task"
	"github.com/mkzilla/koala/pkg/api/technique"
	"github.com/mkzilla/koala/pkg/api/user"
	"github.com/mkzilla/koala/pkg/api/weekly"
	"github.com/mkzilla/koala/pkg/types/config"
)

func InitRouter() {
	if config.Configs.BasicLogin.Enable {
		config.GinEngine.POST("/api/v1beta/login", auth.LoginWithPassword)
	}
	config.GinEngine.GET("/api/v1beta/login", auth.LoginWithOAuth2)
	config.GinEngine.GET("/api/v1beta/config", cfv2.Echo)
	if config.Configs.BasicRegister.Enable {
		config.GinEngine.POST("/api/v1beta/register", auth.Request)
		config.GinEngine.POST("/api/v1beta/password", auth.RegisterAuth, auth.InitPassword)
	}

	apv1 := config.GinEngine.Group("/api/v1").Use(auth.Auth, auth.IsGroupUser, middleware.Page)
	apv1.GET("/userinfo", auth.UserInfo)

	apv1.GET("/tasks", task.List)
	apv1.GET("/tasks/recent", task.Recent)
	apv1.GET("/tasks/kanban", task.Kanban)
	apv1.POST("/task", task.Create)
	apv1.GET("/task/:task_id", task.View)
	apv1.POST("/task/:task_id", task.Update)
	apv1.GET("/task/:task_id/vote", task.ListVote)
	apv1.POST("/task/:task_id/vote", task.Vote)
	apv1.GET("/task/:task_id/watch", task.ListWatch)
	apv1.POST("/task/:task_id/watch", task.Watch)
	apv1.GET("/task/:task_id/comment", task.ListComment)
	apv1.POST("/task/:task_id/comment", task.Comment)
	apv1.GET("/okrs", task.GetOKRByUser)
	apv1.POST("/okr", task.CreateUserOKR)
	apv1.PUT("/okr", task.EditUserOKR)

	apv1.POST("/technique", technique.StartTechnique)
	apv1.GET("/technique", technique.GetRecentTechnique)

	apv1.GET("/weekly/self", weekly.ListSelf)
	apv1.GET("/weekly/groups/:group_id", weekly.ListGroup)

	apv1.GET("/users/assign", user.AssignList)

	apv1.GET("/notifications", notification.List)
	apv1.POST("/notification/:notification_id", notification.DoRead)

	apv1manger := config.GinEngine.Group("/api/v1/manage").Use(auth.Auth, auth.IsGroupUser, middleware.Page)

	apv1manger.POST("/groups", group.Create)
	apv1manger.GET("/groups/:group_id", group.ListUsers)
	apv1manger.GET("/groups", group.ListIn)
	apv1manger.GET("/group/:group_id", group.GetGroup)
	apv1manger.GET("/group/:group_id/okr", group.GetOKRByGroup)
	apv1manger.POST("/group/:group_id/okr", group.CreateGroupOKR)
	apv1manger.PUT("/group/:group_id/okr", group.EditGroupOKR)
	//apv1manger.DELETE("/group/:group_id/okr/:okr_id", group.CreateGroupOKR)

	apv1manger.POST("/groups/:group_id/user/:username", group.AddUser)
	apv1manger.DELETE("/groups/:group_id/user/:username", group.RemoveUser)

	apv2manger := config.GinEngine.Group("/api/v1/manage").Use(auth.Auth, middleware.Page)
	apv2manger.POST("/groups/:group_id/checkin", group.CheckIn)
	apv2manger.POST("/groups/:group_id/checkout", group.CheckOut)

	admpi := apv1manger.Use(auth.IsAdmin)
	admpi.POST("/users", user.Create)
	admpi.GET("/users", user.List)
	admpi.PUT("/users/:user_id", user.UpdateUser)
	admpi.GET("/statistics", statistics.Statistics)
}
