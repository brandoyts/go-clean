start: 
	make stop && docker compose up -d --build 
	
stop:
	docker compose down --volumes

go:
	go run main.go