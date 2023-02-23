import React from 'react'
import {Container} from "react-bootstrap";
import './HomePage.css';

const HomePage = () => {
    return (
        <>
            <Container className="p-0">
            <section class="py-5 text-center container">
    <div class="row py-lg-5">
      <div class="col-lg-6 col-md-8 mx-auto ">
        <h1 class="fw-light">Welcome to the Programming Universe!</h1>
        <p class="lead text-muted">Here you will be able to compete with other users while studying your dedicated programming language!</p>
        <p>
          <a href="/login" class="btn btn-primary my-2">Let me in!</a>
        </p>
      </div>
    </div>
  </section>
            </Container>
        </>
    )
}

export default HomePage;