package server

import (
	"log"
	"strconv"
	"time"

	"github.com/avehun/todo_crud/internal/model"
	"github.com/avehun/todo_crud/internal/repo"
	"github.com/gofiber/fiber"
)

type Server struct {
	repo *repo.Repo
}

func NewHandler(repo *repo.Repo) error {
	server := &Server{repo: repo}
	fb := fiber.New()

	tasks := fb.Group("/tasks")

	tasks.Get("", server.ListTasks)
	tasks.Post("", server.AddTask)
	tasks.Put("/:id", server.UpdateTask)
	tasks.Delete("/:id", server.DeleteTask)

	return fb.Listen(":8080")
}

func (server *Server) ListTasks(c *fiber.Ctx) {
	tasks, err := server.repo.GetTasks()
	if err != nil {
		c.SendStatus(404)
		log.Fatalf("Error fetching tasks from db: %v", err)
	}
	c.Write(tasks)
}
func (server *Server) AddTask(c *fiber.Ctx) {
	c.SendStatus(200)
	task := model.Task{}
	err := c.BodyParser(&task)
	if err != nil {
		c.SendStatus(404)
		log.Fatalf("Error parsing json: %v", err)
	}
	err = server.repo.AddTask(task)
	if err != nil {
		c.SendStatus(404)
		log.Fatalf("Error inserting task to db: %v", err)
	}
}
func (server *Server) UpdateTask(c *fiber.Ctx) {
	c.SendStatus(200)
	task := model.Task{}

	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.SendStatus(404)
		log.Printf("Wrong id: %v", err)
	}
	err = c.BodyParser(&task)
	if err != nil {
		c.SendStatus(404)
		log.Printf("Error parsing json: %v", err)
	}

	task.Id = id
	task.Updated_at = time.Now()
	err = server.repo.UpdateTask(task)
	if err != nil {
		c.SendStatus(404)
		log.Printf("Error parsing json: %v", err)
	}
}
func (server *Server) DeleteTask(c *fiber.Ctx) {
	c.SendStatus(200)
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.SendStatus(404)
	}
	err = server.repo.DeleteTask(id)
	if err != nil {
		c.SendStatus(404)
	}
}
