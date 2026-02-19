package core

import "time"

// estructura base para guardar en la db. Propenso a cambiarse en futuro
// recordar iniciar con mayus los campos para que sean publicos y poder ser usados por el driver de DB
type ShortedURL struct{
	ID int
	code string
	URL string
	LastVisited time.Time
	Visits int32
}