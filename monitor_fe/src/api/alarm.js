import request from '@/utils/request'

export function  getList(params) {
  return request({
    url:'/v1/alarmAInfo',
    method: 'get',
    params
  })
}

export function  getListForId(alarmId) {
  return request({
    url:`/v1/alarmAInfo/${alarmId}`,
    method: 'get',
  })
}

export function  del(id) {
  return request({
    url:`/v1/alarmAInfo/${id}`,
    method: 'delete',
  })
}

export function  create(data) {
  return request({
    url:'/v1/alarmAInfo',
    method: 'post',
    data:data
  })
}

export function  update(id,data) {
  return request({
    url:`/v1/alarmAInfo/${id}`,
    method: 'put',
    data: data
  })
}
