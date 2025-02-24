# Simulación de Mercado de Acciones 🚀

API REST que simula un mercado de valores con capacidad para manejar órdenes de compra/venta, emparejamiento de operaciones y actualización dinámica de precios.

## 📋 Requisitos Previos

- Go 1.21+
- Docker 20.10+ (Opcional)
- Git 2.30+

## 🚀 Instalación Rápida

### Con Docker (Recomendado)
```bash
# Clonar repositorio
git clone https://github.com/nakamavg/GoMarket.git
cd GoMarket

# Construir y ejecutar
docker compose up --build

# Verificar funcionamiento
curl http://localhost:8080/acciones/health
```

### Sin Docker
```bash
# Clonar repositorio
git clone https://github.com/nakamavg/GoMarket.git
cd GoMarket

# Instalar dependencias
go mod tidy

# Compilar y ejecutar
make run
```

## 📚 Documentación de la API

### Crear una Acción
```http
POST /acciones
Content-Type: application/json

{
  "simbolo": "AAPL",
  "precio_inicial": 150.00
}
```

### Crear una Orden de Compra
```http
POST /ordenes
Content-Type: application/json

{
  "usuario_id": 1,
  "tipo": "compra",
  "simbolo": "AAPL",
  "cantidad": 5,
  "precio_limite": 152.00
}
```

### Crear una Orden de Venta
```http
POST /ordenes
Content-Type: application/json

{
  "usuario_id": 2,
  "tipo": "venta",
  "simbolo": "AAPL",
  "cantidad": 5,
  "precio_limite": 150.00
}
```

### Obtener el Libro de Órdenes
```http
GET /ordenes?simbolo=AAPL
```

### Obtener el Precio de una Acción
```http
GET /acciones/AAPL
```

## 🛠️ Estructura del Proyecto

```
market/
├── internal/
│   ├── api/
│   │   ├── order_handler.go
│   │   └── stock_handler.go
│   ├── order/
│   │   ├── datastructures/
│   │   │   └── heaps.go
│   │   ├── engine/
│   │   │   └── matching_engine.go
│   │   ├── model/
│   │   │   └── order.go
│   │   ├── repository/
│   │   │   └── memory_repository.go
│   │   └── service/
│   │       └── order_service.go
│   └── stock/
│       └── memory_repository.go
├── main.go
├── Makefile
└── go.mod
```

## 🧪 Pruebas

Puedes utilizar el archivo `test.http` para realizar pruebas rápidas de la API utilizando una extensión como "REST Client" en VSCode.

```http
# Crear una Acción
POST localhost:3000/acciones
Content-Type: application/json

{
  "simbolo": "AAPL",
  "precio_inicial": 150.00
}

###

# Crear una Orden de Compra
POST localhost:3000/ordenes
Content-Type: application/json

{
  "usuario_id": 1,
  "tipo": "compra",
  "simbolo": "AAPL",
  "cantidad": 5,
  "precio_limite": 152.00
}

###

# Crear una Orden de Venta
POST localhost:3000/ordenes
Content-Type: application/json

{
  "usuario_id": 2,
  "tipo": "venta",
  "simbolo": "AAPL",
  "cantidad": 5,
  "precio_limite": 150.00
}

###

# Obtener el Libro de Órdenes
GET localhost:3000/ordenes?simbolo=AAPL

###

# Obtener el Precio de una Acción
GET localhost:3000/acciones/AAPL
```

## 📄 Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.

