package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/gocarina/gocsv"
	"github.com/samber/lo"
	"github.com/satori/uuid"
)

type Translation struct {
	Translations []struct {
		Text string
	}
}

type Restaurant struct {
	Id          uuid.UUID
	Name        string `csv:"name"`
	Banner      string
	Rating      float64
	Category    string `csv:"category"`
	Adress      string `csv:"full_address"`
	Description string
	RatingCount int
}

type Menu struct {
	Id       uuid.UUID
	Name     string `csv:"name"`
	Price    string `csv:"price"`
	Category string `csv:"category"`
	ImageURL string
	Weight   int
}

const (
	numUsers              = 10
	numRestaurants        = 10
	productsPerRestaurant = 20
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
	// ... список ресторанов без изменений ...
}

func translate(str string) string {
	client := &http.Client{}

	var data = strings.NewReader(fmt.Sprintf("{\"folderId\": \"b1gk7ijg4gjud5f86vmq\",\"texts\": [\"%s\"],\"targetLanguageCode\": \"ru\"}", str))

	req, err := http.NewRequest("POST", "https://translate.api.cloud.yandex.net/translate/v2/translate", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var t Translation
	json.NewDecoder(resp.Body).Decode(&t)
	return t.Translations[0].Text
}

func escapeSQL(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func generateSQL() string {
	var sb strings.Builder

	restFiles, err := os.OpenFile("./data/restaurants.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer restFiles.Close()

	var rest []Restaurant

	if err := gocsv.UnmarshalFile(restFiles, &rest); err != nil {
		log.Fatal(err)
	}
	log.Println(len(rest))

	good := lo.Filter(rest, func(r Restaurant, _ int) bool {
		for _, r := range r.Name {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		return true
	})
	log.Println(len(good))

	fmt.Fprintf(&sb, `INSERT INTO restaurants (id, name, banner_url, address, rating, rating_count, description, working_mode_from, working_mode_to, delivery_time_from, delivery_time_to)
VALUES`)
	sb.WriteString("\n")

	for i := 0; i < numRestaurants-1; i++ {
		good[i].Id = uuid.NewV4()
		good[i].Adress = html.UnescapeString(translate(good[i].Adress))
		good[i].Banner = "default_restaurant.jpg"
		good[i].Rating = math.Round((4+rand.Float64())*10) / 10
		good[i].Category = html.UnescapeString(translate(good[i].Category))
		good[i].RatingCount = rand.Intn(400) + 200
		good[i].Description = fmt.Sprintf("%s, ресторан с вкусной едой и качественным обслуживанием", good[i].Name)

		fmt.Fprintf(&sb, `('%s', '%s', '%s', '%s', %f, %d, '%s', 10, 22, 50, 60),`,
			good[i].Id, escapeSQL(good[i].Name), escapeSQL(good[i].Banner), escapeSQL(good[i].Adress), good[i].Rating, good[i].RatingCount, escapeSQL(good[i].Description))
		sb.WriteString("\n")
	}

	good[numRestaurants-1].Id = uuid.NewV4()
	good[numRestaurants-1].Adress = html.UnescapeString(translate(good[numRestaurants-1].Adress))
	good[numRestaurants-1].Banner = "default_restaurant.jpg"
	good[numRestaurants-1].Rating = math.Round((4+rand.Float64())*10) / 10
	good[numRestaurants-1].Category = html.UnescapeString(translate(good[numRestaurants-1].Category))
	good[numRestaurants-1].RatingCount = rand.Intn(400) + 200
	good[numRestaurants-1].Description = fmt.Sprintf("%s, ресторан с вкусной едой и качественным обслуживанием", good[numRestaurants-1].Name)

	fmt.Fprintf(&sb, `('%s', '%s', '%s', '%s', %f, %d, '%s', 10, 22, 50, 60);`,
		good[numRestaurants-1].Id, escapeSQL(good[numRestaurants-1].Name), escapeSQL(good[numRestaurants-1].Banner), escapeSQL(good[numRestaurants-1].Adress), good[numRestaurants-1].Rating, good[numRestaurants-1].RatingCount, escapeSQL(good[numRestaurants-1].Description))

	sb.WriteString("\n")

	menuFiles, err := os.OpenFile("./data/restaurant-menus_1.csv", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer menuFiles.Close()

	var menu []Menu

	if err := gocsv.UnmarshalFile(menuFiles, &menu); err != nil {
		log.Fatal(err)
	}
	log.Println(len(menu))

	for i := 0; i < numRestaurants; i++ {

		fmt.Fprintf(&sb,
			`INSERT INTO products (restaurant_id, name, price, image_url, weight, category)
VALUES`)
		sb.WriteString("\n")

		base := productsPerRestaurant * i
		for j := 0; j < productsPerRestaurant-1; j++ {
			menu[base+j].Id = uuid.NewV4()
			menu[base+j].Name = html.UnescapeString(translate(menu[base+j].Name))
			menu[base+j].Category = html.UnescapeString(translate(menu[base+j].Category))
			menu[base+j].Price = strings.Replace(menu[base+j].Price, "USD", "", -1)
			priceFloat, err := strconv.ParseFloat(strings.TrimSpace(menu[base+j].Price), 64)
			if err != nil {
				log.Printf("Ошибка преобразования цены '%s': %v", menu[base+j].Price, err)
				priceFloat = 0
			}
			menu[base+j].ImageURL = "default_product.jpg"
			menu[base+j].Weight = rand.Intn(400) + 100

			fmt.Fprintf(&sb,
				`((SELECT id FROM restaurants WHERE name = '%s' ), '%s', %f, '%s', %d, '%s'),`,
				escapeSQL(good[i].Name), escapeSQL(menu[base+j].Name), priceFloat, menu[base+j].ImageURL, menu[base+j].Weight, escapeSQL(menu[base+j].Category))
			sb.WriteString("\n")
		}

		menu[productsPerRestaurant-1].Id = uuid.NewV4()
		menu[productsPerRestaurant-1].Name = html.UnescapeString(translate(menu[productsPerRestaurant-1].Name))
		menu[productsPerRestaurant-1].Category = html.UnescapeString(translate(menu[productsPerRestaurant-1].Category))
		menu[productsPerRestaurant-1].Price = strings.Replace(menu[productsPerRestaurant-1].Price, "USD", "", -1)
		priceFloat, err := strconv.ParseFloat(strings.TrimSpace(menu[productsPerRestaurant-1].Price), 64)
		if err != nil {
			log.Printf("Ошибка преобразования цены '%s': %v", menu[productsPerRestaurant-1].Price, err)
			priceFloat = 0
		}
		menu[productsPerRestaurant-1].ImageURL = "default_product.jpg"
		menu[productsPerRestaurant-1].Weight = rand.Intn(400) + 100

		fmt.Fprintf(&sb,
			`((SELECT id FROM restaurants WHERE name = '%s' ), '%s', %f, '%s', %d, '%s');`,
			escapeSQL(good[i].Name), escapeSQL(menu[productsPerRestaurant-1].Name), priceFloat, menu[productsPerRestaurant-1].ImageURL, menu[productsPerRestaurant-1].Weight, escapeSQL(menu[productsPerRestaurant-1].Category))
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

	_, err = f.WriteString("\n-- Parsed data inserts\n" + sql)
	if err != nil {
		fmt.Printf("Ошибка при записи: %v\n", err)
		return
	}

	fmt.Printf("SQL данные успешно добавлены в %s\n", filePath)
}
