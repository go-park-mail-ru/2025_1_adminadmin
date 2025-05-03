-- Создание словаря для русского языка
CREATE TEXT SEARCH DICTIONARY russian_ispell (
    TEMPLATE = ispell,
    DictFile = russian,
    AffFile = russian,
    StopWords = russian
);

-- Создание конфигурации для русского языка
CREATE TEXT SEARCH CONFIGURATION ru (COPY = russian);

-- Добавление стемминга для русского языка
ALTER TEXT SEARCH CONFIGURATION ru
ALTER MAPPING FOR hword, hword_part, word
WITH russian_ispell, russian_stem;

-- Устанавливаем конфигурацию как дефолтную (по желанию)
-- ALTER SYSTEM SET default_text_search_config = 'ru';