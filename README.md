**Go Based Microservices Demonstration:**

Tech Stack:
  1. Go lang
  2. Postgres DB
  3. Kafka
  4. Docker
  5. Kubernetes
  6. Redis


**Key Concepts to be demonstrated:**
**Phase 1**
RESTful API Development 
1. Routing, controllers, middleware
2. Database relationships (e.g., users, posts, orders)
3. Authentication (JWT, OAuth2)
4. Pagination, filtering, and sorting

**Phase 2**
Microservices Architecture
1. Service-to-service communication (REST, gRPC, Kafka)
2. API gateway
3. Service discovery and orchestration (e.g., Kubernetes)
4. Event-driven architecture


**Use cases Plannned:**

1. Order Processing
Simple use case: Process multiple orders in parallel.
Example: Simulate 50 concurrent orders during a flash sale, where each order updates stock levels and generates an invoice.
Implementation: Use a worker pool to handle order requests and update the database concurrently.

2. Real-Time Notifications
Simple use case: Send order confirmation emails or SMS to multiple customers.
Example: Simulate 100 notifications being sent out at the same time.
Implementation: Use goroutines to send notifications via an API in parallel.

3. Search and Filtering
Simple use case: Allow users to search a product catalog with filters like price or category.
Example: Build a mock product catalog with 10,000 items and simulate concurrent searches.
Implementation: Use goroutines to handle multiple search queries and respond quickly.

4. Inventory Updates
Simple use case: Update stock levels from multiple warehouses simultaneously.
Example: Simulate 10 concurrent stock updates for the same product.
Implementation: Use locks (e.g., sync.Mutex) to manage concurrent updates safely.

5. Recommendation Engine
Simple use case: Generate product recommendations for multiple users concurrently.
Example: Simulate 20 users viewing products and getting personalized recommendations.
Implementation: Use multithreading to process user activity and return recommendations.

6. Analytics Dashboard
Simple use case: Update a dashboard showing total sales or user activity in real time.
Example: Simulate concurrent data aggregation from 5 different sources (e.g., orders, payments).
Implementation: Use channels to collect and process data streams concurrently.

7. Marketplace Vendor Sync
Simple use case: Sync product catalogs from multiple vendors.
Example: Simulate importing product data from 3 vendors with 1,000 products each.
Implementation: Use goroutines to fetch and process vendor data in parallel.
