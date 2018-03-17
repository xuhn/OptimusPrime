package redis_test

import (
	"net"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"../common/redis"
)

var _ = Describe("newConnDialer with bad connection", func() {
	It("should return an error", func() {
		dialer := redis.NewConnDialer(&redis.Options{
			Dialer: func() (net.Conn, error) {
				return &badConn{}, nil
			},
			MaxRetries: 3,
			Password:   "password",
			DB:         1,
		})
		_, err := dialer()
		Expect(err).To(MatchError("bad connection"))
	})
})
