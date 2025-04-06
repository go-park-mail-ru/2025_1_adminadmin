CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,                     
    phone_number TEXT,                              
    first_name TEXT NOT NULL,  
    last_name TEXT NOT NULL,   
    description TEXT DEFAULT '',  
    user_pic TEXT DEFAULT 'default_user.jpg', 
    password_hash BYTEA NOT NULL                   
);

CREATE TABLE IF NOT EXISTS restaurant_tags (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  
    name TEXT NOT NULL UNIQUE                        
);

CREATE TABLE IF NOT EXISTS restaurants (
    id UUID PRIMARY KEY,  
    name TEXT NOT NULL, 
    banner_url TEXT DEFAULT 'default_restaurant.jpg',                           
    address TEXT DEFAULT '',
    rating FLOAT CHECK (rating >= 0 AND rating <= 5), 
    rating_count FLOAT CHECK (rating_count >= 0),
    description TEXT DEFAULT '',                              
    working_mode_from INT DEFAULT 8,  
    working_mode_to INT DEFAULT 23,   
    delivery_time_from INT DEFAULT 50,  
    delivery_time_to INT DEFAULT 60                                       
);

CREATE TABLE IF NOT EXISTS restaurant_tags_relations (
    restaurant_id UUID REFERENCES restaurants(id) ON DELETE CASCADE,
    tag_id UUID REFERENCES restaurant_tags(id) ON DELETE CASCADE,   
    PRIMARY KEY (restaurant_id, tag_id)
);

CREATE TABLE IF NOT EXISTS addresses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),  
    address TEXT NOT NULL,                                                           
    user_id UUID,                                   
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL                                  
);

CREATE TABLE IF NOT EXISTS products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    restaurant_id UUID NOT NULL REFERENCES restaurants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    image_url TEXT DEFAULT 'default_product.jpg',
    weight INT NOT NULL,
    category TEXT NOT NULL
);


CREATE TABLE IF NOT EXISTS orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status TEXT NOT NULL,
    address_id UUID NOT NULL REFERENCES addresses(id) ON DELETE CASCADE,
    order_products TEXT NOT NULL
);

INSERT INTO restaurant_tags (id, name)
VALUES 
  (gen_random_uuid(), 'Итальянский'),
  (gen_random_uuid(), 'Японский'),
  (gen_random_uuid(), 'Фастфуд'),
  (gen_random_uuid(), 'Турецкий'),
  (gen_random_uuid(), 'Вегетарианский'),
  (gen_random_uuid(), 'Американский'),
  (gen_random_uuid(), 'Европейский'),
  (gen_random_uuid(), 'Индийский'),
  (gen_random_uuid(), 'Средиземноморский'),
  (gen_random_uuid(), 'Веганский'),
  (gen_random_uuid(), 'Французский'),
  (gen_random_uuid(), 'Мексиканский'),
  (gen_random_uuid(), 'Китайский'),
  (gen_random_uuid(), 'Немецкий');

INSERT INTO restaurants (id, name, banner_url, address, rating, rating_count, description, working_mode_from, working_mode_to, delivery_time_from, delivery_time_to)
VALUES
(uuid_generate_v4(), 'Паста и Вино', 'default_restaurant.jpg', 'Паста и Вино на улице Мира', 3.4, 654, 'Ресторан Паста и Вино, уютная атмосфера и вкусная еда.', 9, 22, 56, 61),
(uuid_generate_v4(), 'Суши Дрим', 'default_restaurant.jpg', 'Суши Дрим на улице Суворова', 0.2, 550, 'Ресторан Суши Дрим, уютная атмосфера и вкусная еда.', 11, 21, 34, 78),
(uuid_generate_v4(), 'Бургерная Ривьера', 'default_restaurant.jpg', 'Бургерная Ривьера на улице Мира', 0.7, 556, 'Ресторан Бургерная Ривьера, уютная атмосфера и вкусная еда.', 8, 22, 55, 64),
(uuid_generate_v4(), 'Турецкий базар', 'default_restaurant.jpg', 'Турецкий базар на улице Карла Маркса', 0.2, 582, 'Ресторан Турецкий базар, уютная атмосфера и вкусная еда.', 11, 20, 34, 82),
(uuid_generate_v4(), 'Зеленая вилка', 'default_restaurant.jpg', 'Зеленая вилка на улице Карла Маркса', 3.1, 186, 'Ресторан Зеленая вилка, уютная атмосфера и вкусная еда.', 11, 20, 49, 73),
(uuid_generate_v4(), 'Гриль Бар', 'default_restaurant.jpg', 'Гриль Бар на улице Пушкина', 0.8, 445, 'Ресторан Гриль Бар, уютная атмосфера и вкусная еда.', 12, 22, 50, 87),
(uuid_generate_v4(), 'Американская кухня', 'default_restaurant.jpg', 'Американская кухня на улице Пушкина', 0.5, 911, 'Ресторан Американская кухня, уютная атмосфера и вкусная еда.', 9, 21, 59, 75),
(uuid_generate_v4(), 'Ресторан Средиземноморья', 'default_restaurant.jpg', 'Ресторан Средиземноморья на улице Суворова', 3.1, 342, 'Ресторан Ресторан Средиземноморья, уютная атмосфера и вкусная еда.', 9, 20, 39, 75),
(uuid_generate_v4(), 'Индийские специи', 'default_restaurant.jpg', 'Индийские специи на улице Ленина', 3.1, 626, 'Ресторан Индийские специи, уютная атмосфера и вкусная еда.', 11, 23, 52, 88),
(uuid_generate_v4(), 'Веганское счастье', 'default_restaurant.jpg', 'Веганское счастье на улице Пушкина', 2.4, 240, 'Ресторан Веганское счастье, уютная атмосфера и вкусная еда.', 12, 20, 48, 73),
(uuid_generate_v4(), 'Французский уголок', 'default_restaurant.jpg', 'Французский уголок на улице Суворова', 3.1, 325, 'Ресторан Французский уголок, уютная атмосфера и вкусная еда.', 8, 21, 45, 66),
(uuid_generate_v4(), 'Мексиканская пекарня', 'default_restaurant.jpg', 'Мексиканская пекарня на улице Пушкина', 4.9, 615, 'Ресторан Мексиканская пекарня, уютная атмосфера и вкусная еда.', 8, 21, 53, 64),
(uuid_generate_v4(), 'Китайская империя', 'default_restaurant.jpg', 'Китайская империя на улице Карла Маркса', 3.7, 356, 'Ресторан Китайская империя, уютная атмосфера и вкусная еда.', 12, 20, 52, 76),
(uuid_generate_v4(), 'Баварский пивной сад', 'default_restaurant.jpg', 'Баварский пивной сад на улице Пушкина', 3.9, 702, 'Ресторан Баварский пивной сад, уютная атмосфера и вкусная еда.', 11, 23, 37, 87),
(uuid_generate_v4(), 'Морская звезда', 'default_restaurant.jpg', 'Морская звезда на улице Карла Маркса', 3.2, 472, 'Ресторан Морская звезда, уютная атмосфера и вкусная еда.', 10, 21, 40, 80),
(uuid_generate_v4(), 'Шашлыки от Бабая', 'default_restaurant.jpg', 'Шашлыки от Бабая на улице Карла Маркса', 1.5, 774, 'Ресторан Шашлыки от Бабая, уютная атмосфера и вкусная еда.', 11, 21, 37, 74),
(uuid_generate_v4(), 'Скоро будет', 'default_restaurant.jpg', 'Скоро будет на улице Суворова', 0.1, 754, 'Ресторан Скоро будет, уютная атмосфера и вкусная еда.', 11, 20, 31, 81),
(uuid_generate_v4(), 'Восточный базар', 'default_restaurant.jpg', 'Восточный базар на улице Мира', 0.8, 408, 'Ресторан Восточный базар, уютная атмосфера и вкусная еда.', 9, 23, 36, 81),
(uuid_generate_v4(), 'Греческий дворик', 'default_restaurant.jpg', 'Греческий дворик на улице Мира', 4.6, 736, 'Ресторан Греческий дворик, уютная атмосфера и вкусная еда.', 8, 22, 31, 83),
(uuid_generate_v4(), 'Тосканский огонь', 'default_restaurant.jpg', 'Тосканский огонь на улице Суворова', 2.0, 175, 'Ресторан Тосканский огонь, уютная атмосфера и вкусная еда.', 11, 22, 45, 80),
(uuid_generate_v4(), 'Итальянская ривьера', 'default_restaurant.jpg', 'Итальянская ривьера на улице Ленина', 2.8, 926, 'Ресторан Итальянская ривьера, уютная атмосфера и вкусная еда.', 12, 21, 53, 74),
(uuid_generate_v4(), 'Суши Мания', 'default_restaurant.jpg', 'Суши Мания на улице Карла Маркса', 4.3, 8, 'Ресторан Суши Мания, уютная атмосфера и вкусная еда.', 8, 21, 52, 61),
(uuid_generate_v4(), 'Пельмени на углях', 'default_restaurant.jpg', 'Пельмени на углях на улице Пушкина', 2.7, 773, 'Ресторан Пельмени на углях, уютная атмосфера и вкусная еда.', 11, 22, 57, 70),
(uuid_generate_v4(), 'Бургеры по-американски', 'default_restaurant.jpg', 'Бургеры по-американски на улице Карла Маркса', 1.0, 816, 'Ресторан Бургеры по-американски, уютная атмосфера и вкусная еда.', 12, 21, 39, 85),
(uuid_generate_v4(), 'Китайская звезда', 'default_restaurant.jpg', 'Китайская звезда на улице Ленина', 1.2, 660, 'Ресторан Китайская звезда, уютная атмосфера и вкусная еда.', 12, 23, 53, 69),
(uuid_generate_v4(), 'Мексиканская закуска', 'default_restaurant.jpg', 'Мексиканская закуска на улице Мира', 1.7, 770, 'Ресторан Мексиканская закуска, уютная атмосфера и вкусная еда.', 8, 21, 42, 82),
(uuid_generate_v4(), 'Французский бистро', 'default_restaurant.jpg', 'Французский бистро на улице Пушкина', 0.7, 10, 'Ресторан Французский бистро, уютная атмосфера и вкусная еда.', 8, 22, 57, 87),
(uuid_generate_v4(), 'Греческий остров', 'default_restaurant.jpg', 'Греческий остров на улице Карла Маркса', 4.0, 310, 'Ресторан Греческий остров, уютная атмосфера и вкусная еда.', 10, 21, 60, 88),
(uuid_generate_v4(), 'Турецкая радость', 'default_restaurant.jpg', 'Турецкая радость на улице Суворова', 4.8, 83, 'Ресторан Турецкая радость, уютная атмосфера и вкусная еда.', 9, 20, 49, 71),
(uuid_generate_v4(), 'Индийская сказка', 'default_restaurant.jpg', 'Индийская сказка на улице Ленина', 4.2, 78, 'Ресторан Индийская сказка, уютная атмосфера и вкусная еда.', 8, 23, 48, 85),
(uuid_generate_v4(), 'Американская пекарня', 'default_restaurant.jpg', 'Американская пекарня на улице Карла Маркса', 3.1, 832, 'Ресторан Американская пекарня, уютная атмосфера и вкусная еда.', 8, 23, 35, 76),
(uuid_generate_v4(), 'Восточный салат', 'default_restaurant.jpg', 'Восточный салат на улице Мира', 1.6, 648, 'Ресторан Восточный салат, уютная атмосфера и вкусная еда.', 11, 23, 48, 86),
(uuid_generate_v4(), 'Вегетарианский рай', 'default_restaurant.jpg', 'Вегетарианский рай на улице Ленина', 0.8, 472, 'Ресторан Вегетарианский рай, уютная атмосфера и вкусная еда.', 12, 23, 35, 62),
(uuid_generate_v4(), 'Ресторан на воде', 'default_restaurant.jpg', 'Ресторан на воде на улице Суворова', 3.7, 229, 'Ресторан Ресторан на воде, уютная атмосфера и вкусная еда.', 11, 23, 52, 77),
(uuid_generate_v4(), 'Баварская пивоварня', 'default_restaurant.jpg', 'Баварская пивоварня на улице Ленина', 3.9, 239, 'Ресторан Баварская пивоварня, уютная атмосфера и вкусная еда.', 12, 20, 44, 62),
(uuid_generate_v4(), 'Морская лагуна', 'default_restaurant.jpg', 'Морская лагуна на улице Суворова', 0.7, 616, 'Ресторан Морская лагуна, уютная атмосфера и вкусная еда.', 12, 20, 32, 83),
(uuid_generate_v4(), 'Тосканские вечера', 'default_restaurant.jpg', 'Тосканские вечера на улице Мира', 4.6, 749, 'Ресторан Тосканские вечера, уютная атмосфера и вкусная еда.', 10, 20, 59, 88),
(uuid_generate_v4(), 'Суши и роллы', 'default_restaurant.jpg', 'Суши и роллы на улице Ленина', 2.4, 817, 'Ресторан Суши и роллы, уютная атмосфера и вкусная еда.', 8, 22, 59, 89),
(uuid_generate_v4(), 'Вкус Индии', 'default_restaurant.jpg', 'Вкус Индии на улице Карла Маркса', 2.0, 801, 'Ресторан Вкус Индии, уютная атмосфера и вкусная еда.', 11, 23, 42, 90),
(uuid_generate_v4(), 'Мексиканская площадь', 'default_restaurant.jpg', 'Мексиканская площадь на улице Ленина', 3.8, 70, 'Ресторан Мексиканская площадь, уютная атмосфера и вкусная еда.', 9, 20, 34, 75),
(uuid_generate_v4(), 'Греческая таверна', 'default_restaurant.jpg', 'Греческая таверна на улице Мира', 4.2, 297, 'Ресторан Греческая таверна, уютная атмосфера и вкусная еда.', 11, 20, 31, 88),
(uuid_generate_v4(), 'Пивной бар Баварии', 'default_restaurant.jpg', 'Пивной бар Баварии на улице Мира', 3.1, 400, 'Ресторан Пивной бар Баварии, уютная атмосфера и вкусная еда.', 8, 20, 30, 63),
(uuid_generate_v4(), 'Итальянский дворик', 'default_restaurant.jpg', 'Итальянский дворик на улице Ленина', 3.4, 65, 'Ресторан Итальянский дворик, уютная атмосфера и вкусная еда.', 12, 21, 52, 61),
(uuid_generate_v4(), 'Ресторан Печка', 'default_restaurant.jpg', 'Ресторан Печка на улице Ленина', 3.6, 431, 'Ресторан Ресторан Печка, уютная атмосфера и вкусная еда.', 11, 22, 31, 80),
(uuid_generate_v4(), 'Золотая рыба', 'default_restaurant.jpg', 'Золотая рыба на улице Мира', 4.6, 680, 'Ресторан Золотая рыба, уютная атмосфера и вкусная еда.', 12, 23, 38, 73),
(uuid_generate_v4(), 'Красное море', 'default_restaurant.jpg', 'Красное море на улице Ленина', 0.1, 448, 'Ресторан Красное море, уютная атмосфера и вкусная еда.', 12, 22, 59, 69),
(uuid_generate_v4(), 'Ресторан Томат', 'default_restaurant.jpg', 'Ресторан Томат на улице Мира', 4.3, 702, 'Ресторан Ресторан Томат, уютная атмосфера и вкусная еда.', 11, 20, 33, 79),
(uuid_generate_v4(), 'Турецкая кухня', 'default_restaurant.jpg', 'Турецкая кухня на улице Пушкина', 3.1, 185, 'Ресторан Турецкая кухня, уютная атмосфера и вкусная еда.', 8, 22, 46, 66),
(uuid_generate_v4(), 'Вегетарианская кухня', 'default_restaurant.jpg', 'Вегетарианская кухня на улице Мира', 3.2, 916, 'Ресторан Вегетарианская кухня, уютная атмосфера и вкусная еда.', 9, 21, 30, 83),
(uuid_generate_v4(), 'Ресторан Адель', 'default_restaurant.jpg', 'Ресторан Адель на улице Карла Маркса', 1.3, 999, 'Ресторан Ресторан Адель, уютная атмосфера и вкусная еда.', 11, 21, 42, 74),
(uuid_generate_v4(), 'Гриль и мясо', 'default_restaurant.jpg', 'Гриль и мясо на улице Карла Маркса', 4.9, 220, 'Ресторан Гриль и мясо, уютная атмосфера и вкусная еда.', 12, 22, 53, 81),
(uuid_generate_v4(), 'Том Ям', 'default_restaurant.jpg', 'Том Ям на улице Суворова', 2.3, 740, 'Ресторан Том Ям, уютная атмосфера и вкусная еда.', 11, 21, 41, 75),
(uuid_generate_v4(), 'Пельмени по-русски', 'default_restaurant.jpg', 'Пельмени по-русски на улице Пушкина', 0.8, 597, 'Ресторан Пельмени по-русски, уютная атмосфера и вкусная еда.', 10, 22, 31, 67),
(uuid_generate_v4(), 'Китайская кухня', 'default_restaurant.jpg', 'Китайская кухня на улице Пушкина', 1.3, 916, 'Ресторан Китайская кухня, уютная атмосфера и вкусная еда.', 12, 23, 48, 87),
(uuid_generate_v4(), 'Французская кухня', 'default_restaurant.jpg', 'Французская кухня на улице Карла Маркса', 4.9, 163, 'Ресторан Французская кухня, уютная атмосфера и вкусная еда.', 11, 21, 36, 81),
(uuid_generate_v4(), 'Средиземноморский ресторан', 'default_restaurant.jpg', 'Средиземноморский ресторан на улице Пушкина', 3.2, 455, 'Ресторан Средиземноморский ресторан, уютная атмосфера и вкусная еда.', 10, 22, 34, 64),
(uuid_generate_v4(), 'Ресторан Вкуса', 'default_restaurant.jpg', 'Ресторан Вкуса на улице Суворова', 0.7, 13, 'Ресторан Ресторан Вкуса, уютная атмосфера и вкусная еда.', 9, 21, 48, 69),
(uuid_generate_v4(), 'Шашлык-Бар', 'default_restaurant.jpg', 'Шашлык-Бар на улице Карла Маркса', 1.2, 95, 'Ресторан Шашлык-Бар, уютная атмосфера и вкусная еда.', 12, 23, 37, 86),
(uuid_generate_v4(), 'Паста на ужин', 'default_restaurant.jpg', 'Паста на ужин на улице Пушкина', 1.9, 844, 'Ресторан Паста на ужин, уютная атмосфера и вкусная еда.', 12, 23, 59, 66),
(uuid_generate_v4(), 'Веганский уголок', 'default_restaurant.jpg', 'Веганский уголок на улице Ленина', 2.7, 420, 'Ресторан Веганский уголок, уютная атмосфера и вкусная еда.', 8, 20, 52, 73),
(uuid_generate_v4(), 'Бургерная Сити', 'default_restaurant.jpg', 'Бургерная Сити на улице Пушкина', 1.3, 197, 'Ресторан Бургерная Сити, уютная атмосфера и вкусная еда.', 11, 21, 60, 67),
(uuid_generate_v4(), 'Ресторан Эдем', 'default_restaurant.jpg', 'Ресторан Эдем на улице Карла Маркса', 3.9, 937, 'Ресторан Ресторан Эдем, уютная атмосфера и вкусная еда.', 9, 20, 42, 78),
(uuid_generate_v4(), 'Ресторан Лаванда', 'default_restaurant.jpg', 'Ресторан Лаванда на улице Суворова', 3.2, 628, 'Ресторан Ресторан Лаванда, уютная атмосфера и вкусная еда.', 9, 20, 50, 64),
(uuid_generate_v4(), 'Ресторан Капрезе', 'default_restaurant.jpg', 'Ресторан Капрезе на улице Мира', 4.8, 162, 'Ресторан Ресторан Капрезе, уютная атмосфера и вкусная еда.', 8, 20, 32, 82),
(uuid_generate_v4(), 'Греческий зал', 'default_restaurant.jpg', 'Греческий зал на улице Ленина', 4.0, 710, 'Ресторан Греческий зал, уютная атмосфера и вкусная еда.', 9, 23, 40, 81),
(uuid_generate_v4(), 'Пицца и Суши', 'default_restaurant.jpg', 'Пицца и Суши на улице Суворова', 4.0, 937, 'Ресторан Пицца и Суши, уютная атмосфера и вкусная еда.', 11, 21, 52, 71),
(uuid_generate_v4(), 'Турецкий Султан', 'default_restaurant.jpg', 'Турецкий Султан на улице Пушкина', 3.5, 175, 'Ресторан Турецкий Султан, уютная атмосфера и вкусная еда.', 12, 21, 49, 81),
(uuid_generate_v4(), 'Мексиканский уголок', 'default_restaurant.jpg', 'Мексиканский уголок на улице Суворова', 3.6, 290, 'Ресторан Мексиканский уголок, уютная атмосфера и вкусная еда.', 10, 23, 56, 63),
(uuid_generate_v4(), 'Ресторан Мозаика', 'default_restaurant.jpg', 'Ресторан Мозаика на улице Ленина', 2.9, 105, 'Ресторан Ресторан Мозаика, уютная атмосфера и вкусная еда.', 11, 23, 40, 68),
(uuid_generate_v4(), 'Шашлыки по-кавказски', 'default_restaurant.jpg', 'Шашлыки по-кавказски на улице Ленина', 3.6, 323, 'Ресторан Шашлыки по-кавказски, уютная атмосфера и вкусная еда.', 8, 23, 55, 69),
(uuid_generate_v4(), 'Французская кухня на ужин', 'default_restaurant.jpg', 'Французская кухня на ужин на улице Мира', 4.6, 680, 'Ресторан Французская кухня на ужин, уютная атмосфера и вкусная еда.', 12, 20, 37, 63),
(uuid_generate_v4(), 'Мексиканская кухня для всех', 'default_restaurant.jpg', 'Мексиканская кухня для всех на улице Мира', 3.8, 793, 'Ресторан Мексиканская кухня для всех, уютная атмосфера и вкусная еда.', 9, 23, 49, 90),
(uuid_generate_v4(), 'Томаты и Паста', 'default_restaurant.jpg', 'Томаты и Паста на улице Карла Маркса', 1.2, 329, 'Ресторан Томаты и Паста, уютная атмосфера и вкусная еда.', 9, 20, 53, 76);

INSERT INTO restaurant_tags_relations (restaurant_id, tag_id)
VALUES
((SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Американская пекарня'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Суши и роллы'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Мозаика'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Скоро будет'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Веганское счастье'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Морская лагуна'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Турецкая кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Китайская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Суши Дрим'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Средиземноморья'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Итальянская ривьера'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Баварская пивоварня'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Шашлыки от Бабая'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Тосканский огонь'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Турецкая кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Зеленая вилка'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Шашлык-Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Печка'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Красное море'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Пельмени по-русски'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Греческий зал'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Суши Дрим'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Китайская империя'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Тосканские вечера'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Красное море'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Пельмени по-русски'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Греческий зал'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Средиземноморский ресторан'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская кухня для всех'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Скоро будет'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Вегетарианская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Китайская звезда'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Капрезе'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Французская кухня на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Гриль Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Золотая рыба'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Баварский пивной сад'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Баварская пивоварня'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Зеленая вилка'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская пекарня'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Красное море'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Итальянский дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Гриль Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Скоро будет'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская площадь'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Средиземноморья'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Пивной бар Баварии'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская кухня для всех'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Суши Мания'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Тосканский огонь'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Томат'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Баварский пивной сад'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Вкуса'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Вкус Индии'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Суши и роллы'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Вкус Индии'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Морская звезда'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Турецкая радость'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Эдем'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Французская кухня на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Вкус Индии'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Суши Дрим'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Веганский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Капрезе'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Индийские специи'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Итальянский дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Тосканский огонь'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Баварский пивной сад'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Веганский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Вкуса'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Вкус Индии'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Мексиканский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Ресторан на воде'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Вкуса'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Тосканские вечера'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Французский бистро'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Пельмени на углях'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан на воде'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская кухня для всех'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Средиземноморский ресторан'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский'));


INSERT INTO products (restaurant_id, name, price, image_url, weight, category)
VALUES 
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' AND description = 'Японский рамен и суши' AND type = 'Японский'),
        'Рамен с курицей', 
        740, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' AND description = 'Японский рамен и суши' AND type = 'Японский'),
        'Рамен с говядиной', 
        650, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' AND description = 'Японский рамен и суши' AND type = 'Японский'),
        'Суши ассорти', 
        490, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' AND description = 'Японский рамен и суши' AND type = 'Японский'),
        'Тамаго суши', 
        400, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' AND description = 'Японский рамен и суши' AND type = 'Японский'),
        'Ролл с лососем', 
        300, 
        'default_product.jpg', 
        200,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' AND description = 'Японский рамен и суши' AND type = 'Японский'),
        'Гёдза', 
        200, 
        'default_product.jpg', 
        180,
        'Закуски'
    );


INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash)
VALUES (
    uuid_generate_v4(), 
    'testuser', 
    'Dmitriy',  
    'Nagiev', 
    '88005553535', 
    '',
    'default_user.jpg',
    decode('a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6', 'hex')
);

