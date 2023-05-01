package handler

func (s Server) SetupRoutes() {
	// TODO: Add role based authorization
	api := s.App.Group("/api/auth")
	api.Get("/details/:id", s.GetUserById)
	api.Get("/details/:email", s.GetUserByEmail)
	api.Get("/details/:username", s.GetUserByUsername)

	api.Post("/login", s.SignIn)
	api.Post("/signup", s.SignUp)
	api.Get("/logout", s.LogoutUser)

	api.Get("/validate", s.ValidateToken)
	api.Get("/refresh", s.RefreshToken)

}
