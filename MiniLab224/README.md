# MiniLab 2

## [Link to the Postman workspace](https://www.postman.com/science-architect-46593263/workspace/minilab2)

## How to start the server locally

Install docker on your OS. [Windows link](https://docs.docker.com/desktop/setup/install/windows-install/)

Run the command in the console in the root folder (MiniLab224)
```shell
docker-compose up --build
```

Done! The server is listening at http://localhost:8080

## What the app does

The application can:
- get the top coins, get information about a specific coin and list of coin markets using api https://api.coinlore.net
- endorse coins - the list of IDs of endorsed coins is stored in the database, the coin can also be removed from the endorsed ones 