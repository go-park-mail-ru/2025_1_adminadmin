from diffusers import StableDiffusionPipeline
import torch
import pandas as pd
import os

# Загрузка модели (один раз)
pipe = StableDiffusionPipeline.from_pretrained("runwayml/stable-diffusion-v1-5", torch_dtype=torch.float16)
pipe = pipe.to("cuda")

# Загрузка продуктов
df = pd.read_excel("restaurant-2-products-price.csv")  # Или CSV

# Папка для изображений
os.makedirs("product_images", exist_ok=True)

# Генерация
for i, row in df.iterrows():
    prompt = f"Фотография {row['Item Name (RU)']}, на белом фоне, аппетитная подача, блюдо из ресторана"
    image = pipe(prompt).images[0]
    image.save(f"product_images/{row['Item Name (RU)']}.png")
