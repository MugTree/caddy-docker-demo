# Caddy webserver demo

The Docker Compose file sets up three services: Caddy as a reverse proxy with config and data
volumes and access to Docker, plus two test applications, all connected to a shared network called "proxy".

I've written the dummy apps in Go, was just easier than the .net apps as I was confused about the port mapping with Kestrel.

## Basic idea..

### I've set up my hosts file

127.0.0.1 gaming.localhost

A request to

#### https://gaming.localhost/spins

...should proxy content from the app running at 172.18.0.4:8000 on the docker network

#### https://gaming.localhost/poker

... should proxy content from the app running at 172.18.0.3:8001 on the docker network

**! This now works - ignore the certificate warnings... **

### If you look at the Caddyfile you should see...

                gaming.localhost {

                    handle_path /spins {
                        reverse_proxy 172.18.0.4:8000
                    }

                    handle_path /poker {
                        reverse_proxy 172.18.0.3:8001
                    }
                }

As far as my understanding of Caddy goes the `handle_path` directive should strip off a path that matches /spins/
and proxy back content from the root of the app.

If you look at the apps the handler that gets hit in each app is the root handler. Used a little go framework called chi that simplifies the routing through regexes.

### Useful commands

docker compose up -d  
docker compose down
docker network inspect caddy-proxy-demo_proxy
docker network prune
docker run -d -p 9000:8000 poker-test-app
docker build -t poker-test-app .

## Portainer

This looks like it could work
https://www.youtube.com/watch?v=qj45uHP7Jmo
https://www.youtube.com/watch?v=QBNaOdNSsx8
