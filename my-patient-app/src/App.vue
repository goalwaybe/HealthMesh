<script>
export default {
  onLaunch: function() {
    console.log('App Launch');
  },
  onShow: function() {
    console.log('App Show');
  },
  onHide: function() {
    console.log('App Hide');
  },
  onPageNotFound(res) {
    console.log('页面未找到:', res)
    uni.redirectTo({
      url: '/pages/error/not-found?title=页面不存在&message=您访问的页面走丢了'
    })
  }
}
</script>

<style>
/* 注意要写在 style 标签内，不是 script */
view, text, image, button {
  box-sizing: border-box;
}
page {
  background-color: #f5f5f5;
}
</style>