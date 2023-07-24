# keymaster
An AWS Keypair CLI Tool

## Installation
To get the tool, simply install using go:
```sh
go install github.com/SHA65536/keymaster@main
```

## List
Use list to show a list of all your keypairs across given regions. Leaving the region flag empty will look in all regions.
Example `keymaster list -r us-west-1 -r us-west-2`
```
$ keymaster list -h
NAME:
   Keymaster list - list aws keypairs

USAGE:
   Keymaster list [command options] [arguments...]

OPTIONS:
   --regions value, -r value [ --regions value, -r value ]  regions to show keys in
   --help, -h                                               show help    
```

## Copy
Use copy to copy a keypair from one region to another.
Example `keymaster copy --from us-west-1 --to us-east-1 mykeyname`
```
$ keymaster copy -h
NAME:
   Keymaster copy - copy a keypair from one region to another

USAGE:
   Keymaster copy [command options] keyname

OPTIONS:
   --from value, --sr value     name of the source region
   --to value, --dr value       name of the destination region
   --dst-key value, --dk value  name of the destination key
   --help, -h                   show help
```