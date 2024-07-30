package main

import (
	"fmt"

	"github.com/myselfBZ/common/jwt"
)

func main()  {
    token, _ := jwt.GenerateToken(14)
    fmt.Println(token)
}
