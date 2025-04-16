package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	numUsers              = 10
	numRestaurants        = 100
	productsPerRestaurant = 10
	numOrders             = 200
)

var firstNames = []string{"Иван", "Мария", "Алексей", "Дарья", "Николай", "Татьяна", "Сергей", "Ярослав", "Иван", "Алексей", "Владислав", "Никита"}
var lastNames = []string{"Иванов", "Петрова", "Сидоров", "Кузнецова", "Попов", "Смирнова", "Пермякова", "Торетто", "Шипулина", "Ламар"}
var tags = []string{
	"Пицца", "Бургеры", "Суши", "Веган", "Кофе", "Десерты", "Шаурма",
	"Грузинская кухня", "Салаты", "Завтраки", "Стейки", "Морепродукты",
	"Пасты", "Смузи", "Фалафель", "Гриль", "Курица", "Рамен",
	"Корейская кухня", "Пекарня", "Пельмени", "Вьетнамская кухня",
	"Сибирская кухня", "ЗОЖ", "Кето", "Халяль", "Безглютеновое",
}
var categories = []string{
	"Горячее", "Супы",
}

var restaurants = []string{
	"Паста и Вино",
	"Суши Дрим",
	"Бургерная Ривьера",
	"Турецкий базар",
	"Зеленая вилка",
	"Гриль Бар",
	"Американская кухня",
	"Ресторан Средиземноморья",
	"Индийские специи",
	"Веганское счастье",
	"Французский уголок",
	"Мексиканская пекарня",
	"Китайская империя",
	"Баварский пивной сад",
	"Морская звезда",
	"Шашлыки от Бабая",
	"Скоро будет",
	"Восточный базар",
	"Греческий дворик",
	"Тосканский огонь",
	"Итальянская ривьера",
	"Суши Мания",
	"Пельмени на углях",
	"Бургеры по-американски",
	"Китайская звезда",
	"Мексиканская закуска",
	"Французский бистро",
	"Греческий остров",
	"Турецкая радость",
	"Индийская сказка",
	"Американская пекарня",
	"Восточный салат",
	"Вегетарианский рай",
	"Ресторан на воде",
	"Баварская пивоварня",
	"Морская лагуна",
	"Тосканские вечера",
	"Суши и роллы",
	"Вкус Индии",
	"Мексиканская площадь",
	"Греческая таверна",
	"Пивной бар Баварии",
	"Итальянский дворик",
	"Ресторан Печка",
	"Золотая рыба",
	"Красное море",
	"Ресторан Томат",
	"Турецкая кухня",
	"Вегетарианская кухня",
	"Ресторан Адель",
	"Гриль и мясо",
	"Том Ям",
	"Пельмени по-русски",
	"Китайская кухня",
	"Французская кухня",
	"Средиземноморский ресторан",
	"Ресторан Вкуса",
	"Шашлык-Бар",
	"Паста на ужин",
	"Веганский уголок",
	"Бургерная Сити",
	"Ресторан Эдем",
	"Ресторан Лаванда",
	"Ресторан Капрезе",
	"Греческий зал",
	"Пицца и Суши",
	"Турецкий Султан",
	"Мексиканский уголок",
	"Ресторан Мозаика",
	"Шашлыки по-кавказски",
	"Французская кухня на ужин",
	"Мексиканская кухня для всех",
	"Томаты и Паста",
}

func generateSQL() string {
	var sb strings.Builder

	// Генерация строк для каждого ресторана
	for _, restaurantName := range restaurants {
		// Для каждого ресторана генерируем нужные строки
		fmt.Fprintf(&sb,
`INSERT INTO products (restaurant_id, name, price, image_url, weight, category) VALUES
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Рамен с курицей', 740, 'default_product.jpg', 350, 'Закуски'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Рамен с говядиной', 650, 'default_product.jpg', 400, 'Закуски'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Рамен с ананасом', 640, 'default_product.jpg', 350, 'Закуски'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Пельмени с сыром', 550, 'default_product.jpg', 400, 'Закуски'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Суши ассорти', 490, 'default_product.jpg', 250, 'Суши'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Тамаго суши', 400, 'default_product.jpg', 150, 'Суши'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Сырная тарелка', 790, 'default_product.jpg', 250, 'Суши'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Вареники с грибами', 800, 'default_product.jpg', 150, 'Суши'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Ролл с лососем', 300, 'default_product.jpg', 200, 'Суши'),
	((SELECT id FROM restaurants WHERE name = '%s' ), 'Гёдза', 200, 'default_product.jpg', 180, 'Закуски');`,
			restaurantName, restaurantName, restaurantName, restaurantName, restaurantName, restaurantName, restaurantName, restaurantName, restaurantName, restaurantName)
			
			sb.WriteString("\n")
	}

	return sb.String()
}

func main() {
	sql := generateSQL()

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
