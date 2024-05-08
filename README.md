# Caddy webserver demo

The Docker Compose file sets up three services: Caddy as a reverse proxy with config and data
volumes and access to Docker, plus two test applications, all connected to a shared network called "proxy".

I've written the dummy apps in Go, was just easier than the .net apps as I was confused about the port mapping with Kestrel.

## Basic idea..

### I've set up my hosts file

127.0.0.1 gaming.localhost

In theory a request to

#### http://gaming.localhost/spins

...should proxy content from the app running at localhost:8000

#### http://gaming.localhost/poker

... should proxy content from the app running at localhost:9000

It's not working at the moment though, will get to to the bottom of it :)

### If you look at the Caddyfile you should see...

                gaming.localhost {

                    handle_path /spins/* {
                        reverse_proxy localhost:8000
                    }

                    handle_path /poker/* {
                        reverse_proxy localhost:9000
                    }
                }

As far as my understanding of Caddy goes the `handle_path` directive should strip off a path that matches /spins/
and proxy back content from the root of the app.