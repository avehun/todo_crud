package server

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/avehun/todo_crud/internal/model"
	"github.com/avehun/todo_crud/internal/repo"
	"github.com/gofiber/fiber/v2"
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

func (server *Server) ListTasks(c *fiber.Ctx) error {
	tasks, err := server.repo.GetTasks()
	if err != nil {
		log.Printf("Error fetching tasks from db: %v", err)
		return fiber.ErrInternalServerError
	}
	res, err := json.Marshal(tasks)
	c.Send(res)
}
func (server *Server) AddTask(c *fiber.Ctx) error {
	c.SendStatus(200)
	task := model.Task{}
	err := c.BodyParser(&task)
	if err != nil {
		log.Printf("Error parsing json: %v", err)
		return fiber.ErrBadRequest
	}
	err = server.repo.AddTask(task)
	if err != nil {
		log.Printf("Error inserting task to db: %v", err)
		return fiber.ErrBadRequest
	}
}
func (server *Server) UpdateTask(c *fiber.Ctx) error {
	c.SendStatus(200)
	task := model.Task{}

	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Printf("Wrong id: %v", err)
		return fiber.ErrBadRequest
	}
	err = c.BodyParser(&task)
	if err != nil {
		log.Printf("Error parsing json: %v", err)
		return fiber.ErrBadRequest
	}

	task.Id = id
	task.Updated_at = time.Now()
	err = server.repo.UpdateTask(task)
	if err != nil {
		log.Printf("Error parsing json: %v", err)
		return fiber.ErrNotFound
	}
}
func (server *Server) DeleteTask(c *fiber.Ctx) error {
	c.SendStatus(200)
	param := c.Params("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return fiber.ErrBadRequest
	}
	err = server.repo.DeleteTask(id)
	if err != nil {
		return fiber.ErrNotFound
	}
}
