<script setup lang="ts">
// ✅ 使用 import type —— 告诉编译器这只是类型，不要在运行时加载
import type { Department } from '@/api/medi'

const props = defineProps<{
  departments: Department[]
  activeId: number
}>()

const emit = defineEmits<{
  (e: 'change', id: number): void
}>()

const selectDept = (id: number) => {
  emit('change', id)
}
</script>