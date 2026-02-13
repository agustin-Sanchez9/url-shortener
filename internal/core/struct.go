package internal

import "time"

// estructura base para guardar en la db. Propenso a cambiarse en futuro
type ShortedURL struct{
	id int
	code string
	url string
	lastVisited time.Time
	visits int32
}