# Compila el proyecto y genera el ejecutable "app"
build:
	go build -o app main.go

# Ejecuta la compilación y corre la aplicación
run: build
	./app

# Elimina el ejecutable "app"
clean:
	rm -f app
