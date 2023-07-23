# keymaster
An AWS Keypair CLI Tool

## Installation
To get the tool, simply install using go:
```sh
go install github.com/SHA65536/keymaster@main
```

## List
Use list to show a list of all your keypairs across given regions. Leaving the region flag empty will look in all regions.
```
$ keymaster list -h
NAME:
   Keymaster list - list aws keypairs

USAGE:
   Keymaster list [command options] [arguments...]

OPTIONS:
   --regions value, -r value [ --regions value, -r value ]  regions to show keys in
   --help, -h            
```

## Copy
Use copy to copy a keypair from one region to another.
```
$ keymaster copy -h
NAME:
   Keymaster copy - copy a keypair from one region to another

USAGE:
   Keymaster copy [command options] [arguments...]

OPTIONS:
   --src-key value, --sk value     name of the source key
   --src-region value, --sr value  name of the source region
   --dst-key value, --dk value     name of the destination key
   --dst-region value, --dr value  name of the destination region
   --help, -h                      show help
```