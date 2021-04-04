
import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/category',
    method: 'get',
    params: data
  })
}

export function saveCategory(data) {
  return request({
    url: '/category',
    method: 'post',
    data
  })
}

export function editCategory(data) {
  return request({
    url: '/category',
    method: 'put',
    data
  })
}

export function enableCategory(data) {
  return request({
    url: '/category/enable',
    method: 'put',
    data
  })
}

export function deleteCategory(data) {
  return request({
    url: '/category',
    method: 'delete',
    data
  })
}
