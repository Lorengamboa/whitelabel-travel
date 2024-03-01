import React, { useEffect, useState } from "react";
import axios from "axios";

import Destinations from "./components/Destinations";
import DownloadApp from "./components/DownloadApp";
import Footer from "./components/Footer";
import Home from "./components/Home";
import Navbar from "./components/Navbar";
import Offer from "./components/Offer";
import Services from "./components/Services";
import Testimonial from "./components/Testimonial";
import Tours from "./components/Tours";

// if its dev mode, use the proxy
axios.defaults.baseURL = import.meta.env.DEV ? "/api" : import.meta.env.VITE_API_URL;
    
console.log(import.meta.env)
export default function App() {
  const [data, setData] = useState([]);
  
  useEffect(() => {
    const cliendId = import.meta.env.DEV ? "43e73d76-9cbc-4af5-9a1b-fc6d261b1709" : import.meta.env.VITE_CLIENT_ID;
    
    axios.get(`/public/client/${cliendId}`)
      .then(response => {
        setData(response.data);
      })
      .catch(error => {
        console.log(error);
      });
  } , []);
  
  return (
    <div style={{ 
      '--primary-color': import.meta.env.VITE_BRAND_PRIMARY_COLOR,
    }}>
      <Navbar logo={data.logo} title={data.name}/>
      <Home title={data.title} />
      <Services />
      <Destinations />
      <Offer />
      <Tours />
      <Testimonial />
      <DownloadApp />
      <Footer />
    </div>
  );
}
