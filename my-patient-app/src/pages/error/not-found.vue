<template>
  <view class="error-page">
    <image src="/static/error-404.png" class="error-icon" />
    <text class="title">{{ title }}</text>
    <text class="desc">{{ message }}</text>
    <button class="back-btn" @click="goHome">返回首页</button>
  </view>
</template>

<script setup lang="ts">
import { onLaunch } from '@dcloudio/uni-app'

const props = defineProps<{
  title?: string
  message?: string
}>()

const goHome = () => {
  uni.reLaunch({
    url: '/pages/index/index'
  })
}
</script>

<style scoped>
.error-page {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  padding: 40rpx;
  background: #f5f5f5;
}
.error-icon {
  width: 200rpx;
  height: 200rpx;
  margin-bottom: 40rpx;
}
.title {
  font-size: 36rpx;
  font-weight: bold;
  color: #333;
  margin-bottom: 20rpx;
}
.desc {
  font-size: 28rpx;
  color: #666;
  text-align: center;
  margin-bottom: 60rpx;
}
.back-btn {
  width: 200rpx;
  height: 70rpx;
  line-height: 70rpx;
  background: #007AFF;
  color: white;
  border: none;
  border-radius: 10rpx;
}
</style>