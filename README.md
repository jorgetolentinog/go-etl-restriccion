[![Go Report Card](https://goreportcard.com/badge/github.com/jorgetolentinog/go-etl-restriccion)](https://goreportcard.com/report/github.com/jorgetolentinog/go-etl-restriccion)

# ETL Restricción de cliente

Este programa extra la lista de clientes de la base de datos del sistema **ludoplay** y la guarda en la base de datos del sistema **operation-manager**

---

### Instalación

Requiere:
- golang 1.15.6

```
go mod download
```

### Configuración
Clone el archivo `config.example.yml` en `config.yml` y reemplace las variables necesarias.

### Ejecutar en desarrollo
```
go run main.go
```

### Ejecutar pruebas unitarias
```
make test
```

### Compilar para windows
```
make build_windows
```
generara un binario en ./bin/app.exe

### Compilar para linux
```
make build_linux
```
generara un binario en ./bin/app
