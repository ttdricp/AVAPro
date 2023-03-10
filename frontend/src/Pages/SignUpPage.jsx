import React, { useState } from 'react';
import { Container, Form, Button } from 'react-bootstrap';
import {useDispatch, useSelector} from "react-redux";
import { useNavigate } from 'react-router-dom';
import {registration} from '../Controllers/Auth/auth';

const SignUpPage = () => {
    const navigate = useNavigate();
    const dispatch = useDispatch();
  const [first_name, setFirstName] = useState('');
  const [last_name, setLastName] = useState('');

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [phone, setPhone] = useState('');

  const handleFirstNameChange = (event) => setFirstName(event.target.value);
  const handleLastNameChange = (event) => setLastName(event.target.value);
  const handleEmailChange = (event) => setEmail(event.target.value);
  const handlePasswordChange = (event) => setPassword(event.target.value);
  const handlePhoneChange = (event) => setPhone(event.target.value);


  return (
    <Container className="p-0">
        <h1 className="text-center mb-5">Sign Up</h1>
        <Form>

          <Form.Group className="mb-3" controlId="formName">
            <Form.Label>First Name</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter your first name"
              value={first_name}
              onChange={handleFirstNameChange}
              required
            />
            <Form.Text className="text-muted">
            We'll never share your personal data with anyone else.
          </Form.Text>
          </Form.Group>

          <Form.Group className="mb-3" controlId="formName">
            <Form.Label>Last Name</Form.Label>
            <Form.Control
              type="text"
              placeholder="Enter your last name"
              value={last_name}
              onChange={handleLastNameChange}
              required
            />
          </Form.Group>

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

          <Form.Group className="mb-3" controlId="formPassword">
            <Form.Label>Phone</Form.Label>
            <Form.Control
              type="text"
              placeholder="Phone"
              value={phone}
              onChange={handlePhoneChange}
              required
            />
          </Form.Group>
   
          <Button onClick={() => dispatch(registration(first_name, last_name, password, email, phone)).then(()=>navigate('/'))} variant="primary" type="submit">
            Sign Up
          </Button>
        </Form>
    </Container>
  );
};

export default SignUpPage;