run:
	go run main.go

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

