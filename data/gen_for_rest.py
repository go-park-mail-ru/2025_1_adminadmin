import pandas as pd
import random

# Функция для генерации случайного веса
def generate_random_weight():
    return random.randint(200, 500)

# Загружаем данные из Excel
file_path = "test2.xlsx"  # Путь к вашему файлу
df = pd.read_excel(file_path)

# Получаем список ресторанов
restaurants = df['rest'].unique().tolist()

# Оставляем только уникальные товары по полю 'item'
unique_products = df.drop_duplicates(subset=['item']).to_dict(orient='records')

# Генерируем SQL-запросы
sql_queries = []

for restaurant in restaurants:
    # Выбираем 40 случайных товаров (можно менять на sample, если нужны уникальные)
    selected_products = random.choices(unique_products, k=40)

    values = []
    for product in selected_products:
        item = product["item"]
        price = product["price"]
        cat = product["cat"]
        image_url = product["img"]

        if pd.isna(item):  # Пропускаем пустые названия
            continue

        weight = generate_random_weight()
        value = f"((SELECT id FROM restaurants WHERE name = '{restaurant}'), '{item}', {price}, '{image_url}', {weight}, '{cat}')"
        values.append(value)
    
    query = f"INSERT INTO products (restaurant_id, name, price, image_url, weight, category) VALUES\n\t{',\n\t'.join(values)};"
    sql_queries.append(query)

# Сохраняем результат в файл
with open("sql_inserts.txt", "w", encoding="utf-8") as f:
    for query in sql_queries:
        f.write(query + "\n\n")

print(f"✅ SQL-запросы успешно сгенерированы для {len(restaurants)} ресторанов.")
print(f"📦 Каждый ресторан получил по 40 товаров из исходного файла (с возможными дублями между ресторанами).")