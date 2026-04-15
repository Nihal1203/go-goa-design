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

	Method("getPerson", func() {
		Payload(func() {
			Field(1, "id", Int64)
			Required("id")
		})
		Result(Person)
		HTTP(func() {
			GET("/get/person/{id}")
		})
	})

	Method("addPerson", func() {
		Payload(Person)
		Result(AddPersonResponse)

		Error("person_already_exists", PersonAlreadyExists)
		Error("internal_error", InternalError)

		HTTP(func() {
			POST("/add/person")

			Response(StatusOK) // success
			Response("person_already_exists", StatusConflict)
			Response("internal_error", StatusInternalServerError)
		})
	})

	Method("deletePerson", func() {
		Payload(func() {
			Attribute("id", Int32)
		})
		Result(Boolean)
		Error("internal_error", InternalError)

		HTTP(func() {
			DELETE("/delete/person/{id}")
			Response(StatusOK)
			Response("internal_error", StatusInternalServerError)

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

var AddPersonResponse = Type("AddPersonResponse", func() {
	Description("Add person success response")

	Field(1, "success", Boolean)
	Required("success")
})

var PersonAlreadyExists = Type("PersonAlreadyExists", func() {
	Description("Person already exists")

	Field(1, "message", String)
	Required("message")
})

var InternalError = Type("InternalError", func() {
	Description("Internal server error")

	Field(1, "message", String)
	Required("message")
})
