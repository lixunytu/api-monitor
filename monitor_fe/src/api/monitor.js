import request from '@/utils/request'

export function  httpTest(httptest) {
  return request({
    url:'/v1/httptest',
    method: 'post',
    data:httptest
  })
}


export function  getData(getdata) {
  return request({
    url:'/v1/getdata',
    method: 'post',
    data:getdata
  })
}

export function  DataTest(dataTest) {
  return request({
    url:'/v1/datatest',
    method: 'post',
    data:dataTest
  })
}

export function parseCron(express) {
  return request({
    url: `/v1/prasecron/?e=${express}`,
    method: 'get'
  });
}

export function gealarmtList() {
  return request({
    url: `/v1/alarmSampInfo`,
    method: 'get'
  })
}

export function  createMonitor(monitor) {
  return request({
    url:'/v1/monitorinfo',
    method: 'post',
    data:monitor
  })
}

export function  updateMonitor(id,monitor) {
  return request({
    url:`/v1/monitorinfo/${id}`,
    method: 'put',
    data:monitor
  })
}

export function getMonitorList(params) {
  console.log(params)
  return request({
    url: `/v1/monitorinfo`,
    method: 'get',
    params
  });
}

export function getMonitorInfo(params) {
  console.log(params)
  return request({
    url: `/v1/monitorinfo/get`,
    method: 'get',
    params
  });
}

export function  changeMonitorInfoStatus(del,monitorId) {
  return request({
    url:`/v1/monitorinfo/${del}`,
    method: 'post',
    data:monitorId
  })
}

