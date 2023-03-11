import axiosApi from "./axios";


export async function sumbitRegistration(data){
    const response = await axiosApi.post('/users/signup/', data)
    return response     
}

export async function sumbitLogin(data){
        const response = await axiosApi.post('/users/login/', data)
        return response
}

