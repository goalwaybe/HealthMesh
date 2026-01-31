<!-- src/pages/index/index.vue -->
<template>
  <view class="container">
    <view class="header">
      <text class="title">患者挂号系统</text>
    </view>
    <view class="content">
      <text class="section-title">选择科室</text>
      <DepartmentNav 
        :departments="departments" 
        :active-id="currentDeptId"
        @change="handleDeptChange"
      />
      <text class="section-title">医生列表</text>
      <scroll-view scroll-y class="doctor-list">
        <view v-if="doctors.length === 0" class="empty-tip">
          <text>暂无医生信息</text>
        </view>
        <view class="doctor-card" v-for="doc in doctors" :key="doc.id">
          <image :src="doc.avatar || '/static/default-avatar.png'" class="avatar" />
          <view class="info">
            <text class="name">{{ doc.name }}</text>
            <text class="title">{{ doc.title }}</text>
            <button class="book-btn" size="mini" @click="goToBook(doc.id)">预约</button>
          </view>
        </view>
      </scroll-view>
    </view>
  </view>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import DepartmentNav from '@/components/DepartmentNav.vue'

// 👇 拆分导入：值 vs 类型
import { getDepartments, getDoctorsByDept } from '@/api/medi'      // 函数
import type { Department, Doctor } from '@/api/medi'               // 类型


const departments = ref<Department[]>([])
const doctors = ref<Doctor[]>([])
const currentDeptId = ref(0)

const loadDepartments = async () => {
  try {
    const res = await getDepartments()
    if (res.code === 200) {
      departments.value = res.data
      if (departments.value.length > 0) {
        currentDeptId.value = departments.value[0].id
        loadDoctors(currentDeptId.value)
      }
    }else if (res.code === 404) {
      // 跳转到 404 页面
      uni.navigateTo({
        url: '/pages/error/not-found?title=科室未找到&message=暂无可用科室信息'
      })
    } else {
     // 其他错误（如 500）
      uni.navigateTo({
        url: '/pages/error/not-found?title=服务异常&message=请稍后再试'
      })
    }



  } catch (error) {
    console.error(error)
  }
}

const loadDoctors = async (deptId: number) => {
  try {
    const res = await getDoctorsByDept(deptId)
    if (res.code === 200) {
      doctors.value = res.data
    }
  } catch (error) {
    console.error(error)
  }
}

const handleDeptChange = (id: number) => {
  currentDeptId.value = id
  loadDoctors(id)
}

const goToBook = (doctorId: number) => {
  uni.navigateTo({
    url: `/pages/booking/booking?doctorId= $ {doctorId}`
  })
}

onMounted(() => {
  loadDepartments()
})
</script>

<style scoped>
/* 同之前的样式 */
.container { padding: 20rpx; }
.content { display: flex; height: calc(100vh - 140rpx); }
/* ... 其他样式同上 */
</style>