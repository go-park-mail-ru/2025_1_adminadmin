from deep_translator import GoogleTranslator
import pandas as pd

# Загрузка файла
df = pd.read_csv("restaurant-2-products-price.csv")

# Перевод
translator = GoogleTranslator(source='en', target='ru')
df["Item Name (RU)"] = df["Item Name"].apply(lambda x: translator.translate(x))

# Сохранение результата
df.to_excel("translated_products.xlsx", index=False)
