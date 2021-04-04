import request from '@/utils/request'

export function getList(data) {
  return request({
    url: '/article',
    method: 'get',
    data
  })
}

export function getInfo(id) {
  return request({
    url: `/blog/${id}`,
    method: 'get'
  })
}

export function saveCategory(data) {
  return request({
    url: '/article',
    method: 'post',
    data
  })
}

export function editCategory(data) {
  return request({
    url: '/article',
    method: 'put',
    data
  })
}

export function enableCategory(data) {
  return request({
    url: '/article/enable',
    method: 'put',
    data
  })
}

export function deleteCategory(data) {
  return request({
    url: '/article',
    method: 'delete',
    data
  })
}
