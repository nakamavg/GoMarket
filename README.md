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

