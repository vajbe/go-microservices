import psycopg2
import uuid
import random
from faker import Faker
from datetime import datetime, timezone

# Initialize Faker instance
fake = Faker()

# PostgreSQL connection parameters
DB_PARAMS = {
    "host": "localhost",
    "port": "5432",
    "dbname": "admin",  # Update with your database name
    "user": "admin",  # Update with your user
    "password": "admin",  # Update with your password
}

# Number of records to generate
TOTAL_PRODUCTS = 10_000

# Define categories (expandable)
CATEGORIES = ["Electronics", "Books", "Clothing", "Home Appliances", "Toys", "Sports", "Groceries"]

# Set the max attempts for uniqueness
fake.unique.clear()  # Reset previous uniqueness constraints
fake.unique._max_attempts = 5000  # Increase the max attempts (default is 1000)

# Generate a single product record
# Generate a single product record
def generate_product(product_id):
    return {
        "product_id": str(uuid.uuid4()),  # Generate a UUID for product_id
        "name": fake.word().capitalize() + " " + random.choice(["Pro", "Max", "Lite", "Plus", ""]),
        "description": fake.text(max_nb_chars=200),
        "price": round(random.uniform(5.0, 5000.0), 2),
        "stock": random.randint(0, 1000),
        "category": random.choice(CATEGORIES),
        "created_at": int(datetime.now(timezone.utc).timestamp())  # UTC time in Unix timestamp (seconds)
    }

# Connect to PostgreSQL
def connect_to_db():
    return psycopg2.connect(
        host=DB_PARAMS["host"],
        port=DB_PARAMS["port"],
        dbname=DB_PARAMS["dbname"],
        user=DB_PARAMS["user"],
        password=DB_PARAMS["password"]
    )

# Insert a batch of products into the database
def insert_products_batch(products):
    try:
        conn = connect_to_db()
        cursor = conn.cursor()
        
        # Insert product data into 'products' table
        insert_query = """
        INSERT INTO products (id, name, description, price, stock, category, created_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s)
        """
        
        # Executing insert queries in batch
        cursor.executemany(insert_query, products)
        
        # Commit changes
        conn.commit()

        print(f"Inserted {len(products)} products into database.")
        
        # Close cursor and connection
        cursor.close()
        conn.close()
    
    except Exception as e:
        print(f"Error during insertion: {e}")

# Generate and insert data into the database
def generate_and_insert_data(batch_size=1000):
    products_batch = []

    for i in range(1, TOTAL_PRODUCTS + 1):
        product = generate_product(i)
        products_batch.append((
            product["product_id"],
            product["name"],
            product["description"],
            product["price"],
            product["stock"],
            product["category"],
            product["created_at"]
        ))

        # Insert the batch every `batch_size` products
        if i % batch_size == 0:
            insert_products_batch(products_batch)
            products_batch.clear()  # Clear the batch after insertion

        # Print progress
        if i % 100_000 == 0:
            print(f"Generated and inserted {i} products...")

    # Insert remaining products
    if products_batch:
        insert_products_batch(products_batch)
        print(f"Final batch inserted. Total products: {TOTAL_PRODUCTS}")

if __name__ == "__main__":
    generate_and_insert_data()
