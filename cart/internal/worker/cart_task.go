package worker

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/cart/internal/api/grpc/controller"
	"github.com/Go-Marketplace/backend/cart/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"github.com/Go-Marketplace/backend/proto/gen/cart"
	pbProduct "github.com/Go-Marketplace/backend/proto/gen/product"
	"gopkg.in/tomb.v2"
)

type CartTaskWorkerConfig struct {
	CartTTL                time.Duration
	CartTaskWorkerInterval time.Duration
}

type cartTaskWorker struct {
	tomb          tomb.Tomb
	cartTaskRepo  interfaces.CartTaskRepo
	cartUsecase   *usecase.CartUsecase
	productClient pbProduct.ProductClient
	config        CartTaskWorkerConfig
	logger        *logger.Logger
}

func NewCartTaskWorker(
	cartTaskRepo interfaces.CartTaskRepo,
	cartUsecase *usecase.CartUsecase,
	productClient pbProduct.ProductClient,
	config CartTaskWorkerConfig,
	logger *logger.Logger,
) *cartTaskWorker {
	return &cartTaskWorker{
		tomb:          tomb.Tomb{},
		cartTaskRepo:  cartTaskRepo,
		cartUsecase:   cartUsecase,
		productClient: productClient,
		config:        config,
		logger:        logger,
	}
}

func (worker *cartTaskWorker) Run(ctx context.Context) {
	worker.logger.Info("Start cart task worker with %v interval, cart ttl: %v", worker.config.CartTaskWorkerInterval, worker.config.CartTTL)

	worker.tomb.Go(func() error {
		ticker := time.NewTicker(worker.config.CartTaskWorkerInterval)
		defer ticker.Stop()

		for {
			select {
			case <-worker.tomb.Dying():
				worker.logger.Info("Stop running cart task worker")
				return nil

			case <-ticker.C:
				tasks, err := worker.cartTaskRepo.GetCartTasks(ctx, time.Now().Unix())
				if err != nil {
					worker.logger.Error("Cannot get cart tasks: %s", err.Error())
				}

				if len(tasks) != 0 {
					worker.logger.Info("Got %d tasks: %v", len(tasks), tasks)
				}

				for _, task := range tasks {
					if task != nil {
						if err = controller.DeleteCartCartlines(
							ctx,
							worker.cartUsecase,
							worker.productClient,
							&cart.DeleteCartCartlinesRequest{
								UserId: task.UserID.String(),
							},
						); err != nil {
							worker.logger.Error("failed to delete cart %v cartlines: %s", task.UserID, err.Error())
							continue
						}

						cartTask := model.CartTask{
							UserID:    task.UserID,
							Timestamp: time.Now().Add(worker.config.CartTTL).Unix(),
						}

						if err = worker.cartTaskRepo.CreateCartTask(ctx, cartTask); err != nil {
							worker.logger.Error("failed to create cart %v task: %s", task.UserID, err.Error())
							continue
						}
					} else {
						worker.logger.Warn("Got nil task")
					}
				}
			}
		}
	})
}

func (worker *cartTaskWorker) Stop() error {
	worker.logger.Info("Stop cart task worker")
	worker.tomb.Kill(nil)
	return worker.tomb.Wait()
}
