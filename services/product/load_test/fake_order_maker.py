import requests
import random
import concurrent.futures

# API Endpoints
USERS_API = "http://127.0.0.1:8000/users/users"
PRODUCTS_API = "http://127.0.0.1:8000/products/products?offset=1&limit=1000&sort_by=description&order_by=ASC"
ORDERS_API = "http://127.0.0.1:8000/orders"


# Fetch Users
# Fetch Users
# Fetch Users
def get_users():
    response = requests.get(USERS_API)

    try:
        data = response.json()
        print("Users API Response:", data)  # Debugging

        if isinstance(data, dict) and "data" in data and isinstance(data["data"], list):
            return [user.get("id") for user in data["data"] if "id" in user]  # Extract user IDs safely

    except requests.exceptions.JSONDecodeError:
        print("Failed to parse JSON:", response.text)

    print("Failed to fetch users:", response.text)
    return []

# Fetch Products
def get_products():
    response = requests.get(PRODUCTS_API)

    try:
        data = response.json()
        print("Products API Response:", data)  # Debugging

        if isinstance(data, dict) and "data" in data and isinstance(data["data"], list):
            return data  # Assuming each product has "id", "name", and "price"

    except requests.exceptions.JSONDecodeError:
        print("Failed to parse JSON:", response.text)

    print("Failed to fetch products:", response.text)
    return []



# Generate a random order
def create_order(user_id, products):
    selected_products = random.sample(
        products, min(len(products), random.randint(1, 5))
    )  # 1 to 5 products per order

    order_products = [
        {
            "product_id": product["id"],
            "name": product["name"],
            "quantity": random.randint(1, 3),  # Random quantity between 1-3
            "price": product["price"],
        }
        for product in selected_products
    ]

    total_amount = sum(p["quantity"] * p["price"] for p in order_products)

    return {
        "user_id": user_id,
        "order_status": "PENDING",
        "total_amount": round(total_amount, 2),
        "payment_status": "PAID",
        "products": order_products,
    }


# Send Order Request
def send_order(order):
    response = requests.post(ORDERS_API, json=order)
    if response.status_code == 201:
        print(f"Order placed for user {order['user_id']}")
    else:
        print(f"Failed to place order for {order['user_id']}: {response.text}")


# Main Execution
users = get_users()
products = get_products()

if users and products:
    orders = [
        create_order(random.choice(users), products) for _ in range(100)
    ]  # Generate 100 orders

    # Send orders concurrently
    with concurrent.futures.ThreadPoolExecutor(max_workers=10) as executor:
        executor.map(send_order, orders)

    print("100 orders placed concurrently.")
else:
    print("Failed to fetch users or products, cannot place orders.")
