.PHONY: dev

dev:
	@echo "🚀 Checking and killing old process on port 8081..."
	@lsof -ti :8081 | xargs kill -9 || echo " No process to kill"
	@echo "🔄 Starting the server with Air..."
	air
