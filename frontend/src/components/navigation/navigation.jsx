import React, {useState} from 'react';
import {useDispatch, useSelector} from "react-redux";
import { useNavigate } from 'react-router-dom';
import { logoutPage } from '../../Controllers/Auth/auth';

import {
    Container,
    Form,
    Nav,
    Navbar,
    Button
} from "react-bootstrap";
import {Link} from "react-router-dom";

const Navigation= () => {

    const navigate = useNavigate();
    const dispatch = useDispatch();

    const isAuth = useSelector(state => state.client.isAuth);
    const [search, setSearch] = useState('');
    return (
        <Navbar bg="dark" variant="dark" expand="lg">
          <Container>
          <Link to={'/'} className={'navbar-brand'}>
                        AVAPro
                    </Link>
            <Navbar.Toggle aria-controls="basic-navbar-nav" />
            <Navbar.Collapse id="basic-navbar-nav">
              <Nav className="me-auto">
              <Link to={'/'} className={'nav-link'}>Home</Link>
              <Link to={'/tasks'} className={'nav-link'}>Tasks</Link>
              <Link to={'/rankings'} className={'nav-link'}>Rankings</Link>

              </Nav>

             {
                !isAuth && (

                    <Nav className={'ms-auto d-flex flex-row'}>
                            <Link to={'/signup'} className={'btn btn-outline-secondary me-2'}>Sign Up</Link>
                            <Link to={'/login'} className={'btn btn-primary'}>Log In</Link>
                    </Nav>

                )
             }

{
                isAuth && (

                    <Nav className={'ms-auto d-flex flex-row'}>
                            <Button onClick={() => dispatch(logoutPage()).then(()=>navigate('/'))}>Logout</Button>
                    </Nav>
                )
             }
            </Navbar.Collapse>
          </Container>
        </Navbar>
      );
};

export default Navigation;