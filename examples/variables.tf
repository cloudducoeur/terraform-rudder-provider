variable "rudder_api_url" {
  description = "URL de l'API Rudder"
  type        = string
  default     = "https://localhost:443/rudder/api"
}

variable "rudder_api_token" {
  description = "Token d'authentification pour l'API Rudder"
  type        = string
  sensitive   = true
}

variable "node_id" {
  description = "ID du nœud Rudder à configurer"
  type        = string
}
