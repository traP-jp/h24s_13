import { reactive, ref, computed } from 'vue'

export interface User {
  id: string
}

export const test = ref('test')

export const userInfo = ref<User[]>([
  { id: 'masky5859' },
  { id: 'pirosiki' },
  { id: 'cp20' },
  { id: 'toki' },
  { id: 'Series_205' }
])
