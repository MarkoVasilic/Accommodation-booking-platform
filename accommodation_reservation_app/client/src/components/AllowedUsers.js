import React from 'react';
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axiosApi from '../api/axios';


const AllowedUsers = (props) => {
    let navigate = useNavigate();
    const getData = async () => (
        axiosApi.get(`/user/logged`).then((response) => {
            if(!props.userRole.includes(response.data.user.Role)){
                navigate('/');
            }
        }).catch(function (error) {
            if (error.response) {
              // Request made and server responded
              console.log(error.response.data);
              console.log(error.response.status);
              console.log(error.response.headers);
            } else if (error.request) {
              // The request was made but no response was received
              console.log(error.request);
            } else {
              // Something happened in setting up the request that triggered an Error
              console.log('Error', error.message);
            }
            navigate('/');
          })
    )

    useEffect(() => {
        getData();
    });
  return (
    <div></div>
  )
}

export default AllowedUsers