package app

import (
	"fmt"
	"github.com/CracherX/catalog_hist/internal/controller/http/handlers"
	"github.com/CracherX/catalog_hist/internal/controller/http/router"
	"github.com/CracherX/catalog_hist/internal/usecase"
	"github.com/CracherX/catalog_hist/internal/usecase/repository"
	"github.com/CracherX/catalog_hist/pkg/config"
	"github.com/CracherX/catalog_hist/pkg/db"
	"github.com/CracherX/catalog_hist/pkg/logger"
	validation "github.com/CracherX/catalog_hist/pkg/validator"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

type App struct {
	Config    *config.Config
	Logger    handlers.Logger
	DB        *gorm.DB
	Validator handlers.Validator
	Router    *mux.Router
}

func New() (app *App, err error) {
	app = &App{}

	app.Config = config.MustLoad()
	app.Logger = logger.MustInitZap(app.Config.Server.Debug)
	app.DB, err = db.Connect(app.Config, app.Config.Database.Retries)
	if err != nil {
		return nil, err
	}
	app.Validator = validation.NewPlayground()
	app.Router = router.Setup()

	prepo := repository.NewProductRepoGorm(app.DB)
	puc := usecase.NewProductUseCase(prepo)
	ph := handlers.NewProductHandler(puc, app.Validator, app.Logger)

	crepo := repository.NewCategoryRepoGorm(app.DB)
	cuc := usecase.NewCategoryUseCase(crepo)
	ch := handlers.NewCategoryHandler(cuc, app.Validator, app.Logger)

	ctrepo := repository.NewCountryRepoGorm(app.DB)
	ctuc := usecase.NewCountryUseCase(ctrepo)
	cth := handlers.NewCountryHandler(ctuc, app.Validator, app.Logger)

	router.Product(app.Router, ph)
	router.Category(app.Router, ch)
	router.Country(app.Router, cth)
	return app, nil
}

// Run запуск приложения.
func (a *App) Run() {
	a.Logger.Info("Запуск приложения", zap.String("Приложение:", a.Config.Server.AppName))
	a.Logger.Debug("Запущен режим отладки для терминала!")
	err := http.ListenAndServe(a.Config.Server.Port, a.Router)
	if err != nil {
		fmt.Println(err)
		a.Logger.Error("Ошибка запуска сервера")
	}
}
