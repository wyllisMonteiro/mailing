// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/wyllisMonteiro/mailing/api/docs/restapi/operations/pet"
	"github.com/wyllisMonteiro/mailing/api/docs/restapi/operations/store"
	"github.com/wyllisMonteiro/mailing/api/docs/restapi/operations/user"
)

// NewSwaggerPetstoreAPI creates a new SwaggerPetstore instance
func NewSwaggerPetstoreAPI(spec *loads.Document) *SwaggerPetstoreAPI {
	return &SwaggerPetstoreAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer:          runtime.JSONConsumer(),
		MultipartformConsumer: runtime.DiscardConsumer,
		UrlformConsumer:       runtime.DiscardConsumer,
		XMLConsumer:           runtime.XMLConsumer(),

		JSONProducer: runtime.JSONProducer(),
		XMLProducer:  runtime.XMLProducer(),

		PetAddPetHandler: pet.AddPetHandlerFunc(func(params pet.AddPetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.AddPet has not yet been implemented")
		}),
		UserCreateUserHandler: user.CreateUserHandlerFunc(func(params user.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUser has not yet been implemented")
		}),
		UserCreateUsersWithArrayInputHandler: user.CreateUsersWithArrayInputHandlerFunc(func(params user.CreateUsersWithArrayInputParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUsersWithArrayInput has not yet been implemented")
		}),
		UserCreateUsersWithListInputHandler: user.CreateUsersWithListInputHandlerFunc(func(params user.CreateUsersWithListInputParams) middleware.Responder {
			return middleware.NotImplemented("operation user.CreateUsersWithListInput has not yet been implemented")
		}),
		StoreDeleteOrderHandler: store.DeleteOrderHandlerFunc(func(params store.DeleteOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation store.DeleteOrder has not yet been implemented")
		}),
		PetDeletePetHandler: pet.DeletePetHandlerFunc(func(params pet.DeletePetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.DeletePet has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		}),
		PetFindPetsByStatusHandler: pet.FindPetsByStatusHandlerFunc(func(params pet.FindPetsByStatusParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.FindPetsByStatus has not yet been implemented")
		}),
		PetFindPetsByTagsHandler: pet.FindPetsByTagsHandlerFunc(func(params pet.FindPetsByTagsParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.FindPetsByTags has not yet been implemented")
		}),
		StoreGetInventoryHandler: store.GetInventoryHandlerFunc(func(params store.GetInventoryParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation store.GetInventory has not yet been implemented")
		}),
		StoreGetOrderByIDHandler: store.GetOrderByIDHandlerFunc(func(params store.GetOrderByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation store.GetOrderByID has not yet been implemented")
		}),
		PetGetPetByIDHandler: pet.GetPetByIDHandlerFunc(func(params pet.GetPetByIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.GetPetByID has not yet been implemented")
		}),
		UserGetUserByNameHandler: user.GetUserByNameHandlerFunc(func(params user.GetUserByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserByName has not yet been implemented")
		}),
		UserLoginUserHandler: user.LoginUserHandlerFunc(func(params user.LoginUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LoginUser has not yet been implemented")
		}),
		UserLogoutUserHandler: user.LogoutUserHandlerFunc(func(params user.LogoutUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.LogoutUser has not yet been implemented")
		}),
		StorePlaceOrderHandler: store.PlaceOrderHandlerFunc(func(params store.PlaceOrderParams) middleware.Responder {
			return middleware.NotImplemented("operation store.PlaceOrder has not yet been implemented")
		}),
		PetUpdatePetHandler: pet.UpdatePetHandlerFunc(func(params pet.UpdatePetParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.UpdatePet has not yet been implemented")
		}),
		PetUpdatePetWithFormHandler: pet.UpdatePetWithFormHandlerFunc(func(params pet.UpdatePetWithFormParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.UpdatePetWithForm has not yet been implemented")
		}),
		UserUpdateUserHandler: user.UpdateUserHandlerFunc(func(params user.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		}),
		PetUploadFileHandler: pet.UploadFileHandlerFunc(func(params pet.UploadFileParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation pet.UploadFile has not yet been implemented")
		}),

		// Applies when the "api_key" header is set
		APIKeyAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		},
		PetstoreAuthAuth: func(token string, scopes []string) (interface{}, error) {
			return nil, errors.NotImplemented("oauth2 bearer auth (petstore_auth) has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*SwaggerPetstoreAPI This is a sample server Petstore server.  You can find out more about     Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).      For this sample, you can use the api key `special-key` to test the authorization     filters. */
type SwaggerPetstoreAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer
	// MultipartformConsumer registers a consumer for the following mime types:
	//   - multipart/form-data
	MultipartformConsumer runtime.Consumer
	// UrlformConsumer registers a consumer for the following mime types:
	//   - application/x-www-form-urlencoded
	UrlformConsumer runtime.Consumer
	// XMLConsumer registers a consumer for the following mime types:
	//   - application/xml
	XMLConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer
	// XMLProducer registers a producer for the following mime types:
	//   - application/xml
	XMLProducer runtime.Producer

	// APIKeyAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key api_key provided in the header
	APIKeyAuth func(string) (interface{}, error)

	// PetstoreAuthAuth registers a function that takes an access token and a collection of required scopes and returns a principal
	// it performs authentication based on an oauth2 bearer token provided in the request
	PetstoreAuthAuth func(string, []string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// PetAddPetHandler sets the operation handler for the add pet operation
	PetAddPetHandler pet.AddPetHandler
	// UserCreateUserHandler sets the operation handler for the create user operation
	UserCreateUserHandler user.CreateUserHandler
	// UserCreateUsersWithArrayInputHandler sets the operation handler for the create users with array input operation
	UserCreateUsersWithArrayInputHandler user.CreateUsersWithArrayInputHandler
	// UserCreateUsersWithListInputHandler sets the operation handler for the create users with list input operation
	UserCreateUsersWithListInputHandler user.CreateUsersWithListInputHandler
	// StoreDeleteOrderHandler sets the operation handler for the delete order operation
	StoreDeleteOrderHandler store.DeleteOrderHandler
	// PetDeletePetHandler sets the operation handler for the delete pet operation
	PetDeletePetHandler pet.DeletePetHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// PetFindPetsByStatusHandler sets the operation handler for the find pets by status operation
	PetFindPetsByStatusHandler pet.FindPetsByStatusHandler
	// PetFindPetsByTagsHandler sets the operation handler for the find pets by tags operation
	PetFindPetsByTagsHandler pet.FindPetsByTagsHandler
	// StoreGetInventoryHandler sets the operation handler for the get inventory operation
	StoreGetInventoryHandler store.GetInventoryHandler
	// StoreGetOrderByIDHandler sets the operation handler for the get order by Id operation
	StoreGetOrderByIDHandler store.GetOrderByIDHandler
	// PetGetPetByIDHandler sets the operation handler for the get pet by Id operation
	PetGetPetByIDHandler pet.GetPetByIDHandler
	// UserGetUserByNameHandler sets the operation handler for the get user by name operation
	UserGetUserByNameHandler user.GetUserByNameHandler
	// UserLoginUserHandler sets the operation handler for the login user operation
	UserLoginUserHandler user.LoginUserHandler
	// UserLogoutUserHandler sets the operation handler for the logout user operation
	UserLogoutUserHandler user.LogoutUserHandler
	// StorePlaceOrderHandler sets the operation handler for the place order operation
	StorePlaceOrderHandler store.PlaceOrderHandler
	// PetUpdatePetHandler sets the operation handler for the update pet operation
	PetUpdatePetHandler pet.UpdatePetHandler
	// PetUpdatePetWithFormHandler sets the operation handler for the update pet with form operation
	PetUpdatePetWithFormHandler pet.UpdatePetWithFormHandler
	// UserUpdateUserHandler sets the operation handler for the update user operation
	UserUpdateUserHandler user.UpdateUserHandler
	// PetUploadFileHandler sets the operation handler for the upload file operation
	PetUploadFileHandler pet.UploadFileHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SwaggerPetstoreAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SwaggerPetstoreAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SwaggerPetstoreAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SwaggerPetstoreAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SwaggerPetstoreAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SwaggerPetstoreAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SwaggerPetstoreAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SwaggerPetstoreAPI
func (o *SwaggerPetstoreAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}
	if o.MultipartformConsumer == nil {
		unregistered = append(unregistered, "MultipartformConsumer")
	}
	if o.UrlformConsumer == nil {
		unregistered = append(unregistered, "UrlformConsumer")
	}
	if o.XMLConsumer == nil {
		unregistered = append(unregistered, "XMLConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}
	if o.XMLProducer == nil {
		unregistered = append(unregistered, "XMLProducer")
	}

	if o.APIKeyAuth == nil {
		unregistered = append(unregistered, "APIKeyAuth")
	}
	if o.PetstoreAuthAuth == nil {
		unregistered = append(unregistered, "PetstoreAuthAuth")
	}

	if o.PetAddPetHandler == nil {
		unregistered = append(unregistered, "pet.AddPetHandler")
	}
	if o.UserCreateUserHandler == nil {
		unregistered = append(unregistered, "user.CreateUserHandler")
	}
	if o.UserCreateUsersWithArrayInputHandler == nil {
		unregistered = append(unregistered, "user.CreateUsersWithArrayInputHandler")
	}
	if o.UserCreateUsersWithListInputHandler == nil {
		unregistered = append(unregistered, "user.CreateUsersWithListInputHandler")
	}
	if o.StoreDeleteOrderHandler == nil {
		unregistered = append(unregistered, "store.DeleteOrderHandler")
	}
	if o.PetDeletePetHandler == nil {
		unregistered = append(unregistered, "pet.DeletePetHandler")
	}
	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}
	if o.PetFindPetsByStatusHandler == nil {
		unregistered = append(unregistered, "pet.FindPetsByStatusHandler")
	}
	if o.PetFindPetsByTagsHandler == nil {
		unregistered = append(unregistered, "pet.FindPetsByTagsHandler")
	}
	if o.StoreGetInventoryHandler == nil {
		unregistered = append(unregistered, "store.GetInventoryHandler")
	}
	if o.StoreGetOrderByIDHandler == nil {
		unregistered = append(unregistered, "store.GetOrderByIDHandler")
	}
	if o.PetGetPetByIDHandler == nil {
		unregistered = append(unregistered, "pet.GetPetByIDHandler")
	}
	if o.UserGetUserByNameHandler == nil {
		unregistered = append(unregistered, "user.GetUserByNameHandler")
	}
	if o.UserLoginUserHandler == nil {
		unregistered = append(unregistered, "user.LoginUserHandler")
	}
	if o.UserLogoutUserHandler == nil {
		unregistered = append(unregistered, "user.LogoutUserHandler")
	}
	if o.StorePlaceOrderHandler == nil {
		unregistered = append(unregistered, "store.PlaceOrderHandler")
	}
	if o.PetUpdatePetHandler == nil {
		unregistered = append(unregistered, "pet.UpdatePetHandler")
	}
	if o.PetUpdatePetWithFormHandler == nil {
		unregistered = append(unregistered, "pet.UpdatePetWithFormHandler")
	}
	if o.UserUpdateUserHandler == nil {
		unregistered = append(unregistered, "user.UpdateUserHandler")
	}
	if o.PetUploadFileHandler == nil {
		unregistered = append(unregistered, "pet.UploadFileHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SwaggerPetstoreAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SwaggerPetstoreAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api_key":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.APIKeyAuth)

		case "petstore_auth":
			result[name] = o.BearerAuthenticator(name, o.PetstoreAuthAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *SwaggerPetstoreAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerPetstoreAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		case "multipart/form-data":
			result["multipart/form-data"] = o.MultipartformConsumer
		case "application/x-www-form-urlencoded":
			result["application/x-www-form-urlencoded"] = o.UrlformConsumer
		case "application/xml":
			result["application/xml"] = o.XMLConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *SwaggerPetstoreAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		case "application/xml":
			result["application/xml"] = o.XMLProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SwaggerPetstoreAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the swagger petstore API
func (o *SwaggerPetstoreAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SwaggerPetstoreAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet"] = pet.NewAddPet(o.context, o.PetAddPetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user"] = user.NewCreateUser(o.context, o.UserCreateUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/createWithArray"] = user.NewCreateUsersWithArrayInput(o.context, o.UserCreateUsersWithArrayInputHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/createWithList"] = user.NewCreateUsersWithListInput(o.context, o.UserCreateUsersWithListInputHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/store/order/{orderId}"] = store.NewDeleteOrder(o.context, o.StoreDeleteOrderHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/pet/{petId}"] = pet.NewDeletePet(o.context, o.PetDeletePetHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/user/{username}"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/findByStatus"] = pet.NewFindPetsByStatus(o.context, o.PetFindPetsByStatusHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/findByTags"] = pet.NewFindPetsByTags(o.context, o.PetFindPetsByTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/store/inventory"] = store.NewGetInventory(o.context, o.StoreGetInventoryHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/store/order/{orderId}"] = store.NewGetOrderByID(o.context, o.StoreGetOrderByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/pet/{petId}"] = pet.NewGetPetByID(o.context, o.PetGetPetByIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{username}"] = user.NewGetUserByName(o.context, o.UserGetUserByNameHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/login"] = user.NewLoginUser(o.context, o.UserLoginUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/logout"] = user.NewLogoutUser(o.context, o.UserLogoutUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/store/order"] = store.NewPlaceOrder(o.context, o.StorePlaceOrderHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/pet"] = pet.NewUpdatePet(o.context, o.PetUpdatePetHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet/{petId}"] = pet.NewUpdatePetWithForm(o.context, o.PetUpdatePetWithFormHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/{username}"] = user.NewUpdateUser(o.context, o.UserUpdateUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/pet/{petId}/uploadImage"] = pet.NewUploadFile(o.context, o.PetUploadFileHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SwaggerPetstoreAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SwaggerPetstoreAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SwaggerPetstoreAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SwaggerPetstoreAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *SwaggerPetstoreAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
