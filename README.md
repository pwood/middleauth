# middleauth

[![license](https://img.shields.io/github/license/pwood/middleauth.svg)](https://github.com/pwood/middleauth/blob/master/LICENSE)
[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg)](https://github.com/RichardLitt/standard-readme)
[![Actions Status](https://github.com/pwood/middleauth/workflows/main/badge.svg)](https://github.com/pwood/middleauth/actions)

> A configurable authentication middleware for usage with [traefik](https://github.com/traefik/traefik).

## Table of Contents

- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Background

`traefik` is a reverse proxy that provides a convenient way to authenticate and authorise requests through the usage of
a middleware. This is a simple implementation to provide basic functionality.

## Install

It is recommended that this software is installed through the use of Docker containers.

```
docker pull ghcr.io/pwood/middleauth:latest
```

There are containers available for the following architectures:

* `linux/amd64`
* `linux/arm64`
* `linux/arm/v7`
* `linux/arm/v6`

Otherwise, you can install with Go directly using:

`go install github.com/pwood/middleauth`

## Usage

### Configuration

`middleauth` supports configuration through environmental variables, in the future this may expanded to a JSON or YAML
configuration file.

| Name                 | Default | Purpose                                                                                    |
|----------------------|---------|--------------------------------------------------------------------------------------------|
| `CONFIG_MODE`        | `ENV`   | Select between different configuration modes. Currently only ENV supported.                |
| `CONFIG_FILE`        |         | Filename to load configuration from, currently ignored.                                    |
| `SERVER_HOST`        | `[::]`  | IP address to bind listening port to, default all V4 and V6 addresses.                     |
| `SERVER_PORT`        | `8888`  | TCP port to open HTTP server on.                                                           |
| `PERMITTED_NETWORKS` |         | Comma seperated lists of V4/V6 CIDR networks to access to sides protected by `middleauth`. |

### Deploying

It is recommended that you use `docker-compose` to deploy `middleauth` on a Docker host. 

An example `docker-compose.yml` follows, it assumes you have port 80 and 443 open on `traefik` with the entrypoints 
named `web` and `websecure` respectively. It also services the `middleauth` webserver on a subdomain of `auth.example.org`.
This will be used as `middleauth` develops to optionally show a login screen.

```yaml
version: '3.5'
services:
  middleauth:
    restart: unless-stopped
    build: .
    environment:
     - PERMITTED_NETWORKS=10.10.0.0/24
     - SERVER_PORT=9991
    labels:
     - "traefik.enable=true"
     - "traefik.http.routers.auth.rule=Host(`auth.example.org`)"
     - "traefik.http.routers.auth.entrypoints=websecure"
     - "traefik.http.routers.auth.tls.certresolver=myletsencrypt"

     - "traefik.http.routers.auth-http.rule=Host(`auth.example.org`)"
     - "traefik.http.routers.auth-http.entrypoints=web"
     - "traefik.http.routers.auth-http.middlewares=https_redirect"

     - "traefik.http.services.auth.loadbalancer.server.port=9991"

     - "traefik.http.middlewares.https_redirect.redirectscheme.scheme=https"
    networks:
      traefik: {}

networks:
  traefik:
    name: traefik_default
```

### Protecting Other Services

To protect another service with `middleauth` add a label to that services `docker-compose.yml`.

```yaml
services:
  protectedservice:
    labels:
     - "traefik.http.middlewares.auth.forwardauth.address=http://middleauth:9991/api/check"
```

## Maintainers

[@pwood](https://github.com/pwood)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/pwood/middleauth/issues/new) or submit PRs.

This project follows the [Contributor Covenant](https://www.contributor-covenant.org/version/1/4/code-of-conduct/) Code
of Conduct.

## License

Copyright 2022 Peter Wood & Contributors

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
License. You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "
AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific
language governing permissions and limitations under the License.