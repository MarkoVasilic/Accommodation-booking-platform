import axios from 'axios';

const axiosApi = axios.create({
    baseURL: "http://127.0.0.1:8080",
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