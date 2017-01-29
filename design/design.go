package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("template-ad", func() { // API defines the microservice endpoint and
	Title("The virtual wine cellar")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Scheme("http")                      // the design.
	Host("localhost:8080")
	BasePath("/api")
})

var _ = Resource("users", func() { // Resources group related API endpoints
	BasePath("/v1/users")   // together. They map to REST resources for REST
	DefaultMedia(UserMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get user by username") // with its path, parameters (both path
		Routing(GET("/:userName"))          // parameters and querystring values) and payload
		Params(func() {                     // (shape of the request body).
			Param("userName", Integer, "User Name")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

var UserMedia = MediaType("application/vnd.goa.tests.users+json", func() {
	Description("A bottle of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})

var _ = Resource("bottles", func() { // Resources group related API endpoints
	BasePath("/v1/bottles")   // together. They map to REST resources for REST
	DefaultMedia(BottleMedia) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get user by username") // with its path, parameters (both path
		Routing(GET("/:bottles"))           // parameters and querystring values) and payload
		Params(func() {                     // (shape of the request body).
			Param("bottles", Integer, "Bottles ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})

// BottleMedia defines the media type used to render bottles.
var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	Description("A bottle of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})
