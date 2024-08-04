# Tiltfile

# Define the PostgreSQL instance
k8s_yaml("postgres.yaml")

# Define the backend service
k8s_yaml("backend.yaml")

# Set up local development environment
local("docker build -t my-go-app .")  # Build your Go application image
local("kubectl apply -f postgres.yaml")  # Apply Kubernetes manifest for the PostgreSQL instance
local("kubectl apply -f backend.yaml")  # Apply Kubernetes manifest for the backend service

# Watch for changes in your Go code and automatically restart the backend service
k8s_resource("backend").sync("src", "/cmd/web")

# Print URLs for accessing the services
print("Backend Service URL: http://localhost:8080")
print("PostgreSQL URL: postgres://localhost:5432/your_database_name")
