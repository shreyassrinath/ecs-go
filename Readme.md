# Entity Component System in GO

## Description

### What is Entity Component System?

Entity-component system (ECS) is an architectural pattern that follows Composition over inheritance principle

### Why ECS? 

ECS allows greater flexibility in defining entities where every object in a game's scene is an entity (e.g. enemies, bullets, vehicles, etc.). Every Entity consists of one or more components which add additional behavior or functionality. Therefore, the behavior of an entity can be changed at **runtime** by adding or removing components. This eliminates the ambiguity problems of deep and wide inheritance hierarchies that are difficult to understand, maintain and extend. 

## Main concepts

#### Entity
The entity is a general purpose object. Usually, it only consists of a unique id. They "tag every coarse gameobject as a separate item". Implementations typically use a plain integer for this.
#### Component
 The raw data for one aspect of the object, and how it interacts with the world. "Labels the Entity as possessing this particular aspect". Implementations typically use Structs, Classes, or Associative Arrays.
#### System
 "Each System runs continuously (as though each System had its own private thread) and performs global actions on every Entity that possesses a Component of the same aspect as that System."

## Project details
This project includes proof of concept ECS example project. The project sets up Go Dev environment using vagrant. 

### Project folder details
```
EntityComponentSystem   
    | -- ecs
            | -- component.go
            | -- entitymanager.go   
   
    | -- run
            | -- main.go
   
    | -- subsystems
                | -- cleanup.go
                | -- fallback.go
                | -- movement.go
                | -- render.go
                | -- system.go
   
    | -- wireup
            | -- Initialize.go

```
#### ecs

Contains core of ECS. component.go has components info and entitymanager.go manages "entities" and "components". Basically entitymanager.go acts as the Datastore and management of ECS. 

#### run

run main.go to start the project

#### subsystems

subsystems implement System interface from system.go. subsystems run on seperate **go routines**.

#### wireup

Creates new entities, adds components and starts the susbsystems.      

## Installation requirements

* Install [VirtualBox](https://www.virtualbox.org/wiki/Downloads) and [Vagrant](https://www.vagrantup.com/downloads.html)
* Clone this repo

## Commands

Start:

`$ vagrant up`

This does the following:

* Installs Git
* Downloads Go version 1.8.3 and installs it

SSH into the server

`$ vagrant ssh`

Check Go version- This should display the version of go installed.

`$ go version`

cd into project run directory

` $ cd ECS-GO/src/EntityComponentSystem/run/`

run the project

`$ go run main.go`

## Reference:

[Wiki Article](https://en.wikipedia.org/wiki/Entity%E2%80%93component%E2%80%93system)

[Understanding ECS](https://www.gamedev.net/resources/_/technical/game-programming/understanding-component-entity-systems-r3013)

[Implementing ECS](https://www.gamedev.net/resources/_/technical/game-programming/implementing-component-entity-systems-r3382)

[Nice post on implementing ECS in games](https://www.raywenderlich.com/24878/introduction-to-component-based-architecture-in-games)
