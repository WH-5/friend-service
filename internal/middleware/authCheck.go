// Package middleware authCheck.go
// Author: 王辉
// Created: 2025-03-29 01:36
// 不信任jwt token  需要且只用token进行数据库二次校验 使用接口body里传的身份信息进行操作
package middleware

import (
	"context"
	"fmt"
	"github.com/WH-5/friend-service/internal/pkg"
	"github.com/WH-5/friend-service/internal/service"

	"github.com/golang-jwt/jwt/v5"
	"log"
	"strings"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// AuthCheckExist 检查token是否可用 可用会在上下文携带token和用户id
func AuthCheckExist(friendService *service.FriendService) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			log.Println("auth middleware in", req)

			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				log.Println("auth middleware error: transport context not found")
				return nil, fmt.Errorf("missing transport context")
			}

			authHeader := tr.RequestHeader().Get("Authorization")
			log.Printf("Authorization header value: [%q]", authHeader)
			log.Printf("Full request header:\n%+v", tr.RequestHeader())
			if authHeader == "" {
				log.Println("auth middleware error: Authorization header is empty")
				return nil, fmt.Errorf("missing authorization header")
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			log.Printf("Extracted token string: %s", tokenString)

			token, err := pkg.ParseToken(tokenString, friendService.UC.CF.JWT_SECRET_KEY)
			if err != nil {
				log.Printf("auth middleware error: failed to parse token: %v", err)
				return nil, fmt.Errorf("token 解析失败: %v", err)
			}

			log.Printf("Parsed token: %+v", token)
			if !token.Valid {
				log.Println("auth middleware error: token is not valid")
				return nil, fmt.Errorf("token 无效")
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Printf("auth middleware error: failed to cast claims to jwt.MapClaims, got type: %T", token.Claims)
				return nil, fmt.Errorf("无法解析 Claims")
			}

			log.Printf("Token claims: %+v", claims)

			uid := claims["user_id"]
			session := claims["session"]
			log.Printf("Extracted user_id: %v, session: %v", uid, session)

			ctx = context.WithValue(ctx, "user_id", uid)
			ctx = context.WithValue(ctx, "session", session)
			ctx = context.WithValue(ctx, "token", token)

			log.Println("auth middleware completed, forwarding to handler")
			return handler(ctx, req)
		}
	}
}
