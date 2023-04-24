import axiosApi from '../api/axios';
 
export const setAuthToken = token => {
   if (token) {
        axiosApi.defaults.headers.common["Authorization"] = `${token}`;
   }
   else
       delete axiosApi.defaults.headers.common["Authorization"];
}