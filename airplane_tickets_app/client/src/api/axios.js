import axios from 'axios';
import { URL } from '../env_constants';

const axiosApi = axios.create({
    baseURL: URL,
    headers: {
        Accept: 'application/json',
    },
});

axiosApi.interceptors.response.use(function (response) {
    // Any status code that lie within the range of 2xx cause this function to trigger
    // Do something with response data
    return response;
}, function (error) {
    console.log(error)
    if (error.response.request.status == 403) {
        localStorage.removeItem("token");
        delete axiosApi.defaults.headers.common[
            "Authorization"
        ];
        window.location.assign('/')
    }
    return Promise.reject(error);
});

export default axiosApi;