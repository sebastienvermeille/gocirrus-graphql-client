# DEV notes

## Graphql schema for cirrus ci comes from:
https://github.com/cirruslabs/cirrus-ci-web/blob/master/schema.gql

This has to be removed:

```
"""
A 64-bit signed integer
"""
scalar Int
```

## Generate the client:

Run `go run github.com/Khan/genqlient genqlient.yaml`

# Note
All of that is automated via commands in the makefile 
and then reused by the CI.
