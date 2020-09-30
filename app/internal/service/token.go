package service

// GetTokenByUserID 获取用户Token
func (s *Service) GetTokenByUserID(userID int) string {
	return s.dao.GetTokenByUserID(userID)
}
