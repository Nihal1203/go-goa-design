package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("user", func() {
	Method("getUser", func() {
		Payload(String)
		Result(String)

		HTTP(func() {
			GET("/user/{id}")
		})
	})

	Method("printPerson", func() {
		Payload(Person)
		Result(ConfigMap)
		HTTP(func() {
			POST("/printPerson")
		})
	})

	Method("addPerson", func() {
		Payload(Person)
		Result(Bytes)
		HTTP(func() {
			POST("/add/person")
		})
	})
})

var Person = Type("Person", func() {
	Description("Person struct")

	Field(1, "name", String)
	Field(2, "age", Int64)
	Field(3, "mobileNo", String)
	Field(4, "email", String, func() {
		Format(FormatEmail)
	})
	Field(5, "id", Int64)
})

var ConfigMap = MapOf(Int32, Person, func() {

})
