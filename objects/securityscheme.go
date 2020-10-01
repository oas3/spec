package objects

// SecurityScheme defines a security scheme that can be used by the operations.
type SecurityScheme struct {
    // The type of the security scheme.
    Type string `oas3:"REQUIRED"`
    // A short description for security scheme.
    Description string
    // The name of the header, query or cookie parameter to be used.
    Name string `oas3:"REQUIRED"`
    // The location of the API key. Valid values are "query", "header" or "cookie".
    In string `oas3:"REQUIRED"`
    // The name of the HTTP Authorization scheme to be used in the Authorization header
    // as defined in RFC7235.
    Scheme string `oas3:"REQUIRED"`
    // A hint to the client to identify how the bearer token is formatted.
    BearerFormat string `yaml:"bearerFormat"`
    // An object containing configuration information for the flow types supported.
    Flows OAuthFlows `oas3:"REQUIRED"`
    // OpenId Connect URL to discover OAuth2 configuration values.
    OpenIDConnectURL string `yaml:"openIdConnectUrl" oas3:"REQUIRED"`
}

// OAuthFlows are configuration details for a supported OAuth Flow.
type OAuthFlows struct {
    // Configuration for the OAuth Implicit flow
    Implicit OAuthFlow
    // Configuration for the OAuth Resource Owner Password flow.
    Password OAuthFlow
    // Configuration for the OAuth Client Credentials flow. Previously called
    // `application` in OpenAPI 2.0.
    ClientCredentials OAuthFlow `yaml:"clientCredentials"`
    // Configuration for the OAuth Authorization Code flow. Previously called `accessCode`
    // in OpenAPI 2.0.
    AuthorizationCode OAuthFlow `yaml:"authorizationCode"`
}

// OAuthFlow allows configuration of the supported OAuth Flows.
type OAuthFlow struct {
    // The authorization URL to be used for this flow.
    AuthorizationURL string `yaml:"authorizationUrl" oas3:"REQUIRED"`
    // The token URL to be used for this flow.
    TokenURL string `yaml:"tokenUrl" oas3:"REQUIRED"`
    // The URL to be used for obtaining refresh tokens.
    RefreshURL string `yaml:"refreshUrl"`
    // The available scopes for the OAuth2 security scheme. A map between the scope name
    // and a short description for it. The map may be empty.
    Scopes map[string]string `oas3:"REQUIRED"`
}
