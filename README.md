# Delta
The mathematical symbol for change. 

Delta is a turn-based online word game.

This project is for me to learn and pratice my skills.


# Table of Contents
- [Overview](#overview)
- [Gettings Started](#getting-started)
- [Architecture](#architecture)
- [References](#references)

# Overview

# Getting Started

# Architecture
This project uses a microservices architecture

### Data Storage
***Redis*** is used to store hot data:
- Live game data
- Matchmaking queue
- User sessions
- Pub/Sub

***Postgres*** is used to store persistent data:
- Game history
- User credentials and profile

### Services
***Worker*** handles user authentication, matchmaking, fetching game data and validation game moves.

***Parrot*** holds websocket connections and acts as a real-time event delivery system between services and the clients.

***Updater*** ends games when a move has not been made after 100 seconds.


# References
Resources I used while creating this project
- https://redis.io/blog/how-to-create-a-real-time-online-multi-player-strategy-game-using-redis/
- https://github.com/gofiber/recipes
- [Video - How Roblox Keeps Millions of Users up to Date with Redis Pub/Sub](https://www.youtube.com/watch?v=nXTxXSWBayg&list=TLGGrVGwPBoa-gcxOTExMjAyNA)