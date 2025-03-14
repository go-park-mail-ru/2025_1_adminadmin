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
(uuid_generate_v4(), 'Kebab King 267', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 268', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 269', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 270', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 271', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 272', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 273', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 274', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 275', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 276', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 277', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 278', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 279', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 280', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura 281', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House 282', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien 283', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco 284', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok 285', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner 286', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King 287', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 288', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 289', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 290', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 291', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 292', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 293', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 294', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 295', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 296', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 297', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 298', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 299', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 300', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura 301', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House 302', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien 303', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco 304', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok 305', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner 306', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King 307', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 308', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 309', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 310', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 311', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 312', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 313', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 314', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 315', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 316', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 317', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 318', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 319', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 320', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura 321', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House 322', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien 323', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco 324', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok 325', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner 326', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King 327', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 328', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 329', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 330', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 331', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 332', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 333', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 334', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 335', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 336', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 337', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 338', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 339', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 340', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura 341', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House 342', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien 343', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco 344', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok 345', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner 346', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King 347', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 348', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 349', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 350', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 351', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 352', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 353', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 354', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 355', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 356', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 357', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 358', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 359', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 360', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura 361', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House 362', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien 363', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco 364', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok 365', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner 366', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King 367', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 368', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 369', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 370', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 371', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 372', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 373', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 374', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 375', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 376', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 377', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 378', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 379', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 380', 'Итальянская кухня', 'Итальянский', 4.5),
(uuid_generate_v4(), 'Sakura 381', 'Японская кухня', 'Японский', 4.7),
(uuid_generate_v4(), 'Steak House 382', 'Лучшие стейки в городе', 'Американский', 4.6),
(uuid_generate_v4(), 'Bistro Parisien 383', 'Французская кухня', 'Французский', 4.3),
(uuid_generate_v4(), 'Taco Loco 384', 'Мексиканская кухня', 'Мексиканский', 4.2),
(uuid_generate_v4(), 'Dragon Wok 385', 'Китайская кухня', 'Китайский', 4.4),
(uuid_generate_v4(), 'Berlin Döner 386', 'Настоящий немецкий донер', 'Немецкий', 4.1),
(uuid_generate_v4(), 'Kebab King 387', 'Лучший кебаб в городе', 'Турецкий', 4.0),
(uuid_generate_v4(), 'Green Garden 388', 'Вегетарианская кухня', 'Вегетарианский', 4.8),
(uuid_generate_v4(), 'Sea Breeze 389', 'Свежие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'Burger Haven 390', 'Лучшие бургеры', 'Американский', 4.6),
(uuid_generate_v4(), 'Pasta Palace 391', 'Свежая паста каждый день', 'Итальянский', 4.7),
(uuid_generate_v4(), 'Tokyo Bites 392', 'Аутентичные суши и рамен', 'Японский', 4.9),
(uuid_generate_v4(), 'Brunch & Chill 393', 'Лучший завтрак и бранч', 'Европейский', 4.5),
(uuid_generate_v4(), 'Curry Delight 394', 'Настоящая индийская кухня', 'Индийский', 4.6),
(uuid_generate_v4(), 'Mediterraneo 395', 'Средиземноморская кухня', 'Средиземноморский', 4.7),
(uuid_generate_v4(), 'Pizza Express 396', 'Быстрая пицца с дровяной печи', 'Итальянский', 4.4),
(uuid_generate_v4(), 'Tandoori House 397', 'Ароматные специи Индии', 'Индийский', 4.3),
(uuid_generate_v4(), 'Vegan Bliss 398', 'Полностью веганское меню', 'Веганский', 4.8),
(uuid_generate_v4(), 'Ocean Catch 399', 'Свежайшие морепродукты', 'Морепродукты', 4.9),
(uuid_generate_v4(), 'La Piazza 400', 'Итальянская кухня', 'Итальянский', 4.5);

INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash)
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

