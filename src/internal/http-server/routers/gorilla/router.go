package gorilla

import (
	"database/sql"
	"log/slog"

	"github.com/gorilla/mux"
	authServ "github.com/sabrs0/bmstu-web/src/internal/business/services/auth"
	"github.com/sabrs0/bmstu-web/src/internal/controllers"
	"github.com/sabrs0/bmstu-web/src/internal/dataAccess/repositories/postgres"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/auth"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/foundrisings"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/transactions"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/users"
	authMW "github.com/sabrs0/bmstu-web/src/internal/http-server/middleware/auth"
	loggerMW "github.com/sabrs0/bmstu-web/src/internal/http-server/middleware/logger"
)

func SetRouter(conn *sql.DB, logger *slog.Logger) *mux.Router { //мб сюда передать контроллеры ?
	foundationController := controllers.NewFoundationController(
		postgres.NewFoundationRepository(conn),
		postgres.NewTransactionRepository(conn),
		postgres.NewFoundrisingRepository(conn),
	)
	foundrisingController := controllers.NewFoundrisingController(
		postgres.NewFoundrisingRepository(conn),
		postgres.NewFoundationRepository(conn),
	)
	transactionController := controllers.NewTransactionController(
		postgres.NewFoundationRepository(conn),
		postgres.NewFoundrisingRepository(conn),
		postgres.NewUserRepository(conn),
		postgres.NewTransactionRepository(conn),
	)
	userController := controllers.NewUserController(
		postgres.NewFoundationRepository(conn),
		postgres.NewFoundrisingRepository(conn),
		postgres.NewUserRepository(conn),
		postgres.NewTransactionRepository(conn),
	)
	authController := controllers.NewLoginController(
		postgres.NewFoundationRepository(conn),
		postgres.NewUserRepository(conn),
		authServ.NewJWTService(),
	)
	const baseApiRoute = "/api/v1"
	router := mux.NewRouter()

	router.Use(loggerMW.New(logger))

	//NONE AUTH
	unAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()
	unAuthRouter.HandleFunc("/foundations", foundations.GetAll(logger, foundationController)).Methods("GET")
	unAuthRouter.HandleFunc("/foundrisings", foundrisings.GetAll(logger, foundrisingController)).Methods("GET")
	unAuthRouter.HandleFunc("/login", auth.Login(logger, authController)).Methods("POST")

	//FOUNDATION AUTH
	foundationAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()
	foundationAuthRouter.Use(authMW.New([]string{"Foundation"}, authServ.NewJWTService()))
	foundationAuthRouter.HandleFunc(baseApiRoute+
		"/foundations/{id}/donate", foundations.Donate(logger, foundationController)).Methods("POST")

	//FOUNDATION ADMIN AUTH
	foundationAdminAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()

	foundationAdminAuthRouter.Use(authMW.New([]string{"Admin", "Foundation"}, authServ.NewJWTService()))

	foundationAdminAuthRouter.HandleFunc("/foundations",
		foundations.Add(logger, foundationController)).Methods("POST")
	foundationAdminAuthRouter.HandleFunc("/foundations/{id}",
		foundations.GetByID(logger, foundationController)).Methods("PUT")
	foundationAdminAuthRouter.HandleFunc("/foundations/{id}",
		foundations.Delete(logger, foundationController)).Methods("DELETE")
	foundationAdminAuthRouter.HandleFunc("/foundations/{id}",
		foundations.GetByID(logger, foundationController)).Methods("GET")
	foundationAdminAuthRouter.HandleFunc("/foundations/{id}/foundrisings",
		foundations.GetFoundrisings(logger, foundationController)).Methods("GET")

	foundationAdminAuthRouter.HandleFunc(baseApiRoute+"/foundrisings",
		foundrisings.Add(logger, foundrisingController)).Methods("POST")
	foundationAdminAuthRouter.HandleFunc(baseApiRoute+"/foundrisings/{id}",
		foundrisings.Update(logger, foundrisingController)).Methods("PUT")
	foundationAdminAuthRouter.HandleFunc(baseApiRoute+"/foundrisings/{id}",
		foundrisings.UpdateOptional(logger, foundrisingController)).Methods("PATCH")
	foundationAdminAuthRouter.HandleFunc(baseApiRoute+"/foundrisings/{id}",
		foundrisings.Delete(logger, foundrisingController)).Methods("DELETE")

	//FOUNDATION ADMIN USER AUTH
	foundationUserAdminAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()
	foundationUserAdminAuthRouter.Use(authMW.New([]string{"Foundation", "Admin", "User"}, authServ.NewJWTService()))
	foundationUserAdminAuthRouter.HandleFunc(baseApiRoute+"/foundrisings/{id}",
		foundrisings.GetByID(logger, foundrisingController)).Methods("GET")

	//USER ADMIN AUTH
	userAdminAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()
	userAdminAuthRouter.Use(authMW.New([]string{"User", "Admin"}, authServ.NewJWTService()))

	userAdminAuthRouter.HandleFunc("/users", users.Add(logger, userController)).Methods("POST")
	userAdminAuthRouter.HandleFunc("/users/{id}", users.Update(logger, userController)).Methods("PUT")
	userAdminAuthRouter.HandleFunc("/users/{id}", users.Delete(logger, userController)).Methods("DELETE")
	userAdminAuthRouter.HandleFunc("/users/{id}", users.GetByID(logger, userController)).Methods("GET")
	//USER AUTH
	userAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()
	userAuthRouter.Use(authMW.New([]string{"User"}, authServ.NewJWTService()))

	userAuthRouter.HandleFunc(baseApiRoute+"/users/{id}/donate", users.Donate(logger, userController)).Methods("POST")

	//ADMIN AUTH
	adminAuthRouter := router.PathPrefix(baseApiRoute).Subrouter()
	adminAuthRouter.Use(authMW.New([]string{"Admin"}, authServ.NewJWTService()))
	adminAuthRouter.HandleFunc("/users", users.GetAll(logger, userController)).Methods("GET")

	adminAuthRouter.HandleFunc("/transactions", transactions.GetAll(logger, transactionController)).Methods("GET")
	adminAuthRouter.HandleFunc("/transactions/{id}", transactions.GetByID(logger, transactionController)).Methods("GET")
	adminAuthRouter.HandleFunc("/transactions/{id}", transactions.Delete(logger, transactionController)).Methods("DELETE")

	return router

}
