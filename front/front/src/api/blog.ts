import axios from "axios"
import { Blog } from "../types/Types"

const BASE_URL: string = process.env.BASE_API_URL || "http://localhost:8080"

export const GetAllBlog = async () => {
    const res = await axios.get(`${BASE_URL}/v1/blog/list`)
    return res.data
}

export const GetBlogByID = async (id: number) => {
    const res = await axios.get(`${BASE_URL}/v1/blog/detail/${id}`)
    return res.data
}

export const CreateBlog = async (blog: Blog) => {
    const token: string | null = window.localStorage.getItem("AccessToken")
    const headers = {
        'Content-Type':'application/json',
        'Authorization':`Bearer ${token}`
    }
    const res = await axios.post(`${BASE_URL}/admin/blog/create`, blog,{
        headers: headers
    })
    return res.data
}

export const DeleteBlog = async (id: number) => {
    const token: string | null = window.localStorage.getItem("AccessToken")
    const headers = {
        'Content-Type':'application/json',
        'Authorization':`Bearer ${token}`
    }
    await axios.delete(`${BASE_URL}/admin/blog/delete/${id}`, {
        headers: headers
    })
    return id
}

export const UpdateBlog = async (id: number, blog: Blog) => {
    const token: string | null = window.localStorage.getItem("AccessToken")
    const headers = {
        'Content-Type':'application/json',
        'Authorization':`Bearer ${token}`
    }
    const res = await axios.put(`${BASE_URL}/admin/blog/update/${id}`, blog, {
        headers: headers
    })
    return res.data
}