# ER-диаграмма 

```mermaid
erDiagram
    USERS ||--o{ ADDRESSES : "Имеет адреса | user_id | 1:M"
    USERS ||--o{ ORDERS : "Создает заказы | user_id | 1:M"
    USERS ||--o{ REVIEWS : "Оставляет отзывы | user_id | 1:M"
    USERS ||--o{ RECOMMENDATIONS : "Получает рекомендации | user_id | 1:M"
    
    RESTAURANTS ||--o{ RESTAURANT_CATEGORY_RELATIONS : "Принадлежит к категориям | restaurant_id | 1:M"
    RESTAURANTS ||--o{ PRODUCTS : "Предлагает продукты | restaurant_id | 1:M"
    RESTAURANTS ||--o{ ORDERS : "Выполняет заказы | restaurant_id | 1:M"
    RESTAURANTS ||--o{ REVIEWS : "Получает отзывы | restaurant_id | 1:M"
    
    RESTAURANT_CATEGORIES ||--o{ RESTAURANT_CATEGORY_RELATIONS : "Содержит рестораны | category_id | 1:M"
    
    PRODUCT_CATEGORIES ||--o{ PRODUCTS : "Содержит продукты | category_id | 1:M"
    PRODUCT_CATEGORIES ||--o{ PRODUCT_CATEGORY_RELATIONS : "Имеет связи с продуктами | category_id | 1:M"
    
    PRODUCTS ||--o{ PRODUCT_CATEGORY_RELATIONS : "Относится к категориям | product_id | 1:M"
    PRODUCTS ||--o{ ORDER_ITEMS : "Входит в заказы | product_id | 1:M"
    PRODUCTS ||--o{ RECOMMENDATIONS : "Рекомендуется пользователям | product_id | 1:M"
    
    ORDERS ||--o{ ORDER_ITEMS : "Содержит позиции | order_id | 1:M"
    ORDERS }|--|| ADDRESSES : "Доставляется по адресу | address_id | 1:1"
    ORDERS }|--|| RESTAURANTS : "Выполняется рестораном | restaurant_id | 1:1"
    ORDERS }|--|| USERS : "Принадлежит пользователю | user_id | 1:1"

    USERS {
        UUID id PK
        TEXT username
        TEXT email
        TEXT first_name
        TEXT last_name
        TEXT password_hash
        TEXT description
        TEXT user_pic
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    ADDRESSES {
        UUID id PK
        UUID user_id FK
        TEXT address
        TIMESTAMP created_at
    }

    RESTAURANTS {
        UUID id PK
        TEXT name
        TEXT banner_url
        TEXT address
        FLOAT rating
        INT rating_count
        INT working_mode_from
        INT working_mode_to
        INT delivery_time_from
        INT delivery_time_to
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    RESTAURANT_CATEGORIES {
        UUID id PK
        TEXT name
        TIMESTAMP created_at
    }

    RESTAURANT_CATEGORY_RELATIONS {
        UUID restaurant_id PK, FK
        UUID category_id PK, FK
    }

    PRODUCT_CATEGORIES {
        UUID id PK
        TEXT name
        TIMESTAMP created_at
    }

    PRODUCTS {
        UUID id PK
        TEXT name
        NUMERIC(2) price
        TEXT image_url
        INT weight
        UUID category_id FK
        UUID restaurant_id FK
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    PRODUCT_CATEGORY_RELATIONS {
        UUID product_id PK, FK
        UUID category_id PK, FK
    }

    ORDERS {
        UUID id PK
        UUID user_id FK
        UUID restaurant_id FK
        UUID address_id FK
        TEXT status
        NUMERIC(2) total_price
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    ORDER_ITEMS {
        UUID order_id PK, FK
        UUID product_id PK, FK
        INT quantity
        NUMERIC(2) price_at_time
    }

    REVIEWS {
        UUID id PK
        UUID user_id FK
        UUID restaurant_id FK
        INT rating
        TEXT comment
        TIMESTAMP created_at
    }

    PROMO_CODES {
        UUID id PK
        TEXT code
        INT discount_percent
        TIMESTAMP expiration_date
        TIMESTAMP created_at
    }

    RECOMMENDATIONS {
        UUID id PK
        UUID user_id FK
        UUID product_id FK
        TIMESTAMP created_at
    }