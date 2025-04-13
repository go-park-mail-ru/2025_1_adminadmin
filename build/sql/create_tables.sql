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
    address_id TEXT NOT NULL,
    order_products TEXT NOT NULL,

    apartment_or_office TEXT,
    intercom TEXT,
    entrance TEXT,
    floor TEXT,
    courier_comment TEXT,
    leave_at_door BOOLEAN DEFAULT FALSE,
    final_price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
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
(uuid_generate_v4(), 'Паста и Вино', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Мира, д. 193', 3.6, 833, 'Ресторан Паста и Вино, уютная атмосфера и вкусная еда.', 11, 20, 47, 63),
(uuid_generate_v4(), 'Суши Дрим', 'default_restaurant.jpg', 'Новосибирск, ул. Пушкина, д. 112', 4.3, 965, 'Ресторан Суши Дрим, уютная атмосфера и вкусная еда.', 9, 22, 41, 65),
(uuid_generate_v4(), 'Бургерная Ривьера', 'default_restaurant.jpg', 'Москва, ул. Карла Маркса, д. 46', 2.2, 770, 'Ресторан Бургерная Ривьера, уютная атмосфера и вкусная еда.', 8, 22, 51, 66),
(uuid_generate_v4(), 'Турецкий базар', 'default_restaurant.jpg', 'Москва, ул. Карла Маркса, д. 70', 2.4, 830, 'Ресторан Турецкий базар, уютная атмосфера и вкусная еда.', 10, 20, 40, 77),
(uuid_generate_v4(), 'Зеленая вилка', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 141', 0.8, 237, 'Ресторан Зеленая вилка, уютная атмосфера и вкусная еда.', 11, 21, 44, 70),
(uuid_generate_v4(), 'Гриль Бар', 'default_restaurant.jpg', 'Москва, ул. Ленина, д. 195', 2.1, 95, 'Ресторан Гриль Бар, уютная атмосфера и вкусная еда.', 9, 21, 32, 83),
(uuid_generate_v4(), 'Американская кухня', 'default_restaurant.jpg', 'Новосибирск, ул. Ленина, д. 147', 3.4, 902, 'Ресторан Американская кухня, уютная атмосфера и вкусная еда.', 11, 20, 38, 67),
(uuid_generate_v4(), 'Ресторан Средиземноморья', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 128', 2.2, 7, 'Ресторан Ресторан Средиземноморья, уютная атмосфера и вкусная еда.', 8, 23, 57, 88),
(uuid_generate_v4(), 'Индийские специи', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Мира, д. 118', 2.5, 820, 'Ресторан Индийские специи, уютная атмосфера и вкусная еда.', 9, 21, 60, 71),
(uuid_generate_v4(), 'Веганское счастье', 'default_restaurant.jpg', 'Москва, ул. Мира, д. 49', 3.5, 748, 'Ресторан Веганское счастье, уютная атмосфера и вкусная еда.', 12, 23, 51, 67),
(uuid_generate_v4(), 'Французский уголок', 'default_restaurant.jpg', 'Новосибирск, ул. Карла Маркса, д. 78', 3.9, 145, 'Ресторан Французский уголок, уютная атмосфера и вкусная еда.', 10, 23, 51, 83),
(uuid_generate_v4(), 'Мексиканская пекарня', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 101', 0.2, 846, 'Ресторан Мексиканская пекарня, уютная атмосфера и вкусная еда.', 9, 22, 58, 82),
(uuid_generate_v4(), 'Китайская империя', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 129', 2.9, 71, 'Ресторан Китайская империя, уютная атмосфера и вкусная еда.', 9, 20, 50, 88),
(uuid_generate_v4(), 'Баварский пивной сад', 'default_restaurant.jpg', 'Казань, ул. Пушкина, д. 84', 4.7, 121, 'Ресторан Баварский пивной сад, уютная атмосфера и вкусная еда.', 11, 22, 36, 60),
(uuid_generate_v4(), 'Морская звезда', 'default_restaurant.jpg', 'Екатеринбург, ул. Карла Маркса, д. 13', 2.8, 581, 'Ресторан Морская звезда, уютная атмосфера и вкусная еда.', 11, 21, 41, 66),
(uuid_generate_v4(), 'Шашлыки от Бабая', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Мира, д. 60', 1.3, 828, 'Ресторан Шашлыки от Бабая, уютная атмосфера и вкусная еда.', 8, 20, 34, 62),
(uuid_generate_v4(), 'Скоро будет', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 31', 0.4, 790, 'Ресторан Скоро будет, уютная атмосфера и вкусная еда.', 12, 22, 37, 66),
(uuid_generate_v4(), 'Восточный базар', 'default_restaurant.jpg', 'Казань, ул. Пушкина, д. 174', 1.4, 140, 'Ресторан Восточный базар, уютная атмосфера и вкусная еда.', 11, 22, 46, 78),
(uuid_generate_v4(), 'Греческий дворик', 'default_restaurant.jpg', 'Новосибирск, ул. Карла Маркса, д. 140', 4.4, 909, 'Ресторан Греческий дворик, уютная атмосфера и вкусная еда.', 9, 21, 49, 83),
(uuid_generate_v4(), 'Тосканский огонь', 'default_restaurant.jpg', 'Екатеринбург, ул. Суворова, д. 191', 0.1, 76, 'Ресторан Тосканский огонь, уютная атмосфера и вкусная еда.', 12, 20, 32, 63),
(uuid_generate_v4(), 'Итальянская ривьера', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 91', 4.7, 534, 'Ресторан Итальянская ривьера, уютная атмосфера и вкусная еда.', 12, 23, 35, 67),
(uuid_generate_v4(), 'Суши Мания', 'default_restaurant.jpg', 'Новосибирск, ул. Пушкина, д. 128', 0.8, 164, 'Ресторан Суши Мания, уютная атмосфера и вкусная еда.', 8, 22, 55, 65),
(uuid_generate_v4(), 'Пельмени на углях', 'default_restaurant.jpg', 'Екатеринбург, ул. Пушкина, д. 38', 3.0, 538, 'Ресторан Пельмени на углях, уютная атмосфера и вкусная еда.', 8, 20, 59, 70),
(uuid_generate_v4(), 'Бургеры по-американски', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 88', 4.9, 470, 'Ресторан Бургеры по-американски, уютная атмосфера и вкусная еда.', 10, 20, 41, 67),
(uuid_generate_v4(), 'Китайская звезда', 'default_restaurant.jpg', 'Москва, ул. Пушкина, д. 162', 4.5, 549, 'Ресторан Китайская звезда, уютная атмосфера и вкусная еда.', 8, 23, 38, 63),
(uuid_generate_v4(), 'Мексиканская закуска', 'default_restaurant.jpg', 'Казань, ул. Мира, д. 185', 0.6, 173, 'Ресторан Мексиканская закуска, уютная атмосфера и вкусная еда.', 11, 20, 35, 85),
(uuid_generate_v4(), 'Французский бистро', 'default_restaurant.jpg', 'Москва, ул. Карла Маркса, д. 193', 1.4, 611, 'Ресторан Французский бистро, уютная атмосфера и вкусная еда.', 10, 22, 31, 81),
(uuid_generate_v4(), 'Греческий остров', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 64', 1.0, 828, 'Ресторан Греческий остров, уютная атмосфера и вкусная еда.', 10, 21, 59, 62),
(uuid_generate_v4(), 'Турецкая радость', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Ленина, д. 168', 0.3, 13, 'Ресторан Турецкая радость, уютная атмосфера и вкусная еда.', 11, 23, 60, 62),
(uuid_generate_v4(), 'Индийская сказка', 'default_restaurant.jpg', 'Новосибирск, ул. Ленина, д. 2', 2.8, 674, 'Ресторан Индийская сказка, уютная атмосфера и вкусная еда.', 12, 23, 49, 89),
(uuid_generate_v4(), 'Американская пекарня', 'default_restaurant.jpg', 'Казань, ул. Мира, д. 179', 4.5, 391, 'Ресторан Американская пекарня, уютная атмосфера и вкусная еда.', 9, 21, 39, 65),
(uuid_generate_v4(), 'Восточный салат', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 67', 1.1, 991, 'Ресторан Восточный салат, уютная атмосфера и вкусная еда.', 9, 20, 50, 88),
(uuid_generate_v4(), 'Вегетарианский рай', 'default_restaurant.jpg', 'Казань, ул. Пушкина, д. 111', 3.0, 729, 'Ресторан Вегетарианский рай, уютная атмосфера и вкусная еда.', 10, 22, 43, 86),
(uuid_generate_v4(), 'Ресторан на воде', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 130', 0.6, 456, 'Ресторан Ресторан на воде, уютная атмосфера и вкусная еда.', 9, 20, 51, 65),
(uuid_generate_v4(), 'Баварская пивоварня', 'default_restaurant.jpg', 'Казань, ул. Мира, д. 120', 3.0, 794, 'Ресторан Баварская пивоварня, уютная атмосфера и вкусная еда.', 10, 21, 41, 60),
(uuid_generate_v4(), 'Морская лагуна', 'default_restaurant.jpg', 'Екатеринбург, ул. Суворова, д. 136', 0.3, 335, 'Ресторан Морская лагуна, уютная атмосфера и вкусная еда.', 8, 23, 54, 86),
(uuid_generate_v4(), 'Тосканские вечера', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Ленина, д. 21', 3.0, 536, 'Ресторан Тосканские вечера, уютная атмосфера и вкусная еда.', 11, 23, 44, 86),
(uuid_generate_v4(), 'Суши и роллы', 'default_restaurant.jpg', 'Казань, ул. Карла Маркса, д. 34', 0.2, 847, 'Ресторан Суши и роллы, уютная атмосфера и вкусная еда.', 9, 23, 36, 72),
(uuid_generate_v4(), 'Вкус Индии', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Мира, д. 151', 0.7, 203, 'Ресторан Вкус Индии, уютная атмосфера и вкусная еда.', 12, 22, 53, 75),
(uuid_generate_v4(), 'Мексиканская площадь', 'default_restaurant.jpg', 'Москва, ул. Карла Маркса, д. 138', 4.0, 774, 'Ресторан Мексиканская площадь, уютная атмосфера и вкусная еда.', 8, 23, 53, 78),
(uuid_generate_v4(), 'Греческая таверна', 'default_restaurant.jpg', 'Новосибирск, ул. Карла Маркса, д. 171', 2.2, 84, 'Ресторан Греческая таверна, уютная атмосфера и вкусная еда.', 8, 21, 30, 84),
(uuid_generate_v4(), 'Пивной бар Баварии', 'default_restaurant.jpg', 'Новосибирск, ул. Мира, д. 152', 2.9, 465, 'Ресторан Пивной бар Баварии, уютная атмосфера и вкусная еда.', 10, 20, 39, 80),
(uuid_generate_v4(), 'Итальянский дворик', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Мира, д. 164', 2.3, 219, 'Ресторан Итальянский дворик, уютная атмосфера и вкусная еда.', 9, 20, 51, 86),
(uuid_generate_v4(), 'Ресторан Печка', 'default_restaurant.jpg', 'Москва, ул. Ленина, д. 163', 4.4, 938, 'Ресторан Ресторан Печка, уютная атмосфера и вкусная еда.', 10, 22, 45, 76),
(uuid_generate_v4(), 'Золотая рыба', 'default_restaurant.jpg', 'Казань, ул. Пушкина, д. 84', 3.9, 902, 'Ресторан Золотая рыба, уютная атмосфера и вкусная еда.', 9, 22, 50, 85),
(uuid_generate_v4(), 'Красное море', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Ленина, д. 200', 3.6, 554, 'Ресторан Красное море, уютная атмосфера и вкусная еда.', 10, 21, 56, 80),
(uuid_generate_v4(), 'Ресторан Томат', 'default_restaurant.jpg', 'Екатеринбург, ул. Суворова, д. 43', 4.0, 978, 'Ресторан Ресторан Томат, уютная атмосфера и вкусная еда.', 9, 23, 46, 71),
(uuid_generate_v4(), 'Турецкая кухня', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 151', 4.2, 838, 'Ресторан Турецкая кухня, уютная атмосфера и вкусная еда.', 9, 20, 33, 76),
(uuid_generate_v4(), 'Вегетарианская кухня', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Суворова, д. 195', 0.1, 602, 'Ресторан Вегетарианская кухня, уютная атмосфера и вкусная еда.', 8, 22, 50, 67),
(uuid_generate_v4(), 'Ресторан Адель', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 104', 0.0, 829, 'Ресторан Ресторан Адель, уютная атмосфера и вкусная еда.', 12, 23, 45, 75),
(uuid_generate_v4(), 'Гриль и мясо', 'default_restaurant.jpg', 'Москва, ул. Суворова, д. 95', 3.6, 242, 'Ресторан Гриль и мясо, уютная атмосфера и вкусная еда.', 9, 23, 33, 71),
(uuid_generate_v4(), 'Том Ям', 'default_restaurant.jpg', 'Екатеринбург, ул. Карла Маркса, д. 54', 1.1, 510, 'Ресторан Том Ям, уютная атмосфера и вкусная еда.', 12, 21, 37, 71),
(uuid_generate_v4(), 'Пельмени по-русски', 'default_restaurant.jpg', 'Казань, ул. Карла Маркса, д. 157', 0.3, 899, 'Ресторан Пельмени по-русски, уютная атмосфера и вкусная еда.', 10, 22, 39, 82),
(uuid_generate_v4(), 'Китайская кухня', 'default_restaurant.jpg', 'Новосибирск, ул. Суворова, д. 140', 3.0, 817, 'Ресторан Китайская кухня, уютная атмосфера и вкусная еда.', 9, 23, 54, 74),
(uuid_generate_v4(), 'Французская кухня', 'default_restaurant.jpg', 'Москва, ул. Ленина, д. 26', 1.6, 801, 'Ресторан Французская кухня, уютная атмосфера и вкусная еда.', 9, 21, 31, 65),
(uuid_generate_v4(), 'Средиземноморский ресторан', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Ленина, д. 80', 0.3, 889, 'Ресторан Средиземноморский ресторан, уютная атмосфера и вкусная еда.', 9, 21, 48, 79),
(uuid_generate_v4(), 'Ресторан Вкуса', 'default_restaurant.jpg', 'Москва, ул. Пушкина, д. 169', 4.2, 769, 'Ресторан Ресторан Вкуса, уютная атмосфера и вкусная еда.', 9, 20, 50, 64),
(uuid_generate_v4(), 'Шашлык-Бар', 'default_restaurant.jpg', 'Казань, ул. Ленина, д. 119', 3.3, 414, 'Ресторан Шашлык-Бар, уютная атмосфера и вкусная еда.', 12, 21, 44, 66),
(uuid_generate_v4(), 'Паста на ужин', 'default_restaurant.jpg', 'Казань, ул. Мира, д. 17', 0.7, 221, 'Ресторан Паста на ужин, уютная атмосфера и вкусная еда.', 8, 20, 39, 67),
(uuid_generate_v4(), 'Веганский уголок', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 163', 2.8, 326, 'Ресторан Веганский уголок, уютная атмосфера и вкусная еда.', 11, 22, 42, 76),
(uuid_generate_v4(), 'Бургерная Сити', 'default_restaurant.jpg', 'Казань, ул. Ленина, д. 24', 3.8, 532, 'Ресторан Бургерная Сити, уютная атмосфера и вкусная еда.', 9, 23, 35, 77),
(uuid_generate_v4(), 'Ресторан Эдем', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Суворова, д. 150', 0.5, 719, 'Ресторан Ресторан Эдем, уютная атмосфера и вкусная еда.', 10, 22, 46, 80),
(uuid_generate_v4(), 'Ресторан Лаванда', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 83', 3.3, 354, 'Ресторан Ресторан Лаванда, уютная атмосфера и вкусная еда.', 11, 22, 54, 86),
(uuid_generate_v4(), 'Ресторан Капрезе', 'default_restaurant.jpg', 'Казань, ул. Карла Маркса, д. 26', 3.7, 597, 'Ресторан Ресторан Капрезе, уютная атмосфера и вкусная еда.', 8, 21, 44, 61),
(uuid_generate_v4(), 'Греческий зал', 'default_restaurant.jpg', 'Москва, ул. Карла Маркса, д. 84', 1.2, 508, 'Ресторан Греческий зал, уютная атмосфера и вкусная еда.', 12, 22, 35, 87),
(uuid_generate_v4(), 'Пицца и Суши', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Пушкина, д. 63', 4.2, 365, 'Ресторан Пицца и Суши, уютная атмосфера и вкусная еда.', 10, 20, 36, 63),
(uuid_generate_v4(), 'Турецкий Султан', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Карла Маркса, д. 10', 1.4, 819, 'Ресторан Турецкий Султан, уютная атмосфера и вкусная еда.', 12, 21, 37, 79),
(uuid_generate_v4(), 'Мексиканский уголок', 'default_restaurant.jpg', 'Екатеринбург, ул. Ленина, д. 150', 2.4, 752, 'Ресторан Мексиканский уголок, уютная атмосфера и вкусная еда.', 11, 20, 33, 71),
(uuid_generate_v4(), 'Ресторан Мозаика', 'default_restaurant.jpg', 'Екатеринбург, ул. Пушкина, д. 101', 3.7, 44, 'Ресторан Ресторан Мозаика, уютная атмосфера и вкусная еда.', 11, 22, 40, 69),
(uuid_generate_v4(), 'Шашлыки по-кавказски', 'default_restaurant.jpg', 'Москва, ул. Карла Маркса, д. 156', 0.0, 801, 'Ресторан Шашлыки по-кавказски, уютная атмосфера и вкусная еда.', 10, 21, 53, 88),
(uuid_generate_v4(), 'Французская кухня на ужин', 'default_restaurant.jpg', 'Санкт-Петербург, ул. Пушкина, д. 187', 4.5, 771, 'Ресторан Французская кухня на ужин, уютная атмосфера и вкусная еда.', 9, 21, 31, 84),
(uuid_generate_v4(), 'Мексиканская кухня для всех', 'default_restaurant.jpg', 'Екатеринбург, ул. Карла Маркса, д. 182', 2.1, 484, 'Ресторан Мексиканская кухня для всех, уютная атмосфера и вкусная еда.', 12, 23, 39, 79),
(uuid_generate_v4(), 'Томаты и Паста', 'default_restaurant.jpg', 'Екатеринбург, ул. Мира, д. 114', 0.6, 76, 'Ресторан Томаты и Паста, уютная атмосфера и вкусная еда.', 11, 22, 51, 74);

INSERT INTO restaurant_tags_relations (restaurant_id, tag_id)
VALUES
((SELECT id FROM restaurants WHERE name = 'Ресторан Томат'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Китайская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Паста и Вино'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Морская лагуна'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Веганский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Суши и роллы'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Тосканские вечера'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Веганское счастье'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Гриль Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Французский бистро'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Суши и роллы'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Греческий зал'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Эдем'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Зеленая вилка'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Томат'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Гриль Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская площадь'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Пивной бар Баварии'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Вкуса'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Пивной бар Баварии'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Лаванда'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Итальянский дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Пивной бар Баварии'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Греческий дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Гриль и мясо'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Мозаика'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Турецкая радость'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Китайская звезда'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Морская звезда'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Баварская пивоварня'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Бургерная Ривьера'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Пицца и Суши'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Средиземноморский ресторан'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Мозаика'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Веганский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Индийские специи'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Французский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Пельмени по-русски'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Американская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Шашлык-Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Морская звезда'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Пицца и Суши'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Пицца и Суши'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Средиземноморья'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Бургерная Сити'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
((SELECT id FROM restaurants WHERE name = 'Греческий остров'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Пицца и Суши'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская площадь'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Пельмени на углях'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Баварская пивоварня'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Бургерная Сити'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Восточный салат'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Греческий зал'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Индийские специи'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Золотая рыба'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Восточный салат'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Суши и роллы'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Веганское счастье'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Веганский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
((SELECT id FROM restaurants WHERE name = 'Баварская пивоварня'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Греческий остров'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Шашлык-Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Шашлык-Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Китайская империя'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская пекарня'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Гриль Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
((SELECT id FROM restaurants WHERE name = 'Веганское счастье'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Китайская империя'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская площадь'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Греческая таверна'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Греческий дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Мозаика'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Американская пекарня'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Гриль и мясо'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
((SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Средиземноморский ресторан'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Вкуса'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Суши Мания'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
((SELECT id FROM restaurants WHERE name = 'Паста на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Французский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
((SELECT id FROM restaurants WHERE name = 'Итальянская ривьера'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
((SELECT id FROM restaurants WHERE name = 'Ресторан Адель'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
((SELECT id FROM restaurants WHERE name = 'Китайская империя'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
((SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
((SELECT id FROM restaurants WHERE name = 'Красное море'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
((SELECT id FROM restaurants WHERE name = 'Турецкий базар'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
((SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
((SELECT id FROM restaurants WHERE name = 'Американская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский'));

INSERT INTO products (restaurant_id, name, price, image_url, weight, category)
VALUES 
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с курицей', 
        740, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с говядиной', 
        650, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с ананасом', 
        640, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Пельмени с сыром', 
        550, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Суши ассорти', 
        490, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Тамаго суши', 
        400, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Сырная тарелка', 
        790, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Вареники с грибами', 
        800, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Ролл с лососем', 
        300, 
        'default_product.jpg', 
        200,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Гёдза', 
        200, 
        'default_product.jpg', 
        180,
        'Закуски'
    );

INSERT INTO products (restaurant_id, name, price, image_url, weight, category)
VALUES 
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с курицей', 
        740, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с говядиной', 
        650, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с ананасом', 
        640, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Пельмени с сыром', 
        550, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Суши ассорти', 
        490, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Тамаго суши', 
        400, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Сырная тарелка', 
        790, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Вареники с грибами', 
        800, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Ролл с лососем', 
        300, 
        'default_product.jpg', 
        200,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Гёдза', 
        200, 
        'default_product.jpg', 
        180,
        'Закуски'
    );
    
INSERT INTO products (restaurant_id, name, price, image_url, weight, category)
VALUES 
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с курицей', 
        740, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с говядиной', 
        650, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Рамен с ананасом', 
        640, 
        'default_product.jpg', 
        350,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Пельмени с сыром', 
        550, 
        'default_product.jpg', 
        400,
        'Закуски'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Суши ассорти', 
        490, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Тамаго суши', 
        400, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Сырная тарелка', 
        790, 
        'default_product.jpg', 
        250,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Вареники с грибами', 
        800, 
        'default_product.jpg', 
        150,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
        'Ролл с лососем', 
        300, 
        'default_product.jpg', 
        200,
        'Суши'
    ),
    (
        (SELECT id FROM restaurants WHERE name = 'Красное море' ),
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



-- Generated data inserts
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('4c40d5f6-51ce-480f-a566-9fd4f381c8e1', 'user0', '+7-565-790-4256', 'Сергей', 'Шипулина', '', 'default_user.jpg', decode('1ce80a26def5cdbbc0be5cb8e98ea69d2dc48fd1d2f8a4e97fc0f9213762230d', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('3ccb6b26-2f63-446a-9042-39379fe46b09', 'user1', '+7-788-285-5341', 'Владислав', 'Торетто', '', 'default_user.jpg', decode('6ff859539fde609d25b020323a67843e1a54b3eeb48e74d9df04568a2ed091c6', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('a24a0e09-2a12-49a0-807a-22a39e3fb8a7', 'user2', '+7-213-192-1788', 'Алексей', 'Сидоров', '', 'default_user.jpg', decode('253f2f1b17d9ee78155f774d8757ab60de5717739e6e84b547f66056cbb8419c', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('eb7f1f8c-0dc8-4bdf-abe9-5f1ef2a8627e', 'user3', '+7-181-905-9393', 'Владислав', 'Иванов', '', 'default_user.jpg', decode('6f4b2376ba9f58fdd168d541d05bba155ffe991946d2d88a3295a8a4ece10ab4', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('482a75be-e6b5-4f0e-b56d-01602a1c6b36', 'user4', '+7-428-696-5007', 'Дарья', 'Петрова', '', 'default_user.jpg', decode('e1cd10a18c0932bac2a0fff3f3207ad48bab12220131ddcc2e1d7edd2be6f47a', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('dbc82af9-495f-4690-a314-adeb216affa2', 'user5', '+7-706-779-7599', 'Дарья', 'Иванов', '', 'default_user.jpg', decode('4e6892f7b02f172ef4a3542259065d7042494d524260bb39999e9675ee64cab8', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('8421cf5b-a672-4070-938f-2c9f7c24c81b', 'user6', '+7-621-104-8110', 'Иван', 'Смирнова', '', 'default_user.jpg', decode('fc8efdb437d23cf136c21ff2bcd10c43abbed12c8097af11325eb94c0f8b5b8f', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('360d8748-ef26-468c-916f-0f5f68558f2e', 'user7', '+7-968-708-4865', 'Мария', 'Попов', '', 'default_user.jpg', decode('dd2a8d365f6ee8d2391e5d0969fa934c1fe3552b8d67b92cb961107e35493979', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('23c7701e-f25a-4ad6-9b51-33bcc8830c05', 'user8', '+7-626-419-1707', 'Алексей', 'Шипулина', '', 'default_user.jpg', decode('acd34126d0b39b15ea1428542cacdbd98ee4e57a3956f641add0b671bff7927b', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('8c4aa27f-3a80-4b49-8113-7991dba5603e', 'user9', '+7-683-393-2770', 'Никита', 'Шипулина', '', 'default_user.jpg', decode('8c559cb6925bfa587ced9deefdd53433ca1e501d6fec2658d918b17c2b9b820e', 'hex'));
INSERT INTO restaurant_tags (id, name) VALUES ('d64f5818-ad51-4e49-b1db-4b2cfdcdfff0', 'Пицца');
INSERT INTO restaurant_tags (id, name) VALUES ('50f1ce38-fd6a-4edb-b47c-44373fec7d43', 'Бургеры');
INSERT INTO restaurant_tags (id, name) VALUES ('efd0023e-6999-46e7-a6b2-019e8b277b7d', 'Суши');
INSERT INTO restaurant_tags (id, name) VALUES ('aa7a3d7a-982a-494f-88d6-a934588cff44', 'Веган');
INSERT INTO restaurant_tags (id, name) VALUES ('32a78cbb-14df-4f18-9df3-609a24ae0725', 'Кофе');
INSERT INTO restaurant_tags (id, name) VALUES ('d98129e9-ea4d-45e6-b672-636b019f7e67', 'Десерты');
INSERT INTO restaurant_tags (id, name) VALUES ('46df10d3-9ae5-4344-957d-6c02f7d9661e', 'Шаурма');
INSERT INTO restaurant_tags (id, name) VALUES ('9bb1ac2c-10f6-4137-a6f1-32a9132fe144', 'Грузинская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('6eceffc9-4193-4d73-9c0e-e3d8c713d1db', 'Салаты');
INSERT INTO restaurant_tags (id, name) VALUES ('8431640f-6132-48a2-99b2-af81bf6f5ab2', 'Завтраки');
INSERT INTO restaurant_tags (id, name) VALUES ('9fe68db1-832f-45c3-be4f-5354a26e20ca', 'Стейки');
INSERT INTO restaurant_tags (id, name) VALUES ('c625a00b-af8e-47a9-ac89-f8125633021d', 'Морепродукты');
INSERT INTO restaurant_tags (id, name) VALUES ('8b1b44e7-6b73-4033-bc1e-21cab30ecec5', 'Пасты');
INSERT INTO restaurant_tags (id, name) VALUES ('2eef3843-d609-4390-8169-9d51a0c358cd', 'Смузи');
INSERT INTO restaurant_tags (id, name) VALUES ('7ab339c3-b516-4839-968a-a9c0772722ed', 'Фалафель');
INSERT INTO restaurant_tags (id, name) VALUES ('45fdff4d-7ff4-4927-90a5-7352003b6782', 'Гриль');
INSERT INTO restaurant_tags (id, name) VALUES ('7cf2c1ba-9f41-4c3c-8bcc-69a278704332', 'Курица');
INSERT INTO restaurant_tags (id, name) VALUES ('03875339-7720-4247-a383-75a57f5d126e', 'Рамен');
INSERT INTO restaurant_tags (id, name) VALUES ('471e9332-aa8c-4e01-b7ea-645f47a6a5e2', 'Корейская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('f2a8f08b-6ecf-4269-aa24-c18b8c501db8', 'Пекарня');
INSERT INTO restaurant_tags (id, name) VALUES ('50a85385-4cce-4935-bb37-678d05dc79c6', 'Пельмени');
INSERT INTO restaurant_tags (id, name) VALUES ('be6c0418-ab98-4f0d-bfd1-bd2acd67e5a0', 'Вьетнамская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('44720e92-368c-4fbd-9b5c-c9ce3ff00563', 'Сибирская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('f2f3a4cf-3c39-4701-b019-17b81871fe3e', 'ЗОЖ');
INSERT INTO restaurant_tags (id, name) VALUES ('dbd87a66-b43d-4425-a250-2984a359c487', 'Кето');
INSERT INTO restaurant_tags (id, name) VALUES ('7ab99ec9-d2aa-4907-9423-cd19247a60fe', 'Халяль');
INSERT INTO restaurant_tags (id, name) VALUES ('1aa85122-ef44-408a-9a80-4118accf2577', 'Безглютеновое');
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Ресторан 1', 4.7, 55);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('90d24345-56a3-4852-b5fa-793294e8458f', 'Ресторан 2', 3.8, 90);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6c0b9543-4d60-4097-96b8-c8082bca7835', 'Ресторан 3', 3.0, 129);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e42a7054-3445-4e1d-a5db-63987ccac19a', 'Ресторан 4', 4.7, 82);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6b5f5d61-bf01-4224-9c17-44351a234245', 'Ресторан 5', 3.7, 76);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Ресторан 6', 3.7, 143);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Ресторан 7', 4.8, 99);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Ресторан 8', 3.0, 62);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Ресторан 9', 3.5, 110);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Ресторан 10', 3.4, 61);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3ba4133c-7c68-4a38-af8f-1c217107b148', 'Ресторан 11', 3.4, 123);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Ресторан 12', 3.2, 74);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('434798a4-d545-4b7d-9a15-e2428c391f0d', 'Ресторан 13', 3.9, 67);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c368e238-5f15-49b4-91b3-2e5305117e32', 'Ресторан 14', 4.3, 73);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Ресторан 15', 4.2, 79);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1bd888d9-74be-40c7-85dd-c3d531887f25', 'Ресторан 16', 4.3, 67);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Ресторан 17', 4.0, 136);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Ресторан 18', 4.0, 86);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Ресторан 19', 4.5, 142);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Ресторан 20', 4.5, 147);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Ресторан 21', 4.8, 61);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Ресторан 22', 3.0, 107);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Ресторан 23', 3.0, 68);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7114c561-173b-4efa-957f-cd58298d6f7c', 'Ресторан 24', 4.8, 135);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Ресторан 25', 4.1, 143);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Ресторан 26', 4.9, 68);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Ресторан 27', 4.7, 62);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('27ab3cc4-db11-402f-900b-2efcbb393813', 'Ресторан 28', 3.5, 118);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Ресторан 29', 4.0, 136);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Ресторан 30', 3.6, 111);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Ресторан 31', 4.5, 89);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Ресторан 32', 3.5, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Ресторан 33', 3.0, 111);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Ресторан 34', 4.9, 93);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Ресторан 35', 4.8, 150);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1fa15cf7-704b-4665-a926-e9560fa95141', 'Ресторан 36', 4.1, 103);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c08349e7-d213-481c-9198-13c8c098c0f4', 'Ресторан 37', 3.0, 135);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Ресторан 38', 3.6, 140);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Ресторан 39', 4.3, 102);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('96965548-5a6e-41e4-a4f1-2484b512d119', 'Ресторан 40', 3.6, 58);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('67254fce-e401-46a3-8a83-ae8af0833f52', 'Ресторан 41', 3.6, 139);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Ресторан 42', 4.1, 83);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Ресторан 43', 4.0, 96);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Ресторан 44', 4.0, 117);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Ресторан 45', 4.5, 133);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('ed75802c-33e4-49d9-8853-97504a4a51ff', 'Ресторан 46', 3.5, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Ресторан 47', 3.5, 50);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Ресторан 48', 4.9, 129);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('41879f55-4f34-405c-87b0-a446bcc52406', 'Ресторан 49', 4.4, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Ресторан 50', 3.6, 136);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Ресторан 51', 3.5, 134);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('868fae50-e4a0-4ca6-8311-341fb949a47c', 'Ресторан 52', 4.5, 127);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Ресторан 53', 4.9, 114);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Ресторан 54', 4.2, 66);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('f56d68e6-9497-487d-b609-181aefa4d3d1', 'Ресторан 55', 4.5, 62);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Ресторан 56', 4.4, 130);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3976380b-82ac-4922-8d7f-367a0f3590ce', 'Ресторан 57', 3.9, 76);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('17e1cd17-72fe-4191-b899-d500450eae0b', 'Ресторан 58', 4.5, 90);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Ресторан 59', 4.9, 73);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('389b3427-00b0-44c1-96ec-9e778c1f2669', 'Ресторан 60', 3.6, 51);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Ресторан 61', 4.0, 64);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Ресторан 62', 4.9, 133);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Ресторан 63', 4.2, 131);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Ресторан 64', 4.4, 148);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Ресторан 65', 5.0, 120);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Ресторан 66', 3.2, 120);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('83a9a29f-bc9f-425f-b455-62ed269275c8', 'Ресторан 67', 3.6, 78);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Ресторан 68', 3.9, 62);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Ресторан 69', 3.8, 60);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('fce360d8-611f-41f3-aa78-5184271cb3c4', 'Ресторан 70', 4.0, 89);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Ресторан 71', 3.4, 134);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Ресторан 72', 3.1, 145);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Ресторан 73', 3.5, 136);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Ресторан 74', 3.1, 123);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('73b97018-6893-486f-9eea-a20cd02cc92a', 'Ресторан 75', 3.9, 62);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Ресторан 76', 4.1, 147);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Ресторан 77', 4.2, 124);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Ресторан 78', 3.2, 127);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Ресторан 79', 3.1, 127);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('44f256ba-397d-4a93-98ee-e2b652168673', 'Ресторан 80', 3.5, 89);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1f760e04-7c4d-4670-b546-d30c84d9602a', 'Ресторан 81', 3.5, 79);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('42ac54b2-b31d-46d6-b8bc-49e630623822', 'Ресторан 82', 4.5, 104);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('fda1df37-8d3b-4797-9f47-70e644d917c3', 'Ресторан 83', 4.0, 150);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Ресторан 84', 3.1, 106);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Ресторан 85', 4.4, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Ресторан 86', 4.8, 141);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('29e3b299-1d7b-4bd2-8227-35b749973a47', 'Ресторан 87', 3.2, 149);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('91abbb07-b935-4764-a204-bcedde60d172', 'Ресторан 88', 3.5, 69);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d680129a-8752-4aab-8b71-4bc45ffd27eb', 'Ресторан 89', 4.2, 58);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Ресторан 90', 3.1, 62);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Ресторан 91', 3.9, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Ресторан 92', 4.9, 147);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Ресторан 93', 4.2, 149);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('340a81a5-3476-4bc6-8300-57f988ec6929', 'Ресторан 94', 3.6, 117);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0216534b-fd51-48e9-8aa2-8f0530585f76', 'Ресторан 95', 3.8, 59);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('2fa219b3-1160-4798-b528-d8b080727c27', 'Ресторан 96', 3.9, 75);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('bab94b23-76e8-46a5-9949-d17348eef355', 'Ресторан 97', 4.8, 92);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Ресторан 98', 4.7, 61);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d34e760b-0640-4438-b1e4-80c4ae25eeca', 'Ресторан 99', 3.3, 74);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Ресторан 100', 3.6, 111);
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0599533b-074f-4b0d-9ecc-01bba477cf2d', '32a78cbb-14df-4f18-9df3-609a24ae0725');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('90d24345-56a3-4852-b5fa-793294e8458f', 'f2f3a4cf-3c39-4701-b019-17b81871fe3e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6c0b9543-4d60-4097-96b8-c8082bca7835', '9bb1ac2c-10f6-4137-a6f1-32a9132fe144');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e42a7054-3445-4e1d-a5db-63987ccac19a', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6b5f5d61-bf01-4224-9c17-44351a234245', '9fe68db1-832f-45c3-be4f-5354a26e20ca');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'efd0023e-6999-46e7-a6b2-019e8b277b7d');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0360b3d0-1fa1-47ba-b43f-4072688c006d', 'f2a8f08b-6ecf-4269-aa24-c18b8c501db8');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('2a37f4c9-ead4-4f29-9608-29f784b8225b', 'aa7a3d7a-982a-494f-88d6-a934588cff44');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('ff1a4ece-84f5-4b57-9cae-005f972ca622', '9fe68db1-832f-45c3-be4f-5354a26e20ca');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', '50f1ce38-fd6a-4edb-b47c-44373fec7d43');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3ba4133c-7c68-4a38-af8f-1c217107b148', '7ab99ec9-d2aa-4907-9423-cd19247a60fe');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', '1aa85122-ef44-408a-9a80-4118accf2577');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('434798a4-d545-4b7d-9a15-e2428c391f0d', '50a85385-4cce-4935-bb37-678d05dc79c6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c368e238-5f15-49b4-91b3-2e5305117e32', '44720e92-368c-4fbd-9b5c-c9ce3ff00563');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('913fa5e1-8dd4-442e-ae33-58b419e18a03', '8431640f-6132-48a2-99b2-af81bf6f5ab2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1bd888d9-74be-40c7-85dd-c3d531887f25', '471e9332-aa8c-4e01-b7ea-645f47a6a5e2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'aa7a3d7a-982a-494f-88d6-a934588cff44');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0ba88f41-1a8c-4406-b47f-0588ea36350b', '9fe68db1-832f-45c3-be4f-5354a26e20ca');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('bb37cf73-49ce-4ec2-8eb1-892815aadb23', '471e9332-aa8c-4e01-b7ea-645f47a6a5e2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('45fe1526-4609-4d4f-b8c7-07e64d2a6274', '45fdff4d-7ff4-4927-90a5-7352003b6782');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('45b5f477-e7ce-4492-a9c4-d1d25755bf01', '1aa85122-ef44-408a-9a80-4118accf2577');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d3096ed6-5146-4505-89ad-bb7dc8a14f93', '8b1b44e7-6b73-4033-bc1e-21cab30ecec5');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('4019d9c1-b6ca-4377-a28e-0bb1b89930c0', '32a78cbb-14df-4f18-9df3-609a24ae0725');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7114c561-173b-4efa-957f-cd58298d6f7c', 'c625a00b-af8e-47a9-ac89-f8125633021d');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('573e88a1-6ec7-44e7-a654-3ff0223cdcee', '8431640f-6132-48a2-99b2-af81bf6f5ab2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'd64f5818-ad51-4e49-b1db-4b2cfdcdfff0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', '7ab99ec9-d2aa-4907-9423-cd19247a60fe');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('27ab3cc4-db11-402f-900b-2efcbb393813', '6eceffc9-4193-4d73-9c0e-e3d8c713d1db');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'd98129e9-ea4d-45e6-b672-636b019f7e67');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6df647a5-1d1e-4ee0-bd0d-cd09889946a3', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'dbd87a66-b43d-4425-a250-2984a359c487');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('981c1d90-4dd1-44cb-acd3-4089dd55d351', '50f1ce38-fd6a-4edb-b47c-44373fec7d43');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'c625a00b-af8e-47a9-ac89-f8125633021d');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b839bc6b-520f-4b38-b9ae-b0074fbc849e', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d4c4258d-e1a2-40b7-bd3a-9124721f33f1', '471e9332-aa8c-4e01-b7ea-645f47a6a5e2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1fa15cf7-704b-4665-a926-e9560fa95141', '8431640f-6132-48a2-99b2-af81bf6f5ab2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c08349e7-d213-481c-9198-13c8c098c0f4', '9bb1ac2c-10f6-4137-a6f1-32a9132fe144');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('03ed857f-3fa8-41d0-a571-0a1ac3797a2d', '32a78cbb-14df-4f18-9df3-609a24ae0725');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('72467f7b-ae19-43c2-ab94-fea2cc5f321a', '471e9332-aa8c-4e01-b7ea-645f47a6a5e2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('96965548-5a6e-41e4-a4f1-2484b512d119', '50f1ce38-fd6a-4edb-b47c-44373fec7d43');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('67254fce-e401-46a3-8a83-ae8af0833f52', '9bb1ac2c-10f6-4137-a6f1-32a9132fe144');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('265af5aa-e609-4acb-aa2c-c63a26101fa7', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a605f842-ae90-4fd5-88c6-f3cfc987540c', 'd64f5818-ad51-4e49-b1db-4b2cfdcdfff0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('37b3a0ff-cb03-402f-b669-b2ccc9be85ab', '7ab99ec9-d2aa-4907-9423-cd19247a60fe');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('25d23d81-d8ab-42f9-ba45-2aa7423ce599', '2eef3843-d609-4390-8169-9d51a0c358cd');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('ed75802c-33e4-49d9-8853-97504a4a51ff', 'be6c0418-ab98-4f0d-bfd1-bd2acd67e5a0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e8eaaf1c-322d-4477-9dd6-789cbae71489', '1aa85122-ef44-408a-9a80-4118accf2577');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('41879f55-4f34-405c-87b0-a446bcc52406', 'aa7a3d7a-982a-494f-88d6-a934588cff44');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', '7cf2c1ba-9f41-4c3c-8bcc-69a278704332');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('5a0323f7-571f-457b-8ccd-bf89b0e2984d', '45fdff4d-7ff4-4927-90a5-7352003b6782');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('868fae50-e4a0-4ca6-8311-341fb949a47c', '8b1b44e7-6b73-4033-bc1e-21cab30ecec5');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('816d6f74-aaee-4baf-a83a-686f7b352bdf', '1aa85122-ef44-408a-9a80-4118accf2577');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'be6c0418-ab98-4f0d-bfd1-bd2acd67e5a0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('f56d68e6-9497-487d-b609-181aefa4d3d1', '9fe68db1-832f-45c3-be4f-5354a26e20ca');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'dbd87a66-b43d-4425-a250-2984a359c487');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3976380b-82ac-4922-8d7f-367a0f3590ce', 'efd0023e-6999-46e7-a6b2-019e8b277b7d');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('17e1cd17-72fe-4191-b899-d500450eae0b', '50f1ce38-fd6a-4edb-b47c-44373fec7d43');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', '471e9332-aa8c-4e01-b7ea-645f47a6a5e2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('389b3427-00b0-44c1-96ec-9e778c1f2669', 'aa7a3d7a-982a-494f-88d6-a934588cff44');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('05e802d5-ba9b-4a08-9f19-3ca608a416ff', '46df10d3-9ae5-4344-957d-6c02f7d9661e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'd98129e9-ea4d-45e6-b672-636b019f7e67');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('becc93c8-efa3-4c8b-991e-dd9383eb0e8f', '6eceffc9-4193-4d73-9c0e-e3d8c713d1db');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', '50f1ce38-fd6a-4edb-b47c-44373fec7d43');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d4df3541-ec2c-45f7-95f1-6e2a292fc611', '8b1b44e7-6b73-4033-bc1e-21cab30ecec5');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'c625a00b-af8e-47a9-ac89-f8125633021d');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('83a9a29f-bc9f-425f-b455-62ed269275c8', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0b794ca7-d220-406b-ab9d-4abc7b78761d', '8431640f-6132-48a2-99b2-af81bf6f5ab2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'd98129e9-ea4d-45e6-b672-636b019f7e67');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('fce360d8-611f-41f3-aa78-5184271cb3c4', '9bb1ac2c-10f6-4137-a6f1-32a9132fe144');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'dbd87a66-b43d-4425-a250-2984a359c487');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d09b1f01-9fe4-4444-93ba-7430d9e3e492', '46df10d3-9ae5-4344-957d-6c02f7d9661e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'dbd87a66-b43d-4425-a250-2984a359c487');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d2ccf989-6566-482a-8946-78a9bc9ea6e6', '50f1ce38-fd6a-4edb-b47c-44373fec7d43');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('73b97018-6893-486f-9eea-a20cd02cc92a', '9fe68db1-832f-45c3-be4f-5354a26e20ca');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e4516f8b-0c7f-4cf1-bffe-430be402dee5', '7ab339c3-b516-4839-968a-a9c0772722ed');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d4cf9bed-2af5-4bc1-9c88-32973c833b4d', '6eceffc9-4193-4d73-9c0e-e3d8c713d1db');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('009c0df1-e9a8-4be9-8781-78c98decb8f8', '46df10d3-9ae5-4344-957d-6c02f7d9661e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3b22c3db-89a2-403e-b641-19f9f58a70b5', '1aa85122-ef44-408a-9a80-4118accf2577');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('44f256ba-397d-4a93-98ee-e2b652168673', '44720e92-368c-4fbd-9b5c-c9ce3ff00563');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1f760e04-7c4d-4670-b546-d30c84d9602a', 'aa7a3d7a-982a-494f-88d6-a934588cff44');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('42ac54b2-b31d-46d6-b8bc-49e630623822', '2eef3843-d609-4390-8169-9d51a0c358cd');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('fda1df37-8d3b-4797-9f47-70e644d917c3', '7ab99ec9-d2aa-4907-9423-cd19247a60fe');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'd64f5818-ad51-4e49-b1db-4b2cfdcdfff0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0d9117c6-390b-4f4c-9823-9a5753394f2f', '45fdff4d-7ff4-4927-90a5-7352003b6782');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7bc2eec1-2cc9-4c90-a949-7cc05f294549', '03875339-7720-4247-a383-75a57f5d126e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('29e3b299-1d7b-4bd2-8227-35b749973a47', '2eef3843-d609-4390-8169-9d51a0c358cd');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('91abbb07-b935-4764-a204-bcedde60d172', '03875339-7720-4247-a383-75a57f5d126e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d680129a-8752-4aab-8b71-4bc45ffd27eb', 'd98129e9-ea4d-45e6-b672-636b019f7e67');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'aa7a3d7a-982a-494f-88d6-a934588cff44');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d79c7aa2-db06-42bc-ae95-370f0f1f6490', '7ab99ec9-d2aa-4907-9423-cd19247a60fe');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('db6c3a81-3b78-4709-9c79-0b56ddf44910', '03875339-7720-4247-a383-75a57f5d126e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'efd0023e-6999-46e7-a6b2-019e8b277b7d');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('340a81a5-3476-4bc6-8300-57f988ec6929', '45fdff4d-7ff4-4927-90a5-7352003b6782');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0216534b-fd51-48e9-8aa2-8f0530585f76', '45fdff4d-7ff4-4927-90a5-7352003b6782');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('2fa219b3-1160-4798-b528-d8b080727c27', '46df10d3-9ae5-4344-957d-6c02f7d9661e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('bab94b23-76e8-46a5-9949-d17348eef355', 'f2f3a4cf-3c39-4701-b019-17b81871fe3e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'f2a8f08b-6ecf-4269-aa24-c18b8c501db8');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d34e760b-0640-4438-b1e4-80c4ae25eeca', '7ab99ec9-d2aa-4907-9423-cd19247a60fe');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', '6eceffc9-4193-4d73-9c0e-e3d8c713d1db');
INSERT INTO addresses (id, address, user_id) VALUES ('075cdbed-d5b6-46f0-9c89-a0e365853999', 'Улица 73, дом 10', '4c40d5f6-51ce-480f-a566-9fd4f381c8e1');
INSERT INTO addresses (id, address, user_id) VALUES ('c3013153-c323-456f-a8df-e0de85a740f7', 'Улица 7, дом 44', '3ccb6b26-2f63-446a-9042-39379fe46b09');
INSERT INTO addresses (id, address, user_id) VALUES ('f3134222-1646-4b38-af12-3406f87e9eb6', 'Улица 75, дом 11', 'a24a0e09-2a12-49a0-807a-22a39e3fb8a7');
INSERT INTO addresses (id, address, user_id) VALUES ('148c2546-4b01-4721-9eb2-553681d3ea55', 'Улица 60, дом 6', 'eb7f1f8c-0dc8-4bdf-abe9-5f1ef2a8627e');
INSERT INTO addresses (id, address, user_id) VALUES ('e36fe915-3947-495a-868c-e2901c11f8e1', 'Улица 72, дом 1', '482a75be-e6b5-4f0e-b56d-01602a1c6b36');
INSERT INTO addresses (id, address, user_id) VALUES ('5c75c10b-17a9-4e75-82f2-81198bdd5f72', 'Улица 69, дом 11', 'dbc82af9-495f-4690-a314-adeb216affa2');
INSERT INTO addresses (id, address, user_id) VALUES ('cceefa27-c119-42bf-b01b-7f8050dfa030', 'Улица 44, дом 18', '8421cf5b-a672-4070-938f-2c9f7c24c81b');
INSERT INTO addresses (id, address, user_id) VALUES ('d142a220-159c-4adf-8eb5-52c3863f5d9e', 'Улица 77, дом 37', '360d8748-ef26-468c-916f-0f5f68558f2e');
INSERT INTO addresses (id, address, user_id) VALUES ('5f76fa1f-17f3-4cdd-9c84-b0100ecbad78', 'Улица 76, дом 8', '23c7701e-f25a-4ad6-9b51-33bcc8830c05');
INSERT INTO addresses (id, address, user_id) VALUES ('57e8aa9d-7bd3-4ff9-9016-a59e3abd2883', 'Улица 31, дом 22', '8c4aa27f-3a80-4b49-8113-7991dba5603e');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33d683c4-42bc-4f21-a61c-843951f5b312', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 1', 150.80, 322, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cda5df93-b62c-4d1a-8a1f-780cb21484b4', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 2', 274.80, 280, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c68ce583-bc00-4929-83f4-8901fdc2aa58', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 3', 597.49, 303, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('08430780-f666-49e2-af4d-f615b339b258', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 4', 504.10, 443, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9414d71d-46d6-44e3-b269-677deed9eaf9', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 5', 395.08, 398, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b09add0c-a75d-4b90-a8d2-d54571a1d48c', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 6', 403.45, 153, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a05f3c8-a43c-4a7f-9c39-35a98df40d9b', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 7', 198.98, 427, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b1d14f16-3746-45f6-a187-7e2eb65bf519', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 8', 292.78, 416, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('918ea3a0-08cb-4d8e-b8da-53a51b7bcd04', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 9', 495.45, 356, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('290a631b-8036-4267-8a65-af04d253e36b', '0599533b-074f-4b0d-9ecc-01bba477cf2d', 'Блюдо 10', 570.22, 181, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5d1e4f96-6773-4471-930a-d056494d5d8a', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 1', 421.27, 465, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('283b9317-74ac-4403-831e-00c91382dbe7', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 2', 489.87, 422, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('891b6840-321a-473c-8d6a-9d7343e698e1', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 3', 523.60, 368, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d94b596f-0a3a-4bc4-8223-b904fa8b2ca2', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 4', 313.72, 486, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('009c3a57-b8b9-4111-a0e7-7b033ce1a4f7', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 5', 572.15, 311, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c2ba8384-c801-463b-9429-3c072f28b726', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 6', 398.07, 415, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bce39e55-d2ea-4c03-a44b-aa01fa23e246', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 7', 535.34, 290, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ed900eab-8ded-4892-8631-453f29fc6c5f', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 8', 282.03, 267, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9dbea05e-5b09-4ee6-91e3-28e308875f1d', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 9', 187.10, 137, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('77b1e4e3-b7f8-4d0b-bb73-55d51ac69500', '90d24345-56a3-4852-b5fa-793294e8458f', 'Блюдо 10', 163.11, 435, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2fcbd3cf-7e8a-4bcb-aa9b-5c5831d7b118', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 1', 334.78, 468, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('668291a2-a9fe-4b25-8806-82d305829fe5', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 2', 263.12, 197, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('246cfa50-067d-4a6b-9dad-a04f310771bc', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 3', 569.40, 436, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2b2e6305-b032-4246-9c64-21c0a9779daf', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 4', 116.54, 448, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a4e8cf8e-b84a-4b7e-b5ef-1329d8f946e0', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 5', 397.68, 405, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('be34789e-e663-48bf-b03e-b09ed4618ed8', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 6', 423.09, 178, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2e39a678-2562-44cd-be9e-b760514286bc', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 7', 388.09, 322, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('32bdc01a-b35b-4863-a6ca-02addebb52e4', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 8', 547.31, 305, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('28c54ab5-5b62-4d47-8df5-eeaeacd04d1c', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 9', 585.58, 359, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8a42b9ef-e8a1-4fb3-b142-d5b790e61051', '6c0b9543-4d60-4097-96b8-c8082bca7835', 'Блюдо 10', 571.51, 321, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c186a6ff-950d-460b-b010-9f2e9b262485', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 1', 544.35, 472, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('be691b66-77ca-4c27-8845-6a15c4de9b2d', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 2', 108.86, 140, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5e1b95d4-93a4-4492-a699-6389a4dd8d46', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 3', 485.48, 482, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dbbbc9e8-a23e-4786-b16d-77e4eea131ff', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 4', 426.63, 308, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ea1bf115-8401-45ee-9d41-08507ff2d942', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 5', 279.68, 437, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a6e842b2-5014-4972-8785-5bf2c4eaa79f', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 6', 316.96, 101, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9b6780fa-7458-4431-b57d-39fc4d7484bc', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 7', 530.54, 216, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('56e49259-c530-4e36-9dfe-4c1e2aa44766', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 8', 593.01, 280, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b07b1a1d-06b2-4eec-b1e9-f5b13c3d1efd', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 9', 421.69, 191, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8609ac2d-ac1e-4143-a211-59e8b4e39484', 'e42a7054-3445-4e1d-a5db-63987ccac19a', 'Блюдо 10', 365.62, 157, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f9b4d5f6-d8f1-4b1a-ab90-f0f0bc5c94bd', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 1', 293.01, 151, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('794c0ce6-ca08-43b4-89a6-fb33094fd083', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 2', 122.44, 359, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eb61e77f-1b8c-40c7-8f06-b06290eceba7', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 3', 495.59, 311, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('42475ce0-ebc3-463b-ab19-ae75083e1fc3', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 4', 489.88, 401, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6558c0cb-ad6e-4cfa-b370-925c11faffca', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 5', 554.06, 138, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0972b97a-3bcf-485a-8727-3dd06e895e46', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 6', 464.15, 140, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3a62989e-cf04-4fae-9f74-ce4d9b57b3a0', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 7', 330.70, 212, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a5218040-83a4-4526-8e8b-842410c77efb', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 8', 328.91, 410, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b0e2f294-6170-4902-9797-9eb50ba301e5', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 9', 190.78, 157, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5ffada08-6bca-4892-bcea-ad37635b8697', '6b5f5d61-bf01-4224-9c17-44351a234245', 'Блюдо 10', 547.54, 112, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ecf58e05-1127-470f-a2ab-c160301254ac', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 1', 169.20, 284, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d434e622-72f4-469c-a550-587963115907', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 2', 269.57, 482, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('efcec780-5df3-4229-bfb4-8ccf9c3ab274', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 3', 441.98, 176, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6992eba7-7402-484a-b163-80a5095832e9', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 4', 103.86, 467, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('72338e53-a812-495a-9ee5-8b80ccc95a5c', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 5', 379.76, 470, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f5d6c5ec-0eb8-461e-841b-aa781b8a32d3', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 6', 595.56, 206, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('370cb1ce-3a21-4156-be7c-6b9a8f3e7f5c', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 7', 196.81, 224, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d1ab8f3c-2306-47d1-927a-0f36547f69c0', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 8', 220.09, 148, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('86141d5d-23eb-42d8-b87e-b5497e5a134f', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 9', 491.36, 391, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('62a81e1e-83d4-464d-b633-24012aec8f74', 'c0c98713-3f0d-4fb7-aad2-47fba6aa93bb', 'Блюдо 10', 115.46, 230, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('37a203e3-f987-4bcd-9e62-e57e84ef2c21', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 1', 207.71, 261, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc2ba81a-a62b-4f44-953d-e85116fff079', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 2', 438.16, 224, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6e945ed9-0ebc-4543-b620-a784e77e3cfa', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 3', 573.96, 383, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8a97375d-5843-413b-8e44-01ca6e0380b7', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 4', 310.84, 315, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('da4fa753-c845-431b-9cf4-3e39cdcfd2f4', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 5', 578.16, 193, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d5b74b6a-7020-4064-9d74-9559bf6ba402', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 6', 372.49, 208, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2bdfb2be-f5b8-41aa-b131-2d5fea930ba8', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 7', 587.48, 105, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('26aed16d-6e4b-45ad-a525-dcd4b8a87fea', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 8', 561.14, 479, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5cbffade-2fda-48c9-a797-43580088a2b4', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 9', 566.12, 312, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d7c33adb-4b71-4d34-832c-2e91b8084017', '0360b3d0-1fa1-47ba-b43f-4072688c006d', 'Блюдо 10', 272.57, 256, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2dc0dc63-5b3c-477d-95f8-d97b0baa3f1e', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 1', 599.80, 199, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('440f88ab-e5b3-41a9-a35a-3e0e6c32ced1', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 2', 434.07, 354, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a8dc6c28-cdbf-4650-b791-eb12432ecde2', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 3', 277.31, 403, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ffa63bde-3c52-4004-bff3-92a340bcc14b', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 4', 112.05, 370, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b0f0a49-0dba-43aa-b62d-803f9f4d4d96', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 5', 598.10, 412, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fb27a064-daad-41f4-b245-a55f1492bf88', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 6', 152.16, 174, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('53285a21-0bb1-4207-8dad-72581e81ecb8', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 7', 527.82, 214, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eee2025d-b231-42b6-9413-55f2224d6772', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 8', 362.06, 224, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73971ba7-d2e0-49a6-8fbd-5486c63acc36', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 9', 521.36, 339, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('504a4841-42a3-4e6e-ad3b-9f0bf96cd61b', '2a37f4c9-ead4-4f29-9608-29f784b8225b', 'Блюдо 10', 215.93, 327, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1914aa18-e252-4a36-8d81-007b2536f76c', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 1', 214.77, 279, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e153bca3-e66f-449d-8afd-0b44646dd254', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 2', 591.32, 337, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7df959ef-1b57-40c2-a4ac-31df4fac29a2', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 3', 295.29, 236, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9539e4f7-1be1-4416-90ca-58f4949f3228', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 4', 155.83, 320, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e37b0401-8f96-4db2-9b9a-11b2a241b58e', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 5', 290.19, 483, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('27d7633c-3e79-4a32-9898-0ba9531dfb35', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 6', 404.73, 329, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1cb3009b-4f2f-4c15-b9c4-0b36d305a223', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 7', 319.63, 395, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('611491fd-2f9c-427d-b880-2b0901a5cfb6', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 8', 239.35, 119, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('807e6d2c-de07-4a15-bae5-0030fd610f4a', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 9', 372.18, 310, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6b25369c-fef4-4f9b-9f4a-d6f2829d02a9', 'ff1a4ece-84f5-4b57-9cae-005f972ca622', 'Блюдо 10', 233.38, 353, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f00ef8e7-a512-4b73-8084-e46b663df7ed', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 1', 578.29, 303, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3fad3756-9ee1-4a7a-aa43-578da1172a42', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 2', 170.65, 286, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e4eec53f-cdb1-44e4-a4f0-ee5903618acd', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 3', 515.33, 319, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ef2927bc-9e76-4e64-8f9e-c4da6ced9ae3', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 4', 389.95, 239, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c482e39d-7ef5-4781-8d2c-21da94b352f3', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 5', 553.96, 281, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5859d20a-20ff-44e4-b0ec-1840264e495f', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 6', 421.22, 333, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('64dfa5b0-eb22-4b3d-868d-e7ec5101e06e', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 7', 427.54, 318, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8a17df99-a7ba-4383-8324-5a4317599671', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 8', 205.62, 262, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eaac8357-fc9a-4166-9462-778c22e8bb92', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 9', 452.03, 418, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a3936c3d-8c04-432a-926b-4df92b3c632a', '5fd146a8-1fb0-448c-9e7f-3fedb96ef84c', 'Блюдо 10', 171.28, 286, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2a1cd6ed-40c0-4bfc-b3d3-5ae6ab2ea0c8', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 1', 315.04, 158, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c3897fb7-4100-4957-b7d0-e489296c00be', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 2', 536.01, 340, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f5746106-f81d-45d6-b388-1e9bd7e73f4d', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 3', 543.86, 418, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b275311c-d760-481a-99e0-45e60934d9c5', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 4', 573.75, 325, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3a938691-4bd6-4e25-8efb-7e35de16d7d8', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 5', 397.72, 165, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6fdd291f-51ba-48ec-a87e-19cd85b62c93', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 6', 363.77, 374, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('11d092ed-bfc5-4844-9607-161ea27d41ef', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 7', 483.23, 443, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b7194cf8-a60a-4075-b0af-1f2ca92b0094', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 8', 281.23, 482, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e73b4515-ce0b-49b4-890b-7e05187f7756', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 9', 242.94, 326, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f1f5cdfa-bff7-458f-8a52-2743ad5e01ed', '3ba4133c-7c68-4a38-af8f-1c217107b148', 'Блюдо 10', 135.04, 113, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('17e3e5a0-6e8c-464f-9b14-e44effced5ea', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 1', 450.52, 111, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98747f2d-77ad-412f-ae80-e8956267718f', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 2', 293.05, 432, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5d9ca3be-b95f-495b-9ec7-279de0c82b81', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 3', 227.48, 219, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('658e6fd8-1ab6-42c7-a428-c431223d412b', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 4', 499.59, 303, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('94a4ada5-26c4-4a11-9776-fe32332a4ef8', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 5', 455.82, 131, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7325fce5-812f-4b62-9e09-00ff5d2db57b', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 6', 544.58, 433, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73f39520-f8a8-4ec8-a56b-f2a27937eb14', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 7', 117.24, 474, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('727aaa67-c70c-4690-a45b-e9986ccf277f', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 8', 394.01, 425, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dc477ccb-9a33-4d8c-8385-c29de8889ccf', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 9', 261.82, 389, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('908faaab-5bb3-4b6f-847b-643ced199b9f', 'e5eb7dd2-9a47-4e02-823a-adc2d08a39ff', 'Блюдо 10', 360.47, 145, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d9ec1950-03bd-4174-adc2-2447729c2fbf', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 1', 498.63, 453, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('75bdffbf-5605-4225-a05f-7e5f4d2139a6', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 2', 362.46, 298, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4eabeef6-c3e7-418d-89b4-76bb61aaf92d', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 3', 149.70, 149, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9103fbeb-f19c-45f4-a68f-b09f453314d1', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 4', 265.27, 324, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('488e2f01-8a8d-4d82-9fcc-2505fb8adc4a', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 5', 155.28, 487, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0542834e-8443-4044-81ea-4709f0f83a6c', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 6', 150.81, 222, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('80e7bf8b-f07f-44e4-a423-da3124d38b24', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 7', 250.38, 422, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9e8ea1c3-2570-4a31-84b1-ff18e834d4c6', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 8', 505.24, 289, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b698d894-6421-4926-80e0-23db0affd8dd', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 9', 525.77, 306, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ccd01bd0-1e6d-4fc4-b9f5-a7717337377e', '434798a4-d545-4b7d-9a15-e2428c391f0d', 'Блюдо 10', 535.08, 400, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('338c9acf-7049-453d-be78-8162067811ec', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 1', 249.96, 159, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7fe4662b-a3eb-41ac-b5f6-79cdd0b6fda6', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 2', 103.62, 352, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('efd22016-2c6e-4acb-8939-a1bb50ae1e86', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 3', 184.10, 484, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7774021d-4f41-425e-adf6-85da0fc36a08', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 4', 216.30, 260, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3070b72c-e221-4b5e-8811-0541c48286df', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 5', 188.71, 205, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bfa91443-d4c6-449c-9bd4-0285b3d2c163', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 6', 215.46, 280, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e860c8e8-f5d6-4e47-9be1-93006394151f', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 7', 536.52, 192, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4c6e7eee-c5c8-40af-8adf-155a18b65fa1', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 8', 126.56, 131, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('059870f6-f6da-4627-b5ef-402117075893', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 9', 318.16, 110, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('379eabd5-fca7-4f30-aaab-04c6143f0d3e', 'c368e238-5f15-49b4-91b3-2e5305117e32', 'Блюдо 10', 228.79, 381, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('089dbde1-c7e3-488b-a25a-afe27f038323', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 1', 188.24, 282, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8c3d3aea-1ba0-4fe4-9907-22ccb683cfdc', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 2', 378.58, 360, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fafd0f4d-040b-40e8-95bb-491f3bfe2595', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 3', 462.30, 240, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9b081bc2-cc35-49e1-971f-f7860448c78f', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 4', 478.91, 422, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('25f6025c-0e35-4b02-ae7a-acf0e3eebe98', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 5', 533.15, 315, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87b5a732-b340-4c77-8be6-97bd7a125dbd', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 6', 104.19, 196, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f547854c-170e-4d84-8191-f29f775cc092', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 7', 578.57, 164, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f735abba-4ca9-4a70-b068-c148c3d79daa', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 8', 256.23, 368, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3104108d-ca14-42d8-b279-3cc2fe463328', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 9', 102.79, 141, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5c1e3993-5fe3-4c14-99cb-026cbb6d3a7e', '913fa5e1-8dd4-442e-ae33-58b419e18a03', 'Блюдо 10', 128.33, 408, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('abcf0539-b5ee-44da-9fdf-f2a1017b8d0a', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 1', 383.80, 120, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a789fda8-b019-4099-a6a9-52075064fe67', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 2', 243.43, 105, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a236068-9a56-462a-bd7d-803009504702', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 3', 370.85, 492, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('df82d3a0-373c-4804-b256-d4481fdfadf1', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 4', 118.80, 298, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('38947ab8-c6ee-4411-acd4-399b9ad22245', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 5', 406.75, 329, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c359d3c-480a-452a-837f-ad3c92ce4fa4', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 6', 101.15, 227, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('be1de445-a472-458c-a732-0d387791cb6a', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 7', 307.10, 472, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('34651123-8fad-4045-a380-b65f0d403ba6', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 8', 378.35, 375, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7c6df4ec-5127-4389-8b66-737dc40e98c5', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 9', 397.76, 224, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('23e33079-e1fa-4d74-95b0-08af21a143d7', '1bd888d9-74be-40c7-85dd-c3d531887f25', 'Блюдо 10', 301.15, 140, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b10b29db-b265-45bd-9bd8-ef8417445e2d', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 1', 482.24, 122, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4b36910e-c162-439f-9e3b-9baa3eb6c386', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 2', 174.83, 121, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('26a742b8-1d82-421e-8f0d-d45624547fa6', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 3', 240.57, 260, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ac3d099-2cd9-4f23-8b98-650e37fc8a57', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 4', 282.47, 101, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('392217ca-1b3d-4eda-bf38-406513c31caa', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 5', 271.26, 116, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('710debcb-3560-4aa3-aa7f-d9f39de568ad', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 6', 432.97, 390, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0d289904-df91-4407-9638-3892b43b69db', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 7', 239.17, 222, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('25b14020-02a4-4a8c-8f1b-0ffc575abfa8', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 8', 399.47, 194, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('54b6c719-338d-42dc-bfb8-2672c5467273', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 9', 268.69, 177, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('df973b75-df1f-40cf-921e-e5e764fc9531', '3e11a7ed-1847-471e-8e90-7f76f0d272cf', 'Блюдо 10', 249.41, 208, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6cf3c68e-575c-49ca-9534-89167f69abc2', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 1', 436.33, 126, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('14da905d-2e88-4f70-bfb5-e1bca09596fe', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 2', 514.37, 460, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e98a9ffe-dbab-45ea-8f89-90dc5c9c63bc', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 3', 217.42, 445, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c7627ed2-abd0-49e9-8b32-d3155e1769bb', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 4', 438.55, 423, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c23ea2fb-6677-4f42-904c-97f2e679664a', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 5', 563.27, 221, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5ee9d4ae-67b5-47e6-9c45-094e51e9d402', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 6', 548.29, 416, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1721a25-434c-4496-8d98-3e6e207306fa', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 7', 210.89, 213, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('76aa4e00-77cd-401b-a90a-0217e26f351a', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 8', 225.06, 177, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('946a360e-95b4-4bb2-a3c7-9afd5116452f', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 9', 269.89, 122, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ef9fe7e0-1e7a-4e94-a3da-474a95395f9d', '0ba88f41-1a8c-4406-b47f-0588ea36350b', 'Блюдо 10', 126.33, 286, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9693d7df-3526-455b-ab7c-d8b40342173c', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 1', 109.57, 328, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('698fd34f-1506-4fbc-a092-75e4b61d1df0', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 2', 427.61, 405, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0fabc551-39bd-4515-81ab-094669b676e2', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 3', 586.36, 432, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8c09e832-e77b-4dea-be06-39e12d1444a8', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 4', 381.24, 203, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b890f237-3e97-4d18-b9aa-8fc6e3a6ddad', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 5', 417.86, 365, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('691c0bed-ef60-441c-8fa0-066b52a1fe22', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 6', 345.63, 457, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ee2a5600-3bd0-4d8e-baf1-ba5ed329c523', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 7', 259.97, 235, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a3c0602-3a46-4bef-b717-9943b6d9d2de', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 8', 425.77, 372, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1f3fab3d-5fbb-41df-96c8-91ad3745bfba', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 9', 321.56, 196, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1f80ccc-4415-4526-9b04-ccea5d5a6333', 'bb37cf73-49ce-4ec2-8eb1-892815aadb23', 'Блюдо 10', 356.84, 166, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('37b5b657-c9ff-4f24-b3a5-0809a0fdfd46', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 1', 252.05, 373, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b7f630e8-4d7c-4de8-b266-f1e8da284aae', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 2', 298.78, 240, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bc08b2ba-0b13-4f31-a498-a90f49a4ad78', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 3', 481.46, 479, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9749b7af-4ad5-4e4d-97c8-fed5dfb3ac63', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 4', 463.58, 309, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4dd7650f-2681-4458-98f0-03d22eb60856', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 5', 588.58, 426, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('edf92489-5db3-49b8-889a-d8ec58811305', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 6', 206.64, 143, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73ab5ed7-de08-4bf4-b9d0-299e4b5dfa3a', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 7', 389.74, 381, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2efcfe5b-9720-4af5-91de-fc04c3ebf4cd', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 8', 230.22, 156, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aad6cf90-4bad-46b3-87f6-bb62c5d4010d', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 9', 235.52, 206, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('24a3f4f9-f48b-4105-aa85-511198f504fe', '45fe1526-4609-4d4f-b8c7-07e64d2a6274', 'Блюдо 10', 109.42, 287, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f3dc1a6f-100e-429b-a3f0-1e0bb92bc882', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 1', 573.87, 243, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e114df55-c1c3-40b1-9607-6f7dbc47f683', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 2', 578.91, 321, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('745799db-8b7c-4f61-bcd8-55af5c10cee0', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 3', 170.84, 194, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bbcf2190-6900-4a08-aab3-0bca7272d40f', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 4', 209.10, 473, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2b30a9a4-3f38-43e5-bc98-629118333258', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 5', 214.30, 463, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b5dd0cd-0389-43b9-8f1f-6c5c44bd58db', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 6', 176.82, 365, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d076dcc1-0bda-4a5a-afb8-8c877b9f3a7e', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 7', 175.92, 470, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3d1d5a8a-b575-4d93-b288-366114331100', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 8', 570.81, 112, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8838e423-4044-4af3-857d-9050afacb1ae', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 9', 408.54, 326, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2cf227eb-948a-4d14-8149-5ecd17ba4be0', '45b5f477-e7ce-4492-a9c4-d1d25755bf01', 'Блюдо 10', 514.34, 116, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d3b7ab32-a2e9-482a-9416-c54f9ad5f9ac', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 1', 556.73, 198, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fbce8329-5e1b-470b-bf8b-7f142159fa58', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 2', 398.74, 422, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b514dd98-386d-4220-8161-ca9da3642aae', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 3', 395.53, 141, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f717397-4a0d-41a7-8693-d5c2a7c7ebe2', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 4', 592.96, 142, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('38c07693-73c5-48da-907b-ff10472e053f', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 5', 548.10, 390, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2b37b6b7-c461-4654-9fc6-0a615226d60d', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 6', 274.58, 445, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('12045ff8-ba1b-438c-90f5-4f9272009b01', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 7', 443.62, 445, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b0dcc8d-df86-46b1-9209-3147f5b1a0e2', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 8', 514.59, 133, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ab13bae5-8c12-4328-9e2a-68f0f6c78d65', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 9', 270.51, 337, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c65bcf77-b441-420c-8973-3ad4461b4c1c', 'd3096ed6-5146-4505-89ad-bb7dc8a14f93', 'Блюдо 10', 468.36, 443, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ed77a74-f8eb-4451-9778-8e104a135bc4', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 1', 152.65, 228, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5d682a7f-136a-45af-9fb7-3aa9ef325a53', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 2', 263.26, 177, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d8a3de18-3e1b-48db-bc71-7e421042d630', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 3', 289.02, 169, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eaccb70a-de09-4370-9ac6-2cd5553186c1', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 4', 583.23, 370, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2566e9d1-5000-4bdb-9820-125416bebb8a', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 5', 120.19, 347, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5604e8d0-2545-4a1b-a5ee-1a9a90e87e76', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 6', 228.75, 159, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ad68b28-0ca5-4d2d-810a-14167af0997b', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 7', 501.99, 188, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8097980c-9d0f-4114-bc4a-c2b1f1ca13cf', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 8', 189.79, 319, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('36061c40-7cd3-46c6-8b65-f93a46b4de98', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 9', 361.58, 339, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fe397cab-b208-4e2b-aa38-4217abb53600', '4019d9c1-b6ca-4377-a28e-0bb1b89930c0', 'Блюдо 10', 229.89, 287, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e078b337-a15b-4b15-8f92-6ae3010cbb43', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 1', 152.96, 149, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bac48420-36e9-4566-9818-dd2b48364598', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 2', 506.99, 128, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fca200af-5b1c-4f56-a5f2-16d0199b2738', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 3', 213.60, 108, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f5a0ee06-66ef-4dcb-9f28-d46d87d4baef', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 4', 235.22, 219, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('955638cf-5b99-4a90-923b-c82ab6f87e3d', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 5', 364.73, 139, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0d407b51-f8f5-4fe7-826f-edba0abd2527', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 6', 489.00, 467, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a897e1c6-b240-456d-a332-516fd11ddb18', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 7', 401.43, 277, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f506bcfd-9d42-4ec1-b982-a4d0d58bebef', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 8', 312.96, 487, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4434b491-2f38-45be-bf2d-a1c36ff67f78', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 9', 219.17, 452, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3761a91f-d234-4cce-a5d5-4dc48d8e5c91', '7114c561-173b-4efa-957f-cd58298d6f7c', 'Блюдо 10', 402.58, 352, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('264c7233-363e-4cfa-86ed-81b58087ada4', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 1', 450.63, 113, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2a2249ad-1c32-4767-b8f8-0b4bdffb7367', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 2', 211.17, 376, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('450ed262-be0d-444d-aebb-0404069a42ca', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 3', 160.79, 298, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4d607d61-063c-4437-b096-df183ce695fe', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 4', 352.48, 163, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a115488b-ea45-4b2e-a845-874b34f2e196', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 5', 468.97, 350, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a3cb84b3-0395-4889-b28f-c007bd4b097a', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 6', 124.89, 225, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('17d28fff-c7c8-49c0-87d7-b6617b0eccb5', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 7', 194.96, 439, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ad0596a1-3817-4379-afac-b48d2fb37442', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 8', 537.25, 159, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c2700962-e298-45aa-ba77-4dc70e5b37c3', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 9', 241.61, 115, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('99f28730-68c7-45c9-8d65-2459d953c5b9', '573e88a1-6ec7-44e7-a654-3ff0223cdcee', 'Блюдо 10', 344.30, 239, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('25989dfd-a1ef-4a58-962a-9a23e64c7d2d', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 1', 181.66, 200, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bc2dea1f-c13c-41c5-a90d-3ff0b1c0ab1b', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 2', 387.95, 167, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6b4ce8c4-6714-4f4f-9b5e-ed7175e4cc54', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 3', 518.99, 395, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('02450dc7-5f49-4fd9-a098-cc323bfdba64', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 4', 182.28, 236, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('990c1c3f-83d5-497c-b57e-db9b6aa93cea', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 5', 451.82, 418, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e9365185-fc6d-40c0-b49d-c87ab4463ff6', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 6', 556.10, 331, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddcb3889-9663-464e-b781-8c13314e9d12', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 7', 241.40, 272, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e7065541-67e6-4573-aeb3-06321f7c1ac7', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 8', 261.54, 487, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0bc79b92-56fa-4b99-b571-628125c8a37d', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 9', 100.13, 480, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5fc0092d-06f7-4f81-a5bc-da296eb24515', '7d0c43d0-8d84-499b-8fb4-b3ae88125e04', 'Блюдо 10', 249.43, 237, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('78874400-22d7-44f9-9d84-7348b7501a0c', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 1', 427.74, 191, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('81e8273f-fa92-4a7d-8f86-127c631b0552', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 2', 411.43, 262, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2a51dbcf-36bc-44ee-81b8-65b8108dbba3', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 3', 507.09, 129, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ac23468e-a339-4e4e-8350-e86311a9a656', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 4', 161.22, 461, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('40d9a4f9-2bd5-4dbe-bc62-83f4880cd77b', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 5', 416.93, 252, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8bb30cad-075c-4bcb-8690-aa61dc6c847c', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 6', 596.68, 367, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b203a2d-912c-4478-9d48-fefe8f1d792b', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 7', 158.59, 197, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e3d7a612-555f-4a56-8151-9c30f633a351', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 8', 360.80, 306, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b7cfea1b-8bbf-4119-b5a0-b26f70f1faaa', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 9', 109.22, 439, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5f129d08-cf9d-4975-bb66-486fbe628459', '5bc63529-c190-4a1a-b6ef-b28cfb8c2c2b', 'Блюдо 10', 106.49, 351, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3108f5ce-1519-4fad-9666-d3cba2be0ebe', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 1', 274.45, 499, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('26849f8b-17c3-453b-82a2-80ac9601ab3e', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 2', 488.44, 421, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31e5461d-9f0d-4f66-a1ec-d50f522e8aee', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 3', 583.84, 359, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('12de2029-f947-4f5c-b85a-24eb4f971893', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 4', 161.62, 492, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98bee1a6-0d1e-454a-8370-bf4099c47d44', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 5', 238.68, 318, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('841644d3-2c28-43e0-ac9c-8c7dc55c4c60', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 6', 365.21, 344, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('92e9444f-535a-40e8-b936-c14979907a1c', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 7', 585.15, 415, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('484e9233-a03e-4a33-879a-26d4c3e8a59d', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 8', 150.13, 290, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3174b1c4-6b00-4d68-9986-fda3bd0a44ae', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 9', 216.36, 175, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('35cfe23a-df17-45e6-a41d-6f992292db72', '27ab3cc4-db11-402f-900b-2efcbb393813', 'Блюдо 10', 486.75, 212, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('24c57394-7b67-46ee-98cd-4a3ef1f36faa', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 1', 167.14, 474, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bdc0adfd-5d8f-4f9a-892f-24862af4d722', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 2', 231.55, 441, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c5c46766-21c9-4268-b12e-1672c3263ef6', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 3', 256.29, 464, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('71165b4e-45a8-42c9-b455-673003e5525e', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 4', 451.39, 169, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('db2b51d8-c185-4a66-a42a-ec32e9ea72f7', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 5', 490.19, 474, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e874fccf-c572-4668-9917-93424807c1e1', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 6', 246.79, 475, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bd598eb5-2172-4577-8d93-7df27712e4cf', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 7', 118.22, 477, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('081bfd3b-25f9-4b21-9efc-3aecf6b68e80', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 8', 244.33, 456, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1346293-23c9-4552-98ef-c032c94f2a79', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 9', 245.02, 231, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e6505146-c241-4779-a96b-eeb818c5847b', '3fb8784b-88f2-4f99-93b4-5f07396d45d0', 'Блюдо 10', 550.98, 139, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d7964e73-70bd-49d4-b69c-3b2defc65ba9', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 1', 577.11, 306, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('61715aa4-6d1b-42e8-80df-70a5d0697155', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 2', 581.03, 304, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e1247ac9-4fb0-4748-94a9-95de0cdc9176', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 3', 219.17, 358, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9fa86e2b-97ca-4f3e-a500-5c02474cbd5b', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 4', 409.82, 233, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('64ad8edd-d1e9-4590-ab93-60e2cac4afb1', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 5', 496.42, 135, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d5375de6-a689-4492-9d89-95641a9131d0', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 6', 500.63, 172, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('36b44cbd-87d1-496d-9661-b46bfa76ff39', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 7', 251.13, 171, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('10c15deb-05f7-4c9c-9b2d-ec41a32381b0', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 8', 401.46, 437, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b73a07b0-4fa3-4db7-ad59-bda10f7e8954', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 9', 456.39, 389, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8ff45b31-4871-4279-b19a-db3046db6c41', '6df647a5-1d1e-4ee0-bd0d-cd09889946a3', 'Блюдо 10', 124.94, 117, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f4579f8f-0861-47d3-a007-42d5c47ccbe2', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 1', 182.03, 276, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('393194a4-8cc1-43a7-a39a-dbf908659c43', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 2', 382.77, 429, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e3521406-4ac1-4382-ba67-90ff2a92d506', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 3', 245.21, 278, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8695148f-d6f0-4ecc-a5c5-4586f8325f0f', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 4', 488.42, 221, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('454c10e8-c5c8-4ccd-bb01-adcaf9f50f8f', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 5', 132.78, 359, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('372c3584-fa0b-4f8f-981c-97cabec7461f', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 6', 194.64, 302, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7fa73ef0-8577-487a-9b6f-1f9c018a1cc6', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 7', 106.75, 336, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('67ade0f7-5996-40ac-82d3-e7bbb03cb232', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 8', 350.80, 242, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c1ae16d-3fdf-4d0c-a04a-7af3905047f4', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 9', 108.74, 265, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cbbd2171-e129-4609-8f8b-165174bd45ad', 'a37d138b-2f4d-4bc7-bffe-c6866acc386a', 'Блюдо 10', 465.43, 380, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('60106bf5-241f-4025-a773-c623737e1982', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 1', 295.84, 257, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c21e1de0-2041-4c7f-b1d4-0da34929c201', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 2', 254.54, 264, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9e49ed31-e4fb-4503-b03b-c4405c44dca3', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 3', 282.56, 183, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3aeb6719-28ad-42c0-90a4-0be305a4d1c0', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 4', 565.27, 144, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('52a3622a-1908-43e4-92d3-94f3bcd0e145', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 5', 249.30, 412, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8b1f20fc-e814-40d6-8bed-63cc53269b67', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 6', 114.49, 251, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fe4e45eb-0585-488b-8c6a-17140144b5ca', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 7', 382.36, 393, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('188f22f0-8d1c-4584-a5ec-e9d90bbecdf6', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 8', 216.47, 143, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f0ff6f1f-bc25-4d0d-a0de-1e61bc65b891', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 9', 227.65, 423, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ad915629-2f98-4b7d-a502-205846ef74ab', '981c1d90-4dd1-44cb-acd3-4089dd55d351', 'Блюдо 10', 198.54, 101, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f328ddf2-5da9-4c8e-b322-86149e5961b4', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 1', 111.84, 123, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ad313352-b654-47f5-a0ae-d0aba717ca82', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 2', 576.81, 183, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cc6924f0-1115-496d-8385-cd14ac58bd83', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 3', 326.47, 174, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('998f6c01-ae5f-4d22-94a2-75182f9ce54b', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 4', 496.62, 330, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8ba7ca70-f1bd-412d-b0b8-8dcaf1421271', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 5', 169.37, 319, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('40f4a8ff-96e1-4b1d-b83f-5c45719629f4', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 6', 193.26, 295, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('364a4677-f51b-4ad9-a76f-8fde777fa612', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 7', 192.40, 465, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0b461207-51cd-4a05-9cc1-a478156d9873', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 8', 532.53, 387, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('624ad6cb-25e8-4dfe-9e25-176bd99da57b', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 9', 182.96, 351, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('183e8f59-45e4-4228-9bee-ba08487d38a8', 'fee41da9-ea00-4a8c-9bc6-47c4da1afedf', 'Блюдо 10', 182.90, 133, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('908ca183-b94c-47c9-83d1-7ab706aabec0', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 1', 542.94, 443, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fdae0c77-1643-4707-ab5f-cb714770ef2a', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 2', 201.38, 330, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2410da2d-6a56-4e4c-b584-6bd5c56d15cb', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 3', 420.93, 425, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01e221c9-d923-4611-8685-ff8e0563ad8d', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 4', 397.05, 195, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('82d68dce-6642-4f40-a768-13be8535bbc0', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 5', 206.21, 213, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f0598b0c-13bf-445b-a5d7-9b8ec2f5e6b5', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 6', 217.95, 294, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b3ca9545-b053-4ac9-96c3-61c7e8407b33', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 7', 225.44, 278, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7862def0-ef69-497e-bc80-f88e4e436cea', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 8', 401.90, 466, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1fd49570-8919-4c6e-bc56-b2873c650641', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 9', 507.60, 481, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4252d9e2-1a8f-4ac8-b88b-6bd337aa3dae', 'b839bc6b-520f-4b38-b9ae-b0074fbc849e', 'Блюдо 10', 520.71, 223, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('299a4776-df0e-41e0-b340-ff9ec2693278', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 1', 151.15, 143, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6bc9bb99-502a-4d55-a536-55aeaa1db992', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 2', 410.65, 136, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('208f971b-93f9-4402-a6e6-706fdd729762', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 3', 451.43, 414, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('20c7c828-a29c-437f-ae04-8ed57616d8f2', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 4', 423.40, 488, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bd7ccce0-210a-44c3-940c-2e8188c4a388', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 5', 374.93, 453, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8e135e72-14a2-4592-ac4e-d97b34a95b9b', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 6', 540.42, 204, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3cf19c06-6641-47ef-b153-20326d7ff6be', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 7', 470.72, 143, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f7622573-e17e-4b0d-a030-d242c8576d53', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 8', 416.80, 282, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('af772857-cf3e-4649-8365-08d6973c6e8b', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 9', 553.81, 479, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ec2b5c71-1237-4e0a-ad21-359a9db3fa60', 'd4c4258d-e1a2-40b7-bd3a-9124721f33f1', 'Блюдо 10', 171.88, 146, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('06755097-b288-4988-8d05-0d2c9d55a904', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 1', 255.79, 288, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('826fe49a-0a2b-4d6d-a46b-0cb817cab53d', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 2', 300.39, 392, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7e964ca9-2a0d-4db8-83e0-e36e4525d214', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 3', 268.68, 115, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bd2bdf56-d9f4-483c-9f4a-dbd164063a27', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 4', 498.40, 438, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('25fe19cc-1d78-4341-a8dc-20e36bf44f0c', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 5', 447.62, 198, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fe14d24e-b60a-4011-b918-474eca9f7f59', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 6', 352.18, 237, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('59723ade-3db8-48c2-971f-8e78d8766de1', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 7', 249.78, 203, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d0bdf333-b064-48e4-baa8-1506036116bf', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 8', 309.88, 317, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('64119c98-8ea3-48f0-aa88-7149c225ccef', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 9', 223.85, 320, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2b3baf3a-d238-4c89-a4d1-fc5f9b3cb251', '1fa15cf7-704b-4665-a926-e9560fa95141', 'Блюдо 10', 477.64, 424, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4833978e-9df2-4c1a-bb45-a4e48d9536d6', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 1', 586.56, 296, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fef2cf39-217f-434f-88e3-575aae2472bc', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 2', 466.60, 259, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eb361a4d-ec10-4ac5-a21c-669788221f20', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 3', 320.61, 494, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2079fd12-e681-4ae4-8f88-6631ac7aa740', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 4', 296.28, 436, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('963a28df-a951-4104-8ce1-1a0bf1f43fb1', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 5', 389.55, 376, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c1dc43b4-ceaa-4a06-9d74-1d4a554d5b3a', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 6', 496.83, 135, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('92aeaf73-5c86-4462-9a33-9d1e2f25128b', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 7', 525.96, 431, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7585127f-abe4-49a1-a756-bb2f50aed05b', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 8', 475.11, 453, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0e324ddc-6152-4e80-aa73-3f5a40516b0a', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 9', 463.56, 150, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('49553c83-0937-4c07-a1eb-855e8124e5e9', 'c08349e7-d213-481c-9198-13c8c098c0f4', 'Блюдо 10', 557.22, 422, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2e43a474-c4c8-4c37-a725-68c3917e3f36', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 1', 408.29, 241, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d10d4592-ad9a-4b6c-9915-90d5f8c2a969', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 2', 441.56, 411, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6794fea8-86d0-43f6-ad30-ddb81a6ca2fe', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 3', 144.19, 489, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('48978147-b798-45b3-8e66-f3ae273a010c', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 4', 281.74, 420, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2a946fdc-84f2-4b5b-8281-366b58d3d575', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 5', 195.30, 181, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5089bb4d-567d-4292-898b-f8574e9e03b2', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 6', 434.84, 407, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('83f8cca9-8c17-46cd-87a0-334799cd8214', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 7', 566.34, 121, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e36318c1-9d7e-4674-82e8-22e799e65884', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 8', 102.73, 446, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('76ccf321-b1e3-4692-9750-8a74678db5d7', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 9', 347.35, 252, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cec21104-1d2d-4a1d-9047-9c9f01227248', '03ed857f-3fa8-41d0-a571-0a1ac3797a2d', 'Блюдо 10', 463.97, 442, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6ca4aa7d-a80b-4867-90e2-f01bba26b697', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 1', 417.08, 441, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fa6e9da8-8ba8-43bb-bdbe-e4de7bd1b7a8', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 2', 404.34, 103, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a95eddc-7833-4f79-a7dc-ba9e41711c61', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 3', 433.93, 165, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ac068856-3a2d-4513-b550-475836b2d2a5', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 4', 340.80, 301, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('51f9b435-6816-46d4-8f03-fa9e0c08830f', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 5', 435.25, 476, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('75efe384-823c-4862-95b7-68cdf57ed6a0', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 6', 278.91, 264, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('36fb9f42-0c10-4827-8758-a90aca56b9fb', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 7', 404.44, 237, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('413f5a8e-8273-40ad-8db3-014a3a5c3160', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 8', 269.83, 256, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2d5c7214-e002-4d89-bea1-0249f130c6bd', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 9', 492.16, 293, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c75944be-de3f-457a-bc0e-6cc26fe7f1d3', '72467f7b-ae19-43c2-ab94-fea2cc5f321a', 'Блюдо 10', 574.30, 444, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5189560a-9dc2-4d31-bb6d-0ea05b88ff64', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 1', 330.92, 229, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9d33e4d9-0d73-4ee2-8471-f6eb79cbec8e', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 2', 426.42, 434, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('69b4584f-6182-41af-aad2-2905f550a8ad', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 3', 172.74, 431, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33100cb8-cd3d-48d1-9bca-9233e828023f', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 4', 109.95, 223, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('068973d9-029a-4060-885c-38ce28aa92a4', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 5', 178.79, 474, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('24e40814-463d-4887-9e0a-5b9a88243cc6', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 6', 553.24, 415, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a446bb01-1735-40e2-8edd-9736ffea4154', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 7', 264.24, 221, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bd750b29-5c38-484f-9868-09ab48279918', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 8', 153.03, 339, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e91e423d-472e-4908-9b7f-1fdd8ad3c909', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 9', 277.14, 398, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4191768e-7a67-496f-9173-c70b5270f129', '96965548-5a6e-41e4-a4f1-2484b512d119', 'Блюдо 10', 148.85, 472, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('062fdd4c-abe8-44ac-a65b-f432591b9dd3', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 1', 483.64, 379, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4335db09-01db-4df9-bfd4-17199c9519fe', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 2', 442.14, 292, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d318623c-f958-4654-8942-773d1c826676', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 3', 241.95, 371, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('249f661c-6aea-42d2-82a7-3cbd798cd5ab', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 4', 219.60, 217, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('71617319-6931-44e8-801b-e2e24c2e0ae8', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 5', 450.07, 190, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6fa19287-9c6b-4242-86a2-4a05c0caae4d', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 6', 402.95, 404, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b7aceb1-d910-46d7-a33a-d59816bb2ba2', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 7', 497.98, 445, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9936b419-42df-4eee-9b11-123b0495d443', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 8', 516.37, 461, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98825311-5d8f-4e6c-a812-d2ceb6e4c225', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 9', 181.41, 106, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1ee28b32-7ac9-4b73-b977-9c6d1ca5e5c0', '67254fce-e401-46a3-8a83-ae8af0833f52', 'Блюдо 10', 523.32, 490, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7d2f050f-4a64-45ee-a60a-ea64328b5a4a', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 1', 561.75, 415, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c5d91c95-34b3-46b0-872a-196473f8af55', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 2', 496.33, 469, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('56bdfcf7-15fa-4d2f-8a49-591d080e1108', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 3', 543.98, 228, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('77621f00-680f-410b-930c-6db476db21bb', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 4', 375.12, 248, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4cdf962a-0ecd-47ae-b06c-5ff9ed46349d', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 5', 278.85, 261, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('43f8df50-2ccf-40ae-9c26-e8a17c97bcbd', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 6', 596.11, 260, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('657de8a1-0161-49d0-a620-2ebd7297cb3f', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 7', 145.47, 206, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d4ac6524-7a97-4d99-aade-b21b8d6c3618', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 8', 228.48, 440, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7cfd83f3-392e-41e9-ace8-1661ead4fd17', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 9', 152.80, 118, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('38baff3d-f30c-48e5-8db5-70483fc8bd14', '265af5aa-e609-4acb-aa2c-c63a26101fa7', 'Блюдо 10', 543.08, 493, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4d483e55-25da-4557-8845-cd846ec1c9d4', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 1', 351.22, 346, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4f360940-2828-4610-a7d1-a3ce9557b946', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 2', 401.43, 230, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3d4e5c0f-1cca-4b8c-b2d8-ad80dda6a0fe', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 3', 205.86, 302, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('06232c85-035f-4d82-a531-e6ad6b4c77de', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 4', 224.92, 385, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('24def64b-da62-4976-bc73-408ef6fa7cd2', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 5', 599.19, 496, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cb0c84e7-553c-4bd9-9ad7-f64014a7877f', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 6', 579.84, 225, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1af6380-ad70-4bad-81a5-9e23a23372be', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 7', 217.41, 112, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7c072c5a-abbc-496d-9f5e-36d7d6efe54d', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 8', 445.74, 100, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3c886360-53a8-470b-b9c8-1bd47e12e6aa', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 9', 308.21, 257, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('845389a0-d0fc-4701-962f-ecb604ec976c', 'a605f842-ae90-4fd5-88c6-f3cfc987540c', 'Блюдо 10', 147.03, 200, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1844112c-eb3a-4ede-883b-013591ebcc8a', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 1', 367.12, 318, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a179da82-5a27-4f68-b255-279fae450141', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 2', 444.36, 433, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d66d7f64-e810-4d47-b3ec-8f91b7a2400c', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 3', 278.85, 481, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1bcc013e-9823-4f6e-9796-f5f303ac6df2', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 4', 263.35, 180, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0cc61ef7-ca5d-48a0-98d2-854802d01fe3', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 5', 288.38, 408, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f14e21d6-9714-4221-87a6-4dd681c10db6', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 6', 216.17, 278, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8fa61d76-5c4c-4aaa-852b-e1d8f3213ecd', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 7', 183.59, 312, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a61043c-5e62-4b7f-863a-90cd1426059e', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 8', 256.06, 282, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0b44c6c9-a6ec-456f-9025-b7b0b06b7a1c', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 9', 469.78, 108, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('51d13034-5644-4570-85a3-f04e74513e61', '37b3a0ff-cb03-402f-b669-b2ccc9be85ab', 'Блюдо 10', 359.41, 105, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('03b470ab-3d1e-4b98-9690-fb4bce633965', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 1', 243.38, 486, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('326acf76-2181-4f04-ba3f-c349833b231f', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 2', 390.37, 195, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bf90b00e-6c7e-44a9-99a9-dbdda1c59bcf', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 3', 551.20, 234, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1eb1dd14-bc39-41dd-8cd9-f2ab5b1e17ad', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 4', 175.17, 474, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e970a206-7dbb-4fce-827e-1b3db9f2f160', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 5', 534.58, 160, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a5f0906-a2ab-4af1-a224-156bada46493', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 6', 326.62, 347, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5494572a-70d4-4ccf-ba77-7df0fd2fac23', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 7', 494.84, 253, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7d16d3cf-9463-4974-ac05-b69c2ab0f4e0', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 8', 223.84, 409, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a418eaf7-638d-4734-a172-6d932cc631af', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 9', 558.35, 128, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('36668e2e-5b69-4d3b-b3f9-e6e17470c5ef', '25d23d81-d8ab-42f9-ba45-2aa7423ce599', 'Блюдо 10', 483.23, 151, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('894f0a8b-724d-41a5-8e65-256af799014c', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 1', 510.93, 183, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2e53b530-1c15-42cc-8e9f-3fddce5ad659', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 2', 231.97, 350, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('30bd9a40-cce8-45b4-8e26-3a62df35bd79', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 3', 132.03, 222, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5cabe8a-6a39-4dbd-b7e9-3260733a73d8', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 4', 475.33, 250, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5ec37158-94b7-4b21-970a-dccdb4b6a6d9', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 5', 567.12, 412, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('048fb9c6-7f3a-4da4-b448-0f22562c1cd4', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 6', 147.97, 188, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f668087c-52df-48ea-983d-fe260e955cde', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 7', 121.65, 199, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c372dd44-d652-4d7e-8495-51a0f9d04af3', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 8', 307.71, 286, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f54610be-3efe-4714-8ddd-f6079bbe87f2', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 9', 562.60, 194, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5f53c0c8-9820-468f-8072-472b17220b44', 'ed75802c-33e4-49d9-8853-97504a4a51ff', 'Блюдо 10', 448.36, 498, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ed998db5-30be-49ed-b59a-da9488f6821c', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 1', 426.11, 462, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('adba8fab-94c4-4969-b2f1-c881c19fed8b', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 2', 396.76, 216, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c9039f6a-e726-4dda-90a9-e532daf79604', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 3', 269.47, 492, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('91386ddb-9881-4e56-b3c0-4dd7c97b74d1', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 4', 239.15, 348, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dc03d021-2c14-401a-8ad1-df7aafcc173c', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 5', 109.16, 230, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c1792ec-2646-4581-b9a6-b33bab2b4b7a', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 6', 389.49, 317, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9f7cd273-ffbd-4bd7-80e3-c5c5068c216f', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 7', 190.78, 315, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b0b5053d-c683-44b1-ae58-0d7f61ddfb37', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 8', 417.11, 429, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f66e7c6c-243c-4fef-8434-fb8df2423eee', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 9', 549.52, 396, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('af86689a-17de-4dda-94e5-2b01172691af', '4ce4d764-af69-469b-bcc5-0d5f8b08c7d6', 'Блюдо 10', 320.74, 365, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2ea541ae-cd49-4e21-ba4a-8f7674f89dd7', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 1', 503.48, 236, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b75e138d-7aba-4806-8618-d065183bbee3', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 2', 545.27, 312, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8b54f32f-a0be-4be3-adcc-a3b008522435', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 3', 411.25, 462, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('665ea417-4d67-4512-b520-c72cfaed2713', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 4', 330.19, 191, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1888c3da-01e1-44ad-a370-fe77faf2cf61', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 5', 310.39, 115, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddeed5d9-4320-47cd-b815-0bd0d7e27055', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 6', 510.71, 376, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dd43d7a5-a79f-4658-8777-e424b2cedc46', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 7', 526.20, 249, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dd51fe82-c431-424f-8cc4-01f5b1505beb', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 8', 572.01, 475, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('765ada1a-eeab-4e57-8736-96c55c45739c', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 9', 290.15, 153, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eb767266-fc0d-4cb2-8f16-1b6998686745', 'e8eaaf1c-322d-4477-9dd6-789cbae71489', 'Блюдо 10', 469.42, 183, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0be2c14c-80e1-41a9-bdd2-9ae2347b4168', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 1', 397.52, 267, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b9d3fd1-be43-4ddf-ab23-2b08aab33114', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 2', 442.69, 253, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('22a571ad-c515-4d94-98ea-ed5e87544f8b', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 3', 439.88, 490, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c02b4126-5818-483d-b49d-5024fd838515', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 4', 176.45, 249, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('50ed1093-785d-4519-b3bc-10f80de58bc9', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 5', 111.33, 139, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e2bf97fa-5f3f-42d6-a062-aae689901be9', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 6', 424.90, 326, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9f810c5a-44d2-438e-9218-90b68097d70b', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 7', 249.57, 318, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fa0aaa81-024c-48ac-8970-a31359dc52d9', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 8', 123.82, 466, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b3157095-d593-460e-b01f-03283642b099', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 9', 149.25, 125, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b2839f25-f531-4b89-b556-5a89ab93f108', '41879f55-4f34-405c-87b0-a446bcc52406', 'Блюдо 10', 420.17, 248, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e5b4a2a6-8472-4574-bc0c-31a0a4718528', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 1', 186.48, 429, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('76289796-03a5-4c7f-9742-9bdc06d7d246', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 2', 372.62, 107, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('90a85d32-a29b-4d39-a5e0-6c0a958ff5dd', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 3', 386.68, 356, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e52652a8-6cc2-4690-97b6-b66ae2dc1e9c', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 4', 140.48, 150, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9ec18027-c904-4fc1-9948-9814ae6b5a6a', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 5', 520.70, 258, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d03cb50e-733b-4d82-94be-8f1866a5e032', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 6', 428.00, 368, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4ea87ac6-d05c-4793-977d-19fbea3e7c04', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 7', 468.25, 447, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ad14cff6-9ea2-4d9d-bf69-199410040ec4', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 8', 150.95, 381, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('061c3c5a-554d-450a-a583-c6816e0d0f79', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 9', 555.36, 352, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ae246178-e041-49e2-aa55-ec62393465bb', '3f4ec0d8-df1a-4fd2-a163-b09a333f18e2', 'Блюдо 10', 131.16, 267, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e18f30da-1842-401f-91e4-5712560f7e0e', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 1', 343.60, 388, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aeeaf665-b3a3-4df8-b1c9-c7e64c9238e7', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 2', 150.37, 322, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('30c6ba37-2588-4499-a89d-8296987ee585', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 3', 280.53, 491, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fb90ffab-ac20-4004-a35d-1cdc0c0ec219', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 4', 222.52, 232, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a299d491-cfd2-491c-973e-66390ef29937', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 5', 119.33, 245, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e00bcfda-08a2-4cb0-ab44-d807b1b9a513', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 6', 563.58, 207, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('033f5813-c12c-4ca2-aba8-d87468b6ed36', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 7', 316.48, 336, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73a40324-67c8-481c-8323-ce0ae558e25a', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 8', 310.05, 479, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('20599b60-ed7e-4333-9155-7dd66773be14', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 9', 218.48, 296, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ed5494aa-4332-4942-b141-2b2a82ab18e5', '5a0323f7-571f-457b-8ccd-bf89b0e2984d', 'Блюдо 10', 588.75, 338, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('19f3b2ad-1022-400a-b136-5ccd221f18a2', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 1', 538.07, 289, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('161df49b-e349-4cef-91f3-1087061e9ed3', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 2', 468.70, 273, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('88903fe1-ca0d-4f43-987a-522fd2681209', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 3', 155.80, 444, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('53c831e9-1016-4d35-8fc4-737b9d059268', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 4', 348.20, 202, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dc12ec9a-6d9e-4888-8894-86acdb714c63', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 5', 177.17, 355, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cae76f83-e144-41b9-8e46-f7997ee0f55e', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 6', 267.30, 313, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('16338300-ad3d-4f5d-a66e-e2d0db72e953', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 7', 400.10, 487, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('55b78ba4-ad6d-4a1e-8cfe-51bcd961ebd3', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 8', 448.63, 201, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9d32bfdb-b2c0-4d42-b74d-55ee457aad9f', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 9', 347.81, 171, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('df2c444c-a798-466e-959c-e7e798113037', '868fae50-e4a0-4ca6-8311-341fb949a47c', 'Блюдо 10', 158.75, 100, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('55c55fd9-8a80-4bd1-b7fe-f92afe4edd0b', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 1', 258.20, 471, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8495e655-487c-4025-aa5a-cfa842a5c718', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 2', 404.60, 272, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f1d763a0-0c90-4baa-a0a9-12ea2f6ed6be', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 3', 527.22, 112, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('853e8080-603d-49cc-ba6e-4647637c7682', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 4', 224.61, 104, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a59b718c-8a31-4018-9f25-946aab33a1bc', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 5', 591.09, 336, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2cc17a01-cdf4-4a77-a333-d22aec0c7d0b', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 6', 263.15, 185, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a379469-e3c5-40d0-84a1-6e1b29c09d60', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 7', 464.23, 141, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3fb4b4c0-5fbf-4166-946d-5dc08857d14d', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 8', 524.76, 115, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('15f3b798-d345-4415-83fa-8b19531c34d6', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 9', 579.53, 419, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c06c4f19-2cc0-4f56-8d40-840cc89fb714', '816d6f74-aaee-4baf-a83a-686f7b352bdf', 'Блюдо 10', 146.64, 229, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ff9f56d9-90ff-460f-99ff-734b56d55ad5', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 1', 592.63, 383, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('acbdadd9-a6fd-4d09-9812-42bc22e54017', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 2', 443.97, 165, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d4eec9ad-7493-40e3-8bcc-3395315e8f9e', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 3', 547.37, 220, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('34f5159c-1438-47eb-94fa-1e98b818d11a', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 4', 327.01, 265, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9740892d-d4b6-464d-a7a6-b9c190a30b2e', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 5', 549.72, 252, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aba86ebf-3a0a-4b6a-9e28-fa573d9587be', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 6', 594.58, 299, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('703ba382-c9f2-41d8-91a0-ca6034ebe193', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 7', 376.41, 168, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('732ff558-b1bd-4245-a8c0-da86d8c9eeb5', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 8', 200.55, 380, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('72ebb77a-b248-42c4-bfb3-fe0977717ae6', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 9', 150.14, 451, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1a78ed05-2713-4f80-ab3b-80d249363872', 'bfe5850a-53c5-4a11-adf5-9c56b8085f24', 'Блюдо 10', 371.07, 193, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e877b8e0-2b0c-4804-adae-3194d91bd679', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 1', 144.51, 221, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('531a9488-3008-4b8a-890c-fdb76d808c81', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 2', 183.30, 390, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5c926534-32b2-49b5-af7a-3d135113b281', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 3', 450.16, 321, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8eae1a27-4a24-43bb-931e-0c42e2b20a33', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 4', 267.05, 229, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8cdb8ca3-995e-4c99-b2b0-4117bb06ca70', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 5', 300.12, 208, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b4421a77-4836-4496-876c-12158178e16c', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 6', 128.28, 134, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f4b9c17-feb7-482a-869b-78651b310327', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 7', 484.55, 419, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6e2519a7-24cb-422d-8d0d-325e571ffac4', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 8', 204.84, 238, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3e2a6a60-5602-42df-92a4-ff423e477992', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 9', 183.91, 272, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2bcc4e6e-f9b5-4842-a538-e292fab7849a', 'f56d68e6-9497-487d-b609-181aefa4d3d1', 'Блюдо 10', 379.06, 407, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ec544b09-0d28-4781-9a1d-23ea19b9d750', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 1', 162.58, 295, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87922c20-0a07-4bcf-adf5-7ff5259a42ee', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 2', 132.99, 489, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6538554c-0b44-4217-9451-a467f0c78e46', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 3', 337.79, 311, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('310ab7fa-e800-44c1-980f-29207ec75687', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 4', 268.28, 159, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c08d2196-da3e-4581-a1ca-24d77ebd6be9', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 5', 159.82, 425, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c5ddec97-3033-43b0-bd7c-174a21879b3e', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 6', 499.58, 310, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1e44ee0c-c8c8-4147-9dd0-5a6fc3911a2f', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 7', 313.47, 202, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8041e27c-8ecd-4c60-ab2e-ba479fff1116', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 8', 216.71, 413, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c45e228-5cfa-43dc-a91a-80557798b7a5', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 9', 293.29, 121, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dcc48e50-c7da-464e-b86e-729a23da9485', 'caf222c0-4f53-4fcd-bdbb-0c4a959eb876', 'Блюдо 10', 246.63, 351, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3c6c6d67-a398-4511-abfb-573a6cb8012d', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 1', 224.46, 135, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('93759c63-5393-46e4-903f-79bff4e75f09', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 2', 329.39, 200, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('255f5078-12f0-406e-8d10-670725022b93', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 3', 431.32, 380, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('03840695-f6c9-4f39-a9b6-c6fe16eaafd7', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 4', 294.36, 356, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('63155fd9-8e60-477a-9f20-3a46b36642b8', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 5', 395.82, 491, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eb9f50fe-79c6-4669-880d-c161c26c68fc', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 6', 547.44, 395, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cf704147-11c9-40ff-9b1f-a2f2ecb18506', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 7', 257.05, 258, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87be2fcb-ca50-45a3-b5a6-88ce06bf8c37', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 8', 216.79, 259, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('54317cf3-43da-4f42-8d52-aeff1ffea04e', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 9', 474.33, 255, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc67e7a0-a6b5-4ce2-b028-1bb31fd51aef', '3976380b-82ac-4922-8d7f-367a0f3590ce', 'Блюдо 10', 168.70, 363, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a897ef27-78cb-4a3c-a347-24184f0df6a6', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 1', 111.11, 345, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7949cbe9-b3f4-4f3b-9a4d-2ab14dabd26f', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 2', 561.43, 116, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('74d72aa8-e577-4cae-af07-61cfa2439ce2', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 3', 583.25, 169, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ebdff62-c731-490f-8cfb-b47760e98ecd', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 4', 365.07, 162, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('04fcd463-0c49-4dcb-8128-f2e6c113d6ef', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 5', 324.32, 329, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a59e123c-8222-482e-96c6-708f00982c84', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 6', 333.77, 236, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('145553f4-09a3-474d-a24b-172162dce6fa', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 7', 589.41, 113, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('10d48199-11a0-4124-940e-59cf2d92f59f', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 8', 493.10, 153, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('127b72db-5817-4b63-a61d-e47e0389e7aa', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 9', 144.53, 379, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a824b37e-6d63-4b89-a400-af54ec484fcb', '17e1cd17-72fe-4191-b899-d500450eae0b', 'Блюдо 10', 362.22, 124, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3cca002d-dd0f-422b-b181-e619d7c76345', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 1', 184.42, 215, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9f19aa0b-7a14-4e8a-aff4-345079ea3d74', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 2', 538.78, 456, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a3ae212-bd4b-4bbf-85ae-b8f0e5605e0b', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 3', 245.54, 481, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('27684df9-63f5-4910-9853-6759c627887c', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 4', 102.33, 400, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5df15e15-899b-468d-bb85-6c937e4b30b1', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 5', 463.05, 120, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bfe96147-34e5-4483-8efe-d2b5bd42d2a3', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 6', 326.77, 319, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8ed8a87c-7b48-44a0-af54-716ce249abb8', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 7', 331.47, 136, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('02af46c5-c724-4bfd-80c3-89d1b1b547bd', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 8', 211.14, 304, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('871b6bdc-6531-4b58-93a3-41fda6cfdef6', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 9', 397.10, 354, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8db7db11-2d47-4f4f-a94c-5c9f478549f3', '37dbe9a3-d3ea-4b96-a779-4493f5a66b2e', 'Блюдо 10', 490.70, 430, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2aded5b1-f033-40ca-bbf4-284fb9a24c53', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 1', 152.76, 382, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1846d361-e5cd-4c6a-86b4-1df2dc7ba0da', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 2', 147.57, 341, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d0993bfb-d756-4b5d-bd80-54362bf2df40', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 3', 591.89, 367, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d42fc1ae-ee44-4488-b6cc-0cda033e0b6c', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 4', 331.73, 227, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6806b213-adb9-4760-89aa-08ca8e6a9739', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 5', 534.13, 386, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f307aa5-1bef-4d9b-b219-0ebedda008b8', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 6', 209.74, 323, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('69f851d8-f916-448f-9012-f61ea3f2e2f5', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 7', 176.90, 343, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ee3d553e-ca51-4222-becb-7562dab6f5fe', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 8', 504.58, 227, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f6ae0da9-e32a-4bb4-81b7-adf505d90c05', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 9', 310.48, 493, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a13c476a-fea5-4e20-9ee9-5ed7e7873483', '389b3427-00b0-44c1-96ec-9e778c1f2669', 'Блюдо 10', 154.55, 147, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e4b8291c-310d-45f2-a9fa-c7fb7a0dbd8d', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 1', 472.52, 170, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7bc5b6f1-1175-4387-9d84-dda03756e8dd', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 2', 364.28, 387, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6871f939-e2a0-4173-9bbe-2902dd5513cb', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 3', 230.67, 277, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b2d7ef27-4628-415c-a63e-451b9f6f74f1', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 4', 162.65, 383, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6b2299a6-0845-45c5-9263-36974f9d034f', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 5', 209.27, 428, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0a20b215-e99f-43ae-914c-76fc354f4b55', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 6', 238.15, 227, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c816ad11-7e5e-42d6-ad9f-424394af8dc0', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 7', 540.09, 278, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('375176be-4c3a-4cc7-8be4-4df7dd4e2626', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 8', 217.79, 468, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('097d113d-c7b2-420b-9284-ea59b207cfac', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 9', 586.40, 343, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1a14bf8-f91c-4bd1-ad04-a1cd5d6fe43a', '05e802d5-ba9b-4a08-9f19-3ca608a416ff', 'Блюдо 10', 429.14, 294, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dfda441a-71a1-4a70-86b4-1e26fbedd748', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 1', 357.62, 389, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bbeae6a8-77f3-4a0e-ab20-fb9adb43fe75', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 2', 136.17, 344, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f3322b53-6ac1-4dc1-af08-0520b4bee2d4', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 3', 436.60, 140, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('37408803-36d6-4a55-8360-c7baacf31290', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 4', 498.73, 310, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2dafcfda-abd3-43bf-837f-6ccdf289b3c3', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 5', 440.69, 122, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1e2d485c-c026-4acc-b934-eb8f8891dd90', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 6', 330.11, 244, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('30fd6510-a8e3-4830-b291-cdbe84c0168c', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 7', 317.92, 299, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('687addd1-a18c-452c-896c-2b61328d7970', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 8', 440.40, 310, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5b584634-7437-4634-9e5e-71d36224c975', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 9', 453.80, 288, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2ad32105-0a0a-4599-8940-1cc00f259c64', '90b7f956-2c9d-4adf-9c07-f56cf6d5d272', 'Блюдо 10', 364.73, 374, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('57241f9a-8d0f-499a-bca2-d2b0508dc5e0', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 1', 161.08, 186, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bbeb5db7-3cc8-4cae-8c9a-8b49067d2441', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 2', 225.62, 228, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ed60d9b5-1a6d-4e62-98c5-6e3a6ebd488c', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 3', 571.72, 252, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('74c088a3-e77a-48c4-a597-0830d2d724b8', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 4', 162.15, 152, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('22b607e4-7cd0-4bf2-b1c1-bdc5e439ed9c', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 5', 309.84, 348, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8095a8d9-bd0f-4686-8b88-2505e124a248', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 6', 479.42, 191, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d9c1ff89-0209-40e1-84ce-241d83399a2d', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 7', 241.51, 388, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('230a6cab-0f2e-4907-b32d-62c46a64fac2', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 8', 331.77, 174, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73d3f253-c111-4c7b-979a-f6fe1b5953ef', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 9', 132.43, 267, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c4c9f57-0059-4e5d-90bd-fd30c5d2674c', 'becc93c8-efa3-4c8b-991e-dd9383eb0e8f', 'Блюдо 10', 498.10, 217, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('90f3bda4-3192-4fb3-8036-790d00554004', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 1', 264.58, 366, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('65ef8bfc-f681-45e4-ba0a-ce8e62a5d0a8', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 2', 505.37, 435, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e8726c22-9e05-4b56-b830-7f4d60a10f45', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 3', 317.84, 324, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('321b8694-0b62-428b-a977-db21ada81f6b', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 4', 529.23, 380, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ec049b4-5dd5-43fc-8eb8-9dbc4633aa92', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 5', 318.99, 390, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eabfacc1-d734-47b5-b97e-68849c6b02f8', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 6', 280.57, 276, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ac179de-616a-426c-8e0a-9483772a34ac', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 7', 451.27, 273, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('580a06dc-a173-4bfc-b648-8764b827fb9b', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 8', 343.64, 289, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7fd2bb7f-c72f-4f68-b69a-132a38ab6cc0', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 9', 255.28, 292, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0d09b448-91c5-448d-93b0-1f5124f25ff1', '3c5fc35d-4636-4d5c-863e-4d98ece5ab6d', 'Блюдо 10', 540.06, 228, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e5ac7e0a-e730-45b1-ab22-66e155e97fc8', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 1', 251.12, 356, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e8b50951-1e6d-4d70-8317-a9f3aa65aa1a', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 2', 285.91, 277, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33e64e5b-97d1-4bc0-b64f-22fc44158ac5', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 3', 394.80, 454, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8768f831-1d1b-466e-95fa-95807c2bc3e6', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 4', 161.94, 338, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('534c88be-6287-4786-a870-3febf3b97eed', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 5', 431.01, 294, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('613e4c62-0a46-4a24-b4f4-33bc94276d3a', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 6', 227.90, 365, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a976dc6-59a5-418a-adf4-63e38ec9c299', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 7', 218.62, 318, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('caea58e7-c914-4870-b7f0-6fa6ef22beba', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 8', 160.23, 193, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('34138fc1-3c14-43a1-b675-dd7dafd08236', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 9', 535.08, 464, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fec1214c-17ec-4251-8049-0278cdc8d266', 'd4df3541-ec2c-45f7-95f1-6e2a292fc611', 'Блюдо 10', 590.49, 256, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f5e9a518-ecbc-4615-b05e-08f8020d431e', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 1', 583.93, 164, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('44a6ddc7-f214-478f-bb5f-b29266a85c69', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 2', 512.55, 303, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5ca1cb2c-6885-46e2-8a70-bd8e613381ac', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 3', 305.83, 375, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f42310a-c190-4bf6-b322-8462bd297b3d', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 4', 305.91, 358, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9e50c0c1-9d61-4f27-982a-21d5a598b331', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 5', 282.56, 495, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6595cb8b-106e-47e2-92f0-e53005d6fbf0', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 6', 221.61, 226, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ca274f4-86e6-4c68-aabc-45129fbcf710', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 7', 499.50, 403, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b11b4442-9056-48ff-901b-98e501f746df', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 8', 243.45, 259, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('df6e3ff0-dab0-4f35-84d5-2ae255763d8e', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 9', 520.26, 276, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c9fc6e0e-9da5-428f-9a80-64cfe02cc3b1', '4d0eaea3-61f8-4ac2-bd3f-76263955f91c', 'Блюдо 10', 453.39, 309, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('04d95ce1-9cd1-4781-971b-013a0b317ceb', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 1', 409.34, 358, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('66d1db43-b4ea-4688-839c-5c0668afc0bf', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 2', 196.99, 350, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('14a9d1e5-473e-416c-a612-87dbb2b8dbd1', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 3', 561.60, 499, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5fe80752-90eb-4d6c-9339-0454f02253f7', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 4', 349.43, 407, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('431332d5-10f3-4738-8cbe-f4906b737276', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 5', 322.08, 346, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3ca01847-bfa7-4ecc-b619-2f274439f7f6', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 6', 507.72, 163, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9974797a-33ea-4e18-8181-9a5ab9fcd048', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 7', 292.74, 147, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('37646688-d10f-46aa-bd9e-8861d1929497', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 8', 280.01, 348, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('44770c2f-e273-4fe2-b0fc-53b47b37b1c9', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 9', 121.97, 442, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1a582fb6-7914-42af-a398-b3398a1c3657', '83a9a29f-bc9f-425f-b455-62ed269275c8', 'Блюдо 10', 300.33, 244, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('48dbaf80-8e70-4039-8ecd-6a19ea0f2bd0', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 1', 502.77, 103, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e0f21f8f-4e0d-4642-b302-f61f0d7001be', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 2', 548.04, 248, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cd1dc46a-7f4d-45c1-8bb9-b885f99719ee', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 3', 478.25, 117, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bef188ce-dbfe-48c2-b5ef-4cee1d6be20b', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 4', 589.87, 353, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7d858d1b-a51c-4a3a-b8a6-7e9e1d985db7', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 5', 550.22, 454, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5f75b93-0481-44e8-93c4-496de86189b3', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 6', 175.75, 399, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9d93466e-7695-4656-bf03-dc0de11044bb', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 7', 310.39, 460, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33620058-da7d-4d5f-aec6-408c04e116d0', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 8', 160.28, 440, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3816e3cd-6007-4e55-a8f3-77996775e843', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 9', 465.70, 276, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f4430c39-9419-40a5-b394-e8e7d99f663a', '0b794ca7-d220-406b-ab9d-4abc7b78761d', 'Блюдо 10', 413.75, 268, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2e381d52-47c1-4532-8425-15b41815862e', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 1', 334.42, 117, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c115a61f-29da-46e8-99b8-55ba027e493d', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 2', 244.31, 218, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0656ee0a-ae21-4835-87f5-b2e8c0888cdb', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 3', 487.41, 373, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('95af9970-c2b8-4d1a-a543-df51cd6ec137', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 4', 314.15, 266, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a03a16d2-6360-4e2d-9047-f53eb99980fa', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 5', 411.97, 458, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c3e584ef-d13b-436d-9a95-5b2323c56209', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 6', 400.18, 482, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('54062b8d-4095-44ef-b0e3-5322986d2230', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 7', 262.29, 262, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6f5491aa-ff22-4626-b700-230290561961', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 8', 233.59, 195, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5ad86bf-8695-4880-95f2-55b2fae009bd', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 9', 151.73, 425, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4949de1c-9752-41f2-b129-87072cca28c6', 'cfa0f5da-44e5-4c36-9852-e2e73ee4358b', 'Блюдо 10', 342.07, 337, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('841132a0-8189-4502-92d2-69e67f891ea5', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 1', 592.49, 166, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e6101666-692f-48dd-b281-1cd1e0a128cb', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 2', 494.28, 299, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c0eb89a0-d609-4af4-a98b-8151db5abeb3', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 3', 326.45, 166, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a7faaff-dc80-476f-b9dd-5e2342bd04e5', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 4', 356.79, 252, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('287e606a-018a-4412-a5f8-2f0723fcac7b', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 5', 561.33, 323, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c7337c7b-7236-436a-aeac-f6e3841ccd35', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 6', 561.25, 342, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a7b9f642-5623-488c-946d-45d3bfab2a15', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 7', 541.26, 204, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b18d7cb3-c98c-4302-ac23-8190c2681def', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 8', 395.86, 498, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fafd4f02-762b-4ced-b61c-db199ae50813', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 9', 103.58, 211, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c8148d17-79a5-451e-b0b0-1e635a65ec3e', 'fce360d8-611f-41f3-aa78-5184271cb3c4', 'Блюдо 10', 546.04, 213, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('59a53f95-64a9-46ce-b8f7-60aaaea97242', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 1', 479.68, 160, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('141df5cf-3a15-4fd9-9458-eed5f77b4325', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 2', 358.79, 408, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('723d1549-815b-4fcc-bf28-02df2a772055', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 3', 102.17, 101, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b19328a-cc7a-4a33-83bc-73624189efda', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 4', 290.66, 140, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1f368f69-bd73-4a39-af0a-ec2744209bf2', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 5', 464.49, 340, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fe73a844-ecef-4a1f-b5c1-fee8c48e309b', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 6', 279.14, 453, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9e99e529-2793-45f9-b3ea-005e15befc78', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 7', 551.68, 137, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ab3df713-93f4-4b4d-96fd-82d3d9bdfb4e', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 8', 167.73, 362, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('13ba28d8-7799-40f4-872f-eb72f27208e9', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 9', 580.67, 227, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('19400ea4-4d8e-4ab4-83a7-e0bb87ac0427', '1166a550-7a5d-4955-b7a1-23e3b9982e2e', 'Блюдо 10', 414.57, 301, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e56809a0-6350-4833-bb77-a6fff31ed0fc', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 1', 585.41, 277, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('579db788-6b92-44c5-bbd6-f4b62c0929b9', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 2', 310.83, 396, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5249c2c-d5c0-482c-bea9-9c1f422e3c01', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 3', 522.16, 363, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1b72cb3b-2694-4c9f-bb5c-38f0c2902fff', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 4', 509.87, 290, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6bece770-19f2-4ad1-a552-db5f730009e8', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 5', 415.93, 374, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('925d7f19-c014-4dd7-9a52-692c9bfa170d', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 6', 142.62, 259, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('363698d5-647d-4af7-b9fd-877e1d9a3a0f', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 7', 499.32, 400, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a4f2079-28e2-4b65-b4c1-b2958dca57cf', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 8', 325.88, 393, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('308501cc-2587-4571-bc38-6419efa72b0e', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 9', 385.76, 106, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c0e8c2a-da44-4fde-8af2-4d541c365371', 'd09b1f01-9fe4-4444-93ba-7430d9e3e492', 'Блюдо 10', 417.95, 386, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e8c9a34f-7088-4c53-8a05-2d4069857302', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 1', 167.28, 360, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b8785dc6-a1e9-41a9-88d4-03ace3b4b0b0', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 2', 594.70, 325, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c3935823-265b-4891-b809-cc54b2496662', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 3', 510.39, 280, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('81e41094-0fc0-4a84-b317-84d19b1e0569', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 4', 431.80, 408, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('65f5d0d8-dd1e-44d1-87d3-a96170520c73', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 5', 263.18, 454, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('91837e12-2516-4571-919a-e2cc3dc637ae', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 6', 449.66, 290, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3362cd20-91b0-45a2-b1cf-29806f641123', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 7', 177.13, 451, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddf16a06-0316-47ef-be16-35d02ed8aded', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 8', 434.72, 286, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8f6e7c1f-a55d-4eea-acd5-63bf741516e1', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 9', 487.69, 348, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('60f48724-470d-4270-82bc-cc55e407f551', 'eb9f7dc0-0246-4dbf-80bb-ee00d16758e3', 'Блюдо 10', 322.66, 363, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc189063-34a5-4bea-bdd1-56d4d2249843', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 1', 541.42, 475, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6e44eedf-9115-4025-8431-888a6bc35d60', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 2', 234.07, 476, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cbb8298f-7eda-4e3e-918a-a827bab51c3d', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 3', 369.15, 167, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f95610e-9d4f-49e2-a4b3-41eff626501e', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 4', 361.94, 232, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e9f07294-2ce5-4acc-839e-1e39f4fba4a1', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 5', 274.44, 391, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0e8696c7-1100-4b55-8f6b-51774b31a45d', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 6', 183.44, 354, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9c1e86c0-9081-4082-9041-d09adfa8ecfa', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 7', 378.44, 419, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1727b4c3-042a-4ff5-8fa3-1d197aa71cb5', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 8', 416.25, 315, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('058dd273-73cf-4f7e-a920-4f053607cd0f', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 9', 433.40, 239, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e2d616a4-1a15-4727-8a87-86118a31faee', 'd2ccf989-6566-482a-8946-78a9bc9ea6e6', 'Блюдо 10', 230.00, 102, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('afc1b157-07d0-41e7-a837-65de19c88703', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 1', 424.99, 100, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('53862031-1d93-48b3-b997-c79f5dfc31fa', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 2', 150.12, 341, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e3c51bc9-8e96-4b5c-8e56-9dfdd7368fd7', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 3', 300.47, 250, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8716cfbe-cea2-4464-8521-877d9cd41b41', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 4', 369.34, 453, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5093df8-6161-49df-baa4-0842d59f06dd', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 5', 195.63, 230, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f7a93b28-5606-49db-8a2d-f7bfa06addc8', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 6', 594.03, 441, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01897c66-7a57-4517-bc7c-7a5c889a9bc9', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 7', 436.34, 125, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a2d7c783-024b-4bb8-94f6-5eb1b0aaa9e2', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 8', 254.79, 271, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8eee55f8-bcef-4a8b-8251-5a89e0915ab1', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 9', 248.49, 239, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2ec485f6-9664-4234-843e-fadbf4e7efc3', '73b97018-6893-486f-9eea-a20cd02cc92a', 'Блюдо 10', 298.89, 393, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cc479f94-3874-4485-85ab-c326f46ff0bb', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 1', 388.10, 256, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5311385b-65ba-4a21-af4f-d389fbae0435', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 2', 445.24, 354, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('592d2633-44c7-4a4b-a732-9b36a21402fb', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 3', 231.27, 294, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c35ac171-5ad0-41de-bc62-1fdd35b2aa28', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 4', 376.60, 280, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0dad2b84-7fa9-4904-8ee7-fff92fb4b112', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 5', 397.43, 481, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('16c09094-26ff-4c4e-8bde-82a2fc16c9cb', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 6', 313.70, 189, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c43dd1b3-09b4-428c-a404-53a6096fa6a4', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 7', 254.96, 365, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('10fb44bf-28df-4fe3-9661-4cc5620971c7', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 8', 272.02, 294, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7933da25-f199-4bed-8e3c-15751620c955', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 9', 140.15, 286, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('171cd83e-e46a-41c2-8a87-1ae0ff4c37aa', 'e4516f8b-0c7f-4cf1-bffe-430be402dee5', 'Блюдо 10', 407.04, 332, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5db79024-de5c-4b2d-9670-c06ac0e0074a', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 1', 110.03, 164, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d7d7bb79-d6e5-40fe-964f-207c1cbb3105', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 2', 377.64, 293, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dbea61a5-69d2-4d82-85a9-4c8bbf7b31a5', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 3', 512.20, 124, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('900ad8b2-9378-404b-90a3-bcf2740d3fc4', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 4', 211.98, 185, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('81c105a9-d9ab-40fc-9419-5eae91c8aa46', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 5', 183.96, 255, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c2c9ecf4-f40e-4197-b8a1-b3dfd409c3dc', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 6', 528.11, 496, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('45e6066e-7da0-4354-8648-241995622f34', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 7', 479.76, 322, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1e1c787b-f732-4459-8e0c-8ffeff2385af', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 8', 353.79, 256, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5d19bcf6-727c-4876-bc25-49a0280afc7f', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 9', 489.27, 102, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5a7566f1-96ae-43dc-9423-4fedbb3a1a41', 'd4cf9bed-2af5-4bc1-9c88-32973c833b4d', 'Блюдо 10', 372.51, 332, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ea5e2e01-e0cc-48f6-bfad-097781015aea', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 1', 383.35, 125, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bbaea38d-8954-4a7b-8b36-3867021c0419', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 2', 257.81, 277, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a8abf1f0-0d3a-475b-980b-dd6932887745', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 3', 238.79, 114, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ee096c1-6b12-473b-99f7-a6c09c9795ba', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 4', 502.29, 208, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9213e023-d5ed-4367-861d-eeba391a4362', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 5', 116.60, 113, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dc692314-086b-45b9-a160-d5ef367885a7', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 6', 189.95, 464, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('91e5b6a1-fe19-48d0-9a6e-aa6c2c0bcba1', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 7', 470.07, 180, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1c2b14eb-57a4-4114-8f7e-082b13080ead', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 8', 113.31, 490, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5836f20-af64-4b62-9129-cd7b28f7171b', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 9', 148.02, 378, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4d0ba1bd-e100-4113-ab2c-96831cbf64eb', '009c0df1-e9a8-4be9-8781-78c98decb8f8', 'Блюдо 10', 375.35, 173, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b3bd3721-cfef-46d1-9b1d-d404f6099eed', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 1', 243.37, 333, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e1c9415c-1b6b-4703-8548-a695d6c8463a', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 2', 495.61, 427, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7fca0225-f640-4bf6-a383-018da3c71197', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 3', 102.34, 292, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('563e5474-40f2-44f3-91c8-768d592d322d', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 4', 535.47, 494, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a7a02373-e06f-44d7-b884-8180c86acbcc', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 5', 472.09, 131, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c96dd9ba-fafb-438c-8f1b-12bde84a256e', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 6', 521.72, 334, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('af9d4781-9b73-475b-ab55-d90e99fb0621', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 7', 430.47, 482, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('008c42fb-3af2-4042-882a-bf55f3d8f00d', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 8', 596.07, 117, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1de2dfb1-1fab-4fc7-94fa-4cb2356e0e61', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 9', 162.20, 126, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98031c45-1edb-4232-8884-0593a216dc84', '3b22c3db-89a2-403e-b641-19f9f58a70b5', 'Блюдо 10', 543.81, 496, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('108806f1-4ff9-4ffa-bf72-bf23f46cd837', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 1', 108.81, 457, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('15b2cfd2-e4f5-4658-931e-6138138512ae', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 2', 317.27, 436, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a65b04db-b873-48c3-a059-e371b32fdbc7', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 3', 407.89, 114, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('155183a1-5061-43a2-9f0a-b0056abfa451', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 4', 316.15, 132, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fbf17f6f-324d-401a-a7c3-ec2f6a665d49', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 5', 292.54, 497, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('84384616-cc14-4dcd-805a-5a841b5e36ec', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 6', 146.08, 348, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('67b91db7-e210-48b7-951d-1c5b1b780dbf', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 7', 496.79, 316, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f4b4e3ee-2f37-4963-a535-7e53bf7a2135', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 8', 407.85, 241, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7dd9f100-fcd4-4cba-9880-3e2afc438135', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 9', 511.37, 115, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c450d5ba-1f79-4f88-975b-ad4b54c36979', '44f256ba-397d-4a93-98ee-e2b652168673', 'Блюдо 10', 266.52, 174, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1c2a895a-aac6-47f5-8679-9d449ab322c7', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 1', 270.65, 412, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d5c7a238-9c0a-48d3-96b7-f4533289070b', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 2', 230.14, 101, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d72bfc6c-7942-4364-a201-461649e95a8c', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 3', 202.52, 453, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('338aacea-4456-4f45-9d21-c0a446a2aa64', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 4', 271.99, 125, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('db015995-f577-4d7e-aa94-43686d4aae44', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 5', 388.63, 335, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a29c9d4f-6c1d-4012-8599-37c3f2d57c2b', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 6', 444.22, 144, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('18348b89-b384-498b-80f5-f5cf34ec368f', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 7', 334.78, 140, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('93bde623-1489-407f-b2e6-27d1f1baa93f', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 8', 583.27, 175, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6acd0de8-d64c-456d-bece-8028577a2dfc', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 9', 326.62, 486, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0284ca44-5a4e-417c-ac60-06569592df30', '1f760e04-7c4d-4670-b546-d30c84d9602a', 'Блюдо 10', 109.21, 105, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('afcb2859-890c-48ad-aa87-606f95a0513a', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 1', 166.22, 311, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e04cb151-b2cc-4654-b65c-8017de1e32f2', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 2', 272.00, 348, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a057fc3-042e-4ee7-b36e-14b61ed99f45', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 3', 438.26, 241, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eeac2d32-6c83-4ceb-a150-b16e693312a0', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 4', 235.49, 478, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2ed56794-b42a-4881-9e0a-facfbe80b231', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 5', 495.99, 109, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a3ae3454-1ef1-4803-b6d8-e7a31a5df0a3', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 6', 460.55, 370, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('194d704f-708e-4879-9671-f1a45aec6b63', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 7', 573.40, 251, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9cb9f3fb-0980-4fb5-bf7c-46c3a6199524', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 8', 446.54, 386, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3539d90a-0bcc-46f2-b20b-fc966e7bc998', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 9', 377.37, 190, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1dcc9f86-e920-4e3d-95b0-e170b0d1ba16', '42ac54b2-b31d-46d6-b8bc-49e630623822', 'Блюдо 10', 212.23, 143, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ecf62173-6171-48a2-a4da-fa20fa5062a1', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 1', 399.94, 493, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2816ed37-598a-4c0d-a6b5-7a59c0b49b8c', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 2', 136.09, 456, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4bde6845-00b5-4c2c-aaa1-4cdb4ac09d23', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 3', 502.49, 265, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('83207f1b-cfe5-423c-b4fd-52949de615f0', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 4', 255.92, 263, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('babb4365-31d1-40d5-975b-f3b0305b2dd8', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 5', 448.98, 257, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f2b249c8-bf2b-43ec-8a26-7decd035567e', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 6', 345.79, 306, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('80efe28c-20ce-4ae5-becb-117e0a05be7f', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 7', 393.91, 186, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dac44609-6d5a-427d-9361-c4ce5bea8b6b', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 8', 164.11, 315, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7d01ea4c-e9cd-4b13-afd1-6188badd1d0d', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 9', 255.71, 377, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddbb5f27-02e4-49f0-80ff-56ce78fe7111', 'fda1df37-8d3b-4797-9f47-70e644d917c3', 'Блюдо 10', 322.01, 265, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('83a9f932-dd30-4a0b-bd5b-7097e7c1cd53', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 1', 284.68, 349, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d57dd8ad-fb7e-45fc-95d9-f764181b0475', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 2', 316.99, 252, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3a914212-c77d-4773-9b0e-e3ea8e452d22', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 3', 573.95, 191, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e1ff2ba8-f722-4044-a4a5-68b44531d2f4', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 4', 490.23, 465, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3fb1f65b-94ed-4990-838e-aecd5515f137', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 5', 109.69, 157, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7187eea2-85f9-4a08-ac14-72e71beb34bd', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 6', 362.04, 281, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4bb48259-206c-4533-bc8a-dcf2184ca496', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 7', 358.66, 333, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('869b7e92-da77-4154-8324-56e214bd326e', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 8', 592.97, 142, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7af1195e-31f4-4a52-8033-f12201495c55', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 9', 220.31, 429, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7d88bed5-51f3-4ac0-ae13-100b507a235e', 'c6ede17e-f729-4e8a-b79b-7472bf9dbf29', 'Блюдо 10', 385.78, 206, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ca01f78a-5c4e-40a3-a288-63eb1fec832a', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 1', 187.91, 282, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6f375af8-09f4-4b5a-b096-2ea159356743', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 2', 543.20, 440, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5da22ebd-d9b2-4524-b85a-23dda069a842', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 3', 165.38, 366, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('116582e2-d245-4a89-b2f9-94c3e66e160c', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 4', 515.38, 308, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f9aea8d1-8344-4357-bc47-1679e31d2ab5', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 5', 219.39, 246, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('60c14bca-d4b1-4f05-bdb8-854934ba156e', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 6', 346.44, 202, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f25c064-b6ef-43ec-ae42-d2406c397076', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 7', 142.51, 437, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a0c01388-2fb6-49e1-9906-03d084f7f81d', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 8', 416.69, 243, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1ead8f2c-812c-4e59-8ff4-9d659c703c43', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 9', 382.79, 153, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('14137d8c-5b27-4896-ad44-9e1e7a56e2e2', '0d9117c6-390b-4f4c-9823-9a5753394f2f', 'Блюдо 10', 230.60, 379, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('884eab42-4c2f-48ab-a620-c2595d6648b0', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 1', 100.78, 237, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aec169f8-2254-4e61-9527-26e51e9439d7', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 2', 442.58, 438, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8f6d79ab-3074-4b55-b163-2270a2959c6c', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 3', 511.72, 372, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b7c84c49-137d-472c-9fe5-d43851563829', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 4', 171.84, 485, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('27956860-9619-4f0d-901b-ac332f8dfc67', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 5', 177.24, 368, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5c2122f8-b795-499a-8b09-aec957ae869e', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 6', 169.42, 202, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2cd10d7b-b0b3-445d-8292-9131e0933d96', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 7', 280.31, 362, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d4fb3dc9-0134-4ae3-93cd-780aa050eb27', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 8', 456.73, 246, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('83424de8-6ae6-431a-961d-415d3d19c80e', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 9', 244.91, 174, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cb6b1df0-a2aa-420e-9084-e4f0a23506ea', '7bc2eec1-2cc9-4c90-a949-7cc05f294549', 'Блюдо 10', 237.44, 120, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('68155983-19fb-46c9-9f68-3acbf06b19a3', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 1', 513.83, 306, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('381a4a30-900e-4e61-9254-8a8c506bfae9', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 2', 157.73, 147, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a8e4dfd9-2455-4e9c-b477-20ce98887e61', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 3', 216.93, 294, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('609e06d5-f2e3-466b-8271-3877b6610846', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 4', 277.59, 194, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('668d110b-08d6-497e-9c92-cf25584eeb88', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 5', 353.75, 185, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9c9f4ca5-b3ab-4365-918f-6147af91330a', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 6', 499.38, 287, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6ab12d6c-6e83-4da1-8a23-9cbecf5ed324', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 7', 315.08, 428, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d68d9e74-4f5b-4868-af0b-6bdd8f4af26a', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 8', 571.56, 469, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('783f636b-93af-4604-8829-930a32309727', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 9', 574.22, 271, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('065b1462-b9e1-4ea0-89c9-7b528f4241df', '29e3b299-1d7b-4bd2-8227-35b749973a47', 'Блюдо 10', 327.29, 122, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('44bafc7f-d43d-4b1a-8f32-59f13d0b3674', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 1', 236.66, 185, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6365aa1c-e57e-41a1-9815-baffa9af96d5', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 2', 331.30, 187, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98c82f15-0bd3-4f5c-833b-b4c0efe6e712', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 3', 360.56, 398, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c3b587f2-7efd-49bb-ac29-49249f4e519f', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 4', 355.28, 198, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('554d5aa5-859c-41e1-b38e-62fdd1e520de', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 5', 252.13, 377, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('befaad37-852e-4262-9151-22cbc988d345', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 6', 585.82, 302, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ccabf5ea-1b47-43cd-be50-cee064e7d620', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 7', 550.33, 220, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f8e7033-3cfd-45d8-8f4c-9bb1ab4dcdfd', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 8', 491.37, 447, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('457e8d22-2105-4121-909c-b03efa43351f', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 9', 291.02, 334, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('70459b37-c033-48a6-a452-ab532a92bcbb', '91abbb07-b935-4764-a204-bcedde60d172', 'Блюдо 10', 322.96, 477, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c38ebc45-481e-49e5-a039-9a61aa52d92f', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 1', 100.11, 469, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ea77a62a-9451-4bf8-bcd5-803f6d6a1c0e', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 2', 413.36, 350, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7d54e195-45fe-4146-a582-ef58ba7bd982', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 3', 207.15, 422, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bdf03ee4-8757-4944-8223-96f8c61070d8', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 4', 490.07, 225, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bc0dd856-386d-4a8f-8287-cf158c9ee091', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 5', 328.25, 372, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ebd48ffd-6c28-433a-a556-c9be6b8ab885', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 6', 558.61, 439, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4f54a8df-49f9-41ed-aae9-11bae2bbdea2', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 7', 430.23, 349, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a6a6e0c8-a9a5-4c9c-aef6-934f889872b8', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 8', 160.78, 485, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31c02786-49bb-4a61-821d-ee646060f2d0', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 9', 530.18, 402, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5274db65-6899-4e77-b772-d931bba864d6', 'd680129a-8752-4aab-8b71-4bc45ffd27eb', 'Блюдо 10', 505.12, 274, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('30a570bb-4937-436a-ae11-6d11b9d73579', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 1', 349.33, 222, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4dcbf122-fa37-46b5-9794-e4ad63b07cbd', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 2', 412.23, 499, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('28611d1f-620b-472b-a463-909bdbb9a574', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 3', 445.72, 174, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a78fe52-1678-44de-908c-6f6e2a336475', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 4', 354.38, 384, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aebce01c-deef-42b2-a743-5004f32a2fe0', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 5', 119.44, 368, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f7d0afd5-7f4a-40dd-b66b-4bece9618343', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 6', 353.02, 470, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9d15ec4a-d0d4-443c-ac74-0c4f45ef7f2a', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 7', 198.02, 253, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c075452-ef28-4834-914f-48455486ba3d', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 8', 363.44, 148, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6f4ea3c6-6ac4-4337-930d-81cb927753a2', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 9', 212.32, 473, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f05b89e5-c294-4ca3-a767-a5b6c13abe50', '3f180fd1-1516-430a-b976-b9fd3b1bc76c', 'Блюдо 10', 362.88, 175, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c5a478a1-f9d6-4031-a73b-d4a6366796bc', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 1', 311.93, 258, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('575c77b6-03c1-498e-a89b-0701401fde4e', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 2', 273.65, 429, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b6fe95ff-c81c-4b99-882e-36260a7ee8d1', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 3', 212.88, 252, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('22ab2ada-6ec8-49de-aef3-0e2d8e06d5e4', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 4', 446.38, 233, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d59d3cb8-7851-4f6c-91ee-7c1e18574329', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 5', 482.12, 167, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('163a6e67-4bb7-415e-bdfd-57ea51ac1f80', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 6', 366.80, 166, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1c30029e-5ec3-49cd-839c-93e5e63856b3', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 7', 146.29, 400, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4e48d596-35e3-49c6-9fab-82c424fda180', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 8', 272.24, 495, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a36742ba-0b6b-40b9-a0c9-cef14c9efa96', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 9', 371.58, 408, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5bfc8cda-be2b-4234-a8cb-9ca82116218c', 'd79c7aa2-db06-42bc-ae95-370f0f1f6490', 'Блюдо 10', 424.43, 469, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6745875b-98db-4d7d-bc9f-03fb60d89940', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 1', 405.69, 242, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fa99dadf-0a0e-4d06-969a-04b25c460e2c', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 2', 519.90, 189, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e91793f2-8c84-4a65-bd1c-cf2022e42e75', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 3', 418.72, 277, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('239b3956-e35a-466e-ad88-5dbab9bff64b', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 4', 563.12, 311, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8c0ec708-5ac9-4bce-bf3d-6f15ccccf82f', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 5', 197.01, 259, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3ce7e37a-ae5e-4f5d-9e97-1bd8a42fa627', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 6', 164.67, 246, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4ffb4373-d8a2-4c78-b983-ff2513319c6a', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 7', 241.43, 317, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddd857d6-8ca3-42ba-a5c5-383d0c55aac9', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 8', 258.28, 395, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1ad8b8a3-97c0-4511-9ee7-c7a4c536a342', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 9', 239.43, 219, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5b4db7c1-504a-4fe0-8bc1-f1ffa853308b', 'db6c3a81-3b78-4709-9c79-0b56ddf44910', 'Блюдо 10', 419.58, 140, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cc026b9d-7acd-4b27-8d75-3c89baacd8d1', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 1', 291.31, 363, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('37572820-bb38-4fa0-a9ca-0535c5564659', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 2', 348.44, 270, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('789980f8-eb46-48d5-9200-bfe641a1574f', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 3', 280.98, 465, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87b7a374-ef5f-4929-aeb6-3ffa1d06e74b', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 4', 137.79, 310, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('83851942-1708-4e5f-ac98-209b9b3bafbf', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 5', 198.28, 493, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87643f07-b3ee-4877-8e89-7df8838dcc3f', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 6', 154.56, 260, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('43f0b509-a7eb-47dd-83e8-fbf908475868', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 7', 403.70, 371, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1cc46df1-771c-4216-b245-4b61aab57f51', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 8', 588.28, 147, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e9e3e0a9-97fc-4c81-9da4-acf24b9ce40b', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 9', 168.31, 408, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('21499130-dbdf-4dd7-896c-165d6cd16038', 'c5f23fa2-d9c5-413e-90a2-8d1b62219184', 'Блюдо 10', 122.54, 456, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('42f07472-6261-441c-a489-e68d905a72b6', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 1', 551.93, 440, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('acc5b0fe-4078-4252-8fc3-5a7664547265', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 2', 256.69, 115, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('97dbfa67-d6a4-4d14-b5c7-79f82c35c7fe', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 3', 511.94, 257, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('62275096-e67c-4f95-8bf4-906e9f2cbfd5', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 4', 162.17, 311, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1c57c1c9-20a5-4482-abeb-3674107da97e', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 5', 481.83, 182, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ca79dd57-bb25-4646-8e70-f0931f28d32e', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 6', 172.92, 151, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b1d02099-84aa-4440-a09d-d7597a71d491', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 7', 279.36, 498, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3cdaa781-ac85-41b3-afa8-ecf8175d3616', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 8', 251.45, 491, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dab79922-e0f2-4e2c-95ee-b55fc049e489', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 9', 594.29, 379, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b098ca82-fee4-478f-8b28-9be79705ffc1', '340a81a5-3476-4bc6-8300-57f988ec6929', 'Блюдо 10', 445.97, 353, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d180c652-2acd-445d-a64c-331ac994ca9d', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 1', 348.43, 421, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('38bab2a8-28d8-4a14-b126-b6330331d309', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 2', 309.41, 270, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b8832da-7f5c-4149-9b79-6965a39908ac', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 3', 415.94, 459, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('656ec053-85f2-4f66-8a10-82f4113fdb87', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 4', 530.45, 117, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7e59f7d6-1f43-422c-af72-86fca4f087e0', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 5', 328.87, 497, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bcc46d1c-a4c1-41ca-98f8-b46098e07793', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 6', 281.61, 151, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8d97bc34-6123-47a8-a66e-0c9e9994c641', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 7', 466.96, 366, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ec6a8504-b45f-4530-b9e6-be1795ae9791', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 8', 357.97, 400, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b7d59628-7c31-4199-994a-d47483ad7651', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 9', 417.30, 304, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('88ec285d-6ce5-4674-946f-8b0bc72b42e7', '0216534b-fd51-48e9-8aa2-8f0530585f76', 'Блюдо 10', 335.20, 128, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('04a6c27c-f931-48e1-ba7c-eb1264dfc655', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 1', 228.18, 181, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('34f0a395-2db0-41e6-b126-20ac82b20f74', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 2', 164.61, 103, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b1afef1a-3151-499c-ab3e-e6ef70f89c11', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 3', 177.32, 439, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c00ccd1-a8b3-4ade-b4f5-b4540ce25d74', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 4', 582.53, 117, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('688adfa7-8fd4-433b-aba5-c429417bbd8d', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 5', 177.52, 145, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a11ce16-7b31-485e-94e5-25073c5075d5', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 6', 495.58, 284, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4798a205-2d0d-4f29-b448-52c31365ffe0', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 7', 534.36, 340, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8760c602-57cb-428f-bf23-b5121c60bc38', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 8', 443.17, 491, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e22afa9f-8b06-4ba3-985c-5a4e62558861', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 9', 405.02, 296, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('efd1db9a-e242-40fd-9d10-5dd3e680d481', '2fa219b3-1160-4798-b528-d8b080727c27', 'Блюдо 10', 313.71, 119, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d24e4ace-36c8-40f4-bfcc-05310dcb48f5', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 1', 169.80, 214, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('497b387f-8f08-4d8b-b37b-8b20309d8459', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 2', 583.27, 159, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dcf82ad8-7081-49ec-b5e8-1bb03ee5e5d1', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 3', 358.04, 483, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f36d8072-4afe-419b-9e73-4d659442034e', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 4', 337.13, 181, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cb25f787-cfb9-4eb3-92be-2c77b14df09d', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 5', 462.98, 283, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d41f971f-39c3-4795-951b-1a58dfe3849f', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 6', 100.31, 434, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('785b7f35-5a85-45ea-a94f-6e8a72f4cb5e', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 7', 575.60, 242, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('36134fdc-c80a-42a0-9fad-23726fb85b30', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 8', 133.71, 176, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b682081-2f74-4080-bf14-c60b9611a024', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 9', 110.57, 351, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c18ef64-159c-46bd-8034-3907bdcb0ade', 'bab94b23-76e8-46a5-9949-d17348eef355', 'Блюдо 10', 346.72, 355, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('de895550-82a4-4ab7-9d35-e812d67fb079', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 1', 574.43, 434, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cc504d92-54f6-4bb3-a0ce-2593012e9738', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 2', 562.04, 453, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e0ad2a20-42a4-40ae-a9f6-5cc5ee14c406', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 3', 356.54, 375, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('24703654-6514-439e-85a2-38b6cb30605a', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 4', 575.25, 145, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f1449fcb-3f8a-4223-9365-ffeac0ef2c9a', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 5', 480.44, 244, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('975c5270-f361-4d0c-a848-8a749b203062', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 6', 502.04, 328, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a94fde2e-cca5-42a2-874f-aa550b5aacb6', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 7', 546.10, 336, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c8d0dbf7-e209-415b-acc4-40411fce9eae', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 8', 125.25, 134, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('27684e1e-06fc-4dcf-8373-fbfbb9e5cdc8', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 9', 112.66, 472, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('507ec657-3f3e-4dc1-a9fc-fbc76aae5ceb', '0b6a4247-192c-4d6b-9cfc-5ce1c756fc6c', 'Блюдо 10', 462.84, 256, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4627660f-2677-4d41-aa0a-31499afa2870', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 1', 330.72, 313, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('43a5be5b-d6dd-4cdd-93d7-4ad00832a5bc', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 2', 424.40, 404, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('56739bbc-6b62-4359-aae9-4893cc9aa129', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 3', 174.15, 118, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('096acdd1-2b06-457a-b218-e0a628539451', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 4', 198.48, 118, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b7bc0c3-4337-43b9-8a7c-53bd01dea1b2', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 5', 269.13, 439, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b9de6b4e-eeb6-48c9-9a96-7b459a1d5094', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 6', 311.85, 346, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6972c973-a1dd-4445-b5f9-ce971e85114c', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 7', 267.02, 143, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4eafa378-3e92-48b1-a967-be375d201085', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 8', 494.40, 243, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2052ba21-6fe0-4801-bb61-edb857fec3be', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 9', 532.92, 286, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1c99f097-ee9a-4084-8e15-e63676fdc287', 'd34e760b-0640-4438-b1e4-80c4ae25eeca', 'Блюдо 10', 220.39, 158, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5602b22-aa56-4a0a-b30e-0e99f0a1f55d', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 1', 175.47, 365, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bc044bb8-f2fa-45fe-b915-cf911d9dcede', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 2', 329.20, 141, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5dcda023-df12-42d4-b5b0-73cb4ab4885f', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 3', 419.15, 105, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('200976a3-49c8-479f-bee7-74842b98b6a7', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 4', 294.36, 210, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b2081f00-fcfb-4d1d-965d-05564b66a177', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 5', 288.91, 361, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d9abe8d8-9ad6-44f2-8c96-2d2c597df882', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 6', 142.46, 292, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6ad5ad46-99fa-4826-91bb-cade61b71154', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 7', 518.53, 427, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('15b3cf6b-ae1a-44db-9abe-e49aa89c17e4', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 8', 551.31, 322, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fee07cd4-146d-47e4-a7cb-cc85a7e22a90', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 9', 372.62, 462, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('391ba9f0-4fa9-4601-aa5a-231642a8eab8', '7cc2c6b8-0533-4d99-b494-95a9cbf14ef7', 'Блюдо 10', 524.98, 451, 'Супы');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('ad770464-b8fa-440e-8ba1-ed66de9a2cbd', '8c4aa27f-3a80-4b49-8113-7991dba5603e', 'created', 'Улица 77, дом 37', '[{"id":"ef9fe7e0-1e7a-4e94-a3da-474a95395f9d","name":"Блюдо 10","price":126.33376161480263,"image_url":"default_product.jpg","weight":286,"amount":2},{"id":"76aa4e00-77cd-401b-a90a-0217e26f351a","name":"Блюдо 8","price":225.05711301521094,"image_url":"default_product.jpg","weight":177,"amount":1}]', 'кв. 105', '619', '1', '3', 'Не забудьте вилки', false, '477.72');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('2fc28600-6bdc-4b55-9691-81ac5297db65', '8421cf5b-a672-4070-938f-2c9f7c24c81b', 'created', 'Улица 69, дом 11', '[{"id":"c21e1de0-2041-4c7f-b1d4-0da34929c201","name":"Блюдо 2","price":254.53572858782226,"image_url":"default_product.jpg","weight":264,"amount":1}]', 'кв. 102', '714', '2', '15', 'Не забудьте вилки', true, '254.54');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('d05303d9-37ac-427f-8ba1-b8f3061d9b1e', '4c40d5f6-51ce-480f-a566-9fd4f381c8e1', 'created', 'Улица 44, дом 18', '[{"id":"1f3fab3d-5fbb-41df-96c8-91ad3745bfba","name":"Блюдо 9","price":321.55748672758665,"image_url":"default_product.jpg","weight":196,"amount":2},{"id":"691c0bed-ef60-441c-8fa0-066b52a1fe22","name":"Блюдо 6","price":345.63291424362876,"image_url":"default_product.jpg","weight":457,"amount":3}]', 'кв. 148', '578', '1', '9', 'Не забудьте вилки', false, '1680.01');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('2146c2ec-16ad-4cee-b5e5-f9c4f583cdc5', '3ccb6b26-2f63-446a-9042-39379fe46b09', 'created', 'Улица 73, дом 10', '[{"id":"bd750b29-5c38-484f-9868-09ab48279918","name":"Блюдо 8","price":153.02788946085477,"image_url":"default_product.jpg","weight":339,"amount":3},{"id":"4191768e-7a67-496f-9173-c70b5270f129","name":"Блюдо 10","price":148.85338648261765,"image_url":"default_product.jpg","weight":472,"amount":1},{"id":"9d33e4d9-0d73-4ee2-8471-f6eb79cbec8e","name":"Блюдо 2","price":426.42179543106823,"image_url":"default_product.jpg","weight":434,"amount":1}]', 'кв. 198', '422', '2', '3', 'Не забудьте вилки', false, '1034.36');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('ea51bf4d-cbd7-4f40-aed7-0aa35acd67c0', '4c40d5f6-51ce-480f-a566-9fd4f381c8e1', 'created', 'Улица 31, дом 22', '[{"id":"6c075452-ef28-4834-914f-48455486ba3d","name":"Блюдо 8","price":363.4388998287095,"image_url":"default_product.jpg","weight":148,"amount":2},{"id":"9d15ec4a-d0d4-443c-ac74-0c4f45ef7f2a","name":"Блюдо 7","price":198.02197286569486,"image_url":"default_product.jpg","weight":253,"amount":1}]', 'кв. 66', '425', '2', '14', 'Не забудьте вилки', false, '924.90');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('c5197fc1-42d3-496e-94f2-7d6de294c532', 'eb7f1f8c-0dc8-4bdf-abe9-5f1ef2a8627e', 'created', 'Улица 44, дом 18', '[{"id":"fe4e45eb-0585-488b-8c6a-17140144b5ca","name":"Блюдо 7","price":382.3592439769125,"image_url":"default_product.jpg","weight":393,"amount":3},{"id":"60106bf5-241f-4025-a773-c623737e1982","name":"Блюдо 1","price":295.8423904381435,"image_url":"default_product.jpg","weight":257,"amount":2},{"id":"188f22f0-8d1c-4584-a5ec-e9d90bbecdf6","name":"Блюдо 8","price":216.4696858426214,"image_url":"default_product.jpg","weight":143,"amount":3}]', 'кв. 111', '684', '5', '4', 'Не забудьте вилки', false, '2388.17');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('d6502469-2876-4877-9e91-e25417dfd9fb', '8c4aa27f-3a80-4b49-8113-7991dba5603e', 'created', 'Улица 75, дом 11', '[{"id":"7949cbe9-b3f4-4f3b-9a4d-2ab14dabd26f","name":"Блюдо 2","price":561.4316586136624,"image_url":"default_product.jpg","weight":116,"amount":2}]', 'кв. 66', '545', '3', '23', 'Не забудьте вилки', true, '1122.86');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('23f384bf-cfb7-476c-8c4c-0183cd059c3b', 'eb7f1f8c-0dc8-4bdf-abe9-5f1ef2a8627e', 'created', 'Улица 44, дом 18', '[{"id":"fe73a844-ecef-4a1f-b5c1-fee8c48e309b","name":"Блюдо 6","price":279.1416305307008,"image_url":"default_product.jpg","weight":453,"amount":3},{"id":"ab3df713-93f4-4b4d-96fd-82d3d9bdfb4e","name":"Блюдо 8","price":167.72882272968164,"image_url":"default_product.jpg","weight":362,"amount":2},{"id":"59a53f95-64a9-46ce-b8f7-60aaaea97242","name":"Блюдо 1","price":479.67682628749867,"image_url":"default_product.jpg","weight":160,"amount":3}]', 'кв. 73', '782', '4', '3', 'Не забудьте вилки', true, '2611.91');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('dbca9233-84d0-42c4-95a9-fcfb8d2662d5', '360d8748-ef26-468c-916f-0f5f68558f2e', 'created', 'Улица 60, дом 6', '[{"id":"2f95610e-9d4f-49e2-a4b3-41eff626501e","name":"Блюдо 4","price":361.93586027555295,"image_url":"default_product.jpg","weight":232,"amount":2},{"id":"cbb8298f-7eda-4e3e-918a-a827bab51c3d","name":"Блюдо 3","price":369.15052951343984,"image_url":"default_product.jpg","weight":167,"amount":1},{"id":"1727b4c3-042a-4ff5-8fa3-1d197aa71cb5","name":"Блюдо 8","price":416.2461910663344,"image_url":"default_product.jpg","weight":315,"amount":3}]', 'кв. 75', '196', '2', '16', 'Не забудьте вилки', false, '2341.76');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('c04a02fc-d7c8-48f2-ab02-e8558d81f3f3', '482a75be-e6b5-4f0e-b56d-01602a1c6b36', 'created', 'Улица 73, дом 10', '[{"id":"53285a21-0bb1-4207-8dad-72581e81ecb8","name":"Блюдо 7","price":527.8183567003597,"image_url":"default_product.jpg","weight":214,"amount":1},{"id":"a8dc6c28-cdbf-4650-b791-eb12432ecde2","name":"Блюдо 3","price":277.31387987234547,"image_url":"default_product.jpg","weight":403,"amount":1},{"id":"73971ba7-d2e0-49a6-8fbd-5486c63acc36","name":"Блюдо 9","price":521.3627324681818,"image_url":"default_product.jpg","weight":339,"amount":3},{"id":"440f88ab-e5b3-41a9-a35a-3e0e6c32ced1","name":"Блюдо 2","price":434.0727900891551,"image_url":"default_product.jpg","weight":354,"amount":2}]', 'кв. 10', '958', '3', '25', 'Не забудьте вилки', false, '3237.37');
