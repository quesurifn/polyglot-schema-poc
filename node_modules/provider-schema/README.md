# The Polyglot Friendly Provider Schema

## Basic Concept

- Have the schema in `/schema/provider.json`
- Have both package.json and go.mod in root
- Have a github action to update version for both package.json and go.mod based on tagging


## Result

- Schema version locally will be able to be managed without the added complexity of helm by using each languages' package manager


## Caveats 

- The Go based schema will be read via `os.Open` and hence the schema will need to be cached after reading upon server boot