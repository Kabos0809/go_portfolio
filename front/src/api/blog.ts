import axios from 'axios'
import { Blog, Token } from "../types/Types"

const BASE_URL = process.env.BASE_API_URL || "http://localhost:8080"
let token: Token

export const GetBlogList = async () => {
    const res = await axios.get(BASE_URL+"/v1/blog/list")
    return res.data
}

export const GetBlogById = async (id: number) => {
    const res = await axios.get(BASE_URL+`/v1/blog/${id}`)
    return res.data
}

export const CreateBlog = async (blog: Blog) => {
    const header = CreateHeader()
    const res = await axios.get(BASE_URL+`/admin/blog/create`, blog, {
        header: header
    })
    return res.data
}

export const UpdateBlog = async (id: number, blog: Blog) => {
    const header = CreateHeader()
    const res = await axios.put(BASE_URL+`/admin/blog/update/${id}`, blog, {
        header: header
    })
    return res.data
}

export const DeleteBlog = async (id: number) => {
    const header = CreateHeader()
    const res = await axios.delete(BASE_URL+`/admin/blog/delete/${id}`, {
        header: header
    })
    return res.data
}

export const ChangeBlogActive = async (id: number) => {
    const header = CreateHeader()
    const res = await axios.put(BASE_URL+`/admin/blog/changeactive/${id}`, {
        header: header
    })
    return res.data
}

export const CreateHeader = () => {
    const json_token = localStorage.getItem("Token")
    if (typeof json_token === "string") {
        token = JSON.parse(json_token)
    }
    const header = {
        "Content-Type": "application/json",
        "Authorization": `Bearer ${token.AccessToken}`
    }
    return header
}