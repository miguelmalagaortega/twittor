package models

import "time"

type GraboTweet struct {
	// recordar que en bson la esttructura es
	// nombre tipo `bson:"nombreQueBuscaBD" json:"ParaRepresentacionJson"
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
