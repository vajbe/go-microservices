import requests
import json
import concurrent.futures
import time

# API endpoint
API_URL = "http://localhost:8081/products"

# Number of requests to send
NUM_REQUESTS = 10000

# Number of concurrent workers
CONCURRENT_WORKERS = 1000

# Product payload
payload = {
    "name": "Test Product",
    "description": "This is a test product",
    "price": 19.99,
    "stock": 100,
    "category": "Test Category"
}

# Function to make a POST request
def send_request(request_id):
    try:
        response = requests.post(API_URL, json=payload, timeout=5)
        if response.status_code == 200:
            return f"Request {request_id} succeeded with status {response.status_code}"
        else:
            return f"Request {request_id} failed with status {response.status_code}: {response.text}"
    except requests.exceptions.RequestException as e:
        return f"Request {request_id} encountered an error: {e}"

# Main function to run the load test
def main():
    start_time = time.time()

    with concurrent.futures.ThreadPoolExecutor(max_workers=CONCURRENT_WORKERS) as executor:
        futures = {executor.submit(send_request, i): i for i in range(NUM_REQUESTS)}

        for future in concurrent.futures.as_completed(futures):
            print(future.result())

    end_time = time.time()
    print(f"\nCompleted {NUM_REQUESTS} requests in {end_time - start_time:.2f} seconds.")

if __name__ == "__main__":
    main()


##python -m venv env before running