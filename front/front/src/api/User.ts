import axios from "axios"
import { User } from "../types/Types"

const BASE_URL: string = process.env.BASE_API_URL || "http://localhost:8080"

export const UserSignUp = async (user: User) => {
    await axios.post(`${BASE_URL}/admin/signup`, user)
    return
}

export const UserLogin = async (user: User) => {
    await axios.post(`${BASE_URL}/admin/login`, user)
    .then((response) => {
        const jwtToken = response.data.response
    })

}

export const UserLogout = async () => {
    
}