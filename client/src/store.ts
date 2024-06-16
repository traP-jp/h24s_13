import { mdiKnife } from '@mdi/js'
import { reactive, ref, computed, onMounted } from 'vue'

export interface User {
  id: string
  url: string
}

const me = await (await fetch(`/api/users/me`)).json()
const myID = me.id
const randomFriends = await (await fetch(`/api/users/${myID}/random?count=1`)).json()

export const targetId = ref(randomFriends[0])

export const changeableUsers = ref<User[]>([])

export const answerUsers = ref<User[]>([])

export const imageURLs = ref<User[]>([])
