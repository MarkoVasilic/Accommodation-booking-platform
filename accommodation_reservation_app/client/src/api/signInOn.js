import axiosApi from "./axios";


export async function sumbitRegistration(data){
    const response = await axiosApi.post('/user', data)
    return response     
}

export async function sumbitLogin(data){
        const response = await axiosApi.post('/login', data)
        return response
}

