import { defineStore } from 'pinia'
import { Relation, RelationType } from '../models/relation'

import { Request } from './utils/request'
import axios from 'axios'

export const useRelationStore = defineStore({
    id: 'relation',
    state: () => ({
        count: 0,
        relations: [] as Relation[],
        friend_link: '',
        couple_link: ''
    }),


    actions: {
        async fakeFetchRelations() {
            throw new Error('Not implemented')
        },
        async fetchRelations() {
            const req = new Request('/api/relations', null)
            const result = req.get<Relation>()
            if (result === null) {
                throw new Error('Failed to fetch relations')
            }
            this.relations = result
            this.count = this.relations.length
        },

        async getRelationById(id: number) {
            //try to get relation from local storage
            if (this.relations.length > 0) {
                const relation = this.relations.find(relation => relation.id === id)
                if (relation) {
                    return relation
                }
            }

            // if not found, fetch from server
            // becareful, this might be null, just because if target user doesn't has relation with you , server will response a error message
            const req = new Request(`/api/relations/?id=${id}`, null)
            return req.get<Relation>()
        },
        async createRelationLink(relation: Relation) {
            const resp = new Request('/api/relations', relation)
            const result = resp.post<{ url, relation_type }>()
            //todo: implement this
            if (result == null) {
                throw new Error('Failed to create relation link')
            }

            result.then((data) => {
                if (data === null) {
                    throw new Error('Failed to create relation link')
                }
                if (data.relation_type === RelationType.FRIEND) {
                    this.friend_link = data.url
                }
                else if (data.relation_type === RelationType.COUPLE) {
                    this.couple_link = data.url
                }
            }).catch((err) => {
                console.log(err)
            })

        },
        async updateRelation(relation: Relation) {
            /*     const resp = await axios.put<response>('/api/relations', relation)
                const result = resp.data
                if (result.code !== 200) {
                    throw new Error(result.data)
                }
                return result.data as Relation */

            throw new Error('Not implemented')
            const resp = new Request('/api/relations', relation)


        }
    }

})