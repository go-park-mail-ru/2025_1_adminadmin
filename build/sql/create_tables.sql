CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,                     
    phone_number TEXT,                              
    first_name TEXT NOT NULL,  
    last_name TEXT NOT NULL,   
    description TEXT DEFAULT '',  
    user_pic TEXT DEFAULT 'default.jpg', 
    password_hash BYTEA NOT NULL                   
);

CREATE TABLE IF NOT EXISTS restaurants (
    id UUID PRIMARY KEY,  
    name TEXT NOT NULL,                            
    description TEXT,                              
    type TEXT,                                     
    rating FLOAT                                   
);

INSERT INTO restaurants (id, name, description, type, rating) VALUES
(uuid_generate_v4(), 'La Piazza', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Pasta Paradise', 'Паста и пицца', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sushi Master', 'Суши и сашими', 'Японский', 4.7),
(uuid_generate_v4(), 'Burger Joint', 'Бургеры и картофель фри', 'Американский', 4.6),
(uuid_generate_v4(), 'Le Petit Bistro', 'Французские деликатесы', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Fiesta', 'Мексиканские тако', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Golden Wok', 'Китайские блюда', 'Китайский', 4.4),
(uuid_generate_v4(), 'Munich Haus', 'Немецкие колбаски', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab Palace', 'Турецкие кебабы', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Veggie Delight', 'Вегетарианские блюда', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Ocean''s Catch', 'Морепродукты и рыба', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Ristorante Roma', 'Итальянские деликатесы', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Tokyo Sushi', 'Японские суши', 'Японский', 4.7),
(uuid_generate_v4(), 'BBQ Pit', 'Американский барбекю', 'Американский', 4.6),
(uuid_generate_v4(), 'Café de Paris', 'Французская выпечка', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Express', 'Мексиканские закуски', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Peking Garden', 'Китайские деликатесы', 'Китайский', 4.4),
(uuid_generate_v4(), 'Bavarian Inn', 'Немецкие блюда', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab Express', 'Турецкие закуски', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Veggie Heaven', 'Вегетарианские деликатесы', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Seafood Haven', 'Морепродукты и устрицы', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Pasta Factory', 'Итальянские паста и соусы', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura Sushi', 'Японские суши и роллы', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak & Ale', 'Американские стейки', 'Американский', 4.6),
(uuid_generate_v4(), 'Parisian Café', 'Французские десерты', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Time', 'Мексиканские тако и буррито', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Palace', 'Китайские деликатесы', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berliner Haus', 'Немецкие блюда', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab House', 'Турецкие кебабы', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Veggie World', 'Вегетарианские блюда', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Seafood Cove', 'Морепродукты и рыба', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Pasta Palace', 'Итальянские паста и пицца', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sushi World', 'Японские суши и сашими', 'Японский', 4.7),
(uuid_generate_v4(), 'Burger Barn', 'Американские бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Parisian Bistro', 'Французские деликатесы', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Town', 'Мексиканские тако', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Express', 'Китайские блюда', 'Китайский', 4.4),
(uuid_generate_v4(), 'Bavarian House', 'Немецкие колбаски', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab Corner', 'Турецкие кебабы', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Veggie Spot', 'Вегетарианские блюда', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Seafood Shack', 'Морепродукты и рыба', 'Морепродукты', 4.9);

INSERT INTO users (id, login, phone_number, description, user_pic, password_hash)
VALUES (
    uuid_generate_v4(), 
    'testuser', 
    'Dmitriy',  
    'Nagiev', 
    '88005553535', 
    '',
    'default.jpg',
    decode('a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6', 'hex')
);

