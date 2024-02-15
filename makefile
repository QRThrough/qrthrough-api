# mode process
mode := dev

# load environment
ifneq (,$(wildcard config/.env.${mode}))
    include config/.env.${mode}
    export
endif

# define Host service and App name
REMOTE_HOST := qrthrough
APP_NAME := $(APP_NAME)

run: stopair
	go run .
stopair:
	-kill $$(lsof -ti:$(PORT)) 2>/dev/null

dbreset: dbrollback dbmigrate dbmock dbalumni
dbcreate:
	migrate create -ext sql -dir db/migrations/ -seq qrthrough
dbmigrate:
	migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_NAME)?sslmode=disable -path ./db/migrations up
dbrollback:
	migrate -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_NAME)?sslmode=disable -path ./db/migrations down
dbmock:
	mode=${mode} go run ./cmd/mock_db/main.go
dbalumni:
	go run ./cmd/alumni_csv/main.go

# deploy command
build-deploy:
	docker build --platform linux/amd64 -t ${APP_NAME} .
	docker save ${APP_NAME} > ${APP_NAME}.tar
	docker rmi ${APP_NAME}
	scp ./${APP_NAME}.tar ${REMOTE_HOST}:/root/
	rm ./${APP_NAME}.tar
	ssh -t ${REMOTE_HOST} 'docker rm $$(docker ps -aqf "name=${APP_NAME}") -f \
    &&  docker rmi $$(docker images -aqf "reference=${APP_NAME}") \
    &&  docker load < /root/${APP_NAME}.tar \
    &&  rm /root/${APP_NAME}.tar \
    &&  docker run -d -p 8000:8000 --name ${APP_NAME} ${APP_NAME}'
init-deploy:
	docker build --platform linux/amd64 -t ${APP_NAME} .
	docker save ${APP_NAME} > ${APP_NAME}.tar
	docker rmi ${APP_NAME}
	scp ./${APP_NAME}.tar ${REMOTE_HOST}:/root/
	rm ./${APP_NAME}.tar
	ssh -t ${REMOTE_HOST} 'docker load < /root/${APP_NAME}.tar \
    &&  rm /root/${APP_NAME}.tar \
    &&  docker run -d -p 8000:8000 --name ${APP_NAME} ${APP_NAME}'