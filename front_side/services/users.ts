import { defineStore } from "pinia";
import { User } from "../models/user";
import { Request } from "./utils/request";


export const useUserStore = defineStore({
    id: 'user',
    state: () => ({
        user: {
            id: 0,
            username: '',
            // avatar?: '';
            nickname: '',
            role: 0,
            email: '',
            // mbti?: '';
            // birthday?: 0,
            // gender?: '';
            created_at: 0,
            last_login: 0
        } as User

    }),
    getters: {},
    actions: {
        async fakeFetchUser() {
            //todo: add test case 
            throw new Error('Not implemented')
        },

        async fetchUser() {

            //not safe but for now, we can use this
            //just directly fetch user from local storage if it exists
            {
                const json = localStorage.getItem('user')
                if (json) {
                    this.user = JSON.parse(json) as User
                    return
                }
            }

            /*   const response = await axios.get<response>('/api/user')
              const resp = response.data
              if (resp.code !== 200) {
                  throw new Error(resp.data)
              }
              this.user = resp.data as User */
            const req = new Request('/api/user', null)
            const result = req.get<User>()
            if (result === null) {
                throw new Error('Failed to fetch user')
            }

            result.then((data) => {
                if (data === null) {
                    throw new Error('Failed to fetch user')
                }
                this.user = data
            })
        },

        async updateUser(user: User) {
            /*     const response = await axios.put<response>('/api/user', user)
                const resp = response.data
                if (resp.code !== 200) {
                    throw new Error(resp.data)
                }
                this.user = resp.data as User */

            const req = new Request('/api/user', user)
            const result = req.put<User>()
            if (result === null) {
                throw new Error('Failed to update user')
            }
            result.then((data) => {
                if (data === null) {
                    throw new Error('Failed to update user')
                }
                this.user = data
            })
        },

        async login(username: string, password: string) {
            /*  const response = await axios.post<response>('/api/login', { username, password })
             const resp = response.data
             if (resp.code !== 200) {
                 console.log('something goes wrong with login ', resp.data)
                 return false
             }
 
             localStorage.setItem('user', JSON.stringify(resp.data))
             this.user = resp.data as User
             return true */
            const req = new Request('/api/login', { username, password })
            const result = req.post<User>()
            if (result === null) {
                return false
            }
            result.then((data) => {
                if (data === null) {
                    return false
                }
                localStorage.setItem('user', JSON.stringify(data))
                this.user = data
                return true
            }).catch((err) => {
                console.log(err)
                return false
            })

        },

        async logout() {

            /*      const response = await axios.post<response>('/api/logout')
                 const resp = response.data
                 if (resp.code !== 200) {
                     return false
                 }
     
                 // remove user from local storage
                 {
                     localStorage.removeItem('user')
                     this.user = {
                         id: 0,
                         username: '',
                         nickname: '',
                         role: 0,
                         email: '',
                         created_at: 0,
                         last_login: 0
                     }
                 }
                 return true */

            const req = new Request('/api/logout', null)
            const result = req.post<null>()
            if (result === null) {
                return false
            }

            result.then(() => {
                localStorage.removeItem('user')
                this.user = {
                    id: 0,
                    username: '',
                    nickname: '',
                    role: 0,
                    email: '',
                    created_at: 0,
                    last_login: 0
                }
                return true
            }).catch((err) => {
                console.log(err)
                return false
            })
        }

    },
})


const store = useUserStore()
store.fetchUser()
store.login('lyt','123')
console.log(store.user)