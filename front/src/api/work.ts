import { Work } from "../types/Types"

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

export const CreateWork = async (work: Work) => {
    const token = get("AccessToken")
    await fetch(BASE_URL + "/admin/work/create",{
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer" + token
        },
        body: JSON.stringify({
            "title": work.Title,
            "text": work.Text
        })
    }).then(response => {
        if (response.status === 200) {
            window.location.href = "/admin/work"
        } else if (response.status === 500) {
            console.log("Internal Server Error")
            window.location.href = "/500"
        } else if (response.status === 400) {
            console.log("Bad Request")
            window.location.href = "/400"
        } else if (response.status === 401) {
            console.log("UnAuthorized")
            window.location.href = "/admin/user/refresh"
        } else {
            console.log("Error")
            window.location.href = "/error"
        }
    }).catch(error => {
        console.error(error)        
        window.location.href = "/admin/blog/create"
    })
}

export const GetAllWork = async () => {
    await fetch(BASE_URL + "/v1/work/list", {
        method: "GET",
    }).then(response => {
        if (response.status === 200) {
            return response.json()
        } else if (response.status === 500) {
            console.log("Internal Server Error")
            window.location.href = "/500"
        } else if (response.status === 400) {
            console.log("Bad Request")
            window.location.href = "/400"
        } else {
            console.log("Error")
            window.location.href = "/error"
        }
    }).then(data => {
        return data
    }).catch(error => {
        console.error(error)
        window.location.href = "/error"
    })
}

export const GetWorkById = async (id: number) => {
    await fetch(BASE_URL + `/v1/work/detail/${id}`, {
        method: "GET"
    }).then(response => {
        if (response.status === 200) {
            return response.json()
        } else if (response.status === 500) {
            console.log("Internal Server Error")
            window.location.href = "/500"
        } else if (response.status === 400) {
            console.log("Bad Request")
            window.location.href = "/400"
        } else if (response.status === 404) {
            console.log("Not Found")
            window.location.href = "/404"
        } else {
            console.log("Error")
            window.location.href = "/error"
        }
    }).then(data => {
        return data
    }).catch(error => {
        console.error(error)
        window.location.href = "/error"
    })
}

export const UpdateWork = async (id: number, work: Work) => {
    const token = get("AccessToken")
    await fetch(BASE_URL + `/admin/work/update/${id}`, {
        method: "PUT",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer" + token
        },
        body: JSON.stringify({
            "title": work.Title,
            "text": work.Text
        })
    }).then(response => {
        if (response.status === 200) {
            window.location.href = "/admin/work"
        } else if (response.status === 500) {
            console.log("Internal Server Error")
            window.location.href = "/500"
        } else if (response.status === 400) {
            console.log("Bad Request")
            window.location.href = "/400"
        } else if (response.status === 401) {
            console.log("UnAuthorized")
            
        } else if (response.status === 404) {
            console.log("Not Found")
            window.location.href = "/404"
        } else {
            console.log("Error")
            window.location.href = "/error"
        }
    }).catch(error => {
        console.error(error)
        window.location.href = "/error"
    })
}

export const DeleteWork = async (id: number) => {
    const token = get("AccessToken")
    await fetch(BASE_URL + `/admin/work/delete/${id}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer" + token
        }
    }).then(response => {
        if (response.status === 200) {
            window.location.href = "/admin/work"
        } else if (response.status === 500) {
            console.log("Internal Server Error")
            window.location.href = "/500"
        } else if (response.status === 400) {
            console.log("Bad Request")
            window.location.href = "/400"
        } else if (response.status === 401) {
            console.log("UnAuthorized")
            window.location.href = "/admin/refresh"
        } else if (response.status === 404) {
            console.log("Not Found")
            window.location.href = "/404"
        } else {
            console.log("Error")
            window.location.href = "/error"
        }
    }).catch(error => {
        console.error(error)
        window.location.href = "/error"
    })
}