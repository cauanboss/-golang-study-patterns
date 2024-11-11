package auth

//func AuthMiddleware(authService AuthService) func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			token := strings.TrimSpace(r.Header.Get("Authorization"))
//			if _, err := authService.ValidateToken(token); err != nil {
//				http.Error(w, "Forbidden", http.StatusForbidden)
//				return
//			}
//			next.ServeHTTP(w, r)
//		})
//	}
//}
