# [OpenAPI](https://www.openapis.org/) Specification 3.0.3
The OpenAPI Specification (OAS) defines a standard, language-agnostic interface to RESTful 
APIs which allows both humans and computers to discover and understand the capabilities of 
the service without access to source code, documentation, or through network traffic 
inspection. 

[Read more...](https://github.com/OAI/OpenAPI-Specification/blob/3.0.3/versions/3.0.3.md#specification)

## [Schema Objects](/objects)
The schema exposes two types of fields: `Fixed` fields, which have a declared name, and 
`Patterned` fields, which declare a regex pattern for the field name.

All field names in the specification are **case sensitive**. This includes all fields that are 
used as keys in a map, except where explicitly noted that keys are case insensitive.

Includes [examples](/objects/testdata) represented  in JSON and YAML format.

### YAML
Struct fields are unmarshalled using the field name lowercased as the default key.
The Fixed field in the Info Object `termsOfService` would become `termsofservice`.
This is not correct since field names are case sensitive.
Some field tags will be needed to prevent this.
