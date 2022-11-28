import { Token } from "../types/Types"

const BASE_URL = process.env.BASE_API_URL || "http://localhost:8080"

const set = (key:string, value:any) => localStorage.setItem(key, value)
const get = (key: string) => getLocalStorage(key)
const getLocalStorage = (key:string) => {
    const ret = localStorage.getItem(key)
    const retToken = localStorage.getItem('token')
    if (ret && retToken) {
        return ret;
    }
    return null
}

export const signUp = async (username: string, password: string) => {
    await fetch(BASE_URL + "/admin/signup", {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "username":username,
            "password":password
        })
    }).then((response) => {
        if(!response.ok) {
            
        } else {
            
        }
    }).then(() => {
        window.location.href = '/login'
    }).catch((error) => {
        console.log(error.statusText)
    })
}

export const Login = async (username: string, password: string) => {
    await fetch(BASE_URL + "/admin/login", {
        method: 'POST',
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            "username": username,
            "password": password
        })
    }).then((response) => {
        if (response.status === 200) {
            return response.json()
        } else {
            throw {
                status: response.status,
                statusText: response.statusText
            }
        }
    }).then(data => {
        const token: Token = {
            AccessToken: data.AccessToken,
            RefreshToken: data.RefreshToken,
        }
        tokenSet(token)
        set("isLogin", true)
        window.location.href = '/admin'
    }).catch((error) => {
        console.error(error.statusText)
        if (error.status === 400 || error.status === 500 || error.status === 404) {
            window.location.href = `/${error.status}`
        } else {
            window.location.href = `/error`
        }
    })
}

export const Logout = async () => {
    if (!get("isLogin")) {
        throw {
            status: 400,
            statusText: "Bad Request"
        }
    }
    const token = get('AccessToken')
    await fetch(BASE_URL + "/admin/logout", {
        method: 'GET',
        headers: {
            'Content-Type':'application/json',
            'Authorization': 'Bearer' + token
        }
    }).then(response => {
        if (response.status === 200) {
            localStorage.removeItem('AccessToken')
            localStorage.removeItem('RefreshToken')
            set("isLogin", false)
            window.location.href = "/"
        } else {
            throw {
                status: response.status, 
                statusText: response.statusText
            }
        }
    }).catch(error => {
        console.error(error.statusText)
        if (error.status === 400 || error.status === 500) {

        }
        return error
    })
}

export const RefreshToken = async () => {
    const token = get('RefreshToken')
    await fetch(BASE_URL + "/admin/refresh", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer" + token
        }
    }).then(response => {
        if (response.status === 200) {
            localStorage.removeItem("AccessToken")
            localStorage.removeItem("RefreshToken")
            return response.json()
        } else {
            throw {
                status: response.status,
                statusText: response.statusText
            }
        }
    }).then(data => {
        const newToken: Token = {
            AccessToken: data.AccessToken,
            RefreshToken: data.RefreshToken
        }
        tokenSet(newToken)
        window.history.back()
    }).catch(error => {
        console.error(error.statusText)
        if (error.status === 400 || error.status === 500 || error.status === 404) {
            window.location.href = `/${error.status}`
        } else if (error.status === 401) {
            window.location.href = "/login"
        } else {
            window.location.href = "/error"
        }
    })
}

const tokenSet = (token: Token) => {
    if (token.AccessToken && token.RefreshToken) {
        set('AccessToken', token.AccessToken)
        set('RefreshToken', token.RefreshToken)
    }
}
