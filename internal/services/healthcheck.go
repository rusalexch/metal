package services

func NewHealthCheck(repo Repositorier) *HealthCheck {
	return &HealthCheck{
		repo: repo,
	}
}

func (h *HealthCheck) Ping() error {
	return h.repo.Ping()
}
