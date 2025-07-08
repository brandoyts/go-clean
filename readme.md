# A simple todo application built in golang using clean architecture.

Make sure you have your docker installed and running

To run the application
`make start`

To stop the application
`make stop`

### Layers

- domain
- service
- infrastructure
- adapters

### Project tree

- adapters

  - mongodb
  - mysql

- infrastructure
  - rest
- internal
  - controller
  - domain
  - repository
  - service
  - utils
