package handler

func (s Server) SetupRoutes() {
	// TODO: Add role based authorization
	api := s.App.Group("/api/auth")
	api.Get("/details/:id", s.GetUserById)
	api.Get("/:email", s.GetUserByEmail)
	api.Get("/:username", s.GetUserByUsername)

	api.Post("/login", s.SignIn)
	api.Post("/signup", s.SignUp)
	api.Get("/validate", s.ValidateToken)
	api.Get("/refresh", s.RefreshToken)

	userApi := s.App.Group("/api/users")
	userApi.Use(s.DeserializeUser())
	userApi.Get("/me", s.GetMe)
	userApi.Get("/logout", s.DeserializeUser(), s.LogoutUser)
}
