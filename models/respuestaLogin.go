package models

//RespuestaLogin tiene el token que se delvuelve en el login
type RespuestaLogin struct {
	Token string `json:"token,omitempty"`
}
