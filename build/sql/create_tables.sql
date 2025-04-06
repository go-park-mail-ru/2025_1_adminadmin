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

