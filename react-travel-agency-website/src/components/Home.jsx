import React, { useState } from "react";
import styled from "styled-components";
import HeroImage from "../assets/hero.png";

export default function Home(props) {
  const [value, setValue] = useState("$500 - $10,000");
  return (
    <Section>
      <div className="background">
        <img src={HeroImage} alt="Hero" />
      </div>
      <div className="content">
        <div className="info">
          <h1>{props.title}</h1>
        </div>
      </div>
    </Section>
  );
}

const Section = styled.section`
   position: relative;
  .background {
    background-size: cover;
    img {
      width: 100%;
    }
  }
  .content {
    .info {
      position: absolute;
      top: 5rem;
      margin-left: 8rem;
      h1 {
        font-size: 5rem;
        margin-bottom: 2rem;
      }
    }
  @media screen (max-width: 1080px) {
    .background {
      img {
        height: 50vh;
      }
    }
    .content {
      .info {
        margin-left: 2rem;
        h1 {
          font-size: 2rem;
          margin-bottom: 1rem;
        }
      }
    }
  }
`;
