
import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/tag',
    method: 'get',
    params: data
  })
}

export function saveTag(data) {
  return request({
    url: '/tag',
    method: 'post',
    data
  })
}

export function editTag(data) {
  return request({
    url: '/tag',
    method: 'put',
    data
  })
}

export function deleteTag(data) {
  return request({
    url: '/tag',
    method: 'delete',
    data
  })
}
