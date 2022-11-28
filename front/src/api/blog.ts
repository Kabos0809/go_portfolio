import { Blog } from "../types/Types"

const BASE_URL = process.env.BASE_API_URL || "http://localhost:8080"

const get = (key: string) => getLocalStorage(key)
const getLocalStorage = (key:string) => {
    const ret = localStorage.getItem(key)
    const retToken = localStorage.getItem('token')
    if (ret && retToken) {
        return ret;
    }
    return null
}

export const CreateBlog = async (blog: Blog) => {
    const token = JSON.stringify(get("AccessToken"))
}