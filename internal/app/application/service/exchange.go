package service

import "payday-savings/m/internal/app/domain/valueobject"

// IExchange is interface of bitcoin exchange
type IExchange interface {
	Ticker(p valueobject.Pair) valueobject.Ticker
	Ohlc(p valueobject.Pair, t valueobject.Timeunit) []valueobject.CandleStick
}
