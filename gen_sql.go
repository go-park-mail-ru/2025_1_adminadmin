package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/satori/uuid"
)

const (
	numUsers              = 10
	numRestaurants        = 100
	productsPerRestaurant = 6
	numOrders             = 200
)

type Product struct {
	ID         string
	Restaurant string
	Name       string
	Price      float64
	Weight     int
	ImageURL   string
	Category   string
}

type OrderProduct struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
	Weight   int     `json:"weight"`
	Amount   int     `json:"amount"`
}

var firstNames = []string{"Иван", "Мария", "Алексей", "Дарья", "Николай", "Татьяна", "Сергей", "Ярослав", "Иван", "Алексей", "Владислав", "Никита"}
var lastNames = []string{"Иванов", "Петрова", "Сидоров", "Кузнецова", "Попов", "Смирнова", "Пермякова", "Торетто", "Шипулина", "Ламар"}
var tags = []string{
	"Пицца", "Бургеры", "Суши", "Веган", "Кофе", "Десерты", "Шаурма", "Мексиканская кухня",
	"Итальянская кухня", "Грузинская кухня", "Китайская кухня", "Японская кухня", "Американская кухня",
	"Фастфуд", "Салаты", "Завтраки", "Стейки", "Морепродукты", "Пасты", "Смузи", "Фалафель", "Гриль", "Курица", "Рамен",
	"Корейская кухня", "Пекарня", "Пельмени", "Вьетнамская кухня", "Сибирская кухня", "ЗОЖ", "Кето", "Халяль", "Безглютеновое",
}
var categories = []string{
	"Напитки", "Горячее", "Супы", "Десерты", "Салаты", "Завтраки", "Гарниры", "Закуски", "Пицца", "Бургеры",
	"Роллы", "Сашими", "Сэндвичи", "Торты", "Пирожные", "Кофе", "Чай", "Смузи", "Соусы", "Супы дня",
	"Горячие блюда", "Веганские блюда", "Детское меню", "Паста", "Гриль", "Лапша", "Морсы", "Фреши", "Рыба", "Курица",
}

func randChoice(r *rand.Rand, list []string) string {
	return list[r.Intn(len(list))]
}

func randomPhone(r *rand.Rand) string {
	return fmt.Sprintf("+7-%03d-%03d-%04d", r.Intn(900)+100, r.Intn(900)+100, r.Intn(10000))
}

func randomPasswordHash(r *rand.Rand) string {
	b := make([]byte, 32)
	r.Read(b)
	return fmt.Sprintf("%x", b)
}

func generateSQL(r *rand.Rand) string {
	var sb strings.Builder

	// USERS
	var userIDs []string
	for i := 0; i < numUsers; i++ {
		id := uuid.NewV4().String()
		login := fmt.Sprintf("user%d", i)
		first := randChoice(r, firstNames)
		last := randChoice(r, lastNames)
		phone := randomPhone(r)
		pass := randomPasswordHash(r)
		userIDs = append(userIDs, id)

		fmt.Fprintf(&sb,
			"INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) "+
				"VALUES ('%s', '%s', '%s', '%s', '%s', '', 'default_user.jpg', decode('%s', 'hex'));\n",
			id, login, phone, first, last, pass)

	}

	// TAGS
	var tagIDs []string
	for _, tag := range tags {
		id := uuid.NewV4().String()
		tagIDs = append(tagIDs, id)
		fmt.Fprintf(&sb, "INSERT INTO restaurant_tags (id, name) VALUES ('%s', '%s');\n", id, tag)
	}

	// RESTAURANTS
	var restaurantIDs []string
	for i := 0; i < numRestaurants; i++ {
		id := uuid.NewV4().String()
		restaurantIDs = append(restaurantIDs, id)
		name := fmt.Sprintf("Ресторан %d", i+1)
		rating := fmt.Sprintf("%.1f", 3+r.Float64()*2)
		ratingCount := fmt.Sprintf("%.0f", 50+r.Float64()*100)

		fmt.Fprintf(&sb,
			"INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('%s', '%s', %s, %s);\n",
			id, name, rating, ratingCount)
	}

	// TAG RELATIONS
	for _, restID := range restaurantIDs {
		tagID := tagIDs[r.Intn(len(tagIDs))]
		fmt.Fprintf(&sb, "INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('%s', '%s');\n", restID, tagID)
	}

	// ADDRESSES
	var addressIDs []string
	userAddressMap := make(map[string]string)
	for _, uid := range userIDs {
		id := uuid.NewV4().String()
		addr := fmt.Sprintf("Улица %d, дом %d", r.Intn(100), r.Intn(50))
		userAddressMap[id] = addr
		addressIDs = append(addressIDs, id)
		fmt.Fprintf(&sb, "INSERT INTO addresses (id, address, user_id) VALUES ('%s', '%s', '%s');\n", id, addr, uid)
	}

	// PRODUCTS
	restaurantProducts := make(map[string][]Product)
	for _, restID := range restaurantIDs {
		for i := 0; i < productsPerRestaurant; i++ {
			id := uuid.NewV4().String()
			name := fmt.Sprintf("Блюдо %d", i+1)
			price := 100 + r.Float64()*500
			weight := 100 + r.Intn(400)
			category := randChoice(r, categories)
			imageURL := "default_product.jpg"

			p := Product{
				ID:         id,
				Restaurant: restID,
				Name:       name,
				Price:      price,
				Weight:     weight,
				ImageURL:   imageURL,
				Category:   category,
			}
			restaurantProducts[restID] = append(restaurantProducts[restID], p)

			fmt.Fprintf(&sb,
				"INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('%s', '%s', '%s', %.2f, %d, '%s');\n",
				id, restID, name, price, weight, category)
		}
	}

	// ORDERS
	for i := 0; i < 10; i++ {
		id := uuid.NewV4()
		userID := userIDs[r.Intn(len(userIDs))]
		addressID := addressIDs[r.Intn(len(addressIDs))]
		addressText := userAddressMap[addressID]
		status := "created"

		restID := restaurantIDs[r.Intn(len(restaurantIDs))]
		availableProducts := restaurantProducts[restID]

		used := map[string]bool{}
		numProducts := r.Intn(4) + 1
		var orderProducts []OrderProduct
		for len(orderProducts) < numProducts {
			p := availableProducts[r.Intn(len(availableProducts))]
			if used[p.ID] {
				continue
			}
			used[p.ID] = true
			amount := r.Intn(3) + 1
			orderProducts = append(orderProducts, OrderProduct{
				ID:       p.ID,
				Name:     p.Name,
				Price:    p.Price,
				ImageURL: p.ImageURL,
				Weight:   p.Weight,
				Amount:   amount,
			})
		}

		productsJSONBytes, _ := json.Marshal(orderProducts)
		escapedProducts := strings.ReplaceAll(string(productsJSONBytes), "'", "''")

		apartment := fmt.Sprintf("кв. %d", r.Intn(200)+1)
		intercom := fmt.Sprintf("%d%d%d", r.Intn(9)+1, r.Intn(9)+1, r.Intn(9)+1)
		entrance := fmt.Sprintf("%d", r.Intn(5)+1)
		floor := fmt.Sprintf("%d", r.Intn(25)+1)
		comment := "Не забудьте вилки"
		leave := r.Intn(2) == 0
		finalPrice := 0.0
		for _, op := range orderProducts {
			finalPrice += op.Price * float64(op.Amount)
		}


		fmt.Fprintf(&sb,
			"INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) "+
				"VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', %t, '%.2f');\n",
			id, userID, status, addressText, escapedProducts,
			apartment, intercom, entrance, floor, comment, leave, finalPrice)
	}

	return sb.String()
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sql := generateSQL(r)

	filePath := "build/sql/create_tables.sql"

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка при открытии %s: %v\n", filePath, err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("\n-- Generated data inserts\n" + sql)
	if err != nil {
		fmt.Printf("Ошибка при записи: %v\n", err)
		return
	}

	fmt.Printf("SQL данные успешно добавлены в %s\n", filePath)
}
