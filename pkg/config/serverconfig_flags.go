// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (ServerConfig) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (ServerConfig) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in ServerConfig and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg ServerConfig) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("ServerConfig", pflag.ExitOnError)
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "httpPort"), defaultServerConfig.HTTPPort, "On which http port to serve admin")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "grpcPort"), defaultServerConfig.GrpcPort, "On which grpc port to serve admin")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "kube-config"), defaultServerConfig.KubeConfig, "Path to kubernetes client config file.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "master"), defaultServerConfig.Master, "The address of the Kubernetes API server.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.secure"), defaultServerConfig.Security.Secure, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.ssl.certificateFile"), defaultServerConfig.Security.Ssl.CertificateFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.ssl.keyFile"), defaultServerConfig.Security.Ssl.KeyFile, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.useAuth"), defaultServerConfig.Security.UseAuth, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.clientId"), defaultServerConfig.Security.Oauth.ClientId, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.clientSecretFile"), defaultServerConfig.Security.Oauth.ClientSecretFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.baseUrl"), defaultServerConfig.Security.Oauth.BaseURL, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.authorizeUrl"), defaultServerConfig.Security.Oauth.AuthorizeUrl, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.tokenUrl"), defaultServerConfig.Security.Oauth.TokenURL, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.callbackUrl"), defaultServerConfig.Security.Oauth.CallbackURL, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.claims.aud"), defaultServerConfig.Security.Oauth.Claims.Audience, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.claims.iss"), defaultServerConfig.Security.Oauth.Claims.Issuer, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.idpUserInfoEndpoint"), defaultServerConfig.Security.Oauth.IdpUserInfoEndpoint, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.cookieHashKeyFile"), defaultServerConfig.Security.Oauth.CookieHashKeyFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.cookieBlockKeyFile"), defaultServerConfig.Security.Oauth.CookieBlockKeyFile, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.redirectUrl"), defaultServerConfig.Security.Oauth.RedirectURL, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.httpAuthorizationHeader"), defaultServerConfig.Security.Oauth.HTTPAuthorizationHeader, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "security.oauth.grpcAuthorizationHeader"), defaultServerConfig.Security.Oauth.GrpcAuthorizationHeader, "")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "security.allowCors"), defaultServerConfig.Security.AllowCors, "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "security.allowedOrigins"), []string{}, "")
	cmdFlags.StringSlice(fmt.Sprintf("%v%v", prefix, "security.allowedHeaders"), []string{}, "")
	return cmdFlags
}
