# template-project

template-project is a Golang-based open-source project designed to provide a robust template for building scalable and efficient applications. It includes best practices for project structure, configuration management, logging, and testing.

## Features

* Well-structured project layout
* Configurable settings via environment variables and config files
* Built-in logging using logrus
* CI/CD ready with GitHub Actions
* Docker support for containerized deployment

## Run

To run the project directly, execute the following command:

```shell
go run main.go module.go
```

### Environment Variables

Get the list of environment variables that can be used to configure the application:

```shell
go run main.go module.go --print_configs
```

## License

This project is licensed under the Apache License. See the LICENSE file for details.
