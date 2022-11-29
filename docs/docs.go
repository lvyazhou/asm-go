// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/captcha": {
            "get": {
                "description": "验证码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "summary": "验证码",
                "responses": {}
            }
        },
        "/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录请求信息",
                        "name": "requestLogin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/login_handle.requestLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.LoginVo"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登出"
                ],
                "summary": "登出",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/oplog/list/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "查询审计日志",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "日志管理"
                ],
                "summary": "查询审计日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "请求uri",
                        "name": "requestUri",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/refresh": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户续期",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户续期"
                ],
                "summary": "用户续期",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.LoginVo"
                        }
                    }
                }
            }
        },
        "/user/del/{id}": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "根据用户ID删除用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/user/get/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "个人用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "个人用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.UserVO"
                        }
                    }
                }
            }
        },
        "/user/get/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "个人用户信息",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "个人用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        },
        "/user/list/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户分页查询",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户分页查询",
                "parameters": [
                    {
                        "description": "用户查询实体",
                        "name": "userQuery",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserQueryDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.UserVO"
                        }
                    }
                }
            }
        },
        "/user/save/": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "request user",
                        "name": "userDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserDTO": {
            "type": "object",
            "required": [
                "account",
                "email",
                "mobile",
                "name"
            ],
            "properties": {
                "account": {
                    "description": "Account 账户名称",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 4
                },
                "email": {
                    "description": "Email 电子邮件",
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1
                },
                "mobile": {
                    "description": "Mobile 手机号码",
                    "type": "string"
                },
                "name": {
                    "description": "Name 姓名",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 1
                },
                "pass_word": {
                    "description": "密码",
                    "type": "string"
                },
                "role_list": {
                    "description": "角色列表",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "description": "Status 用户状态 @see constants.user_status 1:正常；2：禁用；3：已删除'",
                    "type": "integer",
                    "maximum": 3,
                    "minimum": 0
                }
            }
        },
        "dto.UserQueryDto": {
            "type": "object",
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "page": {
                    "description": "分页页面 从0开始",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "分页大小 默认10",
                    "type": "integer"
                },
                "uname": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "handler.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "login_handle.requestLogin": {
            "type": "object",
            "required": [
                "account",
                "auth_code",
                "password"
            ],
            "properties": {
                "account": {
                    "description": "账户名",
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 6
                },
                "auth_code": {
                    "description": "验证码",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 32,
                    "minLength": 8
                }
            }
        },
        "vo.LoginVo": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "用户名称",
                    "type": "string"
                },
                "token": {
                    "description": "用户token",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户ID",
                    "type": "integer"
                }
            }
        },
        "vo.UserRoleVo": {
            "type": "object",
            "properties": {
                "role_id": {
                    "type": "integer"
                },
                "role_name": {
                    "type": "string"
                }
            }
        },
        "vo.UserVO": {
            "type": "object",
            "properties": {
                "account": {
                    "description": "Account 账户名称",
                    "type": "string"
                },
                "create_time": {
                    "description": "CreateTime 创建时间",
                    "type": "string"
                },
                "create_user": {
                    "description": "CreateUser 创建人ID",
                    "type": "string"
                },
                "email": {
                    "description": "Email 电子邮件",
                    "type": "string"
                },
                "id": {
                    "description": "ID 雪花算法生成",
                    "type": "integer"
                },
                "mobile": {
                    "description": "Mobile 手机号码",
                    "type": "string"
                },
                "name": {
                    "description": "Name 姓名",
                    "type": "string"
                },
                "role_list": {
                    "description": "角色列表",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vo.UserRoleVo"
                    }
                },
                "status_name": {
                    "description": "Status 用户状态 @see constants.user_status 1:正常；2：禁用；3：已删除'",
                    "type": "string"
                },
                "update_time": {
                    "description": "UpdateTime 更新时间",
                    "type": "string"
                },
                "update_user": {
                    "description": "UpdateUser 更新人ID",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "asm_platform-TOKEN",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}