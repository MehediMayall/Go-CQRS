# Go CQRS Without Library
A straightforward Go app that demonstrates the implementation of CQRS (Command Query Responsibility Segregation) without requiring any library.
![CQRS Architecture](https://github.com/user-attachments/assets/8e889640-63b5-4873-a91e-737aa115a5a3)


## Folder Structure
   * [cmd](./cmd)
       * [main.go](./cmd/main.go)
   * [internal](./internal)
       * [app](./internal/app)
         * [app.go](./app/app.go)
         * [dependencies.go](./app/dependencies.go)
         * [routes.go](./app/routes.go)
       * [controllers](./internal/controllers)
         *[movieReadController.go](./controllers/movieReadController.go)
         *[movieWriteController.go](./controllers/movieWriteController.go)
         *[response.go](./controllers/response.go)
       * [database](./internal/database)
       * [entities](./internal/entities)
       * [handlers](./internal/handlers)
       * [repositories](./internal/repositories)

