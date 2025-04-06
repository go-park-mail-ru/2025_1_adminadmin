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
(uuid_generate_v4(), 'Паста и Вино', 'default_restaurant.jpg', 'Паста и Вино на улице Пушкина', 1.5, 850, 'Ресторан Паста и Вино, уютная атмосфера и вкусная еда.', 12, 22, 49, 88),
(uuid_generate_v4(), 'Суши Дрим', 'default_restaurant.jpg', 'Суши Дрим на улице Пушкина', 2.5, 714, 'Ресторан Суши Дрим, уютная атмосфера и вкусная еда.', 11, 23, 36, 71),
(uuid_generate_v4(), 'Бургерная Ривьера', 'default_restaurant.jpg', 'Бургерная Ривьера на улице Суворова', 4.7, 662, 'Ресторан Бургерная Ривьера, уютная атмосфера и вкусная еда.', 9, 22, 55, 65),
(uuid_generate_v4(), 'Турецкий базар', 'default_restaurant.jpg', 'Турецкий базар на улице Ленина', 3.8, 752, 'Ресторан Турецкий базар, уютная атмосфера и вкусная еда.', 9, 22, 47, 89),
(uuid_generate_v4(), 'Зеленая вилка', 'default_restaurant.jpg', 'Зеленая вилка на улице Пушкина', 2.1, 265, 'Ресторан Зеленая вилка, уютная атмосфера и вкусная еда.', 11, 22, 46, 88),
(uuid_generate_v4(), 'Гриль Бар', 'default_restaurant.jpg', 'Гриль Бар на улице Пушкина', 2.7, 941, 'Ресторан Гриль Бар, уютная атмосфера и вкусная еда.', 9, 21, 57, 74),
(uuid_generate_v4(), 'Американская кухня', 'default_restaurant.jpg', 'Американская кухня на улице Пушкина', 4.1, 778, 'Ресторан Американская кухня, уютная атмосфера и вкусная еда.', 8, 23, 54, 66),
(uuid_generate_v4(), 'Ресторан Средиземноморья', 'default_restaurant.jpg', 'Ресторан Средиземноморья на улице Суворова', 1.9, 526, 'Ресторан Ресторан Средиземноморья, уютная атмосфера и вкусная еда.', 8, 23, 57, 78),
(uuid_generate_v4(), 'Индийские специи', 'default_restaurant.jpg', 'Индийские специи на улице Карла Маркса', 4.6, 278, 'Ресторан Индийские специи, уютная атмосфера и вкусная еда.', 11, 21, 50, 80),
(uuid_generate_v4(), 'Веганское счастье', 'default_restaurant.jpg', 'Веганское счастье на улице Ленина', 1.8, 9, 'Ресторан Веганское счастье, уютная атмосфера и вкусная еда.', 12, 20, 48, 66),
(uuid_generate_v4(), 'Французский уголок', 'default_restaurant.jpg', 'Французский уголок на улице Суворова', 1.9, 786, 'Ресторан Французский уголок, уютная атмосфера и вкусная еда.', 11, 20, 53, 70),
(uuid_generate_v4(), 'Мексиканская пекарня', 'default_restaurant.jpg', 'Мексиканская пекарня на улице Карла Маркса', 0.7, 441, 'Ресторан Мексиканская пекарня, уютная атмосфера и вкусная еда.', 10, 22, 33, 86),
(uuid_generate_v4(), 'Китайская империя', 'default_restaurant.jpg', 'Китайская империя на улице Карла Маркса', 4.0, 537, 'Ресторан Китайская империя, уютная атмосфера и вкусная еда.', 8, 21, 57, 70),
(uuid_generate_v4(), 'Баварский пивной сад', 'default_restaurant.jpg', 'Баварский пивной сад на улице Пушкина', 1.5, 774, 'Ресторан Баварский пивной сад, уютная атмосфера и вкусная еда.', 8, 21, 55, 80),
(uuid_generate_v4(), 'Морская звезда', 'default_restaurant.jpg', 'Морская звезда на улице Ленина', 2.1, 302, 'Ресторан Морская звезда, уютная атмосфера и вкусная еда.', 11, 21, 42, 75),
(uuid_generate_v4(), 'Шашлыки от Бабая', 'default_restaurant.jpg', 'Шашлыки от Бабая на улице Карла Маркса', 0.1, 699, 'Ресторан Шашлыки от Бабая, уютная атмосфера и вкусная еда.', 8, 20, 56, 79),
(uuid_generate_v4(), 'Скоро будет', 'default_restaurant.jpg', 'Скоро будет на улице Ленина', 3.9, 211, 'Ресторан Скоро будет, уютная атмосфера и вкусная еда.', 8, 21, 60, 77),
(uuid_generate_v4(), 'Восточный базар', 'default_restaurant.jpg', 'Восточный базар на улице Суворова', 3.3, 426, 'Ресторан Восточный базар, уютная атмосфера и вкусная еда.', 8, 22, 51, 67),
(uuid_generate_v4(), 'Греческий дворик', 'default_restaurant.jpg', 'Греческий дворик на улице Ленина', 4.8, 346, 'Ресторан Греческий дворик, уютная атмосфера и вкусная еда.', 11, 21, 53, 85),
(uuid_generate_v4(), 'Тосканский огонь', 'default_restaurant.jpg', 'Тосканский огонь на улице Ленина', 0.1, 105, 'Ресторан Тосканский огонь, уютная атмосфера и вкусная еда.', 11, 21, 32, 70),
(uuid_generate_v4(), 'Итальянская ривьера', 'default_restaurant.jpg', 'Итальянская ривьера на улице Мира', 1.3, 929, 'Ресторан Итальянская ривьера, уютная атмосфера и вкусная еда.', 10, 23, 54, 60),
(uuid_generate_v4(), 'Суши Мания', 'default_restaurant.jpg', 'Суши Мания на улице Карла Маркса', 4.8, 770, 'Ресторан Суши Мания, уютная атмосфера и вкусная еда.', 12, 22, 33, 88),
(uuid_generate_v4(), 'Пельмени на углях', 'default_restaurant.jpg', 'Пельмени на углях на улице Ленина', 1.8, 655, 'Ресторан Пельмени на углях, уютная атмосфера и вкусная еда.', 8, 20, 33, 73),
(uuid_generate_v4(), 'Бургеры по-американски', 'default_restaurant.jpg', 'Бургеры по-американски на улице Суворова', 3.9, 464, 'Ресторан Бургеры по-американски, уютная атмосфера и вкусная еда.', 11, 20, 51, 83),
(uuid_generate_v4(), 'Китайская звезда', 'default_restaurant.jpg', 'Китайская звезда на улице Мира', 2.7, 205, 'Ресторан Китайская звезда, уютная атмосфера и вкусная еда.', 9, 20, 56, 64),
(uuid_generate_v4(), 'Мексиканская закуска', 'default_restaurant.jpg', 'Мексиканская закуска на улице Пушкина', 3.4, 108, 'Ресторан Мексиканская закуска, уютная атмосфера и вкусная еда.', 11, 23, 30, 70),
(uuid_generate_v4(), 'Французский бистро', 'default_restaurant.jpg', 'Французский бистро на улице Мира', 2.7, 961, 'Ресторан Французский бистро, уютная атмосфера и вкусная еда.', 11, 23, 50, 76),
(uuid_generate_v4(), 'Греческий остров', 'default_restaurant.jpg', 'Греческий остров на улице Пушкина', 4.7, 761, 'Ресторан Греческий остров, уютная атмосфера и вкусная еда.', 8, 23, 33, 90),
(uuid_generate_v4(), 'Турецкая радость', 'default_restaurant.jpg', 'Турецкая радость на улице Суворова', 0.8, 104, 'Ресторан Турецкая радость, уютная атмосфера и вкусная еда.', 8, 22, 59, 60),
(uuid_generate_v4(), 'Индийская сказка', 'default_restaurant.jpg', 'Индийская сказка на улице Пушкина', 4.4, 15, 'Ресторан Индийская сказка, уютная атмосфера и вкусная еда.', 11, 23, 31, 68),
(uuid_generate_v4(), 'Американская пекарня', 'default_restaurant.jpg', 'Американская пекарня на улице Ленина', 4.5, 693, 'Ресторан Американская пекарня, уютная атмосфера и вкусная еда.', 9, 21, 37, 77),
(uuid_generate_v4(), 'Восточный салат', 'default_restaurant.jpg', 'Восточный салат на улице Мира', 0.2, 168, 'Ресторан Восточный салат, уютная атмосфера и вкусная еда.', 12, 20, 45, 85),
(uuid_generate_v4(), 'Вегетарианский рай', 'default_restaurant.jpg', 'Вегетарианский рай на улице Ленина', 3.6, 867, 'Ресторан Вегетарианский рай, уютная атмосфера и вкусная еда.', 11, 22, 55, 70),
(uuid_generate_v4(), 'Ресторан на воде', 'default_restaurant.jpg', 'Ресторан на воде на улице Ленина', 2.3, 971, 'Ресторан Ресторан на воде, уютная атмосфера и вкусная еда.', 9, 21, 58, 62),
(uuid_generate_v4(), 'Баварская пивоварня', 'default_restaurant.jpg', 'Баварская пивоварня на улице Ленина', 3.5, 273, 'Ресторан Баварская пивоварня, уютная атмосфера и вкусная еда.', 8, 23, 47, 62),
(uuid_generate_v4(), 'Морская лагуна', 'default_restaurant.jpg', 'Морская лагуна на улице Пушкина', 5.0, 341, 'Ресторан Морская лагуна, уютная атмосфера и вкусная еда.', 10, 20, 38, 84),
(uuid_generate_v4(), 'Тосканские вечера', 'default_restaurant.jpg', 'Тосканские вечера на улице Мира', 5.0, 350, 'Ресторан Тосканские вечера, уютная атмосфера и вкусная еда.', 9, 23, 54, 63),
(uuid_generate_v4(), 'Суши и роллы', 'default_restaurant.jpg', 'Суши и роллы на улице Пушкина', 3.1, 500, 'Ресторан Суши и роллы, уютная атмосфера и вкусная еда.', 11, 23, 57, 88),
(uuid_generate_v4(), 'Вкус Индии', 'default_restaurant.jpg', 'Вкус Индии на улице Ленина', 2.4, 56, 'Ресторан Вкус Индии, уютная атмосфера и вкусная еда.', 8, 22, 38, 89),
(uuid_generate_v4(), 'Мексиканская площадь', 'default_restaurant.jpg', 'Мексиканская площадь на улице Мира', 1.5, 812, 'Ресторан Мексиканская площадь, уютная атмосфера и вкусная еда.', 12, 21, 34, 80),
(uuid_generate_v4(), 'Греческая таверна', 'default_restaurant.jpg', 'Греческая таверна на улице Пушкина', 1.8, 111, 'Ресторан Греческая таверна, уютная атмосфера и вкусная еда.', 12, 21, 38, 89),
(uuid_generate_v4(), 'Пивной бар Баварии', 'default_restaurant.jpg', 'Пивной бар Баварии на улице Мира', 3.9, 50, 'Ресторан Пивной бар Баварии, уютная атмосфера и вкусная еда.', 12, 21, 47, 63),
(uuid_generate_v4(), 'Итальянский дворик', 'default_restaurant.jpg', 'Итальянский дворик на улице Суворова', 3.7, 162, 'Ресторан Итальянский дворик, уютная атмосфера и вкусная еда.', 10, 23, 30, 69),
(uuid_generate_v4(), 'Ресторан Печка', 'default_restaurant.jpg', 'Ресторан Печка на улице Пушкина', 3.1, 326, 'Ресторан Ресторан Печка, уютная атмосфера и вкусная еда.', 9, 20, 52, 89),
(uuid_generate_v4(), 'Золотая рыба', 'default_restaurant.jpg', 'Золотая рыба на улице Карла Маркса', 3.7, 321, 'Ресторан Золотая рыба, уютная атмосфера и вкусная еда.', 9, 22, 60, 60),
(uuid_generate_v4(), 'Красное море', 'default_restaurant.jpg', 'Красное море на улице Пушкина', 2.2, 386, 'Ресторан Красное море, уютная атмосфера и вкусная еда.', 11, 21, 39, 68),
(uuid_generate_v4(), 'Ресторан Томат', 'default_restaurant.jpg', 'Ресторан Томат на улице Карла Маркса', 4.7, 759, 'Ресторан Ресторан Томат, уютная атмосфера и вкусная еда.', 8, 23, 41, 65),
(uuid_generate_v4(), 'Турецкая кухня', 'default_restaurant.jpg', 'Турецкая кухня на улице Пушкина', 0.6, 655, 'Ресторан Турецкая кухня, уютная атмосфера и вкусная еда.', 10, 23, 53, 81),
(uuid_generate_v4(), 'Вегетарианская кухня', 'default_restaurant.jpg', 'Вегетарианская кухня на улице Пушкина', 3.4, 77, 'Ресторан Вегетарианская кухня, уютная атмосфера и вкусная еда.', 8, 22, 51, 78),
(uuid_generate_v4(), 'Ресторан Адель', 'default_restaurant.jpg', 'Ресторан Адель на улице Мира', 2.0, 330, 'Ресторан Ресторан Адель, уютная атмосфера и вкусная еда.', 11, 20, 53, 84);

INSERT INTO restaurant_tags_relations (restaurant_id, tag_id)
VALUES
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пельмени на углях'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкая кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Вкус Индии'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Гриль и мясо'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Паста и Вино'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Тосканские вечера'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Морская звезда'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Красное море'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Греческий зал'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкая радость'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Суши Дрим'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный базар'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Печка'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Китайская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Тосканские вечера'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Гриль и мясо'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Паста на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Тосканские вечера'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Тосканский огонь'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пивной бар Баварии'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Мексиканский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Скоро будет'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Средиземноморский ресторан'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Средиземноморья'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкий базар'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французский бистро'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Суши Мания'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Индийские специи'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пицца и Суши'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкая радость'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французская кухня на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Зеленая вилка'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Адель'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Веганский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Итальянский дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный салат'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Вегетарианская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный салат'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный салат'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Итальянский дворик'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкая радость'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Средиземноморья'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Шашлык-Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Индийские специи'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Шашлыки от Бабая'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Тосканский огонь'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французская кухня на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкая радость'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный базар'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкий Султан'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Шашлык-Бар'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пивной бар Баварии'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Паста на ужин'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Мексиканская закуска'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Китайская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Китайская империя'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Суши Дрим'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Вкуса'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Индийская сказка'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Мексиканская пекарня'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французский бистро'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Турецкий базар'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Китайская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Гриль и мясо'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Греческий остров'), (SELECT id FROM restaurant_tags WHERE name = 'Турецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан на воде'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный салат'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Средиземноморья'), (SELECT id FROM restaurant_tags WHERE name = 'Фастфуд')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Томат'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Эдем'), (SELECT id FROM restaurant_tags WHERE name = 'Веганский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пицца и Суши'), (SELECT id FROM restaurant_tags WHERE name = 'Вегетарианский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Восточный базар'), (SELECT id FROM restaurant_tags WHERE name = 'Немецкий')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пельмени на углях'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Греческая таверна'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Шашлыки от Бабая'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Бургеры по-американски'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Гриль и мясо'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Мексиканский уголок'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Тосканский огонь'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Бургерная Сити'), (SELECT id FROM restaurant_tags WHERE name = 'Японский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Печка'), (SELECT id FROM restaurant_tags WHERE name = 'Итальянский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Ресторан Томат'), (SELECT id FROM restaurant_tags WHERE name = 'Средиземноморский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Том Ям'), (SELECT id FROM restaurant_tags WHERE name = 'Европейский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Пельмени по-русски'), (SELECT id FROM restaurant_tags WHERE name = 'Китайский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Греческий остров'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Греческая таверна'), (SELECT id FROM restaurant_tags WHERE name = 'Американский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Греческая таверна'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Китайская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Французский бистро'), (SELECT id FROM restaurant_tags WHERE name = 'Французский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Американская кухня'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Томаты и Паста'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Шашлыки по-кавказски'), (SELECT id FROM restaurant_tags WHERE name = 'Индийский')),
(uuid_generate_v4(), (SELECT id FROM restaurants WHERE name = 'Американская пекарня'), (SELECT id FROM restaurant_tags WHERE name = 'Мексиканский'));


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

