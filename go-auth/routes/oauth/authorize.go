package oauth

type AuthorizeRequestQuery struct {
	ResponseType string `query:"response_type"`
	ClientId     string `query:"client_id"`
	RedirectURI  string `query:"redirect_uri"`
	Scope        string `query:"scope"`
}
