// src/api/medi.ts

import { mockDepartments, mockDoctors } from './mock'

const BASE_URL = 'https://your-api-gateway.com/api/v1'

// 是否使用 Mock 数据
const USE_MOCK = true

// 定义数据类型
export interface Department {
  id: number
  name: string
}

export interface Doctor {
  id: number
  name: string
  title: string
  departmentName: string
  avatar?: string
}

// 请求返回类型
interface ApiResponse<T> {
  code: number
  data: T
  msg?: string
}

// 封装 request（带泛型）
function request<T>(options: UniApp.RequestOptions): Promise<ApiResponse<T>> {
  return new Promise((resolve, reject) => {
    uni.request({
      ...options,
      url: BASE_URL + options.url,
      success: (res: any) => {
        if (res.statusCode === 200) {
          resolve(res.data as ApiResponse<T>)
        } else {
          reject(new Error('Request failed'))
        }
      },
      fail: (err) => {
        uni.showToast({ title: '网络错误', icon: 'none' })
        reject(err)
      }
    })
  })
}

// 获取科室
export function getDepartments() {
  if (USE_MOCK) {
    return Promise.resolve({
      code: 200,
      data: mockDepartments,
      msg: 'success'
    })
  }
  return request<Department[]>({
    url: '/departments',
    method: 'GET'
  })
}

// 根据科室ID获取医生
export function getDoctorsByDept(deptId: number) {
  return request<Doctor[]>({
    url: '/doctors',
    method: 'GET',
    data: { deptId }
  })
}