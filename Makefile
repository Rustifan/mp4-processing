include .env

DB_NAME = mp4_db
DB_USER = user
DB_PASSWORD = password
DB_HOST = db
DB_PORT = 5432


api-exec:
	docker exec -it mp4_api bash

api-test:
	docker exec -it mp4_api npm run test

db-connect:
	docker exec -it mp4_db psql postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

db-push: 
	docker exec -it mp4_api npm run db push

db-studio: 
	open https://local.drizzle.studio?port=${DRIZLE_STUDIO_PORT}
	docker exec -it mp4_api npm run db:studio

	
