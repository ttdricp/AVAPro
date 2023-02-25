import React, { useState } from 'react';
import { Container, Form, Button } from 'react-bootstrap';
import {useDispatch, useSelector} from "react-redux";
import { useNavigate } from 'react-router-dom';
import {login} from '../Controllers/Auth/auth';

const LoginPage = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = (event) => setEmail(event.target.value);
  const handlePasswordChange = (event) => setPassword(event.target.value);

  return (
    <Container className="p-0">
      <h1 className="text-center mb-5">Login</h1>
      <Form>
        <Form.Group className="mb-3" controlId="formEmail">
          <Form.Label>Email address</Form.Label>
          <Form.Control
            type="email"
            placeholder="Enter email"
            value={email}
            onChange={handleEmailChange}
            required
          />
        </Form.Group>

        <Form.Group className="mb-3" controlId="formPassword">
          <Form.Label>Password</Form.Label>
          <Form.Control
            type="password"
            placeholder="Password"
            value={password}
            onChange={handlePasswordChange}
            required
          />
        </Form.Group>

        <Button onClick={() => dispatch(login(email, password)).then(()=>navigate('/'))}  variant="primary" type="submit">
          Login
        </Button>
      </Form>
    </Container>
  );
};

export default LoginPage;