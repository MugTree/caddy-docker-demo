# Caddy webserver demo

The Docker Compose file sets up three services: Caddy as a reverse proxy with config and data
volumes and access to Docker, plus two test applications, all connected to a shared network called "proxy".

I've written the dummy apps in Go, was just easier than the .net apps as I was confused about the port mapping with Kestrel.

**!! Opened ports 8000, 9000 on the game containers for debugging purposes !!**

## Basic idea..

### I've set up my hosts file

127.0.0.1 gaming.localhost

In theory a request to

#### http://gaming.localhost/spins

...should proxy content from the app running at 172.18.0.4:8000 on the docker network

#### http://gaming.localhost/poker

... should proxy content from the app running at 172.18.0.3:8000 on the docker network

It's not working at the moment though, will get to to the bottom of it :)

### If you look at the Caddyfile you should see...

                gaming.localhost {

                    handle_path /spins/* {
                        reverse_proxy 172.18.0.4:8000
                    }

                    handle_path /poker/* {
                        reverse_proxy 172.18.0.3:8000
                    }
                }

As far as my understanding of Caddy goes the `handle_path` directive should strip off a path that matches /spins/
and proxy back content from the root of the app.

### Useful commands

docker compose up -d  
docker compose down
docker network inspect caddy-proxy-demo_proxy
docker network prune
