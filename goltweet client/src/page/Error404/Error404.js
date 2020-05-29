import React from "react";
import { Link } from "react-router-dom";
import Error404Img from "../../assets/png/404.png";
import Logo from "../../assets/png/original.png";
import "./Error404.scss";

export default function Error404() {
  return (
    <div className="error404">
      <img src={Logo} alt="golTweet" />
      <img src={Error404Img} alt="Error404Img" />
      <Link to="/">Volver al Inicio</Link>
    </div>
  );
}
