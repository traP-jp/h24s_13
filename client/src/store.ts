import { reactive, ref, computed } from 'vue'

export interface User {
  id: string
  url: string
}

export const targetId = ref('masky5859')

export const changeableUsers = ref<User[]>([
  { id: 'masky5859', url: 'https://q.trap.jp/api/v3/public/icon/masky5859' },
  { id: 'takeno_hito', url: 'https://q.trap.jp/api/v3/public/icon/takeno_hito' },
  { id: 'cp20', url: 'https://q.trap.jp/api/v3/public/icon/cp20' },
  { id: 'toki', url: 'https://q.trap.jp/api/v3/public/icon/toki' },
  { id: 'Series_205', url: 'https://q.trap.jp/api/v3/public/icon/Series_205' }
])

export const answerUsers = ref<User[]>([])
export const API_URL = import.meta.env.VITE_APP_API_URL

export const imageURLs = ref<User[]>([])
