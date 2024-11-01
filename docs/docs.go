// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/captcha": {
            "get": {
                "description": "验证码接口",
                "produces": [
                    "application/json"
                ],
                "summary": "验证码接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/dept/add": {
            "post": {
                "description": "新增部门接口",
                "produces": [
                    "application/json"
                ],
                "summary": "新增部门接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysDept"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/dept/delete": {
            "delete": {
                "description": "根据id删除部门接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id删除部门接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysDeptIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/dept/info": {
            "get": {
                "description": "根据id查询部门接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查询部门接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/dept/list": {
            "get": {
                "description": "查询部门列表接口",
                "produces": [
                    "application/json"
                ],
                "summary": "查询部门列表接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "部门名称",
                        "name": "deptName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "部门状态",
                        "name": "deptStatus",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/dept/update": {
            "put": {
                "description": "修改部门接口",
                "produces": [
                    "application/json"
                ],
                "summary": "修改部门接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysDept"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/dept/vo/list": {
            "get": {
                "description": "部门下拉列表接口",
                "produces": [
                    "application/json"
                ],
                "summary": "部门下拉列表接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "用户登录接口",
                "produces": [
                    "application/json"
                ],
                "summary": "用户登录接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/menu/add": {
            "post": {
                "description": "新增菜单接口",
                "produces": [
                    "application/json"
                ],
                "summary": "新增菜单接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/menu/delete": {
            "delete": {
                "description": "根据id删除菜单接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id删除菜单接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenuIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/menu/info": {
            "get": {
                "description": "根据id查询菜单",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查询菜单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/menu/list": {
            "get": {
                "description": "查询菜单列表",
                "produces": [
                    "application/json"
                ],
                "summary": "查询菜单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜单名称",
                        "name": "menuName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "菜单状态",
                        "name": "menuStatus",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/menu/update": {
            "put": {
                "description": "修改菜单接口",
                "produces": [
                    "application/json"
                ],
                "summary": "修改菜单接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/menu/vo/list": {
            "get": {
                "description": "查询新增选项列表接口",
                "produces": [
                    "application/json"
                ],
                "summary": "查询新增选项列表接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/add": {
            "post": {
                "description": "新增岗位接口",
                "produces": [
                    "application/json"
                ],
                "summary": "新增岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/batch/delete": {
            "delete": {
                "description": "批量删除岗位接口",
                "produces": [
                    "application/json"
                ],
                "summary": "批量删除岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.DelSysPostDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/delete": {
            "delete": {
                "description": "根据id删除岗位接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id删除岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysPostIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/info": {
            "get": {
                "description": "根据id查询岗位",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查询岗位",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/list": {
            "get": {
                "description": "分页查询岗位列表",
                "produces": [
                    "application/json"
                ],
                "summary": "分页查询岗位列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "PageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "PageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "岗位名称",
                        "name": "PostName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "状态：1-\u003e启用,2-\u003e禁用",
                        "name": "PostStatus",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "开始时间",
                        "name": "BeginTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "EndTime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/update": {
            "put": {
                "description": "修改岗位接口",
                "produces": [
                    "application/json"
                ],
                "summary": "修改岗位接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/updateStatus": {
            "put": {
                "description": "岗位状态修改接口",
                "produces": [
                    "application/json"
                ],
                "summary": "岗位状态修改接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysPostStatusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/post/vo/list": {
            "get": {
                "description": "岗位下拉列表",
                "produces": [
                    "application/json"
                ],
                "summary": "岗位下拉列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/add": {
            "post": {
                "description": "新增角色接口",
                "produces": [
                    "application/json"
                ],
                "summary": "新增角色接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddSysRoleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/assignPermissions": {
            "put": {
                "description": "分配权限接口",
                "produces": [
                    "application/json"
                ],
                "summary": "分配权限接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RoleMenu"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/delete": {
            "delete": {
                "description": "根据id删除角色接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id删除角色接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysRoleIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/info": {
            "get": {
                "description": "根据id查询角色接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据id查询角色接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/list": {
            "get": {
                "description": "分页查询角色列表接口",
                "produces": [
                    "application/json"
                ],
                "summary": "分页查询角色列表接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页数",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "roleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "帐号启用状态：1-\u003e启用,2-\u003e禁用",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "开始时间",
                        "name": "beginTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "endTime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/update": {
            "put": {
                "description": "修改角色",
                "produces": [
                    "application/json"
                ],
                "summary": "修改角色",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysRoleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/updateStatus": {
            "put": {
                "description": "角色状态启用/停用接口",
                "produces": [
                    "application/json"
                ],
                "summary": "角色状态启用/停用接口",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysRoleStatusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/vo/idList": {
            "get": {
                "description": "根据角色id查询菜单数据接口",
                "produces": [
                    "application/json"
                ],
                "summary": "根据角色id查询菜单数据接口",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id",
                        "name": "Id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/role/vo/list": {
            "get": {
                "description": "角色下拉列表",
                "produces": [
                    "application/json"
                ],
                "summary": "角色下拉列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/upload": {
            "post": {
                "description": "单图片上传接口",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "单图片上传接口",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddSysRoleDto": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "roleKey": {
                    "description": "角色key",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "model.DelSysPostDto": {
            "type": "object",
            "properties": {
                "ids": {
                    "description": "Id列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.LoginDto": {
            "type": "object",
            "required": [
                "idKey",
                "image",
                "password",
                "username"
            ],
            "properties": {
                "idKey": {
                    "description": "UUID",
                    "type": "string"
                },
                "image": {
                    "description": "验证码",
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 4
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "账号",
                    "type": "string"
                }
            }
        },
        "model.RoleMenu": {
            "type": "object",
            "required": [
                "id",
                "menuIds"
            ],
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "menuIds": {
                    "description": "菜单id列表",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "model.SysDept": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysDept"
                    }
                },
                "createTime": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "deptName": {
                    "description": "部门名称",
                    "type": "string"
                },
                "deptStatus": {
                    "description": "部门状态（1-\u003e正常 2-\u003e停用）",
                    "type": "integer"
                },
                "deptType": {
                    "description": "部门类型（1-\u003e公司, 2-\u003e中心，3-\u003e部门）",
                    "type": "integer"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "parentId": {
                    "description": "父id",
                    "type": "integer"
                }
            }
        },
        "model.SysDeptIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysMenu": {
            "type": "object",
            "properties": {
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysMenu"
                    }
                },
                "createTime": {
                    "description": "排序",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "icon": {
                    "description": "菜单图标",
                    "type": "string"
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "menuName": {
                    "description": "菜单名称",
                    "type": "string"
                },
                "menuStatus": {
                    "description": "启用状态；1-\u003e禁用；2-\u003e启用",
                    "type": "integer"
                },
                "menuType": {
                    "description": "菜单类型：1-\u003e目录；2-\u003e菜单；3-\u003e按钮",
                    "type": "integer"
                },
                "parentId": {
                    "description": "父菜单id",
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "url": {
                    "description": "菜单url",
                    "type": "string"
                },
                "value": {
                    "description": "权限值",
                    "type": "string"
                }
            }
        },
        "model.SysMenuIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysPost": {
            "type": "object",
            "properties": {
                "createTime": {
                    "description": "创建时间",
                    "allOf": [
                        {
                            "$ref": "#/definitions/util.HTime"
                        }
                    ]
                },
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "postCode": {
                    "description": "岗位编码",
                    "type": "string"
                },
                "postName": {
                    "description": "岗位名称",
                    "type": "string"
                },
                "postStatus": {
                    "description": "状态（1-\u003e正常 2-\u003e停用）",
                    "type": "integer"
                },
                "remark": {
                    "description": "备注",
                    "type": "string"
                }
            }
        },
        "model.SysPostIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.SysRoleIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysPostStatusDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "postStatus": {
                    "description": "状态（1-\u003e正常 2-\u003e停用）",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysRoleDto": {
            "type": "object",
            "properties": {
                "description": {
                    "description": "描述",
                    "type": "string"
                },
                "id": {
                    "description": "Id",
                    "type": "integer"
                },
                "roleKey": {
                    "description": "角色key",
                    "type": "string"
                },
                "roleName": {
                    "description": "角色名称",
                    "type": "string"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysRoleStatusDto": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "ID",
                    "type": "integer"
                },
                "status": {
                    "description": "状态：1-\u003e启用,2-\u003e禁用",
                    "type": "integer"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "返回的数据"
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "util.HTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "通用后台管理系统",
	Description:      "后台管理系统API接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
