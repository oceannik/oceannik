# Oceannik

This repository contains the core system of Oceannik, including the CLI application for controlling the system, named `ocean`.

```
ocean: a CLI management tool for Oceannik instances.

Visit https://github.com/oceannik/oceannik for more information.

Usage:
  ocean [command]

Available Commands:
  deployments List and schedule new Deployments
  help        Help about any command
  namespaces  List, create and update Namespaces
  projects    List, create and update Projects
  secrets     List, create and update Secrets in Namespaces
  server      Run an Agent server
  version     Display version

Flags:
  -c, --config-dir string   config directory (default "$HOME/.oceannik/")
  -h, --help                help for ocean
      --host string         host to connect to/host to run the server on
  -n, --namespace string    namespace to use for managing resources on the Agent (default "default")
      --port int            port to connect to/port to run the server on (default 5000)

Use "ocean [command] --help" for more information about a command.
```

## Building the `ocean` binary

To build the binary, issue the following command:

```
make build
```

This will create a new binary under the `bin/ocean` location.

## Generating certificates for mutual authentication

The cert generation script requires `openssl` to be installed.

To generate new keys, run the `scripts/generate-certs.sh` script.  
The Makefile provides a helpful command for this purpose. 

```
make gen-certs
```

This will create a set of certificates under the `generated-certs/` directory. 
If everything went as expected and no errors were thrown, you can copy the certificates to the Oceannik configuration directory.

```
make copy-certs
```

This command will create the `~/.oceannik/certs` directory, and copy all the generated keys there.  
This directory is used by the Agent and the Client by default to make developing locally easier.
This means you can run the Oceannik Agent and Client on localhost, and still use mTLS for authentication between the parties.  
To achieve this, Oceannik overrides the Server Names in the certificates to `host.oceannik.local`.

## Examples

Examples are provided in the [oceannik/examples](https://github.com/oceannik/examples) repository.
