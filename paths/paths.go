package paths

// Стандартные пути для всех сервисов PSDS (health, ready, swagger).
// Используются в приложении и пробах (k8s, load balancer).
const (
	PathHealth  = "/health"
	PathReady   = "/ready"
	PathSwagger = "/swagger"
)
