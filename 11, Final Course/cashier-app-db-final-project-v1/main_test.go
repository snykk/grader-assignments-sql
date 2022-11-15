package main_test

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cashier App with GORM", func() {
	var userRepo repo.UserRepository
	var sessionRepo repo.SessionsRepository
	var productRepo repo.ProductRepository
	var cartRepo repo.CartRepository

	db := db.NewDB()
	dbCredential := model.Credential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "12345678",
		DatabaseName: "my_db",
		Port:         5432,
		Schema:       "public",
	}
	conn, err := db.Connect(&dbCredential)
	Expect(err).ShouldNot(HaveOccurred())

	if err = conn.Migrator().DropTable("users", "sessions", "products", "carts"); err != nil {
		panic("failed droping table:" + err.Error())
	}

	userRepo = repo.NewUserRepository(conn)
	sessionRepo = repo.NewSessionsRepository(conn)
	productRepo = repo.NewProductRepository(conn)
	cartRepo = repo.NewCartRepository(conn)

	BeforeEach(func() {
		conn.AutoMigrate(&model.User{}, &model.Session{}, &model.Product{}, &model.Cart{})

		err := db.Reset(conn, "users")
		err = db.Reset(conn, "sessions")
		err = db.Reset(conn, "products")
		err = db.Reset(conn, "carts")
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
					conn.First(&model.User{}).First(&result)
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
					conn.First(&model.User{}).First(&result)
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
					conn.First(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					token, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(token.Token).To(Equal(session.Token))
					Expect(token.Username).To(Equal(session.Username))

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
					conn.First(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					token, err := sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())
					Expect(token.Token).To(Equal(session.Token))
					Expect(token.Username).To(Equal(session.Username))

					err = sessionRepo.DeleteSessions("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).ShouldNot(HaveOccurred())

					token, err = sessionRepo.SessionAvailToken("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).Should(HaveOccurred())
					Expect(token).To(Equal(model.Session{}))

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
					conn.First(&model.Session{}).First(&result)
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
					conn.First(&model.Session{}).First(&result)
					Expect(result.Token).To(Equal(session.Token))
					Expect(result.Username).To(Equal(session.Username))

					tokenFound, err := sessionRepo.TokenValidity("cc03dbea-4085-47ba-86fe-020f5d01a9d8")
					Expect(err).Should(HaveOccurred())
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
					conn.First(&model.Session{}).First(&result)
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
					conn.First(&model.Session{}).First(&result)
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

		})

		Describe("Products repository", func() {
			When("add product data to products table database postgres", func() {
				It("should save data product to products table database postgres", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}
					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Product{}
					conn.First(&model.Product{}).First(&result)
					Expect(result.Name).To(Equal(product.Name))
					Expect(result.Price).To(Equal(product.Price))
					Expect(result.Stock).To(Equal(product.Stock))
					Expect(result.Discount).To(Equal(product.Discount))
					Expect(result.Type).To(Equal(product.Type))

					err = db.Reset(conn, "products")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read all data product from products table database postgres", func() {
				It("should return a list data products", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}
					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Product{}
					conn.First(&model.Product{}).First(&result)
					Expect(result.Name).To(Equal(product.Name))
					Expect(result.Price).To(Equal(product.Price))
					Expect(result.Stock).To(Equal(product.Stock))
					Expect(result.Discount).To(Equal(product.Discount))
					Expect(result.Type).To(Equal(product.Type))

					res, err := productRepo.ReadProducts()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res[0].Name).To(Equal(product.Name))
					Expect(res[0].Price).To(Equal(product.Price))
					Expect(res[0].Stock).To(Equal(product.Stock))
					Expect(res[0].Discount).To(Equal(product.Discount))
					Expect(res[0].Type).To(Equal(product.Type))

					err = db.Reset(conn, "products")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("update product record from products table database postgres", func() {
				It("should change product record according to target ID", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}
					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Product{}
					conn.First(&model.Product{}).First(&result)
					Expect(result.Name).To(Equal(product.Name))
					Expect(result.Price).To(Equal(product.Price))
					Expect(result.Stock).To(Equal(product.Stock))
					Expect(result.Discount).To(Equal(product.Discount))
					Expect(result.Type).To(Equal(product.Type))

					productUpdate := model.Product{
						Name:     "Apple",
						Price:    10000,
						Stock:    10,
						Discount: 25,
						Type:     "Fruit",
					}

					err = productRepo.UpdateProduct(result.ID, productUpdate)
					Expect(err).ShouldNot(HaveOccurred())

					result = model.Product{}
					conn.First(&model.Product{}).First(&result)
					Expect(result.Name).To(Equal(productUpdate.Name))
					Expect(result.Price).To(Equal(productUpdate.Price))
					Expect(result.Stock).To(Equal(productUpdate.Stock))
					Expect(result.Discount).To(Equal(productUpdate.Discount))
					Expect(result.Type).To(Equal(productUpdate.Type))

					err = db.Reset(conn, "products")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("delete data product from products table database postgres", func() {
				It("should remove data product from products table database postgres according to target", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}
					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					result := model.Product{}
					conn.First(&model.Product{}).First(&result)
					Expect(result.Name).To(Equal(product.Name))
					Expect(result.Price).To(Equal(product.Price))
					Expect(result.Stock).To(Equal(product.Stock))
					Expect(result.Discount).To(Equal(product.Discount))
					Expect(result.Type).To(Equal(product.Type))

					err = productRepo.DeleteProduct(result.ID)
					Expect(err).ShouldNot(HaveOccurred())

					res, err := productRepo.ReadProducts()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(len(res)).To(Equal(0))

					err = db.Reset(conn, "products")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})

		Describe("Carts repository", func() {
			When("add product data to carts table database postgres", func() {
				It("Should save data product to carts table with id product, add quantity and add calculated price with discount. This will also reduce the amount of stock in the products table", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}

					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd := model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))
					Expect(resProd.Price).To(Equal(product.Price))
					Expect(resProd.Stock).To(Equal(product.Stock))
					Expect(resProd.Discount).To(Equal(product.Discount))
					Expect(resProd.Type).To(Equal(product.Type))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resCart := model.Cart{}
					conn.First(&model.Cart{}).First(&resCart)
					Expect(resCart.ProductID).To(Equal(uint(1)))
					Expect(resCart.Quantity).To(Equal(float64(1)))
					Expect(resCart.TotalPrice).To(Equal(float64(10000)))

					resProd = model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 1))

					err = db.Reset(conn, "products")
					err = db.Reset(conn, "carts")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("add product data to carts table database postgres with same product", func() {
				It("Should update data if found the same data in table carts database postgre based on product ID", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}

					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd := model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))
					Expect(resProd.Price).To(Equal(product.Price))
					Expect(resProd.Stock).To(Equal(product.Stock))
					Expect(resProd.Discount).To(Equal(product.Discount))
					Expect(resProd.Type).To(Equal(product.Type))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resProd = model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 1))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resCart := model.Cart{}
					conn.First(&model.Cart{}).First(&resCart)

					Expect(resCart.ProductID).To(Equal(uint(1)))
					Expect(resCart.Quantity).To(Equal(float64(2)))
					Expect(resCart.TotalPrice).To(Equal(float64(20000)))

					resProd = model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 2))

					err = db.Reset(conn, "products")
					err = db.Reset(conn, "carts")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("add product data to carts table database postgres with difference product", func() {
				It("Should create new data cart if not found the same data in table carts database postgre based on product ID", func() {
					product := model.Product{Name: "Watermelon", Price: 20000, Stock: 5, Discount: 50, Type: "Fruit"}

					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd := model.Product{}
					conn.Where("id = ?", uint(1)).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resProd = model.Product{}
					conn.Where("id = ?", uint(1)).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 1))

					resCart := model.Cart{}
					conn.Where("id = ?", uint(1)).First(&resCart)
					Expect(resCart.ProductID).To(Equal(uint(1)))
					Expect(resCart.Quantity).To(Equal(float64(1)))
					Expect(resCart.TotalPrice).To(Equal(float64(10000)))

					// New Product

					product = model.Product{Name: "Coffe", Price: 30000, Stock: 50, Discount: 15, Type: "Drink"}

					err = productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd = model.Product{}
					conn.Where("id = ?", uint(2)).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resProd = model.Product{}
					conn.Where("id = ?", uint(2)).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 1))

					resCart = model.Cart{}
					conn.Where("id = ?", uint(2)).First(&resCart)
					Expect(resCart.ProductID).To(Equal(uint(2)))
					Expect(resCart.Quantity).To(Equal(float64(1)))
					Expect(resCart.TotalPrice).To(Equal(float64(25500)))

					err = db.Reset(conn, "products")
					err = db.Reset(conn, "carts")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("read cart data and product with join method", func() {
				It("should return a list data join with field carts.id, carts.product_id, products.name, carts.quantity, carts.total_price", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}

					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd := model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))
					Expect(resProd.Price).To(Equal(product.Price))
					Expect(resProd.Stock).To(Equal(product.Stock))
					Expect(resProd.Discount).To(Equal(product.Discount))
					Expect(resProd.Type).To(Equal(product.Type))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resCart := model.Cart{}
					conn.First(&model.Cart{}).First(&resCart)
					Expect(resCart.ProductID).To(Equal(uint(1)))
					Expect(resCart.Quantity).To(Equal(float64(1)))
					Expect(resCart.TotalPrice).To(Equal(float64(10000)))

					resProd = model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 1))

					res, err := cartRepo.ReadCart()
					Expect(err).ShouldNot(HaveOccurred())
					Expect(res[0].Id).To(Equal(uint(1)))
					Expect(res[0].Name).To(Equal("Watermelon"))
					Expect(res[0].ProductId).To(Equal(uint(1)))
					Expect(res[0].Quantity).To(Equal(float64(1)))
					Expect(res[0].TotalPrice).To(Equal(float64(10000)))

					err = db.Reset(conn, "products")
					err = db.Reset(conn, "carts")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("update cart record from carts table database postgres", func() {
				It("should change cart record according to target ID", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}

					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd := model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))
					Expect(resProd.Price).To(Equal(product.Price))
					Expect(resProd.Stock).To(Equal(product.Stock))
					Expect(resProd.Discount).To(Equal(product.Discount))
					Expect(resProd.Type).To(Equal(product.Type))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resCart := model.Cart{}
					conn.First(&model.Cart{}).First(&resCart)
					Expect(resCart.ProductID).To(Equal(uint(1)))
					Expect(resCart.Quantity).To(Equal(float64(1)))
					Expect(resCart.TotalPrice).To(Equal(float64(10000)))

					resCart.Quantity = float64(4)

					err = cartRepo.UpdateCart(resCart.ID, resCart)
					Expect(err).ShouldNot(HaveOccurred())

					resCart = model.Cart{}
					conn.First(&model.Cart{}).First(&resCart)
					Expect(resCart.Quantity).To(Equal(float64(4)))

					err = db.Reset(conn, "products")
					err = db.Reset(conn, "carts")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})

			When("delete data cart from carts table database postgres", func() {
				It("should remove data cart from carts table database postgres according to target and return the product stock data according to the quantity cart", func() {
					product := model.Product{
						Name:     "Watermelon",
						Price:    20000,
						Stock:    5,
						Discount: 50,
						Type:     "Fruit",
					}

					err := productRepo.AddProduct(product)
					Expect(err).ShouldNot(HaveOccurred())

					resProd := model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Name).To(Equal(product.Name))

					err = cartRepo.AddCart(resProd)
					Expect(err).ShouldNot(HaveOccurred())

					resCart := model.Cart{}
					conn.First(&model.Cart{}).First(&resCart)
					Expect(resCart.ProductID).To(Equal(uint(1)))
					Expect(resCart.Quantity).To(Equal(float64(1)))
					Expect(resCart.TotalPrice).To(Equal(float64(10000)))

					resProd = model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock - 1))

					err = cartRepo.DeleteCart(uint(1), uint(1))
					Expect(err).ShouldNot(HaveOccurred())

					resProd = model.Product{}
					conn.First(&model.Product{}).First(&resProd)
					Expect(resProd.Stock).To(Equal(product.Stock))

					resCart = model.Cart{}
					conn.Table("carts").Select("*").Scan(&resCart)
					Expect(resCart.DeletedAt.Valid).To(Equal(true))

					err = db.Reset(conn, "products")
					err = db.Reset(conn, "carts")
					Expect(err).ShouldNot(HaveOccurred())
				})
			})
		})
	})
})
