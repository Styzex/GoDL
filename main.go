package main

import (
	"net/http"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
)

var ids []int

func main()  {
	log.SetLevel(log.DebugLevel)
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	// Index route
  router.GET("/", func(ctx *gin.Context) {
		ids = []int{}
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
  })

	// HTMX button clicked route
	router.GET("/add", func(ctx *gin.Context) {
		id := "Task"+AssignTaskId()
		ctx.HTML(http.StatusOK, "component.tmpl", gin.H{
			"id": id,
			"task": "To-do: "+id,
		})
		log.Debug("Current ids slice", "ids", ids)
	})

	// HTMX Remove To-do route
	router.POST("/remove", func(ctx *gin.Context) {
		var IdParam struct	{
			ID string //`json:"taskId" validate:"required"`
		}
		// err := ctx.ShouldBindJSON(&IdParam)
		IdParam.ID = ctx.PostForm("taskId")
		id, err := strconv.Atoi(IdParam.ID[4:])

		if err != nil {
			log.Error("Failed to bind body params to ID.", "err", err)
		} else {
			log.Debug("The id was successfully bound", "ID", IdParam.ID)
		}

		ids = RemoveId(ids, id)
		ctx.HTML(http.StatusOK, "", gin.H{})
	})

  router.Run()
}

func AssignTaskId() (string) {
	id := 0
	// While loop
	for {
		found := false
		for _, i := range ids {
			if id == i {
				found = true
				break
			}
		}
		if !found {
			break
		}
		id++
	}
	ids = append(ids, id)
	return strconv.Itoa(id)
}

func SendConsoleLog(ctx *gin.Context, msg string) {
    ctx.Writer.Write([]byte(`<script>console.log("` + msg +`")</script>`))
}

func RemoveId(slice []int, id int) ([]int) {
	if len(slice) > id {
		slice[id] = slice[len(slice)-1]
		return slice[:len(slice)-1]
	} else {
		return slice
	}
}

// func ConnectToDB() {}
