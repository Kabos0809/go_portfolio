import axios from "axios"
import { Contact } from "../types"

const BASE_URL: string = process.env.BASE_API_URL || "http://localhost:8080"
const token: string | null = window.localStorage.getItem("AccessToken")
const headers = {
        'Content-Type':'application/json',
        'Authorization':`Bearer ${token}`
    }


export const GetContact = async () => {
    const res = await axios.get(BASE_URL+"/admin/contact/list", {
        headers:headers
    })
    return res.data
}

export const GetContactByID = async (id: number) => {
    const res = await axios.get(BASE_URL+`admin/contact/detail/${id}`, {
        headers:headers
    })
    return res.data
}

export const CreateContact = async (c: Contact) => {
    const res = await axios.post(BASE_URL+"/v1/contact", c)
    return res.data
}

export const DeleteContact = async (id: number) => {
    await axios.delete(BASE_URL+`/admin/contact/delete/${id}`, {
        headers:headers
    })
}