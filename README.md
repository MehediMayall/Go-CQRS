# Go CQRS Without Library
A straightforward Go app that demonstrates the implementation of CQRS (Command Query Responsibility Segregation) without requiring any library.
![CQRS Architecture](https://github.com/user-attachments/assets/8e889640-63b5-4873-a91e-737aa115a5a3)


## Folder Structure

├── cmd
│   └── main.go
├── docs
├── go-cqrs.http
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   ├── dependencies.go
│   │   └── routes.go
│   ├── controllers
│   │   ├── movieReadController.go
│   │   ├── movieWriteController.go
│   │   └── response.go
│   ├── database
│   │   └── inMemoryDatabase.go
│   ├── entities
│   │   ├── entityBase.go
│   │   └── movie.go
│   ├── handlers
│   │   ├── abstractions
│   │   │   ├── ICommandHandler.go
│   │   │   └── IQueryHandler.go
│   │   ├── commands
│   │   │   ├── addMovieCommandHandler.go
│   │   │   ├── deleteMovieCommandHandler.go
│   │   │   └── updateMovieCommandHandler.go
│   │   └── queries
│   │       ├── getMovieByIdQueryHandler.go
│   │       └── getMoviesQueryHandler.go
│   └── repositories
│       ├── abstractions
│       │   ├── IReadRepository.go
│       │   └── IWriteRepository.go
│       ├── movieReadRepository.go
│       └── movieWriteRepository.go
