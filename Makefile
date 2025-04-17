.PHONY: dev

dev:
	@echo "🚀 Checking and killing old process on port 8080..."
	@lsof -ti :8080 | xargs kill -9 || echo " No process to kill"
	@echo "🔄 Starting the server with Air..."
	air
