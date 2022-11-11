package main_test

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	repo "a21hc3NpZ25tZW50/repository"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Todo App with GORM", func() {
	var userRepo repo.UserRepository
	var sessionRepo repo.SessionsRepository
	var todosRepo repo.TodoRepository

	db := db.NewDB()
	dbCredential := model.CredentialDB{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := db.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	if err = conn.Migrator().DropTable("users", "todos", "sessions"); err != nil {
		panic("failed droping table:" + err.Error())
	}

	userRepo = repo.NewUserRepository(conn)
	sessionRepo = repo.NewSessionsRepository(conn)
	todosRepo = repo.NewTodoRepository(conn)

	BeforeEach(func() {
		err := conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Todo{})
		err = db.Reset(conn, "users")
		err = db.Reset(conn, "todos")
		err = db.Reset(conn, "sessions")
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Repository", func() {
		Describe("Users repository", func() {
			When("add user data to users table database postgres", func() {
				It("should save data user to users table database postgres", func() {
					user := model.User{
						Username: "aditira",
						Password: "!opensesame",
					}
					err := userRepo.AddUser(user)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.User{}
					conn.Model(&model.User{}).First(&result)
					Expect(result.Username).To(Equal(user.Username))
					Expect(result.Password).To(Equal(user.Password))

					err = db.Reset(conn, "users")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("check user availability in users table database postgres", func() {
				It("return error if present and nil if not present", func() {
					user := model.User{}

					err := userRepo.UserAvail(user)
					Expect(err).Should(HaveOccurred())

					user = model.User{
						Username: "aditira",
						Password: "!opensesame",
					}

					err = userRepo.AddUser(user)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.User{}
					conn.Model(&model.User{}).First(&result)
					Expect(result.Username).To(Equal(user.Username))
					Expect(result.Password).To(Equal(user.Password))

					err = userRepo.UserAvail(user)
					Expect(err).ShouldNot(HaveOccurred())

					err = db.Reset(conn, "users")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Sessions repository", func() {
			When("add session data to sessions table database postgres", func() {
				It("should save data session to sessions table database postgres", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("delete selected session to sessions table database postgres", func() {
				It("should delete data session target from sessions table database postgres", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					err = sessionRepo.DeleteSession("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())

					result = model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result).To(Equal(model.Session{}))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("update selected session to sessions table database postgres", func() {
				It("should update data session target the username field from sessions table database postgres", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					sessionUpdate := model.Session{
						Token:    "cc03dbac-4085-22ba-75fe-103f9a01b6d5",
						Username: "aditira",
						Expiry:   time.Date(2022, 11, 17, 20, 34, 58, 651387237, time.UTC),
					}
					err = sessionRepo.UpdateSessions(sessionUpdate)
					Expect(err).ShouldNot(HaveOccurred())

					result = model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(sessionUpdate.Token))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("check expired session with exprired session is 5 hours from now", func() {
				It("should return a session model with token, name, and expired time", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Now().Add(5 * time.Hour),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					tokenFound, err := sessionRepo.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(tokenFound.Token).To(Equal("cc03dbea-4085-47ba-86fe-020f5d01a9d8"))
					Expect(tokenFound.Username).To(Equal("aditira"))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("check expired session with exprired session is 5 hours ago", func() {
				It("should return error message token is expired and empty session model", func() {
					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Now().Add(-5 * time.Hour),
					}
					err := sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					tokenFound, err := sessionRepo.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).To(Equal(fmt.Errorf("Token is Expired!")))
					Expect(tokenFound).To(Equal(model.Session{}))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("check session availability with name", func() {
				It("return data session with target name", func() {
					_, err := sessionRepo.SessionAvailName("aditira")
					Expect(err).Should(HaveOccurred())

					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Now().Add(5 * time.Hour),
					}
					err = sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					res, err := sessionRepo.SessionAvailName("aditira")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res.Token).To(Equal(session.Token))
					Expect(res.Username).To(Equal(session.Username))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("check session availability with token", func() {
				It("return data session with target token", func() {
					_, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).Should(HaveOccurred())

					session := model.Session{
						Token:    "cc03dbea-4085-47ba-86fe-020f5d01a9d8",
						Username: "aditira",
						Expiry:   time.Now().Add(5 * time.Hour),
					}
					err = sessionRepo.AddSessions(session)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Session{}
					conn.Model(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					res, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res.Token).To(Equal(session.Token))
					Expect(res.Username).To(Equal(session.Username))

					err = db.Reset(conn, "sessions")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

		})

		Describe("Todos repository", func() {
			When("add todo data to todos table database postgres", func() {
				It("should save data todo to todos table database postgres", func() {
					todo := model.Todo{
						Task: "Create a todo app program",
						Done: false,
					}
					err := todosRepo.AddTodo(todo)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result.Task).To(Equal(todo.Task))
					Expect(result.Done).To(Equal(todo.Done))

					err = db.Reset(conn, "todos")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read all data todo from todos table database postgres", func() {
				It("should return a list data todo", func() {
					todo := model.Todo{
						Task: "Create a todo app program",
						Done: false,
					}
					err := todosRepo.AddTodo(todo)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result.Task).To(Equal(todo.Task))
					Expect(result.Done).To(Equal(todo.Done))

					err = db.Reset(conn, "todos")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("update status todo from todos table database postgres", func() {
				It("should change field done true or false", func() {
					todo := model.Todo{
						Task: "Create a todo app program",
						Done: false,
					}
					err := todosRepo.AddTodo(todo)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result.Task).To(Equal(todo.Task))
					Expect(result.Done).To(Equal(todo.Done))

					err = todosRepo.UpdateDone(1, true)
					Expect(err).ShouldNot(HaveOccurred())

					result = model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result.Done).To(Equal(true))

					err = todosRepo.UpdateDone(1, false)
					Expect(err).ShouldNot(HaveOccurred())

					result = model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result.Done).To(Equal(false))

					err = db.Reset(conn, "todos")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("delete data todo from todos table database postgres", func() {
				It("should remove data todo from todos table database postgres according to target", func() {
					todo := model.Todo{
						Task: "Create a todo app program",
						Done: false,
					}
					err := todosRepo.AddTodo(todo)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result.Task).To(Equal(todo.Task))
					Expect(result.Done).To(Equal(todo.Done))

					err = todosRepo.DeleteTodo(1)
					Expect(err).ShouldNot(HaveOccurred())

					result = model.Todo{}
					conn.Model(&model.Todo{}).First(&result)
					Expect(result).To(Equal(model.Todo{}))

					err = db.Reset(conn, "todos")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
