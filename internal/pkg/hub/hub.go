package hub

import (
	"context"
	"sync"
	"time"

	repos "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	"github.com/gorilla/websocket"
)

type Hub struct {
	connect       sync.Map
	currentOffset time.Time
	Repo          *repos.RestaurantRepository
}

func (h *Hub) AddClient(userID string, client *websocket.Conn) {
	h.connect.Store(client, userID)

	//постоянно чекаем, готов ли клиент что-то принимать или он отвалился
	go func() {
		for {
			_, _, err := client.NextReader()
			if err != nil {
				_ = client.Close()
				return
			}
		}
	}()

	//когда клиент ушел - удаляем его
	client.SetCloseHandler(func(code int, text string) error {
		h.connect.Delete(client)
		return nil
	})
}

// тут у нас раннер
func (h *Hub) Run(ctx context.Context) {
	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			h.connect.Range(func(key, value interface{}) bool {
				conn := key.(*websocket.Conn)
				userID := value.(string)

				// Получаем обновление
				order, ok := h.Repo.GetUpdates(ctx, userID, h.currentOffset)
				if !ok {
					return true // продолжаем работу, но не отправляем ничего
				}

				// Отправляем только если есть что отправить
				if err := conn.WriteJSON(order); err != nil {
					_ = conn.Close()
					return false // удаляем соединение
				}

				return true
			})

			// Обновляем offset после проверки
			h.currentOffset = time.Now()

		case <-ctx.Done():
			return
		}
	}
}
