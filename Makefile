# Makefile for GoMall Dev Environment

DC=docker-compose
APP_NAME=go_mall

up:
	@echo "🚀 Starting $(APP_NAME) services..."
	$(DC) up -d

down:
	@echo "🛑 Stopping $(APP_NAME) services..."
	$(DC) down

restart:
	@echo "🔁 Restarting $(APP_NAME) services..."
	$(DC) down
	$(DC) up -d

logs:
	@echo "📜 Showing logs..."
	$(DC) logs -f

ps:
	@echo "📋 Listing running containers..."
	docker ps

redis-cli:
	@echo "💬 Entering Redis CLI..."
	docker exec -it go_mall_redis redis-cli

pg-shell:
	@echo "🧾 Connecting to PostgreSQL..."
	docker exec -it go_mall_postgres psql -U devuser -d go_mall

mq-ui:
	@echo "🌐 Opening RabbitMQ UI at http://localhost:15672"

redis-ui:
	@echo "🌐 Opening Redis Insight at http://localhost:5540"

clean:
	@echo "🧹 Stopping and removing volumes..."
	$(DC) down -v

help:
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:' Makefile | awk -F: '{print " - " $$1}'