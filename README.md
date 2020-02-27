# Creature manager

Creature manager aims to be a simple RESTful API someone could use to keep track
of various creatures, playable or else, in a _Dungeons and Dragons, 5th edition_
game. It provides various endpoints to fetch various information about a
creature or roll dice.

The project is still a work in progress, so expect some of the
functionality/endpoints to be missing, unless documented here or at least
present in `main.go`.

## Installation

To install it type:

    go get -u github.com/aakordas/creature_manager

The only dependency, currently, is GorillaMux, with MongoDB to be added later.

## Documentation

For documentation on the available endpoints, objects and available
functionality, check the
[wiki](https://github.com/aakordas/creature_manager/wiki).
