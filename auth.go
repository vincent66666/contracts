package contracts

type GuardDriver func(name string, config Fields, ctx Context, provider UserProvider) Guard
type UserProviderDriver func(config Fields) UserProvider

type Auth interface {
	ExtendUserProvider(name string, provider UserProviderDriver)
	ExtendGuard(name string, guard GuardDriver)

	Guard(name string, ctx Context) Guard
	UserProvider(name string) UserProvider
}

type Authenticatable interface {
	GetId() string
}

type Guard interface {
	Once(user Authenticatable)
	User() Authenticatable
	GetId() string
	Check() bool
	Guest() bool
	Login(user Authenticatable) interface{}
}

type UserProvider interface {
	RetrieveById(identifier string) Authenticatable
}

type Authorizable interface {
	Can(ability string, arguments ...interface{}) bool
}

type GateChecker func(user Authorizable, data interface{}) bool

type Policy map[string]GateChecker

type Gate interface {

	// Allows determined if the given ability should be granted for the current user.
	Allows(ability string, arguments ...interface{}) bool

	// Denies Determine if the given ability should be denied for the current user.
	Denies(ability string, arguments ...interface{}) bool

	// Check Determine if all the given abilities should be granted for the current user.
	Check(abilities []string, arguments ...interface{}) bool

	// Any Determine if any one of the given abilities should be granted for the current user.
	Any(abilities []string, arguments ...interface{}) bool

	// Authorize Determine if the given ability should be granted for the current user.
	Authorize(ability string, arguments ...interface{})

	// Inspect the user for the given ability.
	Inspect(ability string, arguments ...interface{}) HttpResponse

	// Raw Get the raw result from the authorization callback.
	Raw(ability string, arguments ...interface{}) interface{}

	// ForUser Get a guard instance for the given user.
	ForUser(user Authorizable) Gate
}

type GateFactory interface {

	// Has determined if a given ability has been defined.
	Has(ability string) bool

	// Define a new ability.
	Define(ability string, callback GateChecker) GateFactory

	// Resource define abilities for a resource.
	Resource(name string, class Class, abilities ...string) GateFactory

	// Policy define a policy class for a given class type.
	Policy(class Class, policy Policy) GateFactory

	// Before Register a callback to run before all Gate checks.
	Before(callable GateChecker) GateFactory

	// After Register a callback to run after all Gate checks.
	After(callable GateChecker) GateFactory

	// Abilities Get all the defined abilities.
	Abilities() []string
}
