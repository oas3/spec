package objects_test

import (
    "encoding/json"
    "github.com/oas3/spec/objects"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "testing"
)

var schemesObj = []objects.SecurityScheme{
    {
        Type:   "http",
        Scheme: "basic",
    },
    {
        Type: "apiKey",
        Name: "api_key",
        In:   "header",
    },
    {
        Type:         "http",
        Scheme:       "bearer",
        BearerFormat: "JWT",
    },
    {
        Type: "oauth2",
        Flows: objects.OAuthFlows{
            Implicit: objects.OAuthFlow{
                AuthorizationURL: "https://example.com/api/oauth/dialog",
                Scopes: map[string]string{
                    "write:pets": "modify pets in your account",
                    "read:pets":  "read your pets",
                },
            },
        },
    },
    {
        Type: "oauth2",
        Flows: objects.OAuthFlows{
            Implicit: objects.OAuthFlow{
                AuthorizationURL: "https://example.com/api/oauth/dialog",
                Scopes: map[string]string{
                    "write:pets": "modify pets in your account",
                    "read:pets":  "read your pets",
                },
            },
            AuthorizationCode: objects.OAuthFlow{
                AuthorizationURL: "https://example.com/api/oauth/dialog",
                TokenURL:         "https://example.com/api/oauth/token",
                Scopes: map[string]string{
                    "write:pets": "modify pets in your account",
                    "read:pets":  "read your pets",
                },
            },
        },
    },
}

func TestSecurityScheme(t *testing.T) {
    t.Run("JSON", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/security_schemes.json")
        var schemes []objects.SecurityScheme
        if err := json.Unmarshal(raw, &schemes); err != nil {
            t.Error(err)
        }
        eqSecuritySchemes(t, schemesObj, schemes)
    })

    t.Run("YAML", func(t *testing.T) {
        raw, _ := ioutil.ReadFile("testdata/security_schemes.yml")
        var schemes []objects.SecurityScheme
        if err := yaml.Unmarshal(raw, &schemes); err != nil {
            t.Error(err)
        }
        eqSecuritySchemes(t, schemesObj, schemes)
    })
}

func eqSecuritySchemes(t *testing.T, ss1, ss2 []objects.SecurityScheme) {
    eqInt(t, len(ss1), len(ss2))
    for i, s1 := range ss1 {
        s2 := ss2[i]
        eqSecurityScheme(t, s1, s2)
    }
}

func eqSecurityScheme(t *testing.T, s1, s2 objects.SecurityScheme) {
    eqStr(t, s1.Type, s2.Type)
    eqStr(t, s1.Description, s2.Description)
    eqStr(t, s1.Name, s2.Name)
    eqStr(t, s1.In, s2.In)
    eqStr(t, s1.Scheme, s2.Scheme)
    eqStr(t, s1.BearerFormat, s2.BearerFormat)
    eqOAuthFlows(t, s1.Flows, s2.Flows)
    eqStr(t, s1.OpenIDConnectURL, s2.OpenIDConnectURL)
}

func eqOAuthFlows(t *testing.T, f1, f2 objects.OAuthFlows) {
    eqOAuthFlow(t, f1.Implicit, f2.Implicit)
    eqOAuthFlow(t, f1.Password, f2.Password)
    eqOAuthFlow(t, f1.ClientCredentials, f2.ClientCredentials)
    eqOAuthFlow(t, f1.AuthorizationCode, f2.AuthorizationCode)
}

func eqOAuthFlow(t *testing.T, f1, f2 objects.OAuthFlow) {
    eqStr(t, f1.AuthorizationURL, f2.AuthorizationURL)
    eqStr(t, f1.TokenURL, f2.TokenURL)
    eqStr(t, f1.RefreshURL, f2.RefreshURL)
    eqInt(t, len(f1.Scopes), len(f2.Scopes))
    for k, s1 := range f1.Scopes {
        s2 := f2.Scopes[k]
        eqStr(t, s1, s2)
    }
}
