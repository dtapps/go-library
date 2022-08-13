package gocache

import "log"

func (r *Redis) setLog(key string) {
	if r.config.Debug == true {
		log.Printf("gocache [%s] set\n", key)
	}
}
func (r *Redis) getLog(key string) {
	if r.config.Debug == true {
		log.Printf("gocache [%s] get\n", key)
	}
}
func (r *Redis) delLog(key ...string) {
	if r.config.Debug == true {
		log.Printf("gocache [%s] del\n", key)
	}
}
