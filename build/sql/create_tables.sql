CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    login TEXT NOT NULL UNIQUE,                     
    phone_number TEXT,                             
    description TEXT,                              
    user_pic TEXT,                                 
    password_hash BYTEA NOT NULL                   
);

CREATE TABLE IF NOT EXISTS restaurants (
    id UUID PRIMARY KEY,  
    name TEXT NOT NULL,                            
    description TEXT,                              
    type TEXT,                                     
    rating FLOAT                                   
);

INSERT INTO restaurants (id, name, description, type, rating)
VALUES
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
(uuid_generate_v4(), 'La Piazza', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze', 'Свежие морепродукты', 'Морепродукты', 4.9);

INSERT INTO users (id, login, phone_number, description, user_pic, password_hash)
VALUES (
    uuid_generate_v4(), 
    'testuser', 
    '88005553535', 
    'New User',
    'default.png',
    decode('a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6a1b2c3d4e5f6', 'hex')
);

