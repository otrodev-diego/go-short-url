# DTE Short URL Service

Este proyecto es una prueba técnica desarrollada en Go para crear un servicio que genera enlaces cortos (ShortURL) asociados a Documentos Tributarios Electrónicos (DTE), esto es una prueba.

## 📌 Objetivo

- Generar enlaces cortos únicos para acceder a DTEs.
- Permita redirigir a la URL original simulada.
- Controlar la expiración del enlace.
- Utilizar token de seguridad JWT

---

## ⚙️ Tecnologías utilizadas

- 🐹 Go 1.24.4
- 💡 Arquitectura Hexagonal (Domain / Application / Infrastructure)
- 🗂️ Almacenamiento en memoria
- 🐳 Docker
- 🚀 Git

---

## 📁 Estructura del Proyecto

dte-shortener/
├── cmd/shortener/ # Entrada principal (main.go)
├── internal/
│ ├── application/ # Casos de uso (servicios)
│ ├── domain/ # Entidades del negocio
│ └── infrastructure/
│ └── http/ # Adaptadores HTTP


## 🚀 Cómo levantar el proyecto
