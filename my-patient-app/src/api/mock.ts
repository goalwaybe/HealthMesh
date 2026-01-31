// src/api/mock.ts - Mock 数据

import { Department, Doctor } from './medi'

// Mock 科室数据
export const mockDepartments: Department[] = [
  { id: 1, name: '内科' },
  { id: 2, name: '外科' },
  { id: 3, name: '儿科' },
  { id: 4, name: '妇科' },
  { id: 5, name: '眼科' },
]

// Mock 医生数据 - 按科室分类
export const mockDoctors: Record<number, Doctor[]> = {
  1: [
    {
      id: 101,
      name: '张医生',
      title: '主任医师',
      departmentName: '内科',
      avatar: '/static/default-avatar.png',
    },
    {
      id: 102,
      name: '李医生',
      title: '副主任医师',
      departmentName: '内科',
      avatar: '/static/default-avatar.png',
    },
    {
      id: 103,
      name: '王医生',
      title: '医师',
      departmentName: '内科',
      avatar: '/static/default-avatar.png',
    },
  ],
  2: [
    {
      id: 201,
      name: '陈医生',
      title: '主任医师',
      departmentName: '外科',
      avatar: '/static/default-avatar.png',
    },
    {
      id: 202,
      name: '周医生',
      title: '副主任医师',
      departmentName: '外科',
      avatar: '/static/default-avatar.png',
    },
  ],
  3: [
    {
      id: 301,
      name: '赵医生',
      title: '主任医师',
      departmentName: '儿科',
      avatar: '/static/default-avatar.png',
    },
    {
      id: 302,
      name: '吴医生',
      title: '医师',
      departmentName: '儿科',
      avatar: '/static/default-avatar.png',
    },
  ],
  4: [
    {
      id: 401,
      name: '郑医生',
      title: '主任医师',
      departmentName: '妇科',
      avatar: '/static/default-avatar.png',
    },
  ],
  5: [
    {
      id: 501,
      name: '刘医生',
      title: '主任医师',
      departmentName: '眼科',
      avatar: '/static/default-avatar.png',
    },
    {
      id: 502,
      name: '孙医生',
      title: '副主任医师',
      departmentName: '眼科',
      avatar: '/static/default-avatar.png',
    },
  ],
}
