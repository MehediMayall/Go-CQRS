# Go CQRS Without Library
A straightforward Go app that demonstrates the implementation of CQRS (Command Query Responsibility Segregation) without requiring any library.
![CQRS Architecture](https://github.com/user-attachments/assets/8e889640-63b5-4873-a91e-737aa115a5a3)


### Folder Structure
```md

├── cmd
│   └── main.go
├── docs
├── go-cqrs.http
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── dependencies.go
│   │   └── routes.go
│   ├── controllers
│   │   ├── movieReadController.go
│   │   ├── movieWriteController.go
│   │   └── response.go
│   ├── database
│   │   └── inMemoryDatabase.go
│   ├── entities
│   │   ├── entityBase.go
│   │   └── movie.go
│   ├── handlers
│   │   ├── abstractions
│   │   │   ├── ICommandHandler.go
│   │   │   └── IQueryHandler.go
│   │   ├── commands
│   │   │   ├── addMovieCommandHandler.go
│   │   │   ├── deleteMovieCommandHandler.go
│   │   │   └── updateMovieCommandHandler.go
│   │   └── queries
│   │       ├── getMovieByIdQueryHandler.go
│   │       └── getMoviesQueryHandler.go
│   └── repositories
│       ├── abstractions
│       │   ├── IReadRepository.go
│       │   └── IWriteRepository.go
│       ├── movieReadRepository.go
│       └── movieWriteRepository.go

```

***ReadRepository*** has only one responsibility that is to read data from database.
It is used to transfer data from database to the application as efficient as possible without changing its state

***WriteRepository*** has only one responsibility that is to write data to database.
It is used to transfer data from application to the database in more generic ways. 

***Manual Dependency Injection***  [dependencies.go] file manages the dependencies of the application.
Every functions in this file is used to instantiate the dependent objects and inject them into the parent object through constructor injection.

### Dependency Flow
![Dependency Flow](https://github.com/user-attachments/assets/8c1ca414-d653-4500-8c23-047a80972ce9)
