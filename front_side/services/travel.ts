import { defineStore } from 'pinia'
import { Travel } from '../models/travel'
import { Request } from './utils/request'
import axios from 'axios'

export const useTravelStore = defineStore({
    id: 'travel',
    state: () => ({
        count: 0,
        travels: [] as Travel[]
    }),

    actions: {
        async fakeFetchTravels() {
            //todo: fake fetch travels for develop the front end
            throw new Error('Not implemented')
        },

        async fetchTravels() {
            /*    const resp = await axios.post<response>('/api/travels')
   
               const result = resp.data
               if (result.code !== 200) {
                   throw new Error(result.data)
               }
               this.count = result.meta.count
               this.travels = result.data as Travel[] */

            const req = new Request('/api/travels', null)
            const result = req.get<Travel[]>()
            if (result === null) {
                throw new Error('Failed to fetch travels')
            }

            this.travels = result
        },

        async getTravelById(id: number) {
            //try to get travel from local storage
            if (this.travels.length > 0) {
                const travel = this.travels.find(travel => travel.id === id)
                if (travel) {
                    return travel
                }
            }

            // if not found, fetch from server
            const req = new Request(`/api/travels/?id=${id}`, null)
            return req.get<Travel>()
        },

        async createTravel(travel: Travel) {
            /*  const resp = await axios.post<response>(`/api/travels`, travel)
             const result = resp.data
 
             if (result.code !== 200) {
                 throw new Error(result.data)
             }
 
             return result.data as Travel */
            const req = new Request('/api/travels', travel)
            const result = req.post<Travel>()
            if (result === null) {
                throw new Error('Failed to create travel')
            }
            this.travels.push(result)
            return result
        },

        async updateTravel(travel: Travel) {
            /* const response = await axios.put<response>('/api/travels', travel)
            const resp = response.data

            if (resp.code !== 200) {
                throw new Error(resp.data)
            }

            return resp.data as Travel */
            const req = new Request('/api/travels', travel)
            const result = req.put<Travel>()
            if (result === null) {
                throw new Error('Failed to update travel')
            }


            result.then((data) => {
                if (data === null) {
                    throw new Error('Failed to update travel')
                }

                let index = this.travels.findIndex(travel => travel.id === data.id)
                this.travels[index] = data
            })
        },

        async deleteTravel(id: number) {
            /*   const response = await axios.delete<response>(`/api/travels/${id}`)
              const resp = response.data
  
              if (resp.code !== 200) {
                  throw new Error(resp.data)
              }
  
              return resp.data as Travel */

            const req = new Request(`/api/travels/${id}`, null)
            const result = req.delete<Travel>()
            if (result === null) {
                throw new Error('Failed to delete travel')
            }

            let index = this.travels.findIndex(travel => travel.id === id)
        },

        async getTravelList() {
            /*    const response = await axios.get<response>('/api/travels/list')
               const resp = response.data
               if (resp.code !== 200) {
                   throw new Error(resp.data)
               }
               return resp.data as Travel[] */

            const req = new Request('/api/travels/list', null)
            const resp = req.get<Travel[]>()
            if (resp === null) {
                throw new Error('Failed to fetch travel list')
            }

            resp.then((data) => {
                this.travels = data
            })
        }
    }
})