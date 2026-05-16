package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Nihal", func() {

	Method("AddBorrower", func() {
		Payload(borrower, "Borrower struct coming as a payload")
		Result(func() {
			Field(1, "statusCode", Int32)
			Field(2, "message", String)
		})
		Error("borrower_already_exists", BorrowerAlreadyExists, "Borrower Already Exists")
		Error("internal_server_error", InternalServerError, "Internal Server Error")

		HTTP(func() {
			POST("/add/borrower")
			Response(StatusCreated)
			Response("borrower_already_exists", StatusConflict)
			Response("internal_server_error", StatusInternalServerError)
		})
	})
})

var borrower = Type("Borrower", func() {
	Description("Borrower struct")

	Field(1, "FirstName", String)
	Field(2, "LastName", String)
	Field(3, "Age", Int32)
	Field(4, "Address", String)
	Field(5, "Email", String, func() {
		Format(FormatEmail)
	})

})

var BorrowerAlreadyExists = Type("Borrower_Exists", func() {
	Description("Borrower already exists")

	Field(1, "message", String)
	Field(2, "statusCode", Int32)
	Required("message")
})

var InternalServerError = Type("Internal_Server_Error", func() {
	Description("Internal Server Error")

	Field(1, "message", String)
	Field(2, "statusCode", Int32)
	Required("message")
})
