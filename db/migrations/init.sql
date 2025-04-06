-- Создание таблицы пользователей
CREATE TABLE user_table (
  id UUID PRIMARY KEY,
  username TEXT NOT NULL UNIQUE,
  email TEXT NOT NULL UNIQUE,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  password_hash TEXT NOT NULL,
  description TEXT,
  user_pic TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы адресов
CREATE TABLE addresses (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES user_table(id) ON DELETE CASCADE,
  address TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы категорий ресторанов
CREATE TABLE restaurant_category (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы категорий продуктов
CREATE TABLE product_category (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы ресторанов
CREATE TABLE restaurant (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  banner_url TEXT,
  address TEXT NOT NULL,
  rating FLOAT DEFAULT 0 CHECK (rating BETWEEN 0 AND 5),
  rating_count INT DEFAULT 0 CHECK (rating_count >= 0),
  working_mode_from INT NOT NULL CHECK (working_mode_from BETWEEN 0 AND 23),
  working_mode_to INT NOT NULL CHECK (working_mode_to BETWEEN 0 AND 23),
  delivery_time_from INT NOT NULL CHECK (delivery_time_from BETWEEN 0 AND 23),
  delivery_time_to INT NOT NULL CHECK (delivery_time_to BETWEEN 0 AND 23),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CHECK (working_mode_from < working_mode_to),
  CHECK (delivery_time_from < delivery_time_to)
);

-- Создание таблицы связи ресторанов с категориями
CREATE TABLE restaurant_category_relation (
  restaurant_id UUID NOT NULL REFERENCES restaurant(id) ON DELETE CASCADE,
  category_id UUID NOT NULL REFERENCES restaurant_category(id) ON DELETE CASCADE,
  PRIMARY KEY (restaurant_id, category_id)
);

-- Создание таблицы продуктов
CREATE TABLE product (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  price NUMERIC(2) NOT NULL CHECK (price > 0),
  image_url TEXT,
  weight INT CHECK (weight > 0),
  category_id UUID NOT NULL REFERENCES product_category(id),
  restaurant_id UUID NOT NULL REFERENCES restaurant(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы связи продуктов с категориями
CREATE TABLE product_category_relation (
  product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
  category_id UUID NOT NULL REFERENCES product_category(id) ON DELETE CASCADE,
  PRIMARY KEY (product_id, category_id)
);

-- Создание таблицы заказов
CREATE TABLE order_table (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES user_table(id),
  restaurant_id UUID NOT NULL REFERENCES restaurant(id),
  address_id UUID NOT NULL REFERENCES address(id),
  status TEXT NOT NULL CHECK (status IN ('new', 'preparing', 'delivering', 'delivered', 'canceled')),
  total_price NUMERIC(2) NOT NULL CHECK (total_price > 0),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы элементов заказа
CREATE TABLE order_item (
  order_id UUID NOT NULL REFERENCES order_table(id) ON DELETE CASCADE,
  product_id UUID NOT NULL REFERENCES product(id),
  quantity INT NOT NULL CHECK (quantity > 0),
  price_at_time NUMERIC(2) NOT NULL CHECK (price_at_time > 0),
  PRIMARY KEY (order_id, product_id)
);

-- Создание таблицы отзывов
CREATE TABLE review (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES user_table(id),
  restaurant_id UUID NOT NULL REFERENCES restaurant(id),
  rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),
  comment TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  UNIQUE (user_id, restaurant_id)
);

-- Создание таблицы промокодов
CREATE TABLE promo_code (
  id UUID PRIMARY KEY,
  code TEXT NOT NULL UNIQUE,
  discount_percent INT NOT NULL CHECK (discount_percent BETWEEN 1 AND 100),
  expiration_date TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание таблицы рекомендаций
CREATE TABLE recommendation (
  id UUID PRIMARY KEY,
  user_id UUID NOT NULL REFERENCES user_table(id) ON DELETE CASCADE,
  product_id UUID NOT NULL REFERENCES product(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Создание триггера для обновления времени изменения
CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Добавление триггеров для автоматического обновления updated_at
CREATE TRIGGER update_user_timestamp
BEFORE UPDATE ON user_table
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_restaurant_timestamp
BEFORE UPDATE ON restaurant
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_product_timestamp
BEFORE UPDATE ON product
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

CREATE TRIGGER update_order_timestamp
BEFORE UPDATE ON order_table
FOR EACH ROW EXECUTE FUNCTION update_timestamp();

-- Создание индексов для ускорения запросов
CREATE INDEX idx_address_user ON address(user_id);
CREATE INDEX idx_product_restaurant ON product(restaurant_id);
CREATE INDEX idx_product_category ON product(category_id);
CREATE INDEX idx_order_user ON order_table(user_id);
CREATE INDEX idx_order_restaurant ON order_table(restaurant_id);
CREATE INDEX idx_order_status ON order_table(status);
CREATE INDEX idx_review_restaurant ON review(restaurant_id);
CREATE INDEX idx_review_user ON review(user_id);