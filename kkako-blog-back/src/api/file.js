import request from '@/utils/request'

export function fileupload(data) {
  return request({
    url: '/file/upload',
    method: 'post',
    params: data,
    headers: { 'Content-Type': 'multipart/form-data' }
  })
}
