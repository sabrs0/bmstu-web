package gorilla

import (
	"database/sql"
	"log/slog"

	"github.com/gorilla/mux"
	"github.com/sabrs0/bmstu-web/src/internal/controllers"
	"github.com/sabrs0/bmstu-web/src/internal/dataAccess/repositories/postgres"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/foundations"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/foundrisings"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/transactions"
	"github.com/sabrs0/bmstu-web/src/internal/http-server/handlers/users"
)

func SetRoutes(router *mux.Router, conn *sql.DB, logger *slog.Logger) { //мб сюда передать контроллеры ?
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

	const baseApiRoute = "/api/v1/"
	//FOUNDATION
	router.HandleFunc(baseApiRoute+"foundations", foundations.GetAll(logger, foundationController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"foundations", foundations.Add(logger, foundationController)).Methods("POST")
	router.HandleFunc(baseApiRoute+"foundations/{id}", foundations.GetByID(logger, foundationController)).Methods("PUT")
	router.HandleFunc(baseApiRoute+"foundations/{id}", foundations.Delete(logger, foundationController)).Methods("DELETE")
	router.HandleFunc(baseApiRoute+"foundations/{id}", foundations.GetByID(logger, foundationController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"foundations/{id}/donate", foundations.Donate(logger, foundationController)).Methods("POST")
	router.HandleFunc(baseApiRoute+"foundations/{id}/foundrisings", foundations.GetFoundrisings(logger, foundationController)).Methods("GET")

	//FOUNDRISING
	router.HandleFunc(baseApiRoute+"foundrisings", foundrisings.GetAll(logger, foundrisingController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"foundrisings", foundrisings.Add(logger, foundrisingController)).Methods("POST")
	router.HandleFunc(baseApiRoute+"foundrisings/{id}", foundrisings.Update(logger, foundrisingController)).Methods("PUT")
	router.HandleFunc(baseApiRoute+"foundrisings/{id}", foundrisings.UpdateOptional(logger, foundrisingController)).Methods("PATCH")
	router.HandleFunc(baseApiRoute+"foundrisings/{id}", foundrisings.Delete(logger, foundrisingController)).Methods("DELETE")
	router.HandleFunc(baseApiRoute+"foundrisings/{id}", foundrisings.GetByID(logger, foundrisingController)).Methods("GET")

	//TRANSACTION
	router.HandleFunc(baseApiRoute+"transactions", transactions.GetAll(logger, transactionController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"transactions/{id}", transactions.GetByID(logger, transactionController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"transactions/{id}", transactions.Delete(logger, transactionController)).Methods("DELETE")

	//USERS
	router.HandleFunc(baseApiRoute+"users", users.GetAll(logger, userController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"users", users.Add(logger, userController)).Methods("POST")
	router.HandleFunc(baseApiRoute+"users/{id}", users.Update(logger, userController)).Methods("PUT")
	router.HandleFunc(baseApiRoute+"users/{id}", users.Delete(logger, userController)).Methods("DELETE")
	router.HandleFunc(baseApiRoute+"users/{id}", users.GetByID(logger, userController)).Methods("GET")
	router.HandleFunc(baseApiRoute+"users/{id}/donate", users.Donate(logger, userController)).Methods("POST")

}
