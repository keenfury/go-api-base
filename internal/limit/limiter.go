package limit

import (
	tb "github.com/didip/tollbooth"
	tl "github.com/didip/tollbooth/limiter"
	"github.com/labstack/echo"

	e "github.com/keenfury/go-api-base/internal/api_error"
)

type Limit struct {
	myLimiter    *tl.Limiter
	keyLimitFunc func(c echo.Context) []string
}

func (l *Limit) NewLimiter(reqLimitPerDuration float64, durationInSeconds int64, funcToCall func(c echo.Context) []string) {
	// durInSecs, err := time.ParseDuration(fmt.Sprintf("%ds", durationInSeconds))
	// if err != nil {
	// 	durInSecs = time.Second
	// }
	l.myLimiter = tb.NewLimiter(reqLimitPerDuration, nil)
	l.keyLimitFunc = funcToCall
}

func (l *Limit) SetLimitHandler(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		httpError := tb.LimitByKeys(l.myLimiter, l.keyLimitFunc(c))
		if httpError != nil {
			c.Response().Header().Add("Content-Type", l.myLimiter.GetMessageContentType())
			err := e.LimiterError(httpError)
			return c.JSON(err.StatusCode, err)
		}
		return h(c)
	}
}
