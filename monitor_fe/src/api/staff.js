import request from '@/utils/request'

export function  getDepartmentList() {
  return request({
    url:'/v1/departmentinfo',
    method: 'get',
  })
}

export function  getList(params) {
  return request({
    url:'/v1/employee',
    method: 'get',
    params
  })
}

export function  create(data) {
  return request({
    url:'/v1/employee',
    method: 'post',
    data:data
  })
}

export function  update(id,data) {
  return request({
    url:`/v1/employee/${id}`,
    method: 'put',
    data:data
  })
}

export function  del(data) {
  return request({
    url:`/v1/employee/${data}`,
    method: 'delete',
  })
}
