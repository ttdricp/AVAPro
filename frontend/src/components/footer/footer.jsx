import React from 'react';
import {Col, Row, Nav} from "react-bootstrap";

const Footer = () => {
    return (
        <>
            <footer className='mt-auto'>
                <Row className={'d-flex border-top py-4 my-4'}>
                    <Col className={'row'}>
                        <Col>
                            <h5>Category</h5>
                            <Nav>
                                <Nav.Item>Privacy & Policy</Nav.Item>
                                <Nav.Item>Terms & Conditions</Nav.Item>
                                <Nav.Item>Customer Support</Nav.Item>
                                <Nav.Item>FAQs</Nav.Item>
                                <Nav.Item>About</Nav.Item>
                            </Nav>
                        </Col>
                        <Col>
                            <h5>Useful Links</h5>
                            <Nav>
                                <Nav.Item>Refer A Friend</Nav.Item>
                                <Nav.Item>Tasks</Nav.Item>
                                <Nav.Item>Articles</Nav.Item>
                                <Nav.Item>Books</Nav.Item>
                            </Nav>
                        </Col>
                        <Col>
                            <h5>Contact Us</h5>
                            <Nav>
                                {/* TODO: */}
                                <Nav.Item>+7 777 777 77 77</Nav.Item>
                                <Nav.Item>Astana, Kazakhstan</Nav.Item>
                                <Nav.Item>proglang@gmail.com</Nav.Item>
                            </Nav>
                        </Col>
                    </Col>
                    <Col className={'offset-md-3'}>
                            <p className={"text-muted"}>
                                We are a team of students who want to contribute to the development of society.  We want to explain to the younger generation that learning can be educational and fun.
                            </p>
                    </Col>
                </Row>
            
            </footer>
        </>
    );
};

export default Footer;