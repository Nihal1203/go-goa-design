package design

import . "goa.design/goa/v3/dsl"

var _ = Service("user", func() {
	Method("getUser", func() {
		Payload(String)
		Result(String)

		HTTP(func() {
			GET("/user/{id}")
		})
	})
})