package authn

const (
	// UserDatabase names the mongo database.
	UserDatabase = "user"

	// CredentialsCollection names the credentials collection
	CredentialsCollection = "credentials"
)

type AuthNAPI struct {
	collection DBCollection
}

func NewAuthNAPI(collection DBCollection) *AuthNAPI {
	return &AuthNAPI{
		collection: collection,
	}
}

// API Dependencies

type DBCollection interface {
}
