import axios from "axios";


export type response = {
    code: number;
    meta: {
        count: number;
    }
    data: any;
}

export class Request {
    api: string;
    data: any;

    constructor(api_: string, data_: any) {
        this.api = api_
        this.data = data_
    }

    async get<R>(): Promise<R | null> {
        const resp = await axios.get<response>(this.api)
        const result = resp.data

        if (result === null) {
            return null
        }
        if (result.code !== 200) {
            console.log(result.code, result.data)
            return null
        }
        return result.data
    }

    async post<R>(): Promise<R | null> {
        const resp = await axios.post<response>(this.api, this.data)
        const result = resp.data

        if (result === null) {
            return null
        }
        if (result.code !== 200) {
            console.log(result.code, result.data)
            return null
        }
        return result.data
    }

    async put<R>(): Promise<R | null> {
        const resp = await axios.put<response>(this.api, this.data)
        const result = resp.data

        if (result === null) {
            return null
        }
        if (result.code !== 200) {
            console.log(result.code, result.data)
            return null
        }
        return result.data
    }

    async delete<R>(): Promise<R | null> {
        const resp = await axios.delete<response>(this.api)
        const result = resp.data

        if (result === null) {
            return null
        }
        if (result.code !== 200) {
            console.log(result.code, result.data)
            return null
        }
        return result.data
    }

}