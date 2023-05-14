import React, {  useEffect, useState } from "react";
import axiosApi from "../api/axios";
import { useNavigate } from "react-router-dom";
import Alert from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import CloseIcon from "@mui/icons-material/Close";
import { useForm } from "react-hook-form";
import Box from "@mui/material/Box";


const url = "/accommodation";

function CreateAccommodationForm() {
  const [name, setName] = useState("");
  const [location, setLocation] = useState("");
  const [wifi, setWifi] = useState(false);
  const [kitchen, setKitchen] = useState(false);
  const [ac, setAC] = useState(false);
  const [parkingLot, setParkingLot] = useState(false);
  const [minGuests, setMinGuests] = useState("");
  const [maxGuests, setMaxGuests] = useState("");
  const [images, setImages] = useState([]);
  const [autoAccept, setAutoAccept] = useState(false);
  let navigate = useNavigate();
  const [failedAlert, setFailedAlert] = React.useState(false);
  const {  setError } = useForm();


  const [profile, setProfile] = useState({});

    let getData = async () => {
      try {
        axiosApi
        .get('/user/logged')
        .then((response) => {
            setProfile(response.data.user);
            }).catch(er => {
              
              
          });
      } catch (err) {
        
      }
        
    };

    useEffect(() => {
        getData();
    }, []);

  const handleImageUpload =async (event) => {
    const input = event.target;
    const file = input.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = () => {
        const base64Img = reader.result;
        setImages([...images, base64Img]);
      };
    }
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    console.log(profile)
   const hostId=profile.Id
    const accommodation = {
      name,
      hostId,
      location,
      wifi,
      kitchen,
      ac,
      parkingLot,
      minGuests: parseInt(minGuests),
      maxGuests: parseInt(maxGuests),
      images,
      autoAccept,
    };

    console.log(accommodation)
    
    try {
        const resp = await axiosApi.post(url, accommodation);

        navigate("/accommodations/host");
    } catch (err) {
      const errMes = err.response.data
      setFailedAlert(true)

      for (let key in errMes) {
          setError(key, { message: errMes[key] })
      }
    }


  };

  return (
    <div>
      <h1 style={{display:'flex', flexDirection: 'column', alignItems:'center'}}>
        Create accommodation
      </h1>

    <form onSubmit={handleSubmit} style={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
      <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '10px' }}>
        <label htmlFor="name" style={{ marginBottom: '5px' }}>Name:</label>
        <input
          id="name"
          type="text"
          value={name}
          onChange={(event) => setName(event.target.value)}
          required
          minLength={4}
          maxLength={30}
          style={{ padding: '5px', borderRadius: '5px', border: '1px solid gray', width: '100%', boxSizing: 'border-box' }}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '10px' }}>
        <label htmlFor="location" style={{ marginBottom: '5px' }}>Location:</label>
        <input
          id="location"
          type="text"
          value={location}
          onChange={(event) => setLocation(event.target.value)}
          required
          minLength={4}
          maxLength={100}
          style={{ padding: '5px', borderRadius: '5px', border: '1px solid gray', width: '100%', boxSizing: 'border-box' }}
        />
      </div>
      
      <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '10px' }}>
        <label htmlFor="minGuests" style={{ marginBottom: '5px' }}>Minimum Guests:</label>
        <input
          id="minGuests"
          type="number"
          value={minGuests}
          onChange={(event) => setMinGuests(event.target.value)}
          required
          min={1}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '10px' }}>
        <label htmlFor="maxGuests" style={{ marginBottom: '5px' }}>Minimum Guests:</label>
        <input
          id="maxGuests"
          type="number"
          value={maxGuests}
          onChange={(event) => setMaxGuests(event.target.value)}
          required
          min={2}
        />
      </div>
      <div style={{ display: 'flex', flexDirection: 'column', marginBottom: '10px' }}s>
        <label htmlFor="image" style={{ marginBottom: '5px' }}>Images:</label>
        <input
            id="image"
            type="file"
            accept="image/*"
            onChange={handleImageUpload}
            style={{width:'175px'}}
            multiple
        />
        <div>
    </div>
    </div>
    <div>
        <label htmlFor="wifi" style={{ marginRight: '15px' ,marginBottom:'5px',padding: '5px'}}>WiFi:</label>
        <input
          id="wifi"
          type="checkbox"
          checked={wifi}
          onChange={(event) => setWifi(event.target.checked)}
          
        />
      </div>
      <div >
        <label htmlFor="kitchen" style={{ marginBottom: '5px' ,marginRight:'25px'}}>Kitchen:</label>
        <input
          id="kitchen"
          type="checkbox"
          checked={kitchen}
          onChange={(event) => setKitchen(event.target.checked)}
          
        />
      </div>
      <div >
        <label htmlFor="ac" style={{ marginBottom: '5px',marginRight:'25px' }}>AC:</label>
        <input
          id="ac"
          type="checkbox"
          checked={ac}
          onChange={(event) => setAC(event.target.checked)}
          
        />
      </div>
      <div >
        <label htmlFor="parkingLot" style={{ marginBottom: '5px',marginRight:'15px' }}>Parking Lot:</label>
        <input
          id="parkingLot"
          type="checkbox"
          checked={parkingLot}
          onChange={(event) => setParkingLot(event.target.checked)}
         
        />
      </div>
    <div >
        <label htmlFor="autoAccept" style={{ marginBottom: '5px',marginRight:'15px' }}>Auto accept:</label>
        <input
          id="autoAccept"
          type="checkbox"
          checked={autoAccept}
          onChange={(event) => setAutoAccept(event.target.checked)}
          
        />
      </div>
    <button type="submit" style={{ padding: '10px', borderRadius: '5px', backgroundColor: '#4CAF50', color: 'white', border: 'none', cursor: 'pointer', marginTop: '30px' }}>Create Accommodation</button>
    </form>

    <Box sx={{ width: "100%" }}>
          <Collapse in={failedAlert}>
              <Alert
                  severity="error"
                  action={
                      <IconButton
                          aria-label="close"
                          color="inherit"
                          size="small"
                          onClick={() => {
                              setFailedAlert(false);
                              navigate("/");
                          }}
                      >
                          <CloseIcon fontSize="inherit" />
                      </IconButton>
                  }
                  sx={{ mb: 2 }}
              >
                  Error
              </Alert>
          </Collapse>
      </Box>

    </div>
  )}
  
  export default CreateAccommodationForm;
