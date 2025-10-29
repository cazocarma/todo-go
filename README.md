# todo-go

Una aplicación backend sencilla de lista de tareas (“Todo”) desarrollada con **Go**. Sirve como punto de partida para aprender buenas prácticas de Go, API RESTful, contenedores y arquitectura limpia.

## 🔍 ¿Qué incluye?

- Código en Go estructurado siguiendo principios de arquitectura limpia (o separación de responsabilidades).
- API HTTP para crear, listar, actualizar y eliminar tareas.
- Migraciones o esquema para base de datos (dependiendo de la implementación del repositorio).
- Configuración para desarrollo y despliegue (por ejemplo Docker) si aplica.
- Código modular, limpio y preparado para ser extendido.

## 🚀 Guía de instalación rápida

1. Clona el repositorio  
   ```bash
   git clone https://github.com/cazocarma/todo-go.git
   cd todo-go
   ```

2. Crea o copia tu archivo de configuración (por ejemplo `config.yaml`, `.env`, etc) y define variables como:

   ```text
   DB_HOST=…
   DB_PORT=…
   DB_USER=…
   DB_PASS=…
   DB_NAME=…
   ```

3. (Opcional) Levanta servicios si se usa Docker

   ```bash
   docker-compose up -d
   ```

4. Instala dependencias (Go modules)

   ```bash
   go mod tidy
   ```

5. Ejecuta migraciones / prepara la base de datos

   ```bash
   # ejemplo
   go run cmd/migrate/main.go
   ```

6. Ejecuta la aplicación

   ```bash
   go run main.go
   # o bien
   go build -o todo-go && ./todo-go
   ```

7. Accede al API

   ```text
   http://localhost:8080/api/tasks
   ```

## 🛠 Uso básico

* **GET** `/api/tasks` → Listar todas las tareas.
* **POST** `/api/tasks` → Crear una nueva tarea (payload JSON con título, descripción, etc).
* **PUT** `/api/tasks/{id}` → Actualizar una tarea existente (por ejemplo marcar como completada).
* **DELETE** `/api/tasks/{id}` → Eliminar una tarea.

## 📁 Estructura del proyecto

```
├── cmd/                # Punto de entrada(s) de la aplicación
├── internal/           # Lógica de negocio, servicios, repositorios
├── pkg/                # Paquetes reutilizables
├── migrations/         # Scripts o archivos de migración
├── api/                # Controladores, rutas HTTP
├── go.mod
├── go.sum
└── README.md
```

> Ajusta los nombres según lo que realmente tenga el repositorio.

## ✅ Buenas prácticas aplicadas

* Código limpio: funciones cortas, nombres auto-explicativos, separación de capas.
* Uso de módulos de Go (`go mod`) para gestión de dependencias.
* Manejo de errores consistente, logging básico.
* Variables de entorno/archivo de configuración para datos sensibles (DB, puerto, etc).
* Preparado para escalar: nuevo servicio, nueva ruta, nuevo repositorio se pueden agregar sin gran refactor.

## 🔧 Extensiones posibles

* Añadir autenticación/autorización (JWT, OAuth2) para que cada usuario tenga su propia lista de tareas.
* Implementar filtros en la API (por estado, fecha de vencimiento, usuario).
* Añadir un frontend sencillo (React, Vue, Angular) que consuma esta API.
* Migrar a microservicios: separación de servicio de tareas del servicio de usuarios.
* Añadir WebSockets o eventos para notificaciones en tiempo real cuando una tarea cambia.
* Añadir pruebas unitarias/integración para endpoints y lógica de negocio.
* Implementar CI/CD (por ejemplo con GitHub Actions) para testing y despliegue automático.

## 🎓 Para quién es este proyecto

Este proyecto es ideal para:

* Desarrolladores que quieren aprender Go y cómo construir APIs RESTful.
* Quienes desean una base limpia y mínima para extender o modificar según necesidades.
* Proyectos de prototipo rápido donde se requiere backend de tareas sin mucho overhead.
* Entornos de aprendizaje y demostración de buenas prácticas de código.

## 📄 Licencia

Este proyecto se distribuye bajo la licencia **MIT**. Puedes usar, modificar y distribuir según los términos de MIT.

---

¡Gracias por revisar este proyecto! Si te resulta útil, considera dejar una estrella ⭐ en el repositorio y si haces mejoras o extensiones, puedes abrir un Pull Request o compartir tus ideas.