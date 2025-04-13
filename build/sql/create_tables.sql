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
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('8b72e7df-64d4-4037-9cfd-e79a06a7563f', 'user0', '+7-721-465-1171', 'Татьяна', 'Петрова', '', 'default_user.jpg', decode('988dd63a136a8a4932f13726e829a5d81b929c5b78a498fdd470974e627a420f', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('bc20994d-919a-4d57-aae1-50f9da7e8588', 'user1', '+7-310-732-5075', 'Дарья', 'Смирнова', '', 'default_user.jpg', decode('89c1af09c886fb0f5c095fb68d0e5d75db47624ac76fc86159a6b6bdf1768794', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('9b414523-d7ce-4c90-b22c-a569ae0351b5', 'user2', '+7-622-121-4613', 'Дарья', 'Иванов', '', 'default_user.jpg', decode('7307cccb0488b7c7c7ecd11c9ef00951f3f3f62555370e15c6a18af039e9dc9a', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('6c632bf4-5a56-4b22-a3b0-029100dae7fa', 'user3', '+7-357-904-2363', 'Алексей', 'Попов', '', 'default_user.jpg', decode('1407eff2a277943d7aea202c33c32e10ef01a18df5b28a3f379a0173c5c5e030', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('7ca5d04a-fc86-4716-99d4-26e178b6016b', 'user4', '+7-808-547-9572', 'Иван', 'Кузнецова', '', 'default_user.jpg', decode('ad19ec2d8b485026606d1cfff6e024c27cd44f58e18a49c1c51206293739de55', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('594b5f85-a0c0-4b0b-b729-025e233d2c90', 'user5', '+7-408-353-4960', 'Алексей', 'Торетто', '', 'default_user.jpg', decode('6494f384fc45483e8353e7552835922ea404d9aed886374ccf7a4b05f822103f', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('90bbb029-540a-40e6-b17e-448dd8ddb19f', 'user6', '+7-538-130-8478', 'Мария', 'Кузнецова', '', 'default_user.jpg', decode('d656db0e83f036b2a71429fa401d2e82a558f6c22c0848843a3aff21274b0be8', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('a55b7add-a72f-4ee5-910a-77ac6f18891c', 'user7', '+7-350-874-2288', 'Владислав', 'Пермякова', '', 'default_user.jpg', decode('cd289970d03c6600f66e923b0e6afa5e2a286d372798dde88165c96f7cea9ceb', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('3419e315-99c0-4fd3-9f0b-39e7877e738e', 'user8', '+7-909-578-5549', 'Иван', 'Иванов', '', 'default_user.jpg', decode('3e706dfacf256cfbfd146e9b26ddb51b0a82e7d6ff88c412398d551fb6edf530', 'hex'));
INSERT INTO users (id, login, phone_number, first_name, last_name, description, user_pic, password_hash) VALUES ('5cc101e9-0cc3-41f6-8b56-c5e29de758b1', 'user9', '+7-854-599-5779', 'Иван', 'Пермякова', '', 'default_user.jpg', decode('b93ea231e8bbe248017613e27e1408c21d30bafd64aa960064053a5a59bb4df3', 'hex'));
INSERT INTO restaurant_tags (id, name) VALUES ('57282916-0faf-4644-9d71-df0cd0f8adc7', 'Пицца');
INSERT INTO restaurant_tags (id, name) VALUES ('43291590-b536-425b-88c2-ed51a55b602b', 'Бургеры');
INSERT INTO restaurant_tags (id, name) VALUES ('a76ecb9a-1712-4190-996b-0bf7a65de360', 'Суши');
INSERT INTO restaurant_tags (id, name) VALUES ('bfc2c4d4-dd15-40a8-83d0-ebcbf09c80e0', 'Веган');
INSERT INTO restaurant_tags (id, name) VALUES ('a966defe-324d-484e-b24f-6ddd1ed5c8f6', 'Кофе');
INSERT INTO restaurant_tags (id, name) VALUES ('0e0b0dc8-12e5-4fe2-a6c7-be56728db3a0', 'Десерты');
INSERT INTO restaurant_tags (id, name) VALUES ('be45fc01-60fe-4747-9a07-4648105b3fcb', 'Шаурма');
INSERT INTO restaurant_tags (id, name) VALUES ('32602cab-1dfc-4216-b420-b0e52e7e7ef2', 'Мексиканская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('6a0445f6-63f5-4c62-8593-f5f5996c15db', 'Итальянская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('c2a20de7-9745-4730-8663-9183bc9b6094', 'Грузинская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('4704ada0-e0c7-4757-b2fa-8c0da2b33028', 'Китайская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('6ec95b9c-5a45-4dd3-98e7-1459becc86fc', 'Японская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('b32a6026-82ad-499f-b064-549f5de171c4', 'Американская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('82c158cd-8f28-4bd2-a909-07079c6f5432', 'Фастфуд');
INSERT INTO restaurant_tags (id, name) VALUES ('f59be088-9084-465f-8878-94e49c24fee2', 'Салаты');
INSERT INTO restaurant_tags (id, name) VALUES ('085687b3-7b0e-4fcd-bece-89d67aecaaac', 'Завтраки');
INSERT INTO restaurant_tags (id, name) VALUES ('842f30a1-c95c-45c3-bd75-f3c651c22618', 'Стейки');
INSERT INTO restaurant_tags (id, name) VALUES ('6fba5248-381a-4f71-bbe0-9718ed0d5871', 'Морепродукты');
INSERT INTO restaurant_tags (id, name) VALUES ('dafc199d-8315-4cb3-8a8d-e997c71931b1', 'Пасты');
INSERT INTO restaurant_tags (id, name) VALUES ('60b2c182-4bfa-4b80-94ee-8c7e70ab0690', 'Смузи');
INSERT INTO restaurant_tags (id, name) VALUES ('78c370d2-b3dd-4278-ae03-97976ccf192e', 'Фалафель');
INSERT INTO restaurant_tags (id, name) VALUES ('2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23', 'Гриль');
INSERT INTO restaurant_tags (id, name) VALUES ('299f5f73-c94d-4410-92a2-70cbf91b37ea', 'Курица');
INSERT INTO restaurant_tags (id, name) VALUES ('fd720fe2-943a-496c-a23b-8e66e1c9c067', 'Рамен');
INSERT INTO restaurant_tags (id, name) VALUES ('148d017a-37ef-4a21-a030-1725054e26c7', 'Корейская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('469af9f3-1ca1-451c-aca8-3f51eeb2cbd0', 'Пекарня');
INSERT INTO restaurant_tags (id, name) VALUES ('a8886ce1-681b-4e98-a372-52edb6fa2347', 'Пельмени');
INSERT INTO restaurant_tags (id, name) VALUES ('277e9002-c1ca-4ef0-8bab-f1b7a45789ec', 'Вьетнамская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('6f7e909c-122c-470b-9f86-ee04ad9db1b3', 'Сибирская кухня');
INSERT INTO restaurant_tags (id, name) VALUES ('8383f6c0-f9b1-4406-ad3c-11aa6651b54b', 'ЗОЖ');
INSERT INTO restaurant_tags (id, name) VALUES ('ff290859-b711-4ad0-8db0-ba66f9d56af6', 'Кето');
INSERT INTO restaurant_tags (id, name) VALUES ('557e9da9-5df8-48d7-8aae-1e100f6d38f9', 'Халяль');
INSERT INTO restaurant_tags (id, name) VALUES ('90be5088-f75c-4759-8d42-ddcae792125e', 'Безглютеновое');
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Ресторан 1', 3.8, 69);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Ресторан 2', 3.1, 103);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Ресторан 3', 5.0, 123);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Ресторан 4', 3.9, 107);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1f216a35-23e2-474d-823a-f8c6f5e25545', 'Ресторан 5', 3.9, 88);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('61903dee-b84c-4510-a70f-1c022a87ccfc', 'Ресторан 6', 4.9, 143);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Ресторан 7', 4.5, 60);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('692ba0be-c31e-4525-93e8-e7e6eca57383', 'Ресторан 8', 3.7, 56);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Ресторан 9', 4.5, 77);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Ресторан 10', 3.4, 125);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Ресторан 11', 4.8, 123);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Ресторан 12', 4.2, 88);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1653a159-dc2e-4885-8d44-48b44fed3f36', 'Ресторан 13', 3.5, 110);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1c00b207-f01a-4cfd-9179-723890103f7b', 'Ресторан 14', 4.3, 118);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Ресторан 15', 3.5, 122);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Ресторан 16', 4.8, 129);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Ресторан 17', 3.7, 60);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Ресторан 18', 4.4, 81);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('724b1490-5411-485c-b96e-0a7d553e24b5', 'Ресторан 19', 3.1, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Ресторан 20', 3.8, 108);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('69884407-d5d8-4669-ac7f-5926def2a627', 'Ресторан 21', 3.9, 147);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('18cfdbe1-bb62-437b-8e36-74155878da89', 'Ресторан 22', 3.7, 117);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('01c92ba3-6348-4139-8683-82d5a38cc87e', 'Ресторан 23', 4.8, 144);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Ресторан 24', 3.4, 74);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Ресторан 25', 4.0, 94);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Ресторан 26', 4.3, 55);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('90d0f63a-2b94-4ba2-9907-1b981de65101', 'Ресторан 27', 3.2, 115);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Ресторан 28', 3.8, 92);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('555de0c1-d263-4aa8-80ec-0277fec44759', 'Ресторан 29', 3.6, 128);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Ресторан 30', 3.4, 51);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0e7a76e2-617e-449d-a098-4188f44356fb', 'Ресторан 31', 5.0, 143);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Ресторан 32', 4.4, 101);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Ресторан 33', 4.9, 124);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Ресторан 34', 3.9, 63);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Ресторан 35', 3.2, 59);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('61042683-9841-4b54-a3fc-9d578f38184d', 'Ресторан 36', 3.5, 108);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('18753203-4827-418e-89f5-21ef396f3778', 'Ресторан 37', 3.3, 110);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Ресторан 38', 3.1, 147);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Ресторан 39', 3.8, 101);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('9e5297df-3df1-498b-b166-0a39e9df0452', 'Ресторан 40', 3.7, 58);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Ресторан 41', 3.8, 80);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Ресторан 42', 3.1, 117);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Ресторан 43', 4.0, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('8592634e-56ea-4f62-b938-0066953bbf24', 'Ресторан 44', 3.2, 102);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Ресторан 45', 3.3, 83);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Ресторан 46', 3.8, 126);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('40f56c2c-c39d-426e-8496-dba481e4d09e', 'Ресторан 47', 4.4, 85);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('ac695f69-8672-4a1d-9d11-44284ee80009', 'Ресторан 48', 3.9, 73);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('65b2309d-e638-4473-905f-02beca4b92ac', 'Ресторан 49', 4.3, 50);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Ресторан 50', 3.1, 55);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Ресторан 51', 4.3, 114);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('866ccf38-e290-45c9-bd80-429456c725e9', 'Ресторан 52', 4.4, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('528d1d82-430c-4b19-898c-2037c0716c3c', 'Ресторан 53', 4.4, 121);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Ресторан 54', 4.4, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6b689866-3315-4e55-bbd6-0c61e35df71e', 'Ресторан 55', 4.5, 86);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Ресторан 56', 4.9, 148);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('4b9b57f7-8065-4229-881c-4e0345e10687', 'Ресторан 57', 4.3, 95);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7fdafc7d-9501-48b1-893c-287d8168db1a', 'Ресторан 58', 3.3, 135);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Ресторан 59', 4.7, 133);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('dc701bde-05d8-440e-9603-0f7a5559311e', 'Ресторан 60', 4.1, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Ресторан 61', 5.0, 139);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('18574170-cec9-410d-8896-c15a852f16c1', 'Ресторан 62', 3.2, 137);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Ресторан 63', 4.8, 83);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Ресторан 64', 3.7, 78);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Ресторан 65', 4.8, 124);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Ресторан 66', 3.5, 106);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('5d94821c-2f49-4f91-8288-ab7006402f90', 'Ресторан 67', 4.7, 52);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Ресторан 68', 4.3, 101);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('4a422f05-767f-442f-a3f8-5860fe565cd2', 'Ресторан 69', 4.6, 54);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Ресторан 70', 3.3, 55);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Ресторан 71', 3.9, 94);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Ресторан 72', 4.4, 55);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Ресторан 73', 4.5, 52);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1ca369b1-a772-4743-848a-d38d8ed49a79', 'Ресторан 74', 4.0, 122);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Ресторан 75', 4.4, 79);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('8eba03e4-b4fd-46eb-b565-5e300343f092', 'Ресторан 76', 3.6, 102);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Ресторан 77', 3.6, 107);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Ресторан 78', 3.1, 52);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Ресторан 79', 3.6, 83);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Ресторан 80', 5.0, 135);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1af09413-7078-42e0-a250-ab001ec9e57f', 'Ресторан 81', 4.2, 57);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Ресторан 82', 4.3, 67);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('1f8e556d-1779-411c-847e-422eab71ca1f', 'Ресторан 83', 4.5, 71);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b140129e-3a20-4dd2-86b4-cef4499a031d', 'Ресторан 84', 3.2, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('dd3575d9-d937-4ccb-92e7-237491e6a853', 'Ресторан 85', 3.1, 63);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Ресторан 86', 3.5, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Ресторан 87', 3.1, 64);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('d15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Ресторан 88', 3.0, 141);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('493e5895-17bd-4177-a185-6975a789cc58', 'Ресторан 89', 4.2, 110);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('59d649f5-7921-4249-aef9-860748b46b18', 'Ресторан 90', 3.0, 77);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('095ab469-eefb-450e-8cc4-db0a4aabce45', 'Ресторан 91', 4.0, 82);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('c274c432-3f84-4e11-936a-b6818ed1542c', 'Ресторан 92', 4.8, 102);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7cf710da-960f-4725-b249-e95c419d8d6d', 'Ресторан 93', 5.0, 70);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('168f82bb-cc39-4d4a-b35d-21f555f77308', 'Ресторан 94', 3.3, 96);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Ресторан 95', 4.7, 106);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Ресторан 96', 4.5, 102);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Ресторан 97', 3.8, 132);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('adda4717-f118-45cc-9532-30502637846d', 'Ресторан 98', 3.3, 129);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Ресторан 99', 4.8, 72);
INSERT INTO restaurants (id, name, rating, rating_count) VALUES ('64728b3c-67a7-4107-bac8-f6f998aed794', 'Ресторан 100', 3.5, 147);
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6be155ad-8705-4cc5-8bbe-1d60b279516c', '6a0445f6-63f5-4c62-8593-f5f5996c15db');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('657c3fb1-a388-4ca1-ae15-0b31673056b5', '32602cab-1dfc-4216-b420-b0e52e7e7ef2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7b27253f-26e1-494c-a54c-9e9e0c1ff903', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('97358ff0-ccd3-4817-8a9c-673c8875cb02', '6fba5248-381a-4f71-bbe0-9718ed0d5871');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1f216a35-23e2-474d-823a-f8c6f5e25545', '78c370d2-b3dd-4278-ae03-97976ccf192e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('61903dee-b84c-4510-a70f-1c022a87ccfc', '277e9002-c1ca-4ef0-8bab-f1b7a45789ec');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'c2a20de7-9745-4730-8663-9183bc9b6094');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('692ba0be-c31e-4525-93e8-e7e6eca57383', '277e9002-c1ca-4ef0-8bab-f1b7a45789ec');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'a76ecb9a-1712-4190-996b-0bf7a65de360');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', '60b2c182-4bfa-4b80-94ee-8c7e70ab0690');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('da752af1-9e8d-4058-86a4-b0aa3a200f19', '0e0b0dc8-12e5-4fe2-a6c7-be56728db3a0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1653a159-dc2e-4885-8d44-48b44fed3f36', '277e9002-c1ca-4ef0-8bab-f1b7a45789ec');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1c00b207-f01a-4cfd-9179-723890103f7b', 'b32a6026-82ad-499f-b064-549f5de171c4');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('cf1159fd-8142-4ab3-bbe5-bbb504231f1c', '4704ada0-e0c7-4757-b2fa-8c0da2b33028');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b4bab453-d7c0-46e5-9d35-e17c4ed58be6', '82c158cd-8f28-4bd2-a909-07079c6f5432');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('2c79dd49-4004-41ae-8975-8eb8dfd7286c', '60b2c182-4bfa-4b80-94ee-8c7e70ab0690');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'f59be088-9084-465f-8878-94e49c24fee2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('724b1490-5411-485c-b96e-0a7d553e24b5', '469af9f3-1ca1-451c-aca8-3f51eeb2cbd0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('fc77a8be-dc50-4618-a42f-927f9f0f5131', 'c2a20de7-9745-4730-8663-9183bc9b6094');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('69884407-d5d8-4669-ac7f-5926def2a627', '0e0b0dc8-12e5-4fe2-a6c7-be56728db3a0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('18cfdbe1-bb62-437b-8e36-74155878da89', 'ff290859-b711-4ad0-8db0-ba66f9d56af6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('01c92ba3-6348-4139-8683-82d5a38cc87e', '6ec95b9c-5a45-4dd3-98e7-1459becc86fc');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('994a5f49-b858-4b33-8077-4a4cc9e5fbe6', '469af9f3-1ca1-451c-aca8-3f51eeb2cbd0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'a76ecb9a-1712-4190-996b-0bf7a65de360');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'a8886ce1-681b-4e98-a372-52edb6fa2347');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('90d0f63a-2b94-4ba2-9907-1b981de65101', 'c2a20de7-9745-4730-8663-9183bc9b6094');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('72ecb22b-6449-4daf-addb-e036dc56e5e0', '82c158cd-8f28-4bd2-a909-07079c6f5432');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('555de0c1-d263-4aa8-80ec-0277fec44759', '90be5088-f75c-4759-8d42-ddcae792125e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('227dd093-c6c4-4ebf-82d7-a428e590c8ec', '6a0445f6-63f5-4c62-8593-f5f5996c15db');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0e7a76e2-617e-449d-a098-4188f44356fb', '43291590-b536-425b-88c2-ed51a55b602b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('3c36a789-ea44-40dd-a0f2-514acb85e1f7', '82c158cd-8f28-4bd2-a909-07079c6f5432');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', '299f5f73-c94d-4410-92a2-70cbf91b37ea');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'c2a20de7-9745-4730-8663-9183bc9b6094');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'a966defe-324d-484e-b24f-6ddd1ed5c8f6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('61042683-9841-4b54-a3fc-9d578f38184d', '82c158cd-8f28-4bd2-a909-07079c6f5432');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('18753203-4827-418e-89f5-21ef396f3778', '299f5f73-c94d-4410-92a2-70cbf91b37ea');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'f59be088-9084-465f-8878-94e49c24fee2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('5406bca6-71db-44d1-8d4c-00f94213ed8b', '277e9002-c1ca-4ef0-8bab-f1b7a45789ec');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('9e5297df-3df1-498b-b166-0a39e9df0452', 'a8886ce1-681b-4e98-a372-52edb6fa2347');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a0c70f93-8f26-4507-9ce8-0f64b00d32a4', '2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b1b344a1-d5ed-42e6-821e-4db8a60c516a', '6f7e909c-122c-470b-9f86-ee04ad9db1b3');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('ebb28475-a810-4dcf-8980-45c2a8e931f1', '277e9002-c1ca-4ef0-8bab-f1b7a45789ec');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('8592634e-56ea-4f62-b938-0066953bbf24', '148d017a-37ef-4a21-a030-1725054e26c7');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6841778c-dc4d-4a12-9f23-1f9668314b6d', 'bfc2c4d4-dd15-40a8-83d0-ebcbf09c80e0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('bcad1a4d-3a06-4a61-901e-ce369a384e28', '842f30a1-c95c-45c3-bd75-f3c651c22618');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('40f56c2c-c39d-426e-8496-dba481e4d09e', '90be5088-f75c-4759-8d42-ddcae792125e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('ac695f69-8672-4a1d-9d11-44284ee80009', '085687b3-7b0e-4fcd-bece-89d67aecaaac');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('65b2309d-e638-4473-905f-02beca4b92ac', '2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b0a3abd1-84d7-40ce-a92e-9843041ac7c9', '148d017a-37ef-4a21-a030-1725054e26c7');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('9430bc13-c96c-4bb5-bf9c-61593f96c11e', '2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('866ccf38-e290-45c9-bd80-429456c725e9', '299f5f73-c94d-4410-92a2-70cbf91b37ea');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('528d1d82-430c-4b19-898c-2037c0716c3c', 'b32a6026-82ad-499f-b064-549f5de171c4');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d948d9c2-ece7-4a18-bde9-2202bd92fc24', '2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6b689866-3315-4e55-bbd6-0c61e35df71e', '32602cab-1dfc-4216-b420-b0e52e7e7ef2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('f6d0f388-bf0b-44ef-b689-98c6255ee474', 'fd720fe2-943a-496c-a23b-8e66e1c9c067');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('4b9b57f7-8065-4229-881c-4e0345e10687', '2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7fdafc7d-9501-48b1-893c-287d8168db1a', 'ff290859-b711-4ad0-8db0-ba66f9d56af6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'fd720fe2-943a-496c-a23b-8e66e1c9c067');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('dc701bde-05d8-440e-9603-0f7a5559311e', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b1973ee2-5eeb-4521-a635-67b5b3509a3c', '0e0b0dc8-12e5-4fe2-a6c7-be56728db3a0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('18574170-cec9-410d-8896-c15a852f16c1', 'a8886ce1-681b-4e98-a372-52edb6fa2347');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c1c783a0-eb15-4506-9dfb-996212dd9c77', '43291590-b536-425b-88c2-ed51a55b602b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'b32a6026-82ad-499f-b064-549f5de171c4');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'bfc2c4d4-dd15-40a8-83d0-ebcbf09c80e0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('658b90dc-ac14-41fa-bb4f-45ff67b30d51', '32602cab-1dfc-4216-b420-b0e52e7e7ef2');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('5d94821c-2f49-4f91-8288-ab7006402f90', '6fba5248-381a-4f71-bbe0-9718ed0d5871');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b61a03ec-7bae-4955-b2aa-e5951cecd3ca', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('4a422f05-767f-442f-a3f8-5860fe565cd2', '469af9f3-1ca1-451c-aca8-3f51eeb2cbd0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('677200ad-da55-4a3f-b7ad-2d7ace462e82', 'a966defe-324d-484e-b24f-6ddd1ed5c8f6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'be45fc01-60fe-4747-9a07-4648105b3fcb');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'ff290859-b711-4ad0-8db0-ba66f9d56af6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('6aeb2cd0-c462-486e-adcf-d0bfd74d187e', '82c158cd-8f28-4bd2-a909-07079c6f5432');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1ca369b1-a772-4743-848a-d38d8ed49a79', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('cf129cd4-d4a6-4d14-b65d-d3e410808b90', '6ec95b9c-5a45-4dd3-98e7-1459becc86fc');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('8eba03e4-b4fd-46eb-b565-5e300343f092', 'be45fc01-60fe-4747-9a07-4648105b3fcb');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', '60b2c182-4bfa-4b80-94ee-8c7e70ab0690');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a24ad893-ccd4-48df-8ee2-328f66044b1c', 'ff290859-b711-4ad0-8db0-ba66f9d56af6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('dbf102e1-9b55-4402-84aa-0ea25f59a9e1', '085687b3-7b0e-4fcd-bece-89d67aecaaac');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1af09413-7078-42e0-a250-ab001ec9e57f', 'ff290859-b711-4ad0-8db0-ba66f9d56af6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('ee7d2609-c37e-4bae-abf4-e10338833ea2', 'dafc199d-8315-4cb3-8a8d-e997c71931b1');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('1f8e556d-1779-411c-847e-422eab71ca1f', 'bfc2c4d4-dd15-40a8-83d0-ebcbf09c80e0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b140129e-3a20-4dd2-86b4-cef4499a031d', '277e9002-c1ca-4ef0-8bab-f1b7a45789ec');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('dd3575d9-d937-4ccb-92e7-237491e6a853', '148d017a-37ef-4a21-a030-1725054e26c7');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('b5fa3111-3931-4d9d-88f2-cb609047c0af', '0e0b0dc8-12e5-4fe2-a6c7-be56728db3a0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'c2a20de7-9745-4730-8663-9183bc9b6094');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('d15af0ae-f33d-478d-aea4-d018eb2e4a9b', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('493e5895-17bd-4177-a185-6975a789cc58', '78c370d2-b3dd-4278-ae03-97976ccf192e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('59d649f5-7921-4249-aef9-860748b46b18', '90be5088-f75c-4759-8d42-ddcae792125e');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('095ab469-eefb-450e-8cc4-db0a4aabce45', '085687b3-7b0e-4fcd-bece-89d67aecaaac');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('c274c432-3f84-4e11-936a-b6818ed1542c', '2f3ef674-80fc-4f8e-8dbe-dfb3b91a8f23');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7cf710da-960f-4725-b249-e95c419d8d6d', 'fd720fe2-943a-496c-a23b-8e66e1c9c067');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('168f82bb-cc39-4d4a-b35d-21f555f77308', '469af9f3-1ca1-451c-aca8-3f51eeb2cbd0');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', '8383f6c0-f9b1-4406-ad3c-11aa6651b54b');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('9f03c330-1641-4ae4-bbf4-55462cc35faa', '4704ada0-e0c7-4757-b2fa-8c0da2b33028');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', '57282916-0faf-4644-9d71-df0cd0f8adc7');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('adda4717-f118-45cc-9532-30502637846d', '82c158cd-8f28-4bd2-a909-07079c6f5432');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'ff290859-b711-4ad0-8db0-ba66f9d56af6');
INSERT INTO restaurant_tags_relations (restaurant_id, tag_id) VALUES ('64728b3c-67a7-4107-bac8-f6f998aed794', 'c2a20de7-9745-4730-8663-9183bc9b6094');
INSERT INTO addresses (id, address, user_id) VALUES ('1894fc39-26ef-4580-8ac8-3b21cbd3f92c', 'Улица 56, дом 23', '8b72e7df-64d4-4037-9cfd-e79a06a7563f');
INSERT INTO addresses (id, address, user_id) VALUES ('0654b19d-7828-4ea3-81aa-200c52811eb9', 'Улица 36, дом 7', 'bc20994d-919a-4d57-aae1-50f9da7e8588');
INSERT INTO addresses (id, address, user_id) VALUES ('3468a8d1-18e3-4086-9974-20bd6d3ef065', 'Улица 0, дом 40', '9b414523-d7ce-4c90-b22c-a569ae0351b5');
INSERT INTO addresses (id, address, user_id) VALUES ('eae9a975-28f8-4eb5-85ab-38a62777330b', 'Улица 85, дом 27', '6c632bf4-5a56-4b22-a3b0-029100dae7fa');
INSERT INTO addresses (id, address, user_id) VALUES ('aeb565a0-c731-4206-b3ca-6da283b285a0', 'Улица 98, дом 8', '7ca5d04a-fc86-4716-99d4-26e178b6016b');
INSERT INTO addresses (id, address, user_id) VALUES ('b6a36108-d011-4380-bad4-4644c3476a00', 'Улица 55, дом 42', '594b5f85-a0c0-4b0b-b729-025e233d2c90');
INSERT INTO addresses (id, address, user_id) VALUES ('0ef4bace-f20b-4a77-b068-64399c869444', 'Улица 4, дом 32', '90bbb029-540a-40e6-b17e-448dd8ddb19f');
INSERT INTO addresses (id, address, user_id) VALUES ('dbe62d23-0a42-43ab-bccc-018c74d5d308', 'Улица 12, дом 15', 'a55b7add-a72f-4ee5-910a-77ac6f18891c');
INSERT INTO addresses (id, address, user_id) VALUES ('ece0f559-c56c-47e0-9798-97225f2a8519', 'Улица 13, дом 18', '3419e315-99c0-4fd3-9f0b-39e7877e738e');
INSERT INTO addresses (id, address, user_id) VALUES ('4ab86189-e1e5-474a-bd5f-43ae89a57625', 'Улица 4, дом 18', '5cc101e9-0cc3-41f6-8b56-c5e29de758b1');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a23c1fe8-2be8-4788-baf7-60fb1b9e9fe8', '6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Блюдо 1', 397.77, 435, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('16b5f712-f8c9-4f92-8d6d-dc1f986eed2f', '6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Блюдо 2', 584.67, 329, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('208642f1-fe40-4882-bba3-7a7dc419cc40', '6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Блюдо 3', 129.37, 136, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f3138510-ce53-41c5-a558-41d242cc7b2c', '6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Блюдо 4', 382.43, 269, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('af8e9662-5d71-46da-8c6c-3a3a44d4963b', '6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Блюдо 5', 457.05, 438, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ad530ca9-b269-4843-b4c3-c29bd00982fd', '6be155ad-8705-4cc5-8bbe-1d60b279516c', 'Блюдо 6', 522.25, 129, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d56f4ab4-f6bf-4e30-825b-7481f6421178', '657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Блюдо 1', 419.27, 222, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('54ccd259-2010-47b9-8a32-74fee46ced04', '657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Блюдо 2', 505.05, 473, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('81f34c0b-a3f1-422a-a4b1-3430ccd0dccc', '657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Блюдо 3', 458.32, 495, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9d609456-9532-41d2-8895-1c417cae43d6', '657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Блюдо 4', 233.14, 461, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6aeb1009-b06a-4b52-99ef-713fc9ea131c', '657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Блюдо 5', 154.05, 152, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('020f4d0c-c0f7-48c2-b1cc-25d0e5fc8f1f', '657c3fb1-a388-4ca1-ae15-0b31673056b5', 'Блюдо 6', 123.92, 151, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ced3c06e-d0a2-4789-8921-c86abef3301d', '7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Блюдо 1', 508.34, 188, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a4102d8e-844f-4758-a853-4fd3ca6e5569', '7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Блюдо 2', 358.50, 227, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('80e9b1d2-a0bf-4e3c-a8a9-bb09a7b3f633', '7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Блюдо 3', 329.85, 356, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('617f5fa2-6623-4e46-bb4e-0c97e67fa666', '7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Блюдо 4', 543.34, 426, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('45de2120-a288-4c6d-87dc-47e4e7b436d6', '7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Блюдо 5', 336.27, 160, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('68e81208-eabc-4bd9-848b-f348bc1977cf', '7b27253f-26e1-494c-a54c-9e9e0c1ff903', 'Блюдо 6', 406.51, 244, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33a4d4d1-c902-4ccf-a561-8610c0138c97', '97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Блюдо 1', 233.95, 354, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('35451456-2e49-475f-8198-77935083ff8e', '97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Блюдо 2', 281.80, 166, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('00a2ea10-6be2-4b43-9148-598e4698523d', '97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Блюдо 3', 275.87, 463, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5f16f9c4-08c8-4c11-b119-6bb79f4ca0ec', '97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Блюдо 4', 516.66, 364, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3821f9d4-806b-4726-b002-25ff7034c202', '97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Блюдо 5', 432.08, 105, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5fdd1b8d-ec10-4778-847d-cb62dc895d37', '97358ff0-ccd3-4817-8a9c-673c8875cb02', 'Блюдо 6', 487.01, 123, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f49cc3c1-596b-41a5-89fc-189a8e33648d', '1f216a35-23e2-474d-823a-f8c6f5e25545', 'Блюдо 1', 212.51, 482, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f6131b8f-1175-410f-8d45-6c325413f5d1', '1f216a35-23e2-474d-823a-f8c6f5e25545', 'Блюдо 2', 190.15, 368, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fb81d591-c765-42aa-8031-89d1f3c3a985', '1f216a35-23e2-474d-823a-f8c6f5e25545', 'Блюдо 3', 423.73, 127, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('071888c3-2413-438a-88d6-32fadee78cd4', '1f216a35-23e2-474d-823a-f8c6f5e25545', 'Блюдо 4', 181.80, 128, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6d5c4f33-2334-443c-8eb0-b0f2fb37bf47', '1f216a35-23e2-474d-823a-f8c6f5e25545', 'Блюдо 5', 282.02, 369, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8f128ab4-7a45-417b-ac3b-1b4fbd9d54b1', '1f216a35-23e2-474d-823a-f8c6f5e25545', 'Блюдо 6', 388.46, 311, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9ac6fa86-a459-472e-bd19-c356a2b99a08', '61903dee-b84c-4510-a70f-1c022a87ccfc', 'Блюдо 1', 385.16, 191, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc409491-e36c-4351-90d6-776140efc79b', '61903dee-b84c-4510-a70f-1c022a87ccfc', 'Блюдо 2', 424.23, 336, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('27ee291d-ed54-4c06-bd34-0e51ec7c9deb', '61903dee-b84c-4510-a70f-1c022a87ccfc', 'Блюдо 3', 374.18, 285, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('72722fac-8fea-4c60-84b2-163b52592501', '61903dee-b84c-4510-a70f-1c022a87ccfc', 'Блюдо 4', 240.08, 228, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0ec86fc3-bf6c-43fd-bfb4-8e3405b6b159', '61903dee-b84c-4510-a70f-1c022a87ccfc', 'Блюдо 5', 479.89, 309, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a714a8a5-21a2-4f09-bb02-61673bc0706d', '61903dee-b84c-4510-a70f-1c022a87ccfc', 'Блюдо 6', 205.61, 429, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5ac81189-61d2-421f-ba4b-fa2f34a5b02d', 'ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Блюдо 1', 197.89, 343, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ee94337-b28d-4659-abe4-a1f6695ed8c3', 'ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Блюдо 2', 585.68, 497, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d85d55c3-0634-42de-a326-353ce1d53dca', 'ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Блюдо 3', 576.46, 142, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6823b6df-d410-4871-87f4-01b9e8415541', 'ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Блюдо 4', 266.29, 473, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('28c1d851-67e2-4a68-8832-1e313458743f', 'ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Блюдо 5', 462.79, 190, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8669b91c-087f-4f92-a9e6-d92590bef1c7', 'ad2d23d9-ba59-4533-8ea9-aab6ff0eb8e0', 'Блюдо 6', 262.94, 428, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1058ae8e-6d87-4ece-8842-e93cadc4abf6', '692ba0be-c31e-4525-93e8-e7e6eca57383', 'Блюдо 1', 170.95, 142, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('da50599a-ba54-4bb5-bb3e-1547c46c208f', '692ba0be-c31e-4525-93e8-e7e6eca57383', 'Блюдо 2', 242.28, 207, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4925fb6c-a2c0-4daf-a47a-73a2678f7104', '692ba0be-c31e-4525-93e8-e7e6eca57383', 'Блюдо 3', 231.31, 497, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1559b1e-9101-45d3-8f96-e420549eba12', '692ba0be-c31e-4525-93e8-e7e6eca57383', 'Блюдо 4', 201.36, 468, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddaefca2-69aa-48f7-92cf-18b11605d6f0', '692ba0be-c31e-4525-93e8-e7e6eca57383', 'Блюдо 5', 229.14, 465, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d0f86548-7b0f-492a-99fe-7ec670dc953c', '692ba0be-c31e-4525-93e8-e7e6eca57383', 'Блюдо 6', 579.41, 356, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e4b1fc7a-7850-470e-afb0-c601a3b67ed9', '7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Блюдо 1', 383.51, 463, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('64139e40-c386-4827-8065-6b5253cdc601', '7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Блюдо 2', 597.27, 390, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e7f2c399-6820-4cb0-9c6b-4759ad486e78', '7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Блюдо 3', 242.63, 271, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('72448533-206a-442b-a42a-544aafcd50f6', '7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Блюдо 4', 312.75, 275, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('142c5b6b-5b71-4d94-ab6b-f28be95ee23c', '7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Блюдо 5', 219.31, 267, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('42e0bdfc-03db-4e71-a705-a46e16e7d150', '7bd39ea6-8449-414e-80fb-dcc92a62ba69', 'Блюдо 6', 513.05, 110, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6d29bdb8-c185-42b7-987e-bc141d056a7d', 'a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Блюдо 1', 227.84, 247, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8e8a4063-16da-4528-b231-c72bf65ffd3f', 'a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Блюдо 2', 207.76, 151, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ecfb8f2e-540f-43ce-bef9-979bb81b2d05', 'a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Блюдо 3', 366.48, 244, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ae6fe72a-cbf4-4f42-a791-6e861ab172a1', 'a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Блюдо 4', 568.65, 306, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('874c88e8-6e96-43d1-a989-977169df030f', 'a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Блюдо 5', 574.27, 115, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c6ff55a8-b3e6-439b-9687-e0b6ab2a7ab8', 'a0d25f99-97bd-4232-9fc5-8c9ac098b1f1', 'Блюдо 6', 449.70, 272, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1388fa41-9601-431a-82bd-17313678c254', 'c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Блюдо 1', 112.48, 300, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dda03ccd-b876-4718-a3ea-d299780b29a6', 'c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Блюдо 2', 101.73, 246, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9f2fa4fe-15e1-4482-91ef-23c4fa3798ec', 'c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Блюдо 3', 366.61, 109, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4da36638-f332-434c-b1bf-c709e4ca5660', 'c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Блюдо 4', 246.02, 408, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ea19aa57-dee7-40e1-9dad-e6828f107864', 'c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Блюдо 5', 247.41, 204, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e1b1bd7b-1357-46ac-a47f-7769c38b093e', 'c6a874a0-9fa8-4b5b-a0d7-69caf1e471d8', 'Блюдо 6', 556.55, 360, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('14a466a5-ca9b-4d3b-b4b1-a8c09b3dc60f', 'da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Блюдо 1', 295.63, 202, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c43893ef-8537-4634-bfb3-b0a66b87c603', 'da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Блюдо 2', 497.94, 274, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('72033af4-bb2a-4fe7-87eb-fd944f08e72e', 'da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Блюдо 3', 489.38, 335, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('58cc3d6b-59ac-49c3-a29b-f456c6243095', 'da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Блюдо 4', 140.89, 317, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('227edb60-50b6-4079-9c7f-01ad325226d3', 'da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Блюдо 5', 195.08, 349, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1d372cd8-a233-43d4-984d-c788092d4597', 'da752af1-9e8d-4058-86a4-b0aa3a200f19', 'Блюдо 6', 320.51, 269, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bdc4d038-6c96-440c-9563-78e14cca8f63', '1653a159-dc2e-4885-8d44-48b44fed3f36', 'Блюдо 1', 321.11, 274, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('964b7836-ea24-466e-9150-fbe75595ca06', '1653a159-dc2e-4885-8d44-48b44fed3f36', 'Блюдо 2', 578.64, 209, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7882edd1-81a9-477f-9cb6-e26397466e89', '1653a159-dc2e-4885-8d44-48b44fed3f36', 'Блюдо 3', 434.91, 295, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('49cdffcd-9cfe-4a48-934b-2ce0693b3bf7', '1653a159-dc2e-4885-8d44-48b44fed3f36', 'Блюдо 4', 347.62, 340, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6bb0f9e4-a2fd-415a-96f5-588d38aae861', '1653a159-dc2e-4885-8d44-48b44fed3f36', 'Блюдо 5', 570.83, 190, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a5908f87-c507-403a-9659-f1f105448bba', '1653a159-dc2e-4885-8d44-48b44fed3f36', 'Блюдо 6', 332.49, 397, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('231a35ae-1973-42c2-a2c0-f3bd0c7058a5', '1c00b207-f01a-4cfd-9179-723890103f7b', 'Блюдо 1', 171.56, 114, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0f27fafa-7cc1-4e76-90d8-c32dbca9a039', '1c00b207-f01a-4cfd-9179-723890103f7b', 'Блюдо 2', 189.25, 438, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c704123e-02d5-4f2b-a870-fe56d8397e3a', '1c00b207-f01a-4cfd-9179-723890103f7b', 'Блюдо 3', 138.13, 179, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b062c939-c387-4414-9dab-6d00cff602d9', '1c00b207-f01a-4cfd-9179-723890103f7b', 'Блюдо 4', 531.40, 121, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7f981f55-6953-44e6-85c0-9d63970177de', '1c00b207-f01a-4cfd-9179-723890103f7b', 'Блюдо 5', 499.76, 245, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d3a10712-a014-439d-8e3d-3b4c099432d5', '1c00b207-f01a-4cfd-9179-723890103f7b', 'Блюдо 6', 300.06, 173, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ee307f00-0f57-490d-8d89-aebb398bb33a', 'cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Блюдо 1', 519.50, 334, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('db50403e-520a-4096-97b5-a4ad95895eae', 'cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Блюдо 2', 284.10, 357, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('13202dc8-a01d-4d72-a4dd-c581fdc33e7e', 'cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Блюдо 3', 246.19, 254, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2a6b7f7e-1873-45f1-8a6c-5248b74f0119', 'cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Блюдо 4', 417.08, 146, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0382778b-7a04-4c54-8e70-7d480f10706e', 'cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Блюдо 5', 542.30, 129, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f773a76a-d075-4f5f-8bb4-e4bd56f7560d', 'cf1159fd-8142-4ab3-bbe5-bbb504231f1c', 'Блюдо 6', 522.09, 264, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ab59d2cc-51a1-42e0-a43d-80b6adc79c2d', 'b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Блюдо 1', 383.86, 334, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b2960847-31dc-44a2-8e93-c9e8a5014935', 'b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Блюдо 2', 304.34, 409, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('312bf74b-5514-4c28-91d8-7727afd36c3b', 'b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Блюдо 3', 580.61, 346, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e4a450ba-7af1-47cf-b2ee-ba7015f14912', 'b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Блюдо 4', 211.41, 191, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc91e92f-dfcc-4ba9-807f-2ecafb6ef454', 'b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Блюдо 5', 332.84, 385, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2904f072-6ff4-4eab-b9de-d3a4b03986bc', 'b4bab453-d7c0-46e5-9d35-e17c4ed58be6', 'Блюдо 6', 493.09, 218, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98f91fc7-694b-43a3-a56c-3701eddc081f', '2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Блюдо 1', 116.02, 230, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('146354e8-696c-452e-9db8-664aee65925a', '2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Блюдо 2', 408.17, 401, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33ebcb7a-96c4-49db-94fb-74642453b4ac', '2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Блюдо 3', 288.69, 351, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4afbc279-ec87-4b2a-9a5a-7ecb30f52d4d', '2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Блюдо 4', 439.88, 453, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2489bea9-db30-4469-b598-7d3b525ce0e8', '2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Блюдо 5', 285.62, 456, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3605540c-879b-4593-a19a-63385140ee24', '2c79dd49-4004-41ae-8975-8eb8dfd7286c', 'Блюдо 6', 525.07, 317, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('92e9c5a0-0df0-4869-bf65-15536d778770', '13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Блюдо 1', 498.51, 159, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7e4c4e2e-a391-46ff-a809-f244956849a2', '13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Блюдо 2', 296.18, 280, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98e68e6d-c1fe-4b14-b052-5103c170bc4b', '13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Блюдо 3', 158.74, 426, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e6b76467-aeb0-4f92-9acc-8f444bd52aec', '13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Блюдо 4', 453.97, 377, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('39d2cd7a-7412-4c64-b6dc-a1df27ca524a', '13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Блюдо 5', 440.59, 194, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1682b250-6620-4c1a-b4c4-5d04370d236f', '13d416a8-28f4-48e8-9f75-a9b0034a0f25', 'Блюдо 6', 194.26, 353, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e097d6a9-936b-4090-bbcd-ff7ff990dc1a', '724b1490-5411-485c-b96e-0a7d553e24b5', 'Блюдо 1', 507.31, 156, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('488cd444-7194-4a96-bf41-bf9ea40f6042', '724b1490-5411-485c-b96e-0a7d553e24b5', 'Блюдо 2', 467.42, 352, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('05f3fc62-0dac-42c1-b691-debd55c558a9', '724b1490-5411-485c-b96e-0a7d553e24b5', 'Блюдо 3', 114.49, 291, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('00aaaeaa-5e25-4357-909f-a2b92918faf8', '724b1490-5411-485c-b96e-0a7d553e24b5', 'Блюдо 4', 368.87, 246, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6eb1af1f-5e7d-4041-a150-0117b4d277d6', '724b1490-5411-485c-b96e-0a7d553e24b5', 'Блюдо 5', 238.82, 319, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('305255f2-4290-44d7-9ff1-25dff50646c9', '724b1490-5411-485c-b96e-0a7d553e24b5', 'Блюдо 6', 487.53, 428, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('de443ecb-f7db-4fef-ab98-c3a196b86599', 'fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Блюдо 1', 194.87, 368, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('42735bdd-c090-4cb8-b874-5fe05841a8ea', 'fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Блюдо 2', 319.00, 144, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fcb7427f-8eec-4438-97b3-e6e94cbc2f40', 'fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Блюдо 3', 156.50, 244, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1b8e18d0-b7e2-49c9-80c6-1930d489940d', 'fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Блюдо 4', 404.90, 201, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('07b6e850-e7c3-4c71-8175-4530ef7e753f', 'fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Блюдо 5', 247.05, 109, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d45c9f02-6fc5-4ef8-a504-ed4871c8bc19', 'fc77a8be-dc50-4618-a42f-927f9f0f5131', 'Блюдо 6', 411.17, 163, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ae42b30f-3a5c-4e5c-94ba-3aa3e9aad7f1', '69884407-d5d8-4669-ac7f-5926def2a627', 'Блюдо 1', 188.67, 189, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5f2b1334-7502-4d51-8b2d-04507e6be4d6', '69884407-d5d8-4669-ac7f-5926def2a627', 'Блюдо 2', 595.49, 124, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a104e3eb-a24e-48dd-99c7-96d3a99873e1', '69884407-d5d8-4669-ac7f-5926def2a627', 'Блюдо 3', 465.21, 142, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a91c9361-38ee-4c91-9d5d-6f3bb4388f1b', '69884407-d5d8-4669-ac7f-5926def2a627', 'Блюдо 4', 398.44, 412, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1d89912e-8be7-40d6-b5a8-a3831427ce81', '69884407-d5d8-4669-ac7f-5926def2a627', 'Блюдо 5', 175.44, 382, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8ced67f2-ef18-4d55-b421-4be709382c06', '69884407-d5d8-4669-ac7f-5926def2a627', 'Блюдо 6', 543.88, 131, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1901908d-868c-44b7-a211-f5c819f1c886', '18cfdbe1-bb62-437b-8e36-74155878da89', 'Блюдо 1', 585.33, 292, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('63c32c27-c69c-48c4-99c5-9e6d7a52cde1', '18cfdbe1-bb62-437b-8e36-74155878da89', 'Блюдо 2', 430.58, 341, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('728f3d57-0987-4358-9666-8282c9cff20c', '18cfdbe1-bb62-437b-8e36-74155878da89', 'Блюдо 3', 586.41, 113, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f0414cce-63dd-4f0a-9a68-a3f4fda5f1fa', '18cfdbe1-bb62-437b-8e36-74155878da89', 'Блюдо 4', 406.97, 237, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3868ff4d-67c3-46af-9680-d67fe529b779', '18cfdbe1-bb62-437b-8e36-74155878da89', 'Блюдо 5', 366.81, 183, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0220a0ff-334d-4593-bb4a-25426ba0075b', '18cfdbe1-bb62-437b-8e36-74155878da89', 'Блюдо 6', 318.07, 448, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('44b00889-4f46-47aa-b50e-a98099c843fe', '01c92ba3-6348-4139-8683-82d5a38cc87e', 'Блюдо 1', 440.90, 376, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b93d6ce2-4ed6-4bd2-8dc3-5ef7d4b50243', '01c92ba3-6348-4139-8683-82d5a38cc87e', 'Блюдо 2', 493.08, 297, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ddbb9aed-6295-464e-bc7b-d17c51772b23', '01c92ba3-6348-4139-8683-82d5a38cc87e', 'Блюдо 3', 168.44, 102, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('332184c3-498b-4ce1-93fd-e411c9113d0f', '01c92ba3-6348-4139-8683-82d5a38cc87e', 'Блюдо 4', 510.33, 498, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('79b6e5c3-8e72-4e86-b5ea-83c68aa3c27e', '01c92ba3-6348-4139-8683-82d5a38cc87e', 'Блюдо 5', 144.02, 242, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e6d70fa7-a0e5-4888-ade8-605f2ab7711b', '01c92ba3-6348-4139-8683-82d5a38cc87e', 'Блюдо 6', 183.94, 118, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1d6e4d6a-8372-4a61-b2ed-e8823bfa0fed', '994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Блюдо 1', 220.88, 114, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fbb207f3-0322-48a2-91d0-9694233cebea', '994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Блюдо 2', 515.93, 469, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('41e50a8a-e6c7-4cd7-9f73-2e4cc9ce13b6', '994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Блюдо 3', 105.58, 192, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9c1e08f2-f5c7-482f-b171-f9061f8efec8', '994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Блюдо 4', 150.31, 311, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dd9590e2-d0a9-4dc9-8b33-819b12a8fd4a', '994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Блюдо 5', 337.89, 323, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a0086d3-6349-4766-80cc-1e905fe3859a', '994a5f49-b858-4b33-8077-4a4cc9e5fbe6', 'Блюдо 6', 571.96, 192, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e48553cb-6431-4233-88fa-11ee842ffe94', '65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Блюдо 1', 524.10, 486, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3d93040a-2d76-49e7-8d7b-b8b56cb5ff96', '65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Блюдо 2', 580.96, 420, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d5f7cca3-42bf-480c-8f53-9b11eb3b9bfc', '65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Блюдо 3', 467.00, 177, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c19c1474-7e2a-4a0a-aada-f389dfbbce52', '65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Блюдо 4', 387.79, 392, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b3522e02-8c7f-47c7-b2c0-12ea748bb1a9', '65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Блюдо 5', 231.36, 317, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cd86e59c-63bc-4a8a-b7da-d4471163903d', '65ab8514-4eeb-4621-811f-ae8bc4d9921c', 'Блюдо 6', 364.76, 247, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('29537eee-4a50-492a-b023-1f23cd25c970', 'a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Блюдо 1', 532.87, 311, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('500733b3-4091-4d77-a944-236d6e4b23a7', 'a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Блюдо 2', 161.72, 422, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d5687f99-fbed-4dad-8220-bf8dd835aec2', 'a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Блюдо 3', 165.12, 379, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8b077ae5-78ff-452c-bbc0-c30589d2fc19', 'a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Блюдо 4', 379.42, 271, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5be21c92-b376-4209-9cc3-030329f40658', 'a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Блюдо 5', 203.04, 126, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('63521f2f-12ca-4763-bd61-83fae056e30d', 'a48e3110-7f33-4f17-ac18-04b8c18c87e9', 'Блюдо 6', 252.52, 406, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ef003ee0-d245-4c88-99e1-abfd624aadfa', '90d0f63a-2b94-4ba2-9907-1b981de65101', 'Блюдо 1', 216.94, 294, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9fba8c4f-5e56-4e2b-9fb3-8e950e103477', '90d0f63a-2b94-4ba2-9907-1b981de65101', 'Блюдо 2', 291.73, 476, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('261f1acc-fa56-43a2-a3fa-d436027040f5', '90d0f63a-2b94-4ba2-9907-1b981de65101', 'Блюдо 3', 229.32, 458, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bacc95f8-eee2-4a2b-a8f7-1c9899ab18cd', '90d0f63a-2b94-4ba2-9907-1b981de65101', 'Блюдо 4', 370.26, 449, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a2183cd-ff10-4c24-bba0-67f3e2a6c83d', '90d0f63a-2b94-4ba2-9907-1b981de65101', 'Блюдо 5', 550.94, 174, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('43cd2b85-459a-4cb0-a5a7-e8d68181f65d', '90d0f63a-2b94-4ba2-9907-1b981de65101', 'Блюдо 6', 134.75, 477, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('240c3198-6bc8-40ac-a044-96ea58a68641', '72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Блюдо 1', 303.08, 241, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c06e7220-52b4-4c79-9ee0-746d4dcdea1c', '72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Блюдо 2', 160.97, 349, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('78bd0af9-4678-4704-9051-1d059fa3fde6', '72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Блюдо 3', 175.07, 340, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3528e469-49b1-4eae-836f-59129c921dc9', '72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Блюдо 4', 521.62, 244, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0eddbb66-1737-4828-b4ba-a0db5513d926', '72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Блюдо 5', 430.29, 274, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('49170a74-f4b3-4307-b9a0-eea7a23a6146', '72ecb22b-6449-4daf-addb-e036dc56e5e0', 'Блюдо 6', 381.13, 122, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e0cbb8a5-9dc2-48f8-a4c1-e31feaefa054', '555de0c1-d263-4aa8-80ec-0277fec44759', 'Блюдо 1', 291.81, 166, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('878cd540-41b9-47cd-9cc7-34d87df449ed', '555de0c1-d263-4aa8-80ec-0277fec44759', 'Блюдо 2', 231.71, 479, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4ed75271-ab10-4e4d-a645-6096eaa3dc09', '555de0c1-d263-4aa8-80ec-0277fec44759', 'Блюдо 3', 208.68, 221, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1a7f4e25-397e-4b57-8d48-3144807df2c2', '555de0c1-d263-4aa8-80ec-0277fec44759', 'Блюдо 4', 559.59, 175, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('739d0578-5f3a-4ba5-a0db-c00fef688a59', '555de0c1-d263-4aa8-80ec-0277fec44759', 'Блюдо 5', 544.00, 262, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31335c08-e739-45ba-83a6-c0386da39834', '555de0c1-d263-4aa8-80ec-0277fec44759', 'Блюдо 6', 573.58, 247, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d9081b81-dfa1-4d01-a678-c85124b548fd', '227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Блюдо 1', 516.07, 122, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3ba0012c-dcf2-4723-98f2-3cb83cb00925', '227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Блюдо 2', 537.98, 323, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b9bf423f-327f-40ce-aba1-e2e50fd52f29', '227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Блюдо 3', 439.75, 376, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bd2cbc6e-3e65-4766-bb7e-ea0b717fe708', '227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Блюдо 4', 240.00, 352, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4906fea4-0bdd-4666-9fd3-c53ffa4dbc4c', '227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Блюдо 5', 450.58, 343, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d2684afa-90a0-4689-baa3-d07dab58af37', '227dd093-c6c4-4ebf-82d7-a428e590c8ec', 'Блюдо 6', 500.08, 376, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('926cc11c-eaca-462c-a003-d24061df2a6a', '0e7a76e2-617e-449d-a098-4188f44356fb', 'Блюдо 1', 126.89, 107, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d991728a-87c9-49a4-9afa-13114538abe2', '0e7a76e2-617e-449d-a098-4188f44356fb', 'Блюдо 2', 140.45, 438, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7626e150-428e-4f17-ba2b-624b40b9871f', '0e7a76e2-617e-449d-a098-4188f44356fb', 'Блюдо 3', 586.14, 380, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6cdaf687-5235-4c59-9d16-045faf0abe1d', '0e7a76e2-617e-449d-a098-4188f44356fb', 'Блюдо 4', 247.76, 167, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e5aac9d1-9f7e-465e-ac83-24064437459a', '0e7a76e2-617e-449d-a098-4188f44356fb', 'Блюдо 5', 524.33, 406, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d83f67a4-043f-4d13-9b21-6d4ee2f8f913', '0e7a76e2-617e-449d-a098-4188f44356fb', 'Блюдо 6', 360.20, 371, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c902d756-eb5c-42d8-9054-a81a61540557', '3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Блюдо 1', 315.40, 153, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fa36eacf-7467-41fb-9a94-88ca1fdae42c', '3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Блюдо 2', 405.27, 113, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('639760e8-4f21-4bb9-b403-60eb58e7c497', '3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Блюдо 3', 234.47, 181, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bb612ebf-3faa-4e4c-adb9-b059f93bed84', '3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Блюдо 4', 185.51, 299, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b3cb1d45-b948-4efd-9cc5-7834f911e019', '3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Блюдо 5', 198.81, 146, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('60b466bd-eb58-41c2-be14-2e4f8d6df423', '3c36a789-ea44-40dd-a0f2-514acb85e1f7', 'Блюдо 6', 310.50, 322, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c6f0ac92-7832-4431-b702-d639f44af5fa', '0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Блюдо 1', 533.21, 140, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b923511c-6a0d-4adc-9dac-6c044d7a468c', '0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Блюдо 2', 187.10, 398, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bda391bb-8ba1-4208-b725-dc030894c028', '0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Блюдо 3', 540.21, 306, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0c8a1b73-f3da-4b5e-8472-18259bc23ab1', '0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Блюдо 4', 466.41, 453, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4ad48422-1fa3-4a25-b7c3-9b09ddc19544', '0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Блюдо 5', 281.90, 414, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('368fa16b-5a13-4aa8-ad4b-dc7fae952d41', '0f85a8b2-8144-4ab9-893d-b72f6bcb3e0b', 'Блюдо 6', 139.57, 334, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('910e89a7-bf31-48f1-9d46-3590c94692b7', 'e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Блюдо 1', 489.98, 152, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a386e686-9b68-416b-ab4f-20a8c24bcf32', 'e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Блюдо 2', 292.80, 398, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3d89bd8c-704c-43b7-93de-478182347836', 'e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Блюдо 3', 306.17, 165, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('71ee017d-8763-4977-a998-05f83d39ad15', 'e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Блюдо 4', 393.08, 461, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c5659c95-4821-42bb-b070-50af462458fe', 'e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Блюдо 5', 106.81, 339, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('795de75a-71f9-4aaa-9ab5-5dadfbf799b6', 'e56d72e0-24d8-4da6-9c8e-5a7e3a9f5ad5', 'Блюдо 6', 367.69, 480, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d576ec6b-586e-4c6a-a575-8fa163327682', 'e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Блюдо 1', 595.86, 106, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1df609e7-0645-45a2-9c79-040de83e2732', 'e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Блюдо 2', 142.80, 301, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5afa854e-8b07-4e0c-a855-3502b5e4af15', 'e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Блюдо 3', 252.50, 322, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('28d597f3-aa5b-4636-acd5-d6880bb26bd0', 'e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Блюдо 4', 285.92, 216, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aa424020-a78f-4145-b970-232d86fc7a05', 'e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Блюдо 5', 529.53, 152, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f702b84d-68d7-4623-b786-995fffcb2436', 'e0ac8ae4-495f-4d4f-b2bd-48a0ff7d3411', 'Блюдо 6', 332.62, 350, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('de222178-cc69-4ed7-844c-8bb4f2725573', '61042683-9841-4b54-a3fc-9d578f38184d', 'Блюдо 1', 391.92, 122, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1d704e42-d29b-4833-9772-3facc22fe6e9', '61042683-9841-4b54-a3fc-9d578f38184d', 'Блюдо 2', 261.33, 191, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a40756eb-6dd7-4497-a1bd-c1a146455f76', '61042683-9841-4b54-a3fc-9d578f38184d', 'Блюдо 3', 205.48, 153, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fa24f9ee-c28f-4aba-8c63-ab039b952004', '61042683-9841-4b54-a3fc-9d578f38184d', 'Блюдо 4', 467.44, 487, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c8aec7fd-536b-4aa9-8a6f-c9bc5ffe7cab', '61042683-9841-4b54-a3fc-9d578f38184d', 'Блюдо 5', 397.86, 332, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a2786979-13ec-48a9-ada7-39bd5ece29fe', '61042683-9841-4b54-a3fc-9d578f38184d', 'Блюдо 6', 153.27, 225, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('566b6bf5-6d57-4b65-b157-d269ff02d099', '18753203-4827-418e-89f5-21ef396f3778', 'Блюдо 1', 537.80, 291, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('44860ec6-f2b7-49ad-82e3-ff21311f655d', '18753203-4827-418e-89f5-21ef396f3778', 'Блюдо 2', 513.03, 183, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4de4b523-7d02-4efb-8fe9-66e5f09fa221', '18753203-4827-418e-89f5-21ef396f3778', 'Блюдо 3', 299.65, 176, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('635ed1ba-3eee-4a16-8c30-5ce423da5152', '18753203-4827-418e-89f5-21ef396f3778', 'Блюдо 4', 524.17, 293, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9bcf6311-3c92-4ee2-b2be-e29e94f5729f', '18753203-4827-418e-89f5-21ef396f3778', 'Блюдо 5', 291.41, 198, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('30b5024d-fef7-4338-9747-7109bf87adbc', '18753203-4827-418e-89f5-21ef396f3778', 'Блюдо 6', 589.55, 429, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1025db32-06d7-4df2-a839-346992d1e135', '8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Блюдо 1', 469.75, 425, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9fb9097a-5c6b-41ec-b118-39748b3bd720', '8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Блюдо 2', 266.66, 433, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bff9b1e3-ac22-46f5-a5b3-d73b33be2d9a', '8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Блюдо 3', 232.25, 138, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8df72808-eecd-4b34-8947-79e2ed8f7a33', '8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Блюдо 4', 138.07, 147, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1be37ba9-eb04-4772-91ee-b1b0f1bd6ce9', '8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Блюдо 5', 108.37, 268, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8051a70b-ad80-4d08-abc3-2fe5dab8bfb3', '8617420f-77f0-4139-a8d7-5f6dbf1f1ef8', 'Блюдо 6', 287.23, 225, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8515dd4c-f992-447a-9892-db9080ac2c9b', '5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Блюдо 1', 440.25, 261, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0b9b6beb-7cf4-4a0a-b6c4-b825bcb13503', '5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Блюдо 2', 357.68, 450, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('33f41dfa-821a-4002-a2b6-aab713f3e25e', '5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Блюдо 3', 300.64, 267, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cb151453-9f09-4d05-8c4c-1efbc657f7f7', '5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Блюдо 4', 265.09, 104, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a8a00ff7-99c0-40d5-86b4-49f1f2cdf360', '5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Блюдо 5', 195.69, 432, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6220cf02-4719-42ec-877c-f4b07ef411b9', '5406bca6-71db-44d1-8d4c-00f94213ed8b', 'Блюдо 6', 375.12, 442, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('888a00b9-add1-4ea8-b950-29d3d569475b', '9e5297df-3df1-498b-b166-0a39e9df0452', 'Блюдо 1', 461.90, 437, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ee37732d-3c19-4a25-a940-9f94c210b36f', '9e5297df-3df1-498b-b166-0a39e9df0452', 'Блюдо 2', 497.12, 111, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('03c2566d-2eaf-4c33-824a-1eabb22884a7', '9e5297df-3df1-498b-b166-0a39e9df0452', 'Блюдо 3', 120.54, 201, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b521be09-8318-41f7-831c-677c75d42a13', '9e5297df-3df1-498b-b166-0a39e9df0452', 'Блюдо 4', 228.27, 101, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('47b00609-4051-4db2-96bb-e99030705b10', '9e5297df-3df1-498b-b166-0a39e9df0452', 'Блюдо 5', 532.99, 239, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('68de6701-4526-43e3-8c3e-8a4febf3f2c1', '9e5297df-3df1-498b-b166-0a39e9df0452', 'Блюдо 6', 282.85, 254, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3af2a301-6413-45bb-aa78-89c4411214da', 'a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Блюдо 1', 493.55, 157, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dbd10920-5301-43d8-a77a-e8e7330fab65', 'a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Блюдо 2', 313.36, 162, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f761f064-7c64-4382-849c-1c875e3df88d', 'a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Блюдо 3', 101.78, 242, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4b1f0eeb-80b4-4994-a5bf-f8bedd57c6d7', 'a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Блюдо 4', 482.01, 315, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d1867bda-2d1f-43d5-9716-56ac62045ac9', 'a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Блюдо 5', 325.91, 337, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('81df5f9f-df37-4a07-bacf-839514e21616', 'a0c70f93-8f26-4507-9ce8-0f64b00d32a4', 'Блюдо 6', 293.03, 237, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f474e00-65d9-4abb-9f8c-5af888d55121', 'b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Блюдо 1', 565.40, 410, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8fa7c012-09ef-4147-8cfe-6b629505ae19', 'b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Блюдо 2', 183.48, 151, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e16c683c-0e07-47f6-9ebc-1800fb6588e4', 'b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Блюдо 3', 453.56, 463, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('00f0054d-86a0-4484-8d39-205dcf263686', 'b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Блюдо 4', 212.10, 148, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b4753349-65a6-477c-bcc9-a96f464fed32', 'b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Блюдо 5', 291.74, 122, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('899196f3-c701-4f1e-a63a-c4ab4d627377', 'b1b344a1-d5ed-42e6-821e-4db8a60c516a', 'Блюдо 6', 484.23, 199, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9293f246-8837-47d9-be5b-004c27fb0a98', 'ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Блюдо 1', 139.75, 190, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('707cb1f8-fc64-42b2-9af8-4d96b2d5a64f', 'ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Блюдо 2', 394.30, 100, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('789282a4-3b0f-4694-a25e-ca9420fc2172', 'ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Блюдо 3', 141.45, 206, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b8dd524-7010-49bb-9993-65a898bb0c7a', 'ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Блюдо 4', 311.36, 406, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c049d92a-bc4c-4847-89b2-f59dca9cb2ce', 'ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Блюдо 5', 177.36, 202, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('720c5b27-f0cd-4c50-b9ac-c767dae36045', 'ebb28475-a810-4dcf-8980-45c2a8e931f1', 'Блюдо 6', 525.08, 471, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e25f88c9-b070-4986-84e6-23f44a6a53f8', '8592634e-56ea-4f62-b938-0066953bbf24', 'Блюдо 1', 151.41, 325, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e4803692-32d3-44e0-890b-f5c7172c7950', '8592634e-56ea-4f62-b938-0066953bbf24', 'Блюдо 2', 490.18, 206, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7ebf8eb7-366d-4c3a-9df6-2689614f2873', '8592634e-56ea-4f62-b938-0066953bbf24', 'Блюдо 3', 235.08, 226, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d32e4b42-c4e3-4aba-8e15-3ea6943b9c05', '8592634e-56ea-4f62-b938-0066953bbf24', 'Блюдо 4', 534.20, 148, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0a91b44b-a3a0-4e63-af4f-d0a403630e5c', '8592634e-56ea-4f62-b938-0066953bbf24', 'Блюдо 5', 172.99, 341, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e7c307c7-b771-4316-846e-ca4fb9443bec', '8592634e-56ea-4f62-b938-0066953bbf24', 'Блюдо 6', 366.63, 397, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6fe8674d-5fef-4f8c-97ad-b27889dd9f6f', '6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Блюдо 1', 266.35, 478, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bc990396-e526-413b-82b7-244cfca27a51', '6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Блюдо 2', 141.45, 226, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d60b27ca-2cde-4156-adf2-9fd15b3657e4', '6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Блюдо 3', 342.22, 304, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('25e73c15-f351-47d4-9c21-f58fd53cb888', '6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Блюдо 4', 594.16, 318, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3cc42beb-6d57-4177-83a9-37e435ccb087', '6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Блюдо 5', 168.05, 287, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('db662790-adbf-47d9-ac72-6650439c8233', '6841778c-dc4d-4a12-9f23-1f9668314b6d', 'Блюдо 6', 131.29, 103, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('456f9c28-bbca-4bc3-b010-f74dfbffdb52', 'bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Блюдо 1', 486.38, 308, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('47b881c2-03ab-4426-abb3-298ef0a916ff', 'bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Блюдо 2', 184.97, 396, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d289866c-00a2-4814-a6ff-8c90e74055f6', 'bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Блюдо 3', 137.39, 266, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d5cafed2-c71d-4a13-8cfe-4a5c3ae4f9b7', 'bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Блюдо 4', 176.12, 273, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5d2c9302-2645-43ab-9fc3-2ebc784d1a4e', 'bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Блюдо 5', 395.51, 358, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('682d2661-17b2-4e76-a36a-f2f2d854242e', 'bcad1a4d-3a06-4a61-901e-ce369a384e28', 'Блюдо 6', 573.75, 297, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4ce2347e-bd4e-4e2b-8104-23e17027445a', '40f56c2c-c39d-426e-8496-dba481e4d09e', 'Блюдо 1', 582.19, 462, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9a88c9a2-d2fe-4670-b7cd-32561281457d', '40f56c2c-c39d-426e-8496-dba481e4d09e', 'Блюдо 2', 188.32, 189, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4ee243a5-c4d2-497a-9531-4d54c82e3f92', '40f56c2c-c39d-426e-8496-dba481e4d09e', 'Блюдо 3', 169.29, 350, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73321f70-c16c-4af8-b5a0-213d971b228b', '40f56c2c-c39d-426e-8496-dba481e4d09e', 'Блюдо 4', 262.05, 116, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3d495667-7026-4010-b4e4-422d48c1802e', '40f56c2c-c39d-426e-8496-dba481e4d09e', 'Блюдо 5', 354.87, 220, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('259b61a2-d87b-43b0-8f24-41303c639c38', '40f56c2c-c39d-426e-8496-dba481e4d09e', 'Блюдо 6', 275.37, 271, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b6616677-aa63-4b57-8874-ed549668ef7d', 'ac695f69-8672-4a1d-9d11-44284ee80009', 'Блюдо 1', 418.20, 126, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87df9f3f-a761-47fe-9583-f76dc6580297', 'ac695f69-8672-4a1d-9d11-44284ee80009', 'Блюдо 2', 544.79, 171, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('246a9869-7d71-406c-ab64-b019f7475fcb', 'ac695f69-8672-4a1d-9d11-44284ee80009', 'Блюдо 3', 366.21, 140, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9151917e-3066-422b-9e9f-f65834f8182f', 'ac695f69-8672-4a1d-9d11-44284ee80009', 'Блюдо 4', 210.24, 402, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('84ed2510-956a-4c8d-b57b-65cf942ae912', 'ac695f69-8672-4a1d-9d11-44284ee80009', 'Блюдо 5', 157.34, 458, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('57478c35-7401-4ffa-8bfb-496670688019', 'ac695f69-8672-4a1d-9d11-44284ee80009', 'Блюдо 6', 239.36, 116, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b552275f-2a96-4554-a30f-fe272075bce1', '65b2309d-e638-4473-905f-02beca4b92ac', 'Блюдо 1', 236.63, 261, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01478620-e196-407b-bfdd-ecab1e840190', '65b2309d-e638-4473-905f-02beca4b92ac', 'Блюдо 2', 380.36, 278, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('537ae836-9e81-48fc-bc48-e0f57250dd29', '65b2309d-e638-4473-905f-02beca4b92ac', 'Блюдо 3', 212.65, 444, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('76acd451-acfe-4e61-b7bd-83462d076b54', '65b2309d-e638-4473-905f-02beca4b92ac', 'Блюдо 4', 560.69, 217, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01fb6521-9e1f-425d-9947-40ba45d757d9', '65b2309d-e638-4473-905f-02beca4b92ac', 'Блюдо 5', 542.16, 393, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3ddeb179-ee58-4d7e-87cc-2fc50d0b99ab', '65b2309d-e638-4473-905f-02beca4b92ac', 'Блюдо 6', 481.12, 236, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c66a04d-4052-4689-b43f-ba6ee05348c7', 'b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Блюдо 1', 567.16, 183, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3f5b6529-d641-4288-b4d1-d221d78f33f4', 'b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Блюдо 2', 558.10, 372, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8a61108b-69ee-4c44-94a6-73f62bb2aca8', 'b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Блюдо 3', 139.83, 414, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0c603323-ec75-4fc5-a072-4dfa79262124', 'b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Блюдо 4', 404.71, 311, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4d3ddeba-ed85-4d29-8a67-de43342c6240', 'b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Блюдо 5', 115.87, 302, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f3481f1d-62f4-4fcb-bd89-70815b296153', 'b0a3abd1-84d7-40ce-a92e-9843041ac7c9', 'Блюдо 6', 419.58, 355, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0e79b475-6d7b-41e3-98a4-304c25273db5', '9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Блюдо 1', 151.34, 160, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b29cfb71-20d7-4dd4-88c3-e51c4c5582e1', '9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Блюдо 2', 562.85, 195, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bda85de8-2dbf-4723-94c6-925584995764', '9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Блюдо 3', 196.88, 404, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('21080c4e-3da9-42a6-9794-961cb6bf6789', '9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Блюдо 4', 344.50, 392, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('74b8a577-421a-4093-91e4-06512440055e', '9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Блюдо 5', 572.15, 473, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a31d772-2168-4681-9a71-4d9f92f825b7', '9430bc13-c96c-4bb5-bf9c-61593f96c11e', 'Блюдо 6', 263.47, 123, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01a38879-d5ee-4568-9f5e-f5880862b510', '866ccf38-e290-45c9-bd80-429456c725e9', 'Блюдо 1', 261.71, 446, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('acd4f904-8e22-47ab-aae2-86436761b6ab', '866ccf38-e290-45c9-bd80-429456c725e9', 'Блюдо 2', 372.21, 307, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2da2959a-ab5a-48ac-8321-3b9af8c51bab', '866ccf38-e290-45c9-bd80-429456c725e9', 'Блюдо 3', 393.77, 374, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c33f8fde-4b33-4dec-8fca-9196c2c1cfc6', '866ccf38-e290-45c9-bd80-429456c725e9', 'Блюдо 4', 219.19, 116, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7694a94f-5267-4999-b713-4de8f3a656cf', '866ccf38-e290-45c9-bd80-429456c725e9', 'Блюдо 5', 479.56, 281, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3999f592-5d52-4233-a77c-70430cad9c4d', '866ccf38-e290-45c9-bd80-429456c725e9', 'Блюдо 6', 464.86, 129, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0a9c590e-73bb-47b0-8766-ea3dd2967c48', '528d1d82-430c-4b19-898c-2037c0716c3c', 'Блюдо 1', 496.77, 421, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aad70576-2ffb-4023-9b8a-96bc15d7df6c', '528d1d82-430c-4b19-898c-2037c0716c3c', 'Блюдо 2', 590.32, 439, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('87ce8c9e-87dd-4a36-8ad0-50c62fb152a5', '528d1d82-430c-4b19-898c-2037c0716c3c', 'Блюдо 3', 332.47, 223, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a85c1436-9e25-4028-871a-c402e69f54a2', '528d1d82-430c-4b19-898c-2037c0716c3c', 'Блюдо 4', 497.12, 207, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('af86afc5-1e7c-4f24-ae4d-ee3bb8985fe5', '528d1d82-430c-4b19-898c-2037c0716c3c', 'Блюдо 5', 311.90, 248, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('32a47891-5e74-484f-a85c-9f2acbffb3ea', '528d1d82-430c-4b19-898c-2037c0716c3c', 'Блюдо 6', 595.46, 345, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2853bcfe-6f13-40cf-a725-7633466f3ab8', 'd948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Блюдо 1', 423.10, 254, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d4433dff-bef8-4edf-b711-7ebaaf8ffca8', 'd948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Блюдо 2', 453.15, 165, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8a133ad5-6b90-46c1-85d5-33ccb34ea02d', 'd948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Блюдо 3', 231.39, 252, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('452ae20d-8022-4cc6-bf75-dad0a264f705', 'd948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Блюдо 4', 130.56, 112, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b322cea5-e731-40bb-8c9a-321bda39e5bf', 'd948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Блюдо 5', 491.85, 419, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c1ec3643-6346-4d18-b6dd-c6fd74b7c31b', 'd948d9c2-ece7-4a18-bde9-2202bd92fc24', 'Блюдо 6', 225.45, 380, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('60c06005-82c4-45d7-bd80-aa67ed843c1a', '6b689866-3315-4e55-bbd6-0c61e35df71e', 'Блюдо 1', 578.06, 110, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('88391889-bcc0-4c84-b4f6-268a1a3e22df', '6b689866-3315-4e55-bbd6-0c61e35df71e', 'Блюдо 2', 502.32, 124, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('381f7c60-01f6-43fc-bc1a-b8d8087b85b4', '6b689866-3315-4e55-bbd6-0c61e35df71e', 'Блюдо 3', 230.25, 487, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('05825fe2-4c8c-4a76-b6cf-bfb593bcb77b', '6b689866-3315-4e55-bbd6-0c61e35df71e', 'Блюдо 4', 418.49, 112, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c61d36b0-34ee-49b9-bca1-040efc69e875', '6b689866-3315-4e55-bbd6-0c61e35df71e', 'Блюдо 5', 198.56, 199, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b378708b-e2f3-4066-a920-47676f755ad9', '6b689866-3315-4e55-bbd6-0c61e35df71e', 'Блюдо 6', 353.59, 272, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e18e3855-7028-4cb9-bc27-203f357d4f75', 'f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Блюдо 1', 151.47, 279, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a670bca8-1c0f-497d-861e-aeebf3936d67', 'f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Блюдо 2', 147.17, 488, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('23042881-db27-436b-9555-364cf0402a53', 'f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Блюдо 3', 438.14, 285, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9583e88b-1342-448e-acc5-eba2011e858d', 'f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Блюдо 4', 399.02, 190, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5e7706a5-9c57-4172-a30e-66ae15f8b8c2', 'f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Блюдо 5', 164.28, 466, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0c0ab0fa-aab0-4b88-9c1a-62ffcb4c81c2', 'f6d0f388-bf0b-44ef-b689-98c6255ee474', 'Блюдо 6', 263.58, 202, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e07f646e-d276-42a0-a1aa-f8f892a0366a', '4b9b57f7-8065-4229-881c-4e0345e10687', 'Блюдо 1', 468.17, 246, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e0802346-b976-498b-89b9-5554704e3ef1', '4b9b57f7-8065-4229-881c-4e0345e10687', 'Блюдо 2', 291.93, 268, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0827dfca-2a51-43ee-8706-8d71bbdd202e', '4b9b57f7-8065-4229-881c-4e0345e10687', 'Блюдо 3', 144.64, 320, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e113aa38-b5c9-494a-9e79-863eeef83ac7', '4b9b57f7-8065-4229-881c-4e0345e10687', 'Блюдо 4', 552.23, 208, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6c91b521-eed3-4962-94d9-c8d82a7fe25f', '4b9b57f7-8065-4229-881c-4e0345e10687', 'Блюдо 5', 357.56, 465, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('35043468-ca72-40cf-bf8b-2fa322c2e7a7', '4b9b57f7-8065-4229-881c-4e0345e10687', 'Блюдо 6', 583.58, 236, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ec5d53b6-3b39-4bb7-a3fa-b66e18f827ab', '7fdafc7d-9501-48b1-893c-287d8168db1a', 'Блюдо 1', 347.85, 411, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1bf8b428-db95-47d0-8872-e5981b51df7d', '7fdafc7d-9501-48b1-893c-287d8168db1a', 'Блюдо 2', 398.42, 437, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('db5442c5-e9ec-4966-885f-8c2b9f584160', '7fdafc7d-9501-48b1-893c-287d8168db1a', 'Блюдо 3', 561.99, 237, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e6814d62-5b10-4bca-87ed-ebabca06e593', '7fdafc7d-9501-48b1-893c-287d8168db1a', 'Блюдо 4', 341.91, 108, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('08205bb2-cdf5-40d6-9dbe-78b5af824663', '7fdafc7d-9501-48b1-893c-287d8168db1a', 'Блюдо 5', 316.81, 271, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ada4873d-23b6-4222-9896-bfa226b864be', '7fdafc7d-9501-48b1-893c-287d8168db1a', 'Блюдо 6', 195.04, 298, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('618c87f7-d35b-430b-bb06-eb0b46105f73', '12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Блюдо 1', 261.92, 226, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b14cbf58-f84c-443c-8fd0-9f77ec428d82', '12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Блюдо 2', 205.85, 199, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0aa851a6-0a08-4528-9271-1165f3d38e92', '12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Блюдо 3', 345.97, 140, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b9d6b353-a605-4542-bc2f-80b70d264bce', '12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Блюдо 4', 164.82, 210, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8f0e6c13-d2cb-4e36-899f-bff7676e588e', '12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Блюдо 5', 195.19, 457, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6d0b5615-da2e-40f7-bcae-380d17490d35', '12f44a19-be2d-4ae6-82f4-a9c169d4a42b', 'Блюдо 6', 392.76, 243, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e10138c3-fd74-4db7-8bd7-cb39752ab2af', 'dc701bde-05d8-440e-9603-0f7a5559311e', 'Блюдо 1', 207.99, 341, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d9b86d44-1308-4e2c-b7e4-6c66dc637650', 'dc701bde-05d8-440e-9603-0f7a5559311e', 'Блюдо 2', 559.11, 346, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cba9b57d-d20f-4eea-aeec-6338a0779414', 'dc701bde-05d8-440e-9603-0f7a5559311e', 'Блюдо 3', 274.59, 242, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d3f03d3c-d3e1-456c-bbcd-bafcef3325b8', 'dc701bde-05d8-440e-9603-0f7a5559311e', 'Блюдо 4', 445.51, 355, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1f885d47-c38a-49f2-9505-f612529abec9', 'dc701bde-05d8-440e-9603-0f7a5559311e', 'Блюдо 5', 508.84, 274, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ee30bac4-ef1d-4a5e-a768-7f8ff8c106fd', 'dc701bde-05d8-440e-9603-0f7a5559311e', 'Блюдо 6', 357.95, 231, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1f75e65a-0d9e-4a16-8a12-c7245087fc9f', 'b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Блюдо 1', 380.50, 286, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f0d67fdc-f256-439c-bd52-a97fe8b45a94', 'b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Блюдо 2', 251.81, 473, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d12aca1b-d9f7-420b-9188-f18fe5d9b3e9', 'b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Блюдо 3', 438.25, 295, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('133e6e17-7b06-430a-a156-058e5fc4fcde', 'b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Блюдо 4', 527.43, 458, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0136c4cf-49d2-44ce-9d7a-726c508ca42c', 'b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Блюдо 5', 397.48, 429, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d0e184fb-7cff-4311-bb53-d33443dbc09c', 'b1973ee2-5eeb-4521-a635-67b5b3509a3c', 'Блюдо 6', 393.50, 160, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('510203f8-6aa1-4b3a-9b31-4797ee493581', '18574170-cec9-410d-8896-c15a852f16c1', 'Блюдо 1', 294.57, 208, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a6088f9-603a-4ade-9bac-e6c1c7b9b4f4', '18574170-cec9-410d-8896-c15a852f16c1', 'Блюдо 2', 548.33, 274, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('22ab5562-19c7-4f9a-ac40-cb1c7863a54e', '18574170-cec9-410d-8896-c15a852f16c1', 'Блюдо 3', 538.53, 158, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('03fcdf22-c9ab-447a-bc5d-c03802baae6a', '18574170-cec9-410d-8896-c15a852f16c1', 'Блюдо 4', 104.02, 480, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e6164e42-97c6-4bf4-b4d5-ebe0ec7d0f4e', '18574170-cec9-410d-8896-c15a852f16c1', 'Блюдо 5', 281.78, 480, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c30fd44-a515-4016-b4ff-4e2dc5c89ffa', '18574170-cec9-410d-8896-c15a852f16c1', 'Блюдо 6', 213.84, 271, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2e395108-7366-43d0-a005-1d261adff3b5', 'c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Блюдо 1', 197.31, 328, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('117c73d4-a75e-44c3-a7c5-c1ac8c5662b7', 'c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Блюдо 2', 343.39, 111, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3fa2a1c3-7c82-428c-8a77-1e01bea485d4', 'c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Блюдо 3', 407.00, 350, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b27e0c98-d319-4658-ba18-161fc27043ae', 'c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Блюдо 4', 268.25, 277, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5df99e04-3762-42e9-b44d-0796412f881d', 'c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Блюдо 5', 421.04, 477, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a01119d-b234-422c-9a88-763c56dabd8e', 'c1c783a0-eb15-4506-9dfb-996212dd9c77', 'Блюдо 6', 209.49, 497, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6f963d17-68f0-4ddf-909d-21b914d2980d', 'be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Блюдо 1', 244.51, 100, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4c3c47cd-e729-4b41-a41c-36011c28621c', 'be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Блюдо 2', 109.73, 494, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('36a6660b-6ca5-4280-8740-957b06b05a62', 'be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Блюдо 3', 104.85, 396, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a2727dba-2d1e-459b-8abb-bdd74c161fa1', 'be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Блюдо 4', 384.66, 308, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6f2da302-17b2-4be1-a77e-346b76403d86', 'be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Блюдо 5', 371.12, 202, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0f1d2dc9-7b28-4c69-a6cd-1246117959a0', 'be0a7fc5-c2b2-447e-aebf-48e2639bdc48', 'Блюдо 6', 518.89, 158, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b5b8f822-e107-49ce-ba7f-f8e34402e7e3', 'defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Блюдо 1', 106.34, 288, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('53ad3bb4-9d8d-407f-b15c-a33424aa3733', 'defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Блюдо 2', 126.13, 266, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6f2a6d8c-8931-42e9-8d2b-bad2bc4b6e5b', 'defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Блюдо 3', 487.75, 140, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('94d80228-7d59-43b0-857e-93efd5393250', 'defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Блюдо 4', 401.72, 192, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('68569826-5032-4be9-a7b0-ff0671b76209', 'defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Блюдо 5', 263.64, 404, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1880ecf4-50fe-45c8-9b25-a4184908196a', 'defdccc7-01d5-4215-a6bd-b78d4bb29fc8', 'Блюдо 6', 578.80, 304, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c6c470b5-1516-45cd-9fde-20c16b6a297d', '658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Блюдо 1', 529.73, 192, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b411e3c-71d8-4890-add3-039c945377f1', '658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Блюдо 2', 596.17, 322, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5f2e625c-9909-4742-8a45-121af8b89fe5', '658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Блюдо 3', 251.58, 337, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dc644c05-7be0-4969-bccc-db0ad52ba21f', '658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Блюдо 4', 583.66, 349, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e07efcd7-8674-44ce-a0e8-f009d84a8037', '658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Блюдо 5', 108.72, 334, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('da672838-ebee-4c97-afb7-92d2cecaa922', '658b90dc-ac14-41fa-bb4f-45ff67b30d51', 'Блюдо 6', 388.55, 424, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a9468e4c-c789-4f39-af32-8a8f6554b91b', '5d94821c-2f49-4f91-8288-ab7006402f90', 'Блюдо 1', 521.96, 414, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('72a7a398-a823-43dc-9e6a-1188d35aff98', '5d94821c-2f49-4f91-8288-ab7006402f90', 'Блюдо 2', 239.48, 441, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f159a91f-c425-4683-bbe5-3da144ca80c0', '5d94821c-2f49-4f91-8288-ab7006402f90', 'Блюдо 3', 260.68, 347, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c561ed42-f2d7-46ba-aaf3-5d6dd5314261', '5d94821c-2f49-4f91-8288-ab7006402f90', 'Блюдо 4', 175.16, 365, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d3f7150e-9b46-4a47-a601-594a66f408a8', '5d94821c-2f49-4f91-8288-ab7006402f90', 'Блюдо 5', 531.80, 136, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fdddf611-7fb4-47fd-82de-fb2aa39c6310', '5d94821c-2f49-4f91-8288-ab7006402f90', 'Блюдо 6', 472.21, 128, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9ac4c3c5-b2dc-43a7-aeee-1ce77baefed7', 'b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Блюдо 1', 597.15, 293, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bba310fd-bbc3-426b-ad61-e247468817c4', 'b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Блюдо 2', 355.66, 375, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ae0b8c3e-5cae-4d1f-a9d0-b933511fb03c', 'b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Блюдо 3', 430.81, 375, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3e5a7316-6b7e-4138-b93d-6e3332d9106c', 'b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Блюдо 4', 139.49, 200, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('359edffe-789d-49d7-82cc-9864c9131400', 'b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Блюдо 5', 270.94, 234, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('739e03fb-5341-4201-9d59-0ebb998edbbd', 'b61a03ec-7bae-4955-b2aa-e5951cecd3ca', 'Блюдо 6', 594.68, 107, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2f0efb78-1887-46ee-8192-0ef217e0236e', '4a422f05-767f-442f-a3f8-5860fe565cd2', 'Блюдо 1', 340.22, 471, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bed905da-c469-472c-99dc-c6b6c321801c', '4a422f05-767f-442f-a3f8-5860fe565cd2', 'Блюдо 2', 194.66, 158, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6748e2af-99ca-409d-87ae-7fc44ce8892d', '4a422f05-767f-442f-a3f8-5860fe565cd2', 'Блюдо 3', 242.84, 105, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('369d8dc6-f831-4bce-aa98-9ce90b838926', '4a422f05-767f-442f-a3f8-5860fe565cd2', 'Блюдо 4', 163.34, 390, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9f52df8f-de79-4a1e-9ca5-45a2149ab6ca', '4a422f05-767f-442f-a3f8-5860fe565cd2', 'Блюдо 5', 390.94, 391, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('94b07916-b2d5-4c14-acfe-3fedd5039a85', '4a422f05-767f-442f-a3f8-5860fe565cd2', 'Блюдо 6', 292.45, 248, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8bfeea54-e79d-4bf6-a9a2-553fceb6dffd', '677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Блюдо 1', 548.87, 216, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fa18bb96-591c-4acd-8404-947188c5de30', '677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Блюдо 2', 509.94, 152, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('38873d12-80f9-4ba6-8e55-c1f5ea278b24', '677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Блюдо 3', 580.47, 311, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eb6bc889-5c09-4a1f-a708-6fea9b106ddd', '677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Блюдо 4', 171.80, 261, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ca469d27-5f15-4df0-9b99-b6af77639d48', '677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Блюдо 5', 146.06, 494, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d4ceb3aa-6a30-4c2d-ab82-5038661bdc29', '677200ad-da55-4a3f-b7ad-2d7ace462e82', 'Блюдо 6', 520.08, 428, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1246aabc-90a3-46a3-b9d8-3396bc033ee4', 'e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Блюдо 1', 176.79, 409, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('874bbe6b-7acd-4011-93a0-d10194e3f467', 'e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Блюдо 2', 495.89, 377, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3815d824-f5ad-43be-81c1-61708ed3eece', 'e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Блюдо 3', 596.21, 417, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31755e40-a8b7-4327-9b35-7417e1ebac88', 'e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Блюдо 4', 143.69, 100, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('dea099cf-29a1-46f4-8872-628680e0c976', 'e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Блюдо 5', 150.86, 418, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('95abe246-4cfa-4a07-a7a3-cacc28c38f3b', 'e52d1bda-6daf-4c88-9122-f1d16e52b5c8', 'Блюдо 6', 175.17, 424, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b61f9792-761b-4ed5-a7c4-1ad77dc66760', 'b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Блюдо 1', 132.57, 427, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c3934a4c-a684-4e4c-ae55-b1de954ee5b8', 'b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Блюдо 2', 225.87, 355, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c8aa6ddc-a024-4f30-8dea-c153e1b3f329', 'b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Блюдо 3', 171.73, 145, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0a26b796-cbaf-4405-a5ed-a3b14747aa44', 'b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Блюдо 4', 395.99, 236, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9e7b3e3e-858f-4d96-9a71-60e65a364bea', 'b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Блюдо 5', 252.55, 371, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d7c037cb-ae3c-408b-ae1e-fbf835c3a150', 'b667c3e9-dc05-438f-99bf-9bc4a08b98ce', 'Блюдо 6', 432.16, 236, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('976f19dd-e8ad-4716-ac53-5c96e67c21e0', '6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Блюдо 1', 463.52, 407, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0285b6f2-392d-4bd6-9671-75fba1d03bb1', '6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Блюдо 2', 204.05, 486, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bb27c647-e6f7-4f83-a93a-e825b8767629', '6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Блюдо 3', 129.34, 197, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e8f521bf-25e0-4529-92b9-76a04a05993d', '6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Блюдо 4', 458.64, 135, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d812fcf2-b0d6-4ca7-bc5a-09c6557d961d', '6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Блюдо 5', 500.31, 222, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31fd80f9-cef1-4268-ab38-7b5c567f9630', '6aeb2cd0-c462-486e-adcf-d0bfd74d187e', 'Блюдо 6', 297.35, 219, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('03ba3d25-68b6-49d2-aae6-c0fe919ae0d1', '1ca369b1-a772-4743-848a-d38d8ed49a79', 'Блюдо 1', 299.84, 200, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('32337e5a-2dbb-49f7-a157-eb0dcbeeb766', '1ca369b1-a772-4743-848a-d38d8ed49a79', 'Блюдо 2', 392.02, 295, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cce26c21-3e4d-4e31-b9eb-8146b7e77724', '1ca369b1-a772-4743-848a-d38d8ed49a79', 'Блюдо 3', 349.73, 208, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('852cabef-0f69-40f2-81c5-e8c6e38bd11b', '1ca369b1-a772-4743-848a-d38d8ed49a79', 'Блюдо 4', 395.86, 227, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c66f2a6-5d84-4cff-bfec-6200aa4f72ba', '1ca369b1-a772-4743-848a-d38d8ed49a79', 'Блюдо 5', 307.17, 172, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a00409f1-01fd-4b23-9f25-8e829d341b16', '1ca369b1-a772-4743-848a-d38d8ed49a79', 'Блюдо 6', 553.93, 415, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('934486e8-fac1-47a4-9eb5-cb5fac510161', 'cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Блюдо 1', 561.73, 179, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b7de2ca-12fb-45bc-97d9-531bf8ecfa9a', 'cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Блюдо 2', 120.10, 220, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5b957b46-ef94-4fe5-8c56-2529697c2130', 'cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Блюдо 3', 125.21, 131, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9063f0c1-49a9-4491-afca-a4b7c86f52bd', 'cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Блюдо 4', 547.93, 346, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6a57c2ca-a1ea-41d4-b011-e67ca67e9a8b', 'cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Блюдо 5', 563.94, 429, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('26326116-2b3c-42b4-ad5c-c4e98c10990b', 'cf129cd4-d4a6-4d14-b65d-d3e410808b90', 'Блюдо 6', 259.45, 238, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a1709df6-7785-4339-a9de-61c88d58f7d5', '8eba03e4-b4fd-46eb-b565-5e300343f092', 'Блюдо 1', 131.60, 337, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6bb39e0d-9b7c-4de2-b2ad-98e8c06a77bf', '8eba03e4-b4fd-46eb-b565-5e300343f092', 'Блюдо 2', 427.37, 254, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b3cee152-7098-4b6f-b6ad-a04ea3dfbe6c', '8eba03e4-b4fd-46eb-b565-5e300343f092', 'Блюдо 3', 221.25, 149, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('47b0a319-e371-476d-a33a-9fe2403c2c4c', '8eba03e4-b4fd-46eb-b565-5e300343f092', 'Блюдо 4', 122.88, 485, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aa1c12b2-8cd4-44ba-8314-8938c193cc70', '8eba03e4-b4fd-46eb-b565-5e300343f092', 'Блюдо 5', 466.58, 135, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e308a8fa-d8d8-4521-ba34-0fe64b3d9fc7', '8eba03e4-b4fd-46eb-b565-5e300343f092', 'Блюдо 6', 587.99, 356, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8844ad3d-7127-498c-a441-b57ab92db41f', '2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Блюдо 1', 490.44, 380, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('555f1b44-fd14-40bc-a162-56a5c4b1ed23', '2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Блюдо 2', 418.92, 148, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('046a8fa0-a98d-4419-a0f9-a7b5dc5c4c8d', '2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Блюдо 3', 422.52, 329, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('13c597cf-b3d3-4910-8efd-35dcaa013bc4', '2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Блюдо 4', 232.83, 348, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('adf6941f-5e12-470a-8331-51ca0d329437', '2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Блюдо 5', 592.42, 486, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3b502c12-d3fb-4dfe-b58c-83ac0cc71c07', '2e41c69b-455d-4ee6-a0da-0c7222b2f5ad', 'Блюдо 6', 297.60, 322, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('05aa4067-65bb-406a-8ca7-80d46ca3fceb', 'a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Блюдо 1', 152.55, 206, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b1b08889-96cd-41ae-8788-014ac01e8515', 'a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Блюдо 2', 324.66, 495, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('924dc4fb-5976-4f0b-a519-942e831ddcbe', 'a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Блюдо 3', 293.97, 242, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a4d08ff1-fe15-4597-acdd-eaeb6f0aaebf', 'a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Блюдо 4', 441.31, 284, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('adf87dc2-6c4a-4873-a243-9acff6ce220d', 'a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Блюдо 5', 557.57, 123, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('47050f52-da24-4758-8f85-eb62217b070a', 'a24ad893-ccd4-48df-8ee2-328f66044b1c', 'Блюдо 6', 175.30, 243, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a06c0a3-c2b0-4983-91ab-fd03400316ff', '018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Блюдо 1', 168.18, 408, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7f3df326-9a7b-4872-85a5-6dab52a1a98b', '018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Блюдо 2', 527.45, 185, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('73692e4d-3c1f-44fa-b9e5-a1caaad37948', '018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Блюдо 3', 253.18, 464, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('efb5bc2e-7b3b-46ac-932a-c409b1a6f4c8', '018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Блюдо 4', 410.74, 245, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b94dc661-a093-41e2-a14d-df4996b419ac', '018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Блюдо 5', 276.28, 302, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c9c8a765-860a-4d5d-b710-e41efbfe031d', '018a67ff-1b2b-40f5-aaeb-a5e2e6cc66fe', 'Блюдо 6', 433.79, 120, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0c251992-f2ef-4451-899d-21a6fb19c6df', 'dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Блюдо 1', 281.78, 492, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a748a447-3de7-41f8-a537-8c0b07aa6e3e', 'dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Блюдо 2', 333.63, 267, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6fcd27a7-7a08-4b7b-bfc2-1122864b584d', 'dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Блюдо 3', 575.25, 445, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aeff070f-b524-4209-bc64-d07d8ac87c60', 'dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Блюдо 4', 381.48, 407, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d1cb1204-98cf-44b2-8717-9bb46aeecfe5', 'dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Блюдо 5', 535.30, 227, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4012edbc-99bf-4cdb-85f5-7e0d07d74245', 'dbf102e1-9b55-4402-84aa-0ea25f59a9e1', 'Блюдо 6', 359.49, 115, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b4529b2e-db85-4739-b05e-76319ac1cdb3', '1af09413-7078-42e0-a250-ab001ec9e57f', 'Блюдо 1', 350.73, 339, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1f4618f6-c329-4484-95b0-4da06542b732', '1af09413-7078-42e0-a250-ab001ec9e57f', 'Блюдо 2', 291.77, 140, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc2ad316-e790-43dc-850b-4776d001ecfe', '1af09413-7078-42e0-a250-ab001ec9e57f', 'Блюдо 3', 293.95, 436, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3381ee27-72cf-4fd4-8923-65fce220c96a', '1af09413-7078-42e0-a250-ab001ec9e57f', 'Блюдо 4', 110.47, 483, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5a4e5e5e-226c-4fd6-af8b-ad883b5c363f', '1af09413-7078-42e0-a250-ab001ec9e57f', 'Блюдо 5', 403.68, 167, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3a741ef9-2fc6-4fcc-904b-e1fd2ea0aa34', '1af09413-7078-42e0-a250-ab001ec9e57f', 'Блюдо 6', 531.05, 480, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bb207a10-df22-4d76-9546-8d800445ead3', 'ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Блюдо 1', 499.61, 156, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3760d1b5-903d-4051-8091-305ab66f1e78', 'ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Блюдо 2', 540.38, 203, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1270e22e-af8b-4ca9-9116-b05b5a66b672', 'ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Блюдо 3', 397.43, 150, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2099bcd4-58b8-4756-90f1-cce7ae9c2305', 'ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Блюдо 4', 229.24, 355, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e8abd2d6-0116-4b6b-8090-dc639efb4035', 'ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Блюдо 5', 479.16, 169, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('17273a47-0703-42a7-a14e-216ec5d9459d', 'ee7d2609-c37e-4bae-abf4-e10338833ea2', 'Блюдо 6', 549.53, 234, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('767f078d-b711-40ae-8209-eed989d3c494', '1f8e556d-1779-411c-847e-422eab71ca1f', 'Блюдо 1', 204.16, 315, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c78135db-796f-4248-b9ec-0f3e46a52057', '1f8e556d-1779-411c-847e-422eab71ca1f', 'Блюдо 2', 599.02, 343, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31902775-8312-40f7-9776-d4367f4d25cb', '1f8e556d-1779-411c-847e-422eab71ca1f', 'Блюдо 3', 437.46, 256, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c73f7ae3-dcba-4763-8aa8-e5d9040797b4', '1f8e556d-1779-411c-847e-422eab71ca1f', 'Блюдо 4', 422.00, 180, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98fb7f05-7d37-4a13-bf92-1de807f1e42a', '1f8e556d-1779-411c-847e-422eab71ca1f', 'Блюдо 5', 363.63, 260, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3cc00229-6a35-4b43-9477-5f1e51d80284', '1f8e556d-1779-411c-847e-422eab71ca1f', 'Блюдо 6', 384.68, 121, 'Курица');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('89b69870-0461-49e2-afa8-a681dfa162e2', 'b140129e-3a20-4dd2-86b4-cef4499a031d', 'Блюдо 1', 548.57, 278, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('703d9b63-8f75-4149-85c6-b425196d3c70', 'b140129e-3a20-4dd2-86b4-cef4499a031d', 'Блюдо 2', 418.82, 150, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('077e8371-9e1d-470f-afd4-ce727d9fa7fe', 'b140129e-3a20-4dd2-86b4-cef4499a031d', 'Блюдо 3', 156.12, 105, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e050cedd-bab9-476d-be7f-22c3596b87bd', 'b140129e-3a20-4dd2-86b4-cef4499a031d', 'Блюдо 4', 324.06, 160, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7622fb91-ed2b-47f6-a179-c79369014cc4', 'b140129e-3a20-4dd2-86b4-cef4499a031d', 'Блюдо 5', 517.43, 427, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f31c0008-b829-41f7-98a3-82046fd75062', 'b140129e-3a20-4dd2-86b4-cef4499a031d', 'Блюдо 6', 317.08, 368, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2c981812-b03b-4b16-a741-3aaac33fc035', 'dd3575d9-d937-4ccb-92e7-237491e6a853', 'Блюдо 1', 130.14, 173, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('195f552c-6d35-4c79-b3dd-d97a7e37b97d', 'dd3575d9-d937-4ccb-92e7-237491e6a853', 'Блюдо 2', 347.20, 305, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aee64045-8973-46fe-9e3a-5c4711a9e072', 'dd3575d9-d937-4ccb-92e7-237491e6a853', 'Блюдо 3', 307.61, 309, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('68bd09b9-80db-46d8-81ea-832240a1dd83', 'dd3575d9-d937-4ccb-92e7-237491e6a853', 'Блюдо 4', 147.43, 233, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a5e91d5e-06ba-4245-b86c-ed100f7d7078', 'dd3575d9-d937-4ccb-92e7-237491e6a853', 'Блюдо 5', 511.43, 186, 'Чай');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c057429f-8973-4bc6-82aa-f0c0cfce1129', 'dd3575d9-d937-4ccb-92e7-237491e6a853', 'Блюдо 6', 119.74, 312, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('71c10513-571e-496b-9061-074fec81daab', 'b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Блюдо 1', 550.89, 328, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8d175648-720e-4e10-a135-63684fd2fe5f', 'b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Блюдо 2', 188.72, 428, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('074dd67f-3d83-4f25-b19d-f6b055975a23', 'b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Блюдо 3', 470.91, 303, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7651c7ef-5a2a-4399-afba-5cc29e3fd6fc', 'b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Блюдо 4', 533.84, 122, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('85b68434-77c9-48e3-a5e9-bff6070d0dee', 'b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Блюдо 5', 415.41, 394, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5b7adb19-c2fa-401a-aa30-0170050b712e', 'b5fa3111-3931-4d9d-88f2-cb609047c0af', 'Блюдо 6', 302.81, 113, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a59e7457-e704-426a-970b-025f5ae838e4', '18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Блюдо 1', 134.71, 197, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('07ea6f4a-1b49-4ff1-a199-4a90a90c34d8', '18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Блюдо 2', 246.48, 224, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e9b8126e-06ed-4ff0-8ccf-3919f27783ac', '18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Блюдо 3', 390.04, 256, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6474b5d8-2574-4bee-8c45-21d89d14a8b4', '18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Блюдо 4', 354.94, 401, 'Морсы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4489cc2d-c546-4083-b658-b2568e9391f4', '18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Блюдо 5', 440.20, 164, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d4eaec63-a436-40ff-a2d7-e241f50989d9', '18e71863-2711-4fdd-9d9b-ea98ad22aea9', 'Блюдо 6', 394.83, 316, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3558c822-4607-44b0-bba3-5a9c578a04af', 'd15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Блюдо 1', 308.02, 212, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('891bc2dc-9097-48eb-8a2b-2a82222a186c', 'd15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Блюдо 2', 484.83, 468, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('56574d9b-f2ff-4d0e-bdb1-2824de305ce6', 'd15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Блюдо 3', 340.56, 370, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a0cbe264-a615-4759-80a6-3fdb15594eba', 'd15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Блюдо 4', 351.60, 132, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31ff7ace-9cc4-4bdc-a65b-ff64f010e9da', 'd15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Блюдо 5', 340.01, 245, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b51faf12-f7f8-41e9-ae20-beb5df64bd21', 'd15af0ae-f33d-478d-aea4-d018eb2e4a9b', 'Блюдо 6', 323.59, 316, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01198235-ef97-47c0-af08-12a6552f5e97', '493e5895-17bd-4177-a185-6975a789cc58', 'Блюдо 1', 587.64, 436, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('d0f4e630-a76d-4b9e-8886-35520b57f1e3', '493e5895-17bd-4177-a185-6975a789cc58', 'Блюдо 2', 538.64, 345, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('b27f0d82-973a-46ed-b5ed-3d75b8f2a8b7', '493e5895-17bd-4177-a185-6975a789cc58', 'Блюдо 3', 581.32, 206, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('eab91739-e1ab-4d19-8719-f5e325d594f5', '493e5895-17bd-4177-a185-6975a789cc58', 'Блюдо 4', 427.18, 139, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6d071549-626a-4d88-8250-b0aea94b0195', '493e5895-17bd-4177-a185-6975a789cc58', 'Блюдо 5', 253.34, 352, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('98486645-06aa-4ae4-89f1-4b795002f2b9', '493e5895-17bd-4177-a185-6975a789cc58', 'Блюдо 6', 176.39, 291, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31d5b63e-2c6c-49f9-a75a-587419efbf10', '59d649f5-7921-4249-aef9-860748b46b18', 'Блюдо 1', 432.85, 208, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('22aeac69-01e6-4cf3-9f47-1b2147bde47b', '59d649f5-7921-4249-aef9-860748b46b18', 'Блюдо 2', 535.94, 251, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('552eb37e-5671-463d-a297-21ff471add24', '59d649f5-7921-4249-aef9-860748b46b18', 'Блюдо 3', 540.23, 275, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5b2b9cdb-6786-40de-bf06-4b51f545daf9', '59d649f5-7921-4249-aef9-860748b46b18', 'Блюдо 4', 448.15, 120, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('39cfd4f5-c22b-4bff-86bc-157c0eb94cde', '59d649f5-7921-4249-aef9-860748b46b18', 'Блюдо 5', 185.78, 464, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('67e263b0-2ad1-4881-9f23-c0cbcc321368', '59d649f5-7921-4249-aef9-860748b46b18', 'Блюдо 6', 131.61, 137, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f4de191f-28af-4241-add6-28d2d174461a', '095ab469-eefb-450e-8cc4-db0a4aabce45', 'Блюдо 1', 444.26, 156, 'Паста');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('43c8f0bd-88f1-4aec-a884-191ab000efea', '095ab469-eefb-450e-8cc4-db0a4aabce45', 'Блюдо 2', 539.45, 123, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('668b8196-741f-4fe2-95bc-77328e16b44f', '095ab469-eefb-450e-8cc4-db0a4aabce45', 'Блюдо 3', 437.46, 483, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('826bf7f4-9cf8-4f1b-90af-8c67c58f8750', '095ab469-eefb-450e-8cc4-db0a4aabce45', 'Блюдо 4', 350.92, 456, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('015e24c5-e81b-4c81-a0f5-ac5192be06c5', '095ab469-eefb-450e-8cc4-db0a4aabce45', 'Блюдо 5', 518.45, 440, 'Десерты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('677b6054-0924-4c17-afeb-27590cfa9588', '095ab469-eefb-450e-8cc4-db0a4aabce45', 'Блюдо 6', 268.47, 233, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6abbbca9-e297-42a2-b65b-5fc285c28732', 'c274c432-3f84-4e11-936a-b6818ed1542c', 'Блюдо 1', 266.10, 254, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('cfe93e31-8303-4adf-a85d-bbd787f5e68f', 'c274c432-3f84-4e11-936a-b6818ed1542c', 'Блюдо 2', 399.40, 251, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('8f2212a1-9498-412e-99fa-7d5a4045a8a1', 'c274c432-3f84-4e11-936a-b6818ed1542c', 'Блюдо 3', 473.27, 222, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b58778c-e9a4-400f-ac43-f9059623f0f6', 'c274c432-3f84-4e11-936a-b6818ed1542c', 'Блюдо 4', 448.47, 356, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('47ee4845-2252-4cf7-ac9e-dae63188dc97', 'c274c432-3f84-4e11-936a-b6818ed1542c', 'Блюдо 5', 415.00, 337, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9393bc06-9543-48f0-8f94-7615c94c893a', 'c274c432-3f84-4e11-936a-b6818ed1542c', 'Блюдо 6', 275.17, 185, 'Кофе');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('59916d43-5d1d-48e3-a48e-96c0109bbe8d', '7cf710da-960f-4725-b249-e95c419d8d6d', 'Блюдо 1', 317.08, 480, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e3af488d-7077-4547-a648-417d2d2ec8d7', '7cf710da-960f-4725-b249-e95c419d8d6d', 'Блюдо 2', 168.18, 417, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7762c529-2ab6-4c8e-b84d-59a73904ae4f', '7cf710da-960f-4725-b249-e95c419d8d6d', 'Блюдо 3', 193.39, 293, 'Лапша');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9ae21412-98a6-45fe-93f9-2afb2c86811a', '7cf710da-960f-4725-b249-e95c419d8d6d', 'Блюдо 4', 141.04, 198, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fc3f7c56-90fe-40bb-b8f8-0a42ddbd247a', '7cf710da-960f-4725-b249-e95c419d8d6d', 'Блюдо 5', 247.81, 432, 'Напитки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9627503c-2322-48aa-9387-7ce934d97bb8', '7cf710da-960f-4725-b249-e95c419d8d6d', 'Блюдо 6', 356.65, 383, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4a33a30b-18de-466c-9410-cc8cf3ba6fbd', '168f82bb-cc39-4d4a-b35d-21f555f77308', 'Блюдо 1', 544.57, 120, 'Завтраки');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('a6ae1ec8-1152-4f26-975f-d2645357f9ad', '168f82bb-cc39-4d4a-b35d-21f555f77308', 'Блюдо 2', 322.56, 479, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('59df1cff-8359-4bc3-8295-8484e4730f55', '168f82bb-cc39-4d4a-b35d-21f555f77308', 'Блюдо 3', 221.78, 444, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('5dfef60f-25d7-4217-afe7-d039e7025fb4', '168f82bb-cc39-4d4a-b35d-21f555f77308', 'Блюдо 4', 112.93, 274, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1e04b830-93f5-4bcd-b3fd-6ac387555bba', '168f82bb-cc39-4d4a-b35d-21f555f77308', 'Блюдо 5', 399.40, 133, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7a39fc19-3cd5-415c-a0f0-a44946026b39', '168f82bb-cc39-4d4a-b35d-21f555f77308', 'Блюдо 6', 323.91, 303, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('50dc05de-2830-4d53-90e9-0baedd37c83e', 'a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Блюдо 1', 508.67, 440, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('47f0178d-efe1-486a-aa64-e8117a7b6b50', 'a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Блюдо 2', 395.25, 222, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('63312442-6fff-4e14-b782-49655de4e3d9', 'a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Блюдо 3', 273.27, 369, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ce166818-1eb4-4364-a885-259da8c68425', 'a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Блюдо 4', 236.30, 317, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('4e16e2ec-30e6-4eab-b354-2a7e472beb9f', 'a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Блюдо 5', 306.21, 426, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('bcc42041-5226-4c0c-ada7-f9fee604a728', 'a6e492ee-4afb-441b-8ef5-bb7b6e8c6b80', 'Блюдо 6', 129.95, 200, 'Закуски');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('580e7b36-0fa7-4e10-926d-ab966e9d2ea2', '9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Блюдо 1', 386.97, 300, 'Гарниры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0b8a0741-0264-4e5a-ad39-b6c8996e0d47', '9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Блюдо 2', 131.37, 392, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('620d4512-68e8-4e4b-8833-c3dd3390a85e', '9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Блюдо 3', 584.03, 467, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('aef37123-2ce6-4dd6-bdf8-86030197faf5', '9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Блюдо 4', 247.62, 220, 'Рыба');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('7b5d9946-b8ee-454c-ab23-073880974bb6', '9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Блюдо 5', 313.44, 262, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e10c6d0a-149d-4d6f-a858-0034c08e5c5a', '9f03c330-1641-4ae4-bbf4-55462cc35faa', 'Блюдо 6', 319.49, 410, 'Супы дня');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('2a4c3923-a2cd-4c71-bea4-69f286256dc8', '7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Блюдо 1', 346.95, 298, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('01b4aa4f-da55-46de-9342-359f8bbfcb47', '7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Блюдо 2', 225.19, 496, 'Фреши');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('0e8df8b0-719b-4521-aeb3-18724ad80b94', '7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Блюдо 3', 487.39, 439, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6478d8d7-fe6c-48f7-8c44-24fa4a4870f9', '7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Блюдо 4', 349.62, 460, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('9c230299-267d-4286-ab4a-ea5230197774', '7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Блюдо 5', 425.50, 165, 'Торты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('ac094c89-f14a-43ff-a103-1748164037d9', '7bbfad9d-381e-47c8-b063-c9d4a1f9ac82', 'Блюдо 6', 385.81, 104, 'Пицца');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('3685a729-2521-4324-9337-749aeb48d304', 'adda4717-f118-45cc-9532-30502637846d', 'Блюдо 1', 438.85, 480, 'Гриль');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('069f16a8-22f7-4596-ad7c-cd83ece73cd1', 'adda4717-f118-45cc-9532-30502637846d', 'Блюдо 2', 556.29, 264, 'Сэндвичи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e1c5b5d8-0371-4ab5-9a88-5d0bcf9fb320', 'adda4717-f118-45cc-9532-30502637846d', 'Блюдо 3', 558.66, 288, 'Веганские блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('6da602b7-74fe-4033-9fe5-5b6bd8d4c999', 'adda4717-f118-45cc-9532-30502637846d', 'Блюдо 4', 178.17, 155, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('062eabde-2bc8-4f4c-971f-33d6ccb4a44e', 'adda4717-f118-45cc-9532-30502637846d', 'Блюдо 5', 461.58, 257, 'Смузи');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1f95d4a9-20bc-4cdc-934e-e8f27194f066', 'adda4717-f118-45cc-9532-30502637846d', 'Блюдо 6', 277.84, 203, 'Пирожные');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('866a70b6-924d-4b48-99d2-8921f329cd96', 'f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Блюдо 1', 373.46, 408, 'Супы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('464a0fdb-fb68-4fb8-b0f7-f8abd69cabdb', 'f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Блюдо 2', 335.70, 176, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('fbad266c-9a31-4f57-abaf-73b33575b9c2', 'f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Блюдо 3', 590.03, 122, 'Детское меню');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('31aa11aa-b40e-4860-a3cb-e9063c0b94ba', 'f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Блюдо 4', 322.75, 268, 'Сашими');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('e140fad3-212e-4550-9d2f-adce2bfcb85d', 'f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Блюдо 5', 356.04, 152, 'Горячее');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('66539334-9665-4fc9-8aa7-44ae45259257', 'f2c70579-a83b-4c0b-97ea-4dd7f59b3c39', 'Блюдо 6', 532.65, 307, 'Салаты');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('052bc925-0905-40f6-9729-0b49aa384cc1', '64728b3c-67a7-4107-bac8-f6f998aed794', 'Блюдо 1', 380.39, 493, 'Соусы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('48b7b9ea-e1d1-4760-b0c5-8a5a201c18cd', '64728b3c-67a7-4107-bac8-f6f998aed794', 'Блюдо 2', 250.12, 183, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('1726a696-cdee-4012-801c-9787578edf85', '64728b3c-67a7-4107-bac8-f6f998aed794', 'Блюдо 3', 230.12, 466, 'Горячие блюда');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('c2051d35-1d6a-42f0-b4c5-e1a3bf637dd7', '64728b3c-67a7-4107-bac8-f6f998aed794', 'Блюдо 4', 124.35, 131, 'Бургеры');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('190dabf5-6a83-476d-b633-258ae5f96e26', '64728b3c-67a7-4107-bac8-f6f998aed794', 'Блюдо 5', 354.80, 263, 'Роллы');
INSERT INTO products (id, restaurant_id, name, price, weight, category) VALUES ('f6174fb9-f78c-4e87-91c8-811a38cca6b8', '64728b3c-67a7-4107-bac8-f6f998aed794', 'Блюдо 6', 537.34, 389, 'Соусы');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('a56915b5-3abf-40da-9166-eac7484e719b', '7ca5d04a-fc86-4716-99d4-26e178b6016b', 'created', 'Улица 36, дом 7', '[{"id":"eab91739-e1ab-4d19-8719-f5e325d594f5","name":"Блюдо 4","price":427.17536392812076,"image_url":"default_product.jpg","weight":139,"amount":2},{"id":"98486645-06aa-4ae4-89f1-4b795002f2b9","name":"Блюдо 6","price":176.38504441088858,"image_url":"default_product.jpg","weight":291,"amount":3},{"id":"6d071549-626a-4d88-8250-b0aea94b0195","name":"Блюдо 5","price":253.33928449423664,"image_url":"default_product.jpg","weight":352,"amount":1},{"id":"d0f4e630-a76d-4b9e-8886-35520b57f1e3","name":"Блюдо 2","price":538.6357618503907,"image_url":"default_product.jpg","weight":345,"amount":3}]', 'кв. 84', '179', '4', '13', 'Не забудьте вилки', true, '3252.75');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('2d868244-0f5b-4ff5-ab29-92dfdd984092', '3419e315-99c0-4fd3-9f0b-39e7877e738e', 'created', 'Улица 98, дом 8', '[{"id":"976f19dd-e8ad-4716-ac53-5c96e67c21e0","name":"Блюдо 1","price":463.52021040182603,"image_url":"default_product.jpg","weight":407,"amount":1},{"id":"31fd80f9-cef1-4268-ab38-7b5c567f9630","name":"Блюдо 6","price":297.3512345219293,"image_url":"default_product.jpg","weight":219,"amount":1}]', 'кв. 60', '682', '4', '15', 'Не забудьте вилки', false, '760.87');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('605ad692-55e9-4a8e-8a14-6a7ed628716e', '3419e315-99c0-4fd3-9f0b-39e7877e738e', 'created', 'Улица 13, дом 18', '[{"id":"8bfeea54-e79d-4bf6-a9a2-553fceb6dffd","name":"Блюдо 1","price":548.8655422950802,"image_url":"default_product.jpg","weight":216,"amount":1},{"id":"eb6bc889-5c09-4a1f-a708-6fea9b106ddd","name":"Блюдо 4","price":171.79792176968255,"image_url":"default_product.jpg","weight":261,"amount":1},{"id":"ca469d27-5f15-4df0-9b99-b6af77639d48","name":"Блюдо 5","price":146.06228777251377,"image_url":"default_product.jpg","weight":494,"amount":1},{"id":"fa18bb96-591c-4acd-8404-947188c5de30","name":"Блюдо 2","price":509.9425910933774,"image_url":"default_product.jpg","weight":152,"amount":2}]', 'кв. 196', '417', '4', '16', 'Не забудьте вилки', false, '1886.61');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('ece9adae-f57d-4ac7-96d1-99b1de84005a', '594b5f85-a0c0-4b0b-b729-025e233d2c90', 'created', 'Улица 0, дом 40', '[{"id":"0c603323-ec75-4fc5-a072-4dfa79262124","name":"Блюдо 4","price":404.7062380149449,"image_url":"default_product.jpg","weight":311,"amount":3},{"id":"2c66a04d-4052-4689-b43f-ba6ee05348c7","name":"Блюдо 1","price":567.1563645210315,"image_url":"default_product.jpg","weight":183,"amount":2},{"id":"8a61108b-69ee-4c44-94a6-73f62bb2aca8","name":"Блюдо 3","price":139.83271744994883,"image_url":"default_product.jpg","weight":414,"amount":1},{"id":"4d3ddeba-ed85-4d29-8a67-de43342c6240","name":"Блюдо 5","price":115.87298395522272,"image_url":"default_product.jpg","weight":302,"amount":2}]', 'кв. 197', '253', '1', '11', 'Не забудьте вилки', false, '2720.01');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('141f9843-df9b-4cbb-84e6-e7cc136bb8c1', '594b5f85-a0c0-4b0b-b729-025e233d2c90', 'created', 'Улица 4, дом 18', '[{"id":"30b5024d-fef7-4338-9747-7109bf87adbc","name":"Блюдо 6","price":589.5545222192536,"image_url":"default_product.jpg","weight":429,"amount":3}]', 'кв. 183', '625', '3', '8', 'Не забудьте вилки', false, '1768.66');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('482e0857-e59c-4ef9-94ae-9e7e1ae863e7', 'bc20994d-919a-4d57-aae1-50f9da7e8588', 'created', 'Улица 12, дом 15', '[{"id":"e07efcd7-8674-44ce-a0e8-f009d84a8037","name":"Блюдо 5","price":108.71591298725394,"image_url":"default_product.jpg","weight":334,"amount":2}]', 'кв. 124', '197', '1', '24', 'Не забудьте вилки', true, '217.43');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('2381ebb3-49f6-4ae4-a7e4-0314d16653c9', '594b5f85-a0c0-4b0b-b729-025e233d2c90', 'created', 'Улица 0, дом 40', '[{"id":"c6ff55a8-b3e6-439b-9687-e0b6ab2a7ab8","name":"Блюдо 6","price":449.69791933229266,"image_url":"default_product.jpg","weight":272,"amount":2},{"id":"ecfb8f2e-540f-43ce-bef9-979bb81b2d05","name":"Блюдо 3","price":366.48030697747674,"image_url":"default_product.jpg","weight":244,"amount":3},{"id":"874c88e8-6e96-43d1-a989-977169df030f","name":"Блюдо 5","price":574.270944218811,"image_url":"default_product.jpg","weight":115,"amount":3}]', 'кв. 16', '622', '5', '5', 'Не забудьте вилки', true, '3721.65');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('e043ad17-ef0d-434b-92dc-3c214183c9b1', '5cc101e9-0cc3-41f6-8b56-c5e29de758b1', 'created', 'Улица 55, дом 42', '[{"id":"9bcf6311-3c92-4ee2-b2be-e29e94f5729f","name":"Блюдо 5","price":291.4057714711759,"image_url":"default_product.jpg","weight":198,"amount":1},{"id":"30b5024d-fef7-4338-9747-7109bf87adbc","name":"Блюдо 6","price":589.5545222192536,"image_url":"default_product.jpg","weight":429,"amount":1}]', 'кв. 145', '193', '2', '13', 'Не забудьте вилки', true, '880.96');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('2bfb9559-6b58-4243-bee2-63554600105c', '8b72e7df-64d4-4037-9cfd-e79a06a7563f', 'created', 'Улица 12, дом 15', '[{"id":"08205bb2-cdf5-40d6-9dbe-78b5af824663","name":"Блюдо 5","price":316.8147065979378,"image_url":"default_product.jpg","weight":271,"amount":2}]', 'кв. 62', '872', '5', '21', 'Не забудьте вилки', true, '633.63');
INSERT INTO orders (id, user_id, status, address_id, order_products, apartment_or_office, intercom, entrance, floor, courier_comment, leave_at_door) VALUES ('2fa7872b-37fb-4d4e-ae39-160a01bdac9a', '90bbb029-540a-40e6-b17e-448dd8ddb19f', 'created', 'Улица 12, дом 15', '[{"id":"6fe8674d-5fef-4f8c-97ad-b27889dd9f6f","name":"Блюдо 1","price":266.34572427800686,"image_url":"default_product.jpg","weight":478,"amount":1},{"id":"db662790-adbf-47d9-ac72-6650439c8233","name":"Блюдо 6","price":131.28821242093355,"image_url":"default_product.jpg","weight":103,"amount":3}]', 'кв. 87', '294', '3', '20', 'Не забудьте вилки', false, '660.21');
