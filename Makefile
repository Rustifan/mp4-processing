DB_NAME = mp4_db
DB_USER = user
DB_PASSWORD = password
DB_HOST = db
DB_PORT = 5432

db-connect:
	docker exec -it mp4_db psql postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

db-push: 
	docker exec -it mp4_api npm run db push