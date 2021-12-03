build-plugins:
	go build --buildmode=plugin -o ./plugins/karmada.so ./plugins/karmada
	go build --buildmode=plugin -o ./plugins/clusterapi.so ./plugins/clusterapi

run: build-plugins
	go build -o globalscheme main.go
	./globalscheme
