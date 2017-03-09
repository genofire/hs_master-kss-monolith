# Getting Started

1.  Install Docker.
2.  Set up an environment variable `LAB_MONOLITH_DB_HOST` and assign the docker host's IP,
    set `LAB_MONOLITH_DB_PORT` to some free port, e.g. 5432.

    For docker machine type `docker-machine env` and the value of `$DOCKER_HOST`,
    otherwise `localhost` is probably correct.

3.  Set up an environment variable `LAB_MONOLITH_DB_PASSWORD` and assign any password. 

4.  `./start-db.sh`

    Starts a PostgreSQL database using the password from `LAB_MONOLITH_DB_PASSWORD`.

5.  `gradle bootRun`
6.  Open http://localhost:8080
