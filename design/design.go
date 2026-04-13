package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("hello", func() {
	Description("A simple service that says hello.")

	Method("sayHello", func() {
		Payload(String, "Name to greet")
		Result(String, "A greeting message")

		HTTP(func() {
			GET("/hello/{name}")
		})
	})
})
