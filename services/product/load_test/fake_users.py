import requests
from faker import Faker

fake = Faker()

API_URL = "http://127.0.0.1:8080/users"  # Replace with your actual endpoint

for _ in range(100):
    user = {
        "name": fake.name(),
        "email": fake.email(),
        "password": "AuthMind@2024",  # Keeping a fixed password
        "phone_number": fake.numerify("##########")  # Generates a 10-digit number
    }

    response = requests.post(API_URL, json=user)
    
    if response.status_code == 200:
        print(f"User {user['email']} created successfully")
    else:
        print(f"Failed to create user {user['email']}: {response.text}")
