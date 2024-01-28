all: app

app:
	go run cmd/app/main.go

run:
	go run cmd/jungkurt/main.go

run1:
	go run cmd/signintech/main.go

tidy:
	go mod tidy
	go mod vendor

clean:
	rm output/*