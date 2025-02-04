#!/bin/bash

echo "🚀 Stopping any existing containers..."
docker-compose down

echo "🔄 Removing old Docker images..."
docker rmi -f go-app || true

echo "🐳 Building and running the application..."
docker-compose up --build