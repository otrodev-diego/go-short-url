# DTE Short URL Service

Este proyecto es una prueba técnica desarrollada en Go para crear un servicio que genera enlaces cortos (ShortURL) asociados a Documentos Tributarios Electrónicos (DTE), esto es una prueba.

## 📌 Objetivo

- Generar enlaces cortos únicos para acceder a DTEs.
- Permita redirigir a la URL original simulada.
- Controlar la expiración del enlace.
- Utilizar token de seguridad JWT

## 🔍 Características
- ✨ Generación de URLs cortas únicas
- 🔒 Autenticación mediante JWT
- ⏱️ Control de expiración de enlaces
- 📊 Contador de accesos
- 🔄 Redirección a URL original
---

## ⚙️ Tecnologías utilizadas

- 🐹 Go 1.24.4
- 💡 Arquitectura Hexagonal (Domain / Application / Infrastructure)
- 🗂️ Almacenamiento en memoria
- 🐳 Docker
- 🚀 Git

---
## 📁 Estructura del Proyecto
```bash
dte-shortener/
├── cmd/shortener/ # Punto de entrada
│   └── main.go # Archivo principal
├── internal/ # Directorio interno
│   ├── domain/ # Entidades y reglas de negocio
│   ├── application/ # Casos de uso
│   └── infrastructure/ # Adaptadores (HTTP, seguridad)
│       ├── http/
│       └── security/
├── Dockerfile # Configuración de Docker
└── go.mod # Dependencias del proyecto
```
---
## 🚀 Cómo levantar el proyecto

### Clonar proyecto 
https://github.com/otrodev-diego/go-short-url.git
### Construir Imagen
docker build -t go-short-url .
### Ejecutar contenedor
docker run -p 8081:8081 go-short-url
#### el servicio estará disponible en http://localhost:8081/


## 🧪 Pruebas
*1. Crear una URL corta*
```markdown
curl -X POST http://localhost:8081/shorten \
-H "Content-Type: application/json" \
-d '{"dte_id":"DTE001","expire_in_minutes":60}'
```

*Respuesta JSON*
```markdown
{
    "short_code": "abc123",
    "token": "eyJhbGciOiJIUzI1..."
}
```

*2.Prueba redirección*
 Reemplaza ABC123 con el short_code que recibiste y TOKEN con el token

```markdown
curl -v -H "Authorization: Bearer TOKEN" http://localhost:8081/s/ABC123
```