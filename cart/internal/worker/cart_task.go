package worker

import (
	"context"
	"time"

	"github.com/Go-Marketplace/backend/cart/internal/infrastructure/interfaces"
	"github.com/Go-Marketplace/backend/cart/internal/model"
	"github.com/Go-Marketplace/backend/cart/internal/usecase"
	"github.com/Go-Marketplace/backend/pkg/logger"
	"gopkg.in/tomb.v2"
)

const (
	cartTTL                = 5 * time.Minute
	cartTaskWorkerInterval = time.Second
)

type cartTaskWorker struct {
	tomb         tomb.Tomb
	cartTaskRepo interfaces.CartTaskRepo
	cartUsecase  *usecase.CartUsecase
	logger       *logger.Logger
}

func NewCartTaskWorker(cartTaskRepo interfaces.CartTaskRepo, cartUsecase *usecase.CartUsecase, logger *logger.Logger) *cartTaskWorker {
	return &cartTaskWorker{
		tomb:         tomb.Tomb{},
		cartTaskRepo: cartTaskRepo,
		cartUsecase:  cartUsecase,
		logger:       logger,
	}
}

func (worker *cartTaskWorker) Run(ctx context.Context) {
	worker.logger.Info("Start cart task worker")

	worker.tomb.Go(func() error {
		ticker := time.NewTicker(cartTaskWorkerInterval)
		defer ticker.Stop()

		for {
			select {
			case <-worker.tomb.Dying():
				worker.logger.Info("Stop running cart task worker")
				return nil

			case <-ticker.C:
				tasks, err := worker.cartTaskRepo.GetCartTasks(ctx, time.Now().Unix())
				if err != nil {
					worker.logger.Error("Cannot get cart tasks: %w", err)
				}

				if len(tasks) != 0 {
					worker.logger.Info("Got %d tasks: %v", len(tasks), tasks)
				}

				for _, task := range tasks {
					if task != nil {
						if _, err = worker.cartUsecase.DeleteCartCartlines(ctx, task.UserID); err != nil {
							worker.logger.Error("failed to delete cart %v cartlines: %w", task.UserID, err)
							continue
						}

						cartTask := model.CartTask{
							UserID:    task.UserID,
							Timestamp: time.Now().Add(cartTTL).Unix(),
						}

						if err = worker.cartTaskRepo.CreateCartTask(ctx, cartTask); err != nil {
							worker.logger.Error("failed to create cart %v task: %w", task.UserID, err)
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
